package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Username       string `json:"db_username"`
	Password       string `json:"db_password"`
	Host           string `json:"host"`
	Port           string `json:"db_port"`
	DBName         string `json:"db_name"`
	PoolMaxOpen    int    `json:"pool_max_open"`
	PoolMaxIdle    int    `json:"pool_max_idle"`
	MaxIdleSeconds int    `json:"max_idle_seconds"`
	MaxLifeSeconds int    `json:"max_life_seconds"`
}

type dbs struct {
	ChatGroupDB *sql.DB
}

var (
	globalDBs = new(dbs)
)

func GetDBs() *dbs {
	return globalDBs
}

func initDBs() {
	globalDBs.ChatGroupDB = initDB(GetConfig().ChatGroupDB)
}

func initDB(cfg *Database) *sql.DB {
	pool, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s)/%s?parseTime=true",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.DBName))
	if err != nil {
		log.Fatal("start db error", err.Error())
	}

	if err = pool.Ping(); err != nil {
		log.Fatal("reach db error", err.Error())
	}

	pool.SetMaxIdleConns(cfg.PoolMaxIdle)
	pool.SetMaxOpenConns(cfg.PoolMaxOpen)
	pool.SetConnMaxIdleTime(
		time.Duration(cfg.MaxIdleSeconds) * time.Second)
	pool.SetConnMaxLifetime(
		time.Duration(cfg.MaxLifeSeconds) * time.Second)

	return pool
}
