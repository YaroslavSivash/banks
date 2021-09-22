package http

import (
	"bank/banks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Handler struct {
	useCase banks.UseCase
}

func NewHandler (useCase banks.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}


func (h *Handler) AllBanksHandler(c echo.Context) error  {

	data, err :=h.useCase.AllBanks(c)
	if err != nil{
		log.Error(err)
		log.Info("123")
		return echo.NewHTTPError(http.StatusInternalServerError, "error in allbankshandler")
	}

	return c.JSON(http.StatusOK, data)
}

func (h *Handler) CreateBankHandler(c echo.Context) error {
	return nil
}

func (h *Handler) UpdateBankHandler(c echo.Context) error {
	return nil
}

func (h *Handler) DeleteBankHandler(c echo.Context) error {
	return nil
}