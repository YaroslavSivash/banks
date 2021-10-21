package http

import (
	"bank/banks"
	_ "bank/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterHttpEndPointsBanks(e *echo.Echo, uc banks.UseCase) {
	h := NewHandler(uc)

	//e.GET("/", h.AllBanksHandler)
	e.POST("/create-bank", h.CreateBankHandler)
	e.PUT("/update-bank", h.UpdateBankHandler)
	e.DELETE("/delete-bank", h.DeleteBankHandler)
	e.GET("/calculate", h.CalculateHandler)
	// Routes
	e.GET("/", h.HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
