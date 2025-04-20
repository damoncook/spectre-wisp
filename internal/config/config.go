package config

import (
	"os"

	"wisp/internal/logging"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	WebAddr string `yaml:"WebAddr"`
	WebPort string `yaml:"WebPort"`
	LogType string `yaml:"LogType"`
}

func LoadConfigFromYAML(f string) (*ServerConfig, error) {
	// setup logging for the config reader; always STDOUT
	var logConfig = logging.InitLogger("wispconfig", "console")

	file, err := os.ReadFile(f)
	if err != nil {
		logging.PrintErr(logConfig, err)
		return nil, err
	}

	var sc ServerConfig

	err = yaml.Unmarshal(file, &sc)
	if err != nil {
		logging.PrintErr(logConfig, err)
		return nil, err
	}

	logging.PrintGreen(logConfig, "✓ configuration loaded from "+f)
	return &sc, nil
}
