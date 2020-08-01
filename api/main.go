package main

import (
	"flag"
	"influ-dojo/api/api"
	dataModel "influ-dojo/api/infrastructure/persistence/model"
	appLog "influ-dojo/api/log"
	"log"
	"math/rand"
	"runtime"
	"time"
)

var configFilePath string

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&configFilePath, "config", "./api/config.yml", "config file path")
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	cfg, err := api.LoadConfig(configFilePath)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	appLog.NewLogger(cfg.Log.Dir)

	log.Print("starting web server")

	db, err := api.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close mdm database: %+v", err)
			return
		}
		log.Printf("close mdm database")
	}()

	db.AutoMigrate(
		&dataModel.User{},
		&dataModel.DailyWork{},
		&dataModel.DailyResult{},
		&dataModel.WeeklyWork{},
		&dataModel.WeeklyResult{},
		&dataModel.MonthlyWork{},
		&dataModel.MonthlyResult{},
	)

	rd, err := api.ConnectRedis(cfg)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	dependency, err := api.Inject(cfg, db, rd)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	api.NewCron(dependency)

	e := api.NewWebServer(dependency)
	if err := api.StartWebServer(e, cfg); err != nil {
		log.Fatalf("%+v", err)
	}
}
