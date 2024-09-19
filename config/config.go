package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Server   ServerConfig
	JWT      JWTConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Environment string
}

type ServerConfig struct {
	Port     string
	Debug    bool
	TimeZone string
}

type JWTConfig struct {
	AccessTokenSecret  string
	AccessTokenTTL     int
	RefreshTokenSecret string
	RefreshTokenTTL    int
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

const configFilePath = "config/config.yaml"

func LoadConfig() (Config, error) {
	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
