package input

import (
	"influ-dojo/api/domain/repository"
)

type CalcPoint struct {
	Path     string
	WorkRepo repository.Work
}

const (
	Daily     = "daily"
	Weekly    = "weekly"
	Monthly   = "monthly"
	TweetBase = 5.
	RepBase   = 10.
	FavBase   = 500.
)

func (in *CalcPoint) CalcPoint() error {
	works, err := in.WorkRepo.Load()
	if err != nil {
		return err
	}

	for _, work := range works {
		switch in.Path {
		case Daily:
			work.CalcPoint(TweetBase, RepBase, FavBase)
		case Weekly:
			work.CalcPoint(TweetBase*7, RepBase*7, FavBase*7)
		case Monthly:
			//TODO 月に合わせて日数計算
			work.CalcPoint(TweetBase*30, RepBase*30, FavBase*30)
		}

		if err := in.WorkRepo.Save(work); err != nil {
			return err
		}
	}

	return nil
}
