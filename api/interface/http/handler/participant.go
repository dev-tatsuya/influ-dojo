package handler

import (
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeParticipantHandler(user repository.User) echo.HandlerFunc {
	return func(c echo.Context) error {
		in, err := newParticipantInput(c)
		if err != nil {
			return err
		}
		in.UserRepo = user

		out, err := in.Participant()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, out)
	}
}

func newParticipantInput(c echo.Context) (*input.Participant, error) {
	in := new(input.Participant)
	if err := c.Bind(in); err != nil {
		return nil, err
	}

	return in, nil
}
