package banks

import (
	"bank/model"
	"github.com/labstack/echo/v4"
)

type BankRepository interface {
	AllBanksDB (c echo.Context) ([]model.Banks, error)
	CreateBankDB (c echo.Context, bank *model.Banks) (int, error)
	UpdateBankDB (c echo.Context, bank *model.Banks) (*model.Banks, error)
	DeleteBankDB (c echo.Context, bank *model.Banks) error
	CalculatePaymentsDB (c echo.Context, calculation *model.CalculationBorrowed) (int, error)
}