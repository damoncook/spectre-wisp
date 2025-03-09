package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	WebAddr string `yaml:"WebAddr"`
	WebPort string `yaml:"WebPort"`
	LogType string `yaml:"LogType"`
}

func LoadConfigFromYAML(f string) (*ServerConfig, error) {
	file, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var sc ServerConfig

	err = yaml.Unmarshal(file, &sc)
	if err != nil {
		return nil, err
	}

	return &sc, nil
}
