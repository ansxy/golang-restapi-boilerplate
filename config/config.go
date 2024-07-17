package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App      App
	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	URI      string
}

type App struct {
	Port string
	Host string
	Name string
}

func NewConfig() *Config {

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed-Read-Envrionment-File \n%v\n", err.Error())
	}

	v := viper.GetViper()

	viper.AutomaticEnv()

	return &Config{
		App: App{
			Port: v.GetString("APP_PORT"),
			Host: v.GetString("APP_HOST"),
			Name: v.GetString("APP_NAME"),
		},
		Postgres: PostgresConfig{
			Host:     v.GetString(""),
			Port:     v.GetString(""),
			Database: v.GetString(""),
			User:     v.GetString(""),
			Password: v.GetString(""),
			URI:      v.GetString(""),
		},
	}
}
