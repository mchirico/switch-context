package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mchirico/switch-context/config"
	"github.com/mchirico/switch-context/file"
	"github.com/mchirico/switch-context/logger"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var p *Profile

func init() {
	p = New()
}

type Profile struct {
	path string
	last string
}

func New() *Profile {
	p := &Profile{
		path: filepath.Join(HomeDirectory(), ".switchcontext/profiles.yaml"),
		last: filepath.Join(HomeDirectory(), ".switchcontext/lastUsed.json"),
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

type Last struct {
	Key       string `json:"key"`
	TimeEnter string `json:"time_enter"`
}

func (p *Profile) putLastUsed(key string) error {
	l := &Last{Key: key, TimeEnter: time.Now().Format(time.RFC3339)}
	fmt.Println(l)
	data, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(p.last, data, 0644)
}

func LastKey() (string, error) {
	return p.lastUsed()
}

func (p *Profile) lastUsed() (string, error) {
	data, err := os.ReadFile(p.last)
	if err != nil {
		return "", err
	}
	l := &Last{}
	err = json.Unmarshal(data, l)
	if err != nil {
		return "", err
	}
	return l.Key, nil
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
	_ = p.putLastUsed(key)
	for _, v := range env {
		fmt.Print(v)
		out += v
	}
	for _, v := range ps1 {
		fmt.Print(v)
		out += v
	}

	argo, err := ProfileArgoExports(key)
	if err != nil {
		return out, err
	}
	for _, v := range argo {
		fmt.Print(v)
		out += v
	}

	files, err := ProfileFileExports(key)
	if err != nil {
		return out, err
	}
	for _, v := range files {
		if src, ok := v["src"]; ok {
			if dst, ok := v["dst"]; ok {
				_, err = file.Copy(src, dst)
				if err != nil {
					fmt.Println("Error copying file:", err)
				}
			}
		}

	}

	return out, nil
}

func ProfilesAvailable() []string {
	return profilesAvailable()
}

func profilesAvailable() []string {
	m := config.GetMap("profiles")

	out := []string{}
	for k := range m {
		out = append(out, k)
	}
	return out
}

// TODO: Save last profile used
func ProfileEnvExports(key string) ([]string, error) {
	p.log("ProfileEnvExports: profiles." + key + ".env")
	return p.exports("profiles." + key + ".env")
}

func ProfileArgoExports(key string) ([]string, error) {
	p.log("ProfileArgoExports: profiles." + key + ".argo")
	s, err := p.exports("profiles." + key + ".argo")
	if err != nil {
		return []string{}, nil
	}
	return s, nil

}

func ProfileAliasExports(key string) ([]string, error) {
	p.log("ProfileArgoExports: profiles." + key + ".alias")
	s, err := p.exportsAlias("profiles." + key + ".alias")
	if err != nil {
		return []string{}, nil
	}
	return s, nil
}

func ProfilePS1Exports(key string) ([]string, error) {
	s, err := p.exports("profiles."+key+".bash", "PS1")
	if err != nil {
		return []string{}, nil
	}
	return s, nil
}

func ProfileFileExports(key string) ([]map[string]string, error) {
	s, err := p.exportsFile("profiles." + key + ".file")
	if err != nil {
		return nil, nil
	}
	return s, nil
}

func (p *Profile) exportsAlias(key string, opt ...string) ([]string, error) {
	out := []string{}
	for k, v := range config.GetMap(key) {

		if val, ok := v.(string); ok && val == "unalias" {
			out = append(out, fmt.Sprintf("unalias %s\n", strings.ToLower(k)))
			continue
		}
		out = append(out, fmt.Sprintf("alias %s='%s'\n", strings.ToLower(k), v))
	}
	if len(out) == 0 {
		p.log("No exports found for key: " + key)
		return nil, errors.New("no profile found for " + key)
	}
	p.log("exports output returned: " + key + " " + fmt.Sprintf("\n%v\n", out))
	return out, nil
}

func (p *Profile) exportsFile(key string, opt ...string) ([]map[string]string, error) {
	out := []map[string]string{}
	for _, v := range config.GetMap(key) {
		if _, ok := v.(map[string]interface{}); !ok {
			continue
		}
		for _, v := range v.(map[string]interface{}) {
			if val, ok := v.(map[string]interface{}); ok {
				m := map[string]string{}
				if src, ok := val["src"].(string); ok {
					m["src"] = src
				}
				if dst, ok := val["dst"].(string); ok {
					m["dst"] = dst
				}
				out = append(out, m)
			}

		}

	}
	if len(out) == 0 {
		p.log("No exports found for key: " + key)
		return nil, errors.New("no profile found for " + key)
	}
	p.log("exports output returned: " + key + " " + fmt.Sprintf("\n%v\n", out))
	return out, nil
}

func (p *Profile) exports(key string, opt ...string) ([]string, error) {
	out := []string{}
	for k, v := range config.GetMap(key) {
		if len(opt) > 0 && opt[0] == "PS1" {
			out = append(out, fmt.Sprintf("export %s='%s'\n", strings.ToUpper(k), v))
		} else {
			if val, ok := v.(string); ok && val == "unset" {
				out = append(out, fmt.Sprintf("unset %s\n", strings.ToUpper(k)))
				continue
			}

			if val, ok := v.(bool); ok {
				out = append(out, fmt.Sprintf("export %s=%v\n", strings.ToUpper(k), val))
				continue
			}
			out = append(out, fmt.Sprintf("export %s=%q\n", strings.ToUpper(k), v))
		}

	}
	if len(out) == 0 {
		p.log("No exports found for key: " + key)
		return nil, errors.New("no profile found for " + key)
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
