package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type configuration struct {
	MaxLimit int64
	Auth     *AuthConfiguration
	Database *DatabaseConfiguration
}

type AuthConfiguration struct {
	SigningKey string
	TokenTTL   int64
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

	value, ok = os.LookupEnv("ACCOUNT_MAX_LIMIT")
	maxLimit, err := strconv.ParseInt(value, 10, 64)
	if !ok || err != nil {
		log.Println("No max limit passed. Using default 5.000.000,00 limit")
		maxLimit = 5_000_000_00
	}

	value, ok = os.LookupEnv("TOKEN_TTL")
	tokenTTL, err := strconv.ParseInt(value, 10, 64)
	if !ok || err != nil {
		log.Println("No token ttl passed. Using default 86400 ttl")
		tokenTTL = 86400
	}

	return &configuration{
		MaxLimit: maxLimit,
		Auth: &AuthConfiguration{
			SigningKey: os.Getenv("SIGNING_KEY"),
			TokenTTL:   tokenTTL,
		},
		Database: &DatabaseConfiguration{
			Address:  os.Getenv("PG_IP"),
			Port:     pgPort,
			Name:     os.Getenv("PG_DATABASE"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
		},
	}
}
