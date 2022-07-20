package config

type Config struct{}

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
}
