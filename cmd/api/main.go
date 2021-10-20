package main

import (
	"bank/config"
	"bank/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(":" + viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
