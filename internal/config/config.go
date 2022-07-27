package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Config struct {
	ChatGroupDB *Database `json:"chat_group_db"`
}

const (
	configFilePath = "./internal/config/config.json"
)

var (
	GlobalConfig *Config
)

func GetConfig() *Config {
	return GlobalConfig
}

func InitConfig() {
	GlobalConfig = new(Config)
	fetchConfig()
	initDBs()
}

func fetchConfig() {
	DBUrl := os.Getenv("DATABASE_URL")
	if DBUrl != "" {
		split := strings.Split(DBUrl, ":")
		dbUsername := split[0]
		dbPassword := split[1]
		dbHost := split[2]
		dbName := split[3]

		GlobalConfig.ChatGroupDB = &Database{
			Username: dbUsername,
			Password: dbPassword,
			DBName:   dbName,
			Host:     dbHost,
		}
	} else {
		configFile, _ := os.Open("./internal/config/config.json")
		defer configFile.Close()

		decoder := json.NewDecoder(configFile)

		var config Config
		err := decoder.Decode(&config)
		if err != nil {
			log.Fatal("failed to load configs ", err.Error())
		}

		GlobalConfig = &config
	}

	if GlobalConfig == nil {
		log.Fatal("failed to fetch configs")
	}
}
