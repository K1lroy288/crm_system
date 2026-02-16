package config

import (
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	AppHost          string
	AppPort          string
	JwtSecret        string
	UserServiceHost  string
	UserServicePort  string
	VisitServiceHost string
	VisitServicePort string
}

var (
	instance *Config
	once     sync.Once
)

func loadConfig() *Config {
	once.Do(func() {
		instance = &Config{
			AppHost:          os.Getenv("APP_HOST"),
			AppPort:          os.Getenv("APP_PORT"),
			JwtSecret:        os.Getenv("JWT_SECRET"),
			UserServiceHost:  os.Getenv("USER_SERVICE_HOST"),
			UserServicePort:  os.Getenv("USER_SERVICE_PORT"),
			VisitServiceHost: os.Getenv("VISIT_SERVICE_HOST"),
			VisitServicePort: os.Getenv("VISIT_SERVICE_PORT"),
		}
	})

	return instance
}

func GetConfig() *Config {
	if instance == nil {
		return loadConfig()
	}

	return instance
}
