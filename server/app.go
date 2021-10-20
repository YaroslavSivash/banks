package server

import (
	"bank/banks"
	"bank/banks/delivery/http"
	"bank/banks/repository/postgres"
	"bank/banks/usecase"
	_ "bank/docs"
	"bank/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type App struct {
	httpserver *echo.Echo
	bankUC     banks.UseCase
}

func NewApp() *App {
	connectToPostgres := services.NewDbConnect()

	repoBanks := postgres.NewBankRepository(connectToPostgres)
	return &App{
		bankUC: usecase.NewBankUseCase(repoBanks),
	}
}

// @title Echo Swagger Example API
// @version 1.0
// @ описание Это примерный сервер.
// @termsOfService http://swagger.io/terms/

// @ contact.name Поддержка API
// @ contact.url http://www.swagger.io/support
// @ contact.email support@swagger.io

// @ license.name Apache 2.0
// @ license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost: 9001
// @BasePath /
// @schemes http
func (a *App) Run(port string) error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	http.RegisterHttpEndPointsBanks(e, a.bankUC)
	e.Logger.Fatal(e.Start(":" + viper.GetString("port")))

	return nil
}
