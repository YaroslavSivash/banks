package server

import (
	"bank/banks"
	"bank/banks/delivery/http"
	"bank/banks/repository/postgres"
	"bank/banks/usecase"
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

func (a *App) Run(port string) error {
	e := echo.New()
	// Промежуточное ПО
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	http.RegisterHttpEndPointsBanks(e, a.bankUC)
	e.Logger.Fatal(e.Start(":" + viper.GetString("port")))

	return nil
}
