package usecase

import (
	"bank/banks"
	"bank/model"
	"bank/utils"
	"github.com/labstack/echo/v4"
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

func (u *BankUseCase) AllBanks (c echo.Context) ([]model.Banks, error) {
	return u.repo.AllBanksDB(c)
}

func (u *BankUseCase) CreateBank (c echo.Context, banks *model.Banks ) (int, error) {
	return u.repo.CreateBankDB(c, banks)
}

func (u *BankUseCase) UpdateBank (c echo.Context, banks *model.Banks ) (*model.Banks, error) {
	return u.repo.UpdateBankDB(c, banks)
}

func (u *BankUseCase) DeleteBank (c echo.Context, bank *model.Banks )  error {
	return u.repo.DeleteBankDB(c, bank)
}
func (u *BankUseCase) CalculatePayments (c echo.Context, calculation *model.CalculationBorrowed) (float64, error) {
	var bank []model.Banks

	bank, err := u.repo.AllBanksDB(c)
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