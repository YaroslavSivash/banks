package http

import (
	"bank/banks"
	"github.com/labstack/echo/v4"
)

func RegisterHttpEndPointsBanks(e *echo.Echo, uc banks.UseCase) {
	h := NewHandler(uc)

	e.GET("/", h.AllBanksHandler)
	e.POST("/createBank", h.CreateBankHandler)
	e.PUT("/updateBank", h.UpdateBankHandler)
	e.DELETE("/deleteBank", h.DeleteBankHandler)
}

