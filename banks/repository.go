package banks

import (
	"bank/model"
	"context"
)

type BankRepository interface {
	AllBanksDB (ctx context.Context) ([]model.Banks, error)
	CreateBankDB (ctx context.Context, bank *model.Banks) (int, error)
	UpdateBankDB (ctx context.Context, bank *model.Banks) (*model.Banks, error)
	DeleteBankDB (ctx context.Context, bank *model.Banks) error

}