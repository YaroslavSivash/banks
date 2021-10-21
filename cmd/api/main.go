package main

import (
	"bank/config"
	_ "bank/docs"
	"bank/server"
	"github.com/spf13/viper"
	"log"
)

// @title Echo Swagger Banks API
// @version 2.0
// @ описание Это примерный сервер.
// @termsOfService http://swagger.io/terms/

// @ contact.name Поддержка API
// @ contact.url http://www.swagger.io/support
// @ contact.email support@swagger.io

// @ license.name Apache 2.0
// @ license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Access-token
// @host localhost: 9001
// @BasePath /
// @schemes http
func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(":" + viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
