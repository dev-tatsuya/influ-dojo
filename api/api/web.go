package api

import (
	"influ-dojo/api/interface/http/handler"
	myMiddleware "influ-dojo/api/interface/http/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/xerrors"
)

func NewWebServer(dependency *Dependency) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	setRouting(dependency, e)

	return e
}

func setRouting(dependency *Dependency, e *echo.Echo) {
	e.Use(middleware.Recover(), middleware.LoggerWithConfig(middleware.LoggerConfig{}))

	webAPI := e.Group("/api", myMiddleware.MakeErrorHandlerMiddleware())
	webAPI.GET("/hello", handler.MakeHelloHandler())
	webAPI.GET("/followers", handler.MakeFollowersHandler(dependency.FollowerClient))

	rank := webAPI.Group("/rank", myMiddleware.MakeRankHandlerMiddleware())
	rank.GET("/daily", handler.MakeDailyRankHandler(
		dependency.FollowerClient, dependency.UserRepo, dependency.DailyWorkRepo, dependency.DailyResultRepo))
}

func StartWebServer(e *echo.Echo, cfg *Config) error {
	if err := e.Start(cfg.Server.Listener); err != nil {
		return xerrors.Errorf("failed to start web server: %w", err)
	}

	return nil
}
