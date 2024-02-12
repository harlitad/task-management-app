package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const JWTSECRETKEY = "testing"

type Config struct {
	PostgreHost     string `env:"POSTGRE_HOST"`
	PostgreUsername string `env:"POSTGRE_USERNAME"`
	PostgrePassword string `env:"POSTGRE_PASSWORD"`
	PostgrePort     string `env:"POSTGRE_PORT"`
	PostgreDBName   string `env:"POSTGRE_DB_NAME"`
	BaseUrl         string `env:"BASE_URL" envDefault:"/task-management"`
}

func ParseConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}

	err = env.Parse(&config)
	if err != nil {
		log.Fatal("unable to parse environment variables. ", err)
	}

	return &config
}
