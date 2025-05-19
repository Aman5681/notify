package config

import (
	"os"
)

type Config struct {
	Host     string
	User     string
	Password string
	Port     string
	DBName   string
	SSLMode  string
}

func LoadConfig() Config {
	config := Config{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Port:     os.Getenv("PORT"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}

	return config
}
