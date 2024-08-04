package main

import (
	"fmt"

	"github.com/zeindevs/gospotify/internal/config"
	"github.com/zeindevs/gospotify/internal/service"
)

func main() {
	cfg := config.NewConfig()
	auth := service.NewAuthService(cfg)

	res, err := auth.ClientLogin()
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
