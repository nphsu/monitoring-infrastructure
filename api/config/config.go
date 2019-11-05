package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var config *Config

type Config struct {
	Dbhost string `yaml:"dbhost"`
}

func Load(path string) error {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(result, &config)
}

func Get() *Config {
	return config
}
