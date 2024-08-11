package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port       string `env:"PORT"`
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbName     string `env:"DB_NAME"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	JWTSecret  string `env:"JWT_SECRET"`
}

func GetConfig() Configuration {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading dotenv file", err)
	}
	configuration := Configuration{}
	err = gonfig.GetConf("", &configuration)
	if err != nil {
		fmt.Println("error in config:", err)
	}
	return configuration
}
