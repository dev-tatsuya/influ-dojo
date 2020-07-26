package handler

import (
	"fmt"
	"influ-dojo/api/domain/client"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeFollowersHandler(follower client.Follower) echo.HandlerFunc {
	return func(c echo.Context) error {
		count, err := follower.CountFollowers()
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, fmt.Sprintf("number of followers: %d\n", count))
	}
}
