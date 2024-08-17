package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/zeindevs/gospotify/internal/util"
)

var (
	port = 5001
)

type Config struct {
	PORT          int
	CLIENT_ID     string
	CLIENT_SECRET string
}

func NewConfig() *Config {
	var err error
	portStr, ok := os.LookupEnv("PORT")
	if ok {
		port, err = strconv.Atoi(portStr)
		util.ErrorPanic(err)
	}

	return &Config{
		PORT:          port,
		CLIENT_ID:     os.Getenv("CLIENT_ID"),
		CLIENT_SECRET: os.Getenv("CLIENT_SECRET"),
	}
}
