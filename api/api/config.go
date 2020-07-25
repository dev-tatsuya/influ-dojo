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
		Dir     string `yaml:"dir"`
		IsDebug bool   `yaml:"is_debug"`
	} `yaml:"log"`
	Twitter struct {
		AccessToken       string `yaml:"access_token"`
		AccessTokenSecret string `yaml:"access_token_secret"`
		ConsumerKey       string `yaml:"consumer_key"`
		ConsumerSecret    string `yaml:"consumer_secret"`
	} `yaml:"twitter"`
	DB struct {
		Host     string `yaml:"host" validate:"required"`
		Port     int    `yaml:"port" validate:"required"`
		User     string `yaml:"user" validate:"required"`
		Password string `yaml:"password" validate:"required"`
		Database string `yaml:"database" validate:"required"`
	} `yaml:"db"`
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
