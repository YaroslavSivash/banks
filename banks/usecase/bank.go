package usecase

import (
	"bank/banks"
	"bank/model"
	"bank/utils"
	"context"
	"github.com/labstack/gommon/log"
)

type BankUseCase struct {
	repo banks.BankRepository
}

func NewBankUseCase (repo banks.BankRepository) *BankUseCase {
	return &BankUseCase{
		repo: repo,
	}
}

func (u *BankUseCase) AllBanks (ctx context.Context) ([]model.Banks, error) {
	return u.repo.AllBanksDB(ctx)
}

func (u *BankUseCase) CreateBank (ctx context.Context, banks *model.Banks ) (int, error) {
	return u.repo.CreateBankDB(ctx, banks)
}

func (u *BankUseCase) UpdateBank (ctx context.Context, banks *model.Banks ) (*model.Banks, error) {
	return u.repo.UpdateBankDB(ctx, banks)
}

func (u *BankUseCase) DeleteBank (ctx context.Context, bank *model.Banks )  error {
	return u.repo.DeleteBankDB(ctx, bank)
}
func (u *BankUseCase) CalculatePayments (ctx context.Context, calculation *model.CalculationBorrowed) (float64, error) {
	var bank []model.Banks

	bank, err := u.repo.AllBanksDB(ctx)
	if err != nil {
		log.Error(err)
	}
	log.Info(bank)
	for _, b := range bank {
		if b.BankName == calculation.BankName{
			log.Info(b.BankName)
			if b.MaximumLoan<calculation.InitialLoan || b.MinimumDownPayment > calculation.DownPayment {
				log.Error(err)
				return 0, err
			}
			calc := utils.Calculation(calculation.InitialLoan, b.InterestRate, b.LoanTerm)
			return calc, err

		} else {
			continue
		}

	}
	return 0, nil
}