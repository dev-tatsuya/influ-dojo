package api

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/xerrors"
)

const (
	mySQLMaxIdleConnections = 10
	mySQLMaxOpenConnections = 50
)

var ctx = context.Background()

func ConnectDB(cfg *Config) (*gorm.DB, error) {
	dbCfg := cfg.DB
	db, err := connectMySQL(dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Database, dbCfg.Port, cfg.Log.IsDebug)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectRedis(cfg *Config) (*redis.Client, error) {
	rd := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: "",
		DB:       0,
	})

	if _, err := rd.Ping(ctx).Result(); err != nil {
		return nil, xerrors.Errorf("failed to open redis: %w", err)
	}

	return rd, nil
}

func connectMySQL(user, password, host, database string, port int, isDebug bool) (*gorm.DB, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FTokyo",
		user, password, host, port, database)
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		return nil, xerrors.Errorf("failed to open database: %w", err)
	}

	db.DB().SetMaxIdleConns(mySQLMaxIdleConnections)
	db.DB().SetMaxOpenConns(mySQLMaxOpenConnections)
	//db.LogMode(isDebug)

	return db, nil
}
