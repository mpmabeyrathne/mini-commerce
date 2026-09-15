package main

import (
	"os"
	"strconv"
)

type Config struct {
	Port int
	Env  string
}

func LoadConfig() (*Config, error) {
	envStr := os.Getenv("APP_ENV")
	if envStr == "" {
		envStr = "development"
	}
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	return &Config{
		Port: port,
		Env:  envStr,
	}, nil
}
