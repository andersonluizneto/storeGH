package configs

import "os"

var cfg *config

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Password string
	Database string	
}

type config struct {
	API APIConfig
	DB DBConfig
}

func Load() {
	cfg = new(config)

	cfg.API = APIConfig{
		Port: os.Getenv("SERVER_PORT"),
	}

	cfg.DB = DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}