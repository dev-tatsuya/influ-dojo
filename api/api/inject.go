package api

import (
	domainClient "influ-dojo/api/domain/client"
	infraClient "influ-dojo/api/infrastructure/client"
)

type Dependency struct {
	FollowerClient domainClient.Follower
}

func Inject(cfg *Config) (*Dependency, error) {
	return &Dependency{
		FollowerClient: infraClient.NewFollower(
			cfg.Twitter.AccessToken,
			cfg.Twitter.AccessTokenSecret,
			cfg.Twitter.ConsumerKey,
			cfg.Twitter.ConsumerSecret,
		),
	}, nil
}
