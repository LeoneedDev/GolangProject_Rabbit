package configs

import (
	"os"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

type BotConfig struct {
	Token string
}

type Config struct {
	Bot BotConfig
}

func New() *Config {
	return &Config{
		Bot: BotConfig{
			Token: os.Getenv("TOKEN"),
		},
	}
}