package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

var c *Config

func init() {
	c = New()
}

type Config struct {
	name       string
	configType string
	path       string
	addPath    bool
}

func New() *Config {
	c := new(Config)
	c.name = "profiles"
	c.configType = "yaml"
	c.path = "$HOME/.switchcontext"
	c.addPath = false
	if err := c.readConfigFile(); err != nil {
		fmt.Println("Error reading config in home file:", err)
	}
	return c
}

func (c *Config) SetName(name string) {
	c.name = name
}

func (c *Config) SetConfigType(configType string) {
	c.configType = configType
}

func SetPath(path string) error {
	return c.setPath(path)
}
func Path() error {
	if c.path == "" {
		return errors.New("path not set")
	}
	return nil
}

func (c *Config) setPath(path string) error {
	c.path = path
	return c.readConfigFile()
}

func ReadConfigFile() error {
	return c.readConfigFile()
}

func (c *Config) readConfigFile() error {
	viper.SetConfigName(c.name)
	viper.SetConfigType(c.configType)
	viper.AddConfigPath(c.path)
	if c.addPath {
		viper.AddConfigPath(".")
	}

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) string {
	return c.get(key)
}

func GetInt(key string) int {
	return c.getInt(key)
}

func (c *Config) get(key string) string {
	return viper.GetString(key)
}
func (c *Config) getInt(key string) int {
	return viper.GetInt(key)
}

func GetStringSlice(key string) []string {
	return c.getStringSlice(key)
}

func (c *Config) getStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}
func GetMap(key string) map[string]interface{} {
	return c.getMap(key)
}

func (c *Config) getMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}
