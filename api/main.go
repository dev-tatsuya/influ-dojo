package main

import (
	"flag"
	"influ-dojo/api/api"
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

	dependency, err := api.Inject(cfg, db)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	e := api.NewWebServer(dependency)
	if err := api.StartWebServer(e, cfg); err != nil {
		log.Fatalf("%+v", err)
	}
}
