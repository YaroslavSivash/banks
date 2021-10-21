package http

import (
	"bank/banks"
	_ "bank/docs"
	"bank/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Handler struct {
	useCase banks.UseCase
}

func NewHandler(useCase banks.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// ShowAccount godoc
// @Security ApiKeyAuth
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Banks
// @Header 200 {string} Token "qwerty"
// @Router /accounts/{id} [get]
func (h *Handler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Сервер запущен и работает",
	})
}

func (h *Handler) AllBanksHandler(c echo.Context) error {

	data, err := h.useCase.AllBanks(c.Request().Context())
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot read json")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"banks": data,
	})
}

func (h *Handler) CreateBankHandler(c echo.Context) error {

	bank := &model.Banks{}
	if err := c.Bind(bank); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "cannot read json")
	}

	id, err := h.useCase.CreateBank(c.Request().Context(), bank)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}

func (h *Handler) UpdateBankHandler(c echo.Context) error {

	bank := &model.Banks{}
	if err := c.Bind(bank); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "cannot read json")
	}

	updateBank, err := h.useCase.UpdateBank(c.Request().Context(), bank)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"new bank": updateBank,
	})
}

func (h *Handler) DeleteBankHandler(c echo.Context) error {

	bank := &model.Banks{}
	if err := c.Bind(bank); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "cannot read json")
	}

	err := h.useCase.DeleteBank(c.Request().Context(), &model.Banks{Id: bank.Id})
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Bank successful delete")
}

func (h *Handler) CalculateHandler(c echo.Context) error {

	calculation := &model.CalculationBorrowed{}
	if err := c.Bind(calculation); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "cannot read json")
	}

	month, err := h.useCase.CalculatePayments(c.Request().Context(), calculation)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Payment": month,
	})
}
