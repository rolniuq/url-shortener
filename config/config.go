package config

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type Config struct {
	Port string
	DB   DB
}

func NewConfig() *Config {
	return &Config{
		Port: "8080",
	}
}
