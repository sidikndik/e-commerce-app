package util

import (
	"e-commerce-app/helper"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	AppName  string
	Debug    bool
	Port     string
	DBConfig DBConfig
}

type DBConfig struct {
	DBName     string
	DBUsername string
	DBPassword string
	DBHost     string
	TimeZone   string
}

func ReadConfig() (Configuration, error) {
	err := godotenv.Load()
	if err != nil {
		return Configuration{}, err
	}
	return Configuration{
		AppName: os.Getenv("APP_NAME"),
		Debug:   helper.StringToBool(os.Getenv("DEBUG")),
		Port:    os.Getenv("PORT"),
		DBConfig: DBConfig{
			DBName:     os.Getenv("DB_NAME"),
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
		},
	}, nil
}
