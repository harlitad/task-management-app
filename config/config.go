package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const JWTSECRETKEY = "testing"

type Config struct {
	PostgreHost     string `env:"POSTGRE_HOST" envDefault:"localhost"`
	PostgreUsername string `env:"POSTGRE_USERNAME"`
	PostgrePassword string `env:"POSTGRE_PASSWORD"`
	PostgrePort     string `env:"POSTGRE_PORT"`
	PostgreDBName   string `env:"POSTGRE_DB_NAME"`
	BaseUrl         string `env:"BASE_URL" envDefault:"/task-management"`
	LogLevel        string `env:"LOG_LEVEL" envDefault:"debug"`
	UseMongo        bool   `env:"USE_MONGO" envDefault:"false"`
	MongoDB         string `env:"MONGO_DATABASE" envDefault:"tma"`
	MongoDBUsername string `env:"MONGO_USERNAME" envDefault:"admin"`
	MongoDBPassword string `env:"MONGO_PASSWORD" envDefault:"admin"`
	MongoDBHost     string `env:"MONGO_HOST" envDefault:"localhost"`
	MongoDBPort     int    `env:"MONGO_PORT" envDefault:"27017"`
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
