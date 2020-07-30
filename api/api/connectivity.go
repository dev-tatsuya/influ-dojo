package api

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/xerrors"
)

const (
	mySQLMaxIdleConnections = 10
	mySQLMaxOpenConnections = 50
)

func ConnectDB(cfg *Config) (*gorm.DB, error) {
	dbCfg := cfg.DB
	db, err := connectMySQL(dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Database, dbCfg.Port, cfg.Log.IsDebug)
	if err != nil {
		return nil, err
	}

	return db, nil
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
