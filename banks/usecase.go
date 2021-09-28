package banks

import (
	"bank/model"
	"context"
)

type UseCase interface {
	AllBanks (ctx context.Context) ([]model.Banks, error)
	CreateBank (ctx context.Context, bank *model.Banks) (int, error)
	UpdateBank (ctx context.Context, bank *model.Banks) (*model.Banks, error)
	DeleteBank (ctx context.Context, bank *model.Banks) error
	CalculatePayments (ctx context.Context, calculation *model.CalculationBorrowed) (float64, error)
}
