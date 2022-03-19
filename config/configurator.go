package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type configuration struct {
	Database *DatabaseConfiguration
}

type DatabaseConfiguration struct {
	Address  string
	Port     int64
	Name     string
	User     string
	Password string
}

var config *configuration

func GetConfig() *configuration {
	if config == nil {
		config = readConfig()
	}

	return config
}

func readConfig() *configuration {
	err := godotenv.Load("./config/dev.env")
	if err != nil {
		log.Println("Can't load config file")
	}

	value, ok := os.LookupEnv("PG_PORT")
	pgPort, err := strconv.ParseInt(value, 10, 64)
	if !ok || err != nil {
		log.Println("No postgres port passed. Using default 5432 PostgreSQL port")
		pgPort = 5432
	}

	return &configuration{
		Database: &DatabaseConfiguration{
			Address:  os.Getenv("PG_IP"),
			Port:     pgPort,
			Name:     os.Getenv("PG_DATABASE"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
		},
	}
}
