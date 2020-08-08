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

func setRouting(d *Dependency, e *echo.Echo) {
	e.Use(middleware.Recover(), middleware.LoggerWithConfig(middleware.LoggerConfig{}))
	e.Use(middleware.CORS())

	webAPI := e.Group("/api", myMiddleware.MakeErrorHandlerMiddleware())
	webAPI.GET("/hello", handler.MakeHelloHandler())
	webAPI.GET("/followers", handler.MakeFollowersHandler(d.FollowerClient))
	webAPI.GET("/participant", handler.MakeParticipantHandler(d.UserRepo))
	webAPI.POST("/favorite", handler.MakeFavoriteHandler(d.BotClient))
	webAPI.POST("/classify/tweets", handler.MakeClassifyTweetsHandler(d.TweetClient, d.DailyWorkRepo, d.WeeklyWorkRepo, d.MonthlyWorkRepo))
	webAPI.POST("/calc/point", handler.MakeCalcPointHandler(d.DailyWorkRepo))
	webAPI.POST("/reset/tweets/weekly", handler.MakeResetTweetsHandler(d.WeeklyWorkRepo))
	webAPI.POST("/reset/tweets/monthly", handler.MakeResetTweetsHandler(d.MonthlyWorkRepo))

	ranking := webAPI.Group("/ranking")
	ranking.POST("/daily", handler.MakeRankingHandler(d.DailyWorkRepo, d.DailyResultRepo))
	ranking.POST("/weekly", handler.MakeRankingHandler(d.WeeklyWorkRepo, d.WeeklyResultRepo))
	ranking.POST("/monthly", handler.MakeRankingHandler(d.MonthlyWorkRepo, d.MonthlyResultRepo))
	ranking.POST("/cache", handler.MakeCacheHandler(d.RankingQuery, d.RankingRepo))
	ranking.GET("/all", handler.MakeRankingAllHandler(d.RankingRepo))

	record := webAPI.Group("/record")
	record.POST("/daily", handler.MakeRecordHandler(d.FollowerClient, d.UserRepo, d.DailyWorkRepo, d.DailyResultRepo))
	record.POST("/weekly", handler.MakeRecordHandler(d.FollowerClient, d.UserRepo, d.WeeklyWorkRepo, d.WeeklyResultRepo))
	record.POST("/monthly", handler.MakeRecordHandler(d.FollowerClient, d.UserRepo, d.MonthlyWorkRepo, d.MonthlyResultRepo))
	record.POST("/users", handler.MakeRecordUsersHandler(d.FollowerClient, d.UserRepo, d.DailyWorkRepo, d.DailyResultRepo,
		d.WeeklyWorkRepo, d.WeeklyResultRepo, d.MonthlyWorkRepo, d.MonthlyResultRepo))

	tweet := webAPI.Group("/tweet")
	tweet.POST("/daily", handler.MakeTweetHandler(d.BotClient, d.DailyWorkRepo, d.DailyResultRepo))
	tweet.POST("/weekly", handler.MakeTweetHandler(d.BotClient, d.WeeklyWorkRepo, d.WeeklyResultRepo))
	tweet.POST("/monthly", handler.MakeTweetHandler(d.BotClient, d.MonthlyWorkRepo, d.MonthlyResultRepo))
}

func StartWebServer(e *echo.Echo, cfg *Config) error {
	if err := e.Start(cfg.Server.Listener); err != nil {
		return xerrors.Errorf("failed to start web server: %w", err)
	}

	return nil
}
