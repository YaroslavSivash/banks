package http

import (
	"bank/banks"
	"bank/model"
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
		//log.Info("123")
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot read json")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"banks": data,
	})
}

func (h *Handler) CreateBankHandler(c echo.Context) error {
	bank := &model.Banks{}
	//log.Info(" до bind:", bank)
	if err:=c.Bind(bank)
	err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "cannot read json")
	}
	//log.Info(bank)

	id, err := h.useCase.CreateBank(c, bank)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error() )
	}
	return c.JSON(http.StatusOK, id)
}

func (h *Handler) UpdateBankHandler(c echo.Context) error {
	bank := &model.Banks{}

	if err:=c.Bind(bank)
		err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "cannot read json")
	}

	updateBank, err := h.useCase.UpdateBank(c, bank)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error() )
	}

	return c.JSON(http.StatusOK, updateBank)
}

func (h *Handler) DeleteBankHandler(c echo.Context) error {

	bank := &model.Banks{}

	if err:=c.Bind(bank)
		err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "cannot read json")
	}

	err := h.useCase.DeleteBank(c, &model.Banks{Id: bank.Id})
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error() )
	}

	return c.JSON(http.StatusOK, "Bank successful delete")
}

func (h *Handler) CalculateHandler(c echo.Context) error {
	calculation := &model.CalculationBorrowed{}

	if err:=c.Bind(calculation)
		err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "cannot read json")
	}

	month, err := h.useCase.CalculatePayments(c, calculation)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error() )
	}

	return c.JSON(http.StatusOK, month)
}

