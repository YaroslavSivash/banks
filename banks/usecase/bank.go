package usecase

import (
	"bank/banks"
	"bank/model"
	"github.com/labstack/echo/v4"
)

type BankUseCase struct {
	repo banks.BankRepository
}

func NewBankUseCase (repo banks.BankRepository) *BankUseCase {
	return &BankUseCase{
		repo: repo,
	}
}

func (u *BankUseCase) AllBanks (c echo.Context) ([]model.Banks, error) {
	return u.repo.AllBanksDB(c)
}

func (u *BankUseCase) CreateBank (c echo.Context, banks *model.Banks ) (*model.Banks, error) {
	return u.repo.CreateBankDB(c, banks)
}

func (u *BankUseCase) UpdateBank (c echo.Context, banks *model.Banks ) (*model.Banks, error) {
	return u.repo.UpdateBankDB(c, banks)
}

func (u *BankUseCase) DeleteBank (c echo.Context, banks *model.Banks )  error {
	return u.repo.DeleteBankDB(c, banks)
}