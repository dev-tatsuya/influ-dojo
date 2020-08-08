package input

import (
	"influ-dojo/api/domain/repository"
)

type CalcPoint struct {
	WorkRepo repository.Work
}

func (in *CalcPoint) CalcPoint() error {
	works, err := in.WorkRepo.Load()
	if err != nil {
		return err
	}

	for _, work := range works {
		work.SetPoint()

		if err := in.WorkRepo.Save(work); err != nil {
			return err
		}
	}

	return nil
}
