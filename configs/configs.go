package configs

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	Environment      string `json: 'environment"`
	HTTPPort         string `josn:"http_port"`
	LogLevel         string `json:"log_level"`
	PostgresDatabase string
	PostgresHost     string
	PostgresPort     uint16
	PostgresUser     string
	PostgresPassword string
	RedisHost        string
	RedisPort        int
	SendgridEmail    string
}

func Load() Config {
	var config Config
	v := viper.New()
	v.AutomaticEnv()

	v.SetDefault("ENVIRONMENT", "development")
	v.SetDefault("LOG_LEVEL", "debug")
	v.SetDefault("HTTP_PORT", ":8081")
	v.SetDefault("POSTGRES_HOST", ":localhost")
	v.SetDefault("POSTGRES_PORT", 5432)
	v.SetDefault("POSTGREST_USER", "postgres")
	v.SetDefault("POSTGREST_PASSWORD", "12345")
	v.SetDefault("POSTGRES_DB", "usersdb")
	v.SetDefault("REDIS_PORT", 6379)

	config.Environment = v.GetString("ENVIRONMENT")
	config.HTTPPort = v.GetString("HTTP_PORT")
	config.LogLevel = v.GetString("LOG_LEVEL")
	config.PostgresDatabase = v.GetString("POSTGRES_DB")
	config.PostgresUser = v.GetString("POSTGRES_USER")
	config.PostgresPassword = viper.GetString("POSTGRES_PASSWORD")
	config.PostgresHost = viper.GetString("POSTGRES_HOST")
	config.PostgresPort = (uint16)(v.GetUint("POSTGRES_PORT"))
	config.RedisHost = v.GetString("REDIS_HOST")
	config.RedisPort = v.GetInt("REDIS_PORT")
	return config

}

func (c *Config) Validate() error {
	if c.HTTPPort == "" {
		return errors.New("http_port required")
	}
	return nil
}
