package api

import (
	"io/ioutil"

	"golang.org/x/xerrors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Listener string `yaml:"listener" validate:"required"`
	} `yaml:"server"`
	Log struct {
		Dir string `yaml:"dir"`
	} `yaml:"log"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, xerrors.Errorf("failed to read config file: %w", err)
	}

	cfg := new(Config)
	if err := yaml.Unmarshal(file, cfg); err != nil {
		return nil, xerrors.Errorf("failed to unmarshal config file: %w", err)
	}

	return cfg, nil
}
