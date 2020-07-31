package input

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
)

type Tweet struct {
	Path       string            `json:"-"`
	Bot        client.Bot        `json:"-"`
	WorkRepo   repository.Work   `json:"-"`
	ResultRepo repository.Result `json:"-"`
}

func (in *Tweet) Tweet() error {
	works, err := in.WorkRepo.LoadTop3()
	if err != nil {
		return err
	}

	results, err := in.ResultRepo.LoadTop3()
	if err != nil {
		return err
	}

	return in.Bot.Tweet(works, results, in.Path)
}
