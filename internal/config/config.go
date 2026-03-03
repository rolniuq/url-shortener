package config

import "fmt"

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type Config struct {
	DB       DB
	Charset  string
	Port     string
	LogLevel string
}

func NewConfig() *Config {
	return &Config{
		Port:     "8080",
		Charset:  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz",
		LogLevel: "INFO",
	}
}

func (cfg *Config) GetDBConnStr() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DbName)
}

func (cfg *Config) GetLogLevel() int {
	switch cfg.LogLevel {
	default:
		return 0
	}
}
