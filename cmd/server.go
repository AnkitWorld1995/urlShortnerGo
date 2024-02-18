package main

import (
	"github.com/joho/godotenv"
	"github.com/urlShortnerGo/config"
	"github.com/urlShortnerGo/pkg/app"
)

func main() {
	err := godotenv.Load("")
	if err != nil {
		return
	}

	configuration := config.Init()
	app.StartApplication(configuration)
}
