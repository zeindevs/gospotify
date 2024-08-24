package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/zeindevs/gospotify/internal/util"
)

var (
	port          = 5001
	market        = "ID"
	client_id     = ""
	client_secret = ""
)

type Config struct {
	PORT          int
	MARKET        string
	CLIENT_ID     string
	CLIENT_SECRET string
}

func NewConfig() *Config {
	var err error
	if pot, ok := os.LookupEnv("PORT"); ok {
		port, err = strconv.Atoi(pot)
		util.ErrorPanic(err)
	}
	if mkt, ok := os.LookupEnv("MARKET"); ok {
		market = mkt
	}
	if cid, ok := os.LookupEnv("CLIENT_ID"); ok {
		client_id = cid
	}
	if csc, ok := os.LookupEnv("CLIENT_SECRET"); ok {
		client_secret = csc
	}

	return &Config{
		PORT:          port,
		MARKET:        market,
		CLIENT_ID:     client_id,
		CLIENT_SECRET: client_secret,
	}
}
