package profile

import (
	"fmt"
	"github.com/mchirico/switch-context/config"
	"github.com/mchirico/switch-context/logger"
	"os"
	"path/filepath"
	"strings"
)

var p *Profile

func init() {
	p = New()
}

type Profile struct {
	path string
}

func New() *Profile {
	p := &Profile{
		path: filepath.Join(HomeDirectory(), ".switchcontext/profiles.yaml"),
	}
	config.SetPath(p.path)
	return p
}
func SetPath(path string) error {
	return p.setPath(path)
}
func (p *Profile) setPath(path string) error {
	p.path = path
	return config.SetPath(p.path)

}

func (p *Profile) log(msg string) {
	logger.Log(msg)
}

func PR(key string) (string, error) {
	out := ""
	env, err := ProfileEnvExports(key)
	if err != nil {
		profiles := profilesAvailable()
		fmt.Println("Profile not found:", key)
		fmt.Println("Available profiles are:")
		for _, p := range profiles {
			fmt.Println(p)
		}
		os.Exit(1)

	}
	ps1, err := ProfilePS1Exports(key)
	if err != nil {
		return out, err
	}
	for _, v := range env {
		fmt.Print(v)
		out += v
	}
	for _, v := range ps1 {
		fmt.Print(v)
		out += v
	}
	return out, nil
}

func profilesAvailable() []string {
	m := config.GetMap("profiles")

	out := []string{}
	for k := range m {
		out = append(out, k)
	}
	return out
}

func ProfileEnvExports(key string) ([]string, error) {
	p.log("ProfileEnvExports: profiles." + key + ".env")
	return p.exports("profiles." + key + ".env")
}
func ProfilePS1Exports(key string) ([]string, error) {
	return p.exports("profiles."+key+".bash", "PS1")
}
func (p *Profile) exports(key string, opt ...string) ([]string, error) {
	out := []string{}
	for k, v := range config.GetMap(key) {
		if len(opt) > 0 && opt[0] == "PS1" {
			out = append(out, fmt.Sprintf("export %s='%s'\n", strings.ToUpper(k), v))
		} else {
			out = append(out, fmt.Sprintf("export %s=%q\n", strings.ToUpper(k), v))
		}

	}
	if len(out) == 0 {
		p.log("No exports found for key: " + key)
		return nil, fmt.Errorf("no profile found for %s", key)
	}
	p.log("exports output returned: " + key + " " + fmt.Sprintf("\n%v\n", out))
	return out, nil
}

func HomeDirectory() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %s", err)
	}
	return home
}
