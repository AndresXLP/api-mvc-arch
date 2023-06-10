package config

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
)

var (
	Once sync.Once
	Cfg  Config
)

type Config struct {
	Port       int    `required:"true" split_words:"true"`
	DbHost     string `required:"true" split_words:"true"`
	DbPort     int    `required:"true" split_words:"true"`
	DbUser     string `required:"true" split_words:"true"`
	DbPassword string `required:"true" split_words:"true"`
	DbName     string `required:"true" split_words:"true"`
}

func Environments() Config {
	Once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		if err = envconfig.Process("", &Cfg); err != nil {
			log.Fatalf("Error parsing environment vars: %v", err)
		}

	})

	return Cfg
}
