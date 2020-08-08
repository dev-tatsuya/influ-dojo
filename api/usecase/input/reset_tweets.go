package input

import (
	"influ-dojo/api/domain/repository"
)

type ResetTweets struct {
	WorkRepo repository.Work
}

func (in *ResetTweets) ResetTweets() error {
	works, err := in.WorkRepo.Load()
	if err != nil {
		return err
	}

	for _, work := range works {
		work.ResetTweetsCount()

		if err := in.WorkRepo.Save(work); err != nil {
			return err
		}
	}

	return nil
}
