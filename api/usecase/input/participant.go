package input

import (
	"influ-dojo/api/domain/apperr"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/output"
)

type Participant struct {
	ScreenName string `json:"screen_name"`

	UserRepo repository.User `json:"-"`
}

func (p *Participant) Participant() (*output.Participant, error) {
	if _, err := p.UserRepo.LoadByScreenName(p.ScreenName); err != nil {
		if err == apperr.ErrRecordNotFound {
			return &output.Participant{IsParticipant: false}, nil
		}

		return nil, err
	}

	return &output.Participant{IsParticipant: true}, nil
}
