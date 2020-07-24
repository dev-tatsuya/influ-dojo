package api

import (
	"net/http"
)

type Dependency struct {}

func Inject(cfg *Config, client *http.Client) (*Dependency, error) {
	return &Dependency{}, nil
}
