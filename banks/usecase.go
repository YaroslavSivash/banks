package banks

import (
	"bank/model"
	"github.com/labstack/echo/v4"
)

type UseCase interface {
	AllBanks (c echo.Context) ([]model.Banks, error)
	CreateBank (c echo.Context, bank *model.Banks) (int, error)
	UpdateBank (c echo.Context, bank *model.Banks) (*model.Banks, error)
	DeleteBank (c echo.Context, bank *model.Banks) error
	CalculatePayments (c echo.Context, calculation *model.CalculationBorrowed) (float64, error)
}
