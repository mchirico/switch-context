package sb

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

func DoSB() {
	profiles, err := ReadConfigFile()
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) != 2 {
		log.Fatal("expected 1 argument. Recieved", len(os.Args)-1)
	}
	name := os.Args[1]

	if profile, ok := profiles[name]; ok {
		for key, value := range profile.Environment {
			log.Printf("%s: %s", key, value)
			fmt.Printf(" export %s=\"%s\"\n", key, value)
		}
		if profile.KubeContext != nil {
			command := exec.Command(
				"kubectl",
				"config",
				"use-context",
				*profile.KubeContext,
			)
			err := command.Run()
			if err != nil {
				log.Println(err)
			} else {
				log.Println("set kube context:", *profile.KubeContext)
			}
		}

	} else {
		fmt.Println("Profile not found:", name)
		fmt.Println("Available profiles are:")
		for key := range profiles {
			fmt.Println(key)
		}
	}

}

type Profiles map[string]Profile
type Profile struct {
	Environment map[string]string `yaml:"env"`
	KubeContext *string           `yaml:"kube"`
}

func ReadConfigFile() (Profiles, error) {
	configFile := filepath.Join(HomeDirectory(), ".switchcontext/profiles.yaml")

	if fileExists(configFile) {
		data, err := os.ReadFile(configFile)
		if err != nil {
			log.Fatalln("could not read config file: ", err)
		}
		type Config struct {
			Profiles Profiles `yaml:"profiles"`
		}
		var config Config
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			return nil, err
		}
		return config.Profiles, nil
	}
	log.Println("file not found:", configFile)
	return nil, fmt.Errorf("file not found: %s", configFile)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) || info.IsDir() {
		return false
	}
	return true
}

func HomeDirectory() string {
	u, _ := user.Current()
	return u.HomeDir
}
