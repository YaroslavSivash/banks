package main

import (
	"bank/config"
	"bank/server"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.Init(); err != nil{
		log.Fatalf("%s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading env file: %s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(":"+viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
