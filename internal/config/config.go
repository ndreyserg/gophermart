package config

import (
	"flag"
	"os"
)

type Config struct {
	RunAddress           string
	DatabaseURI          string
	AccrualSystemAderess string
	SecretKey            string
}

func NewConfig() *Config {
	conf := Config{}
	flag.StringVar(&conf.RunAddress, "a", "localhost:8080", "run address")
	flag.StringVar(&conf.DatabaseURI, "d",
		"postgres://gophermart:gophermart@localhost:5432/gophermart?sslmode=disable",
		"database uri",
	)
	flag.StringVar(&conf.AccrualSystemAderess, "r", "/", "accrual system aderess")
	flag.StringVar(&conf.SecretKey, "s", "secret_key", "secret key")
	flag.Parse()

	if envRunAddres := os.Getenv("RUN_ADDRESS"); envRunAddres != "" {
		conf.RunAddress = envRunAddres
	}
	if envDatabaseURI := os.Getenv("DATABASE_URI"); envDatabaseURI != "" {
		conf.DatabaseURI = envDatabaseURI
	}
	if envAccrualSystemAderess := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envAccrualSystemAderess != "" {
		conf.AccrualSystemAderess = envAccrualSystemAderess
	}

	return &conf
}
