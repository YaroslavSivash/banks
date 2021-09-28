package postgres

import (
	"bank/model"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type BankRepository struct {
	db *pg.DB
}

func NewBankRepository(db *pg.DB) *BankRepository {
	return &BankRepository{
		db: db,
	}
}

func (r *BankRepository) AllBanksDB(c echo.Context) ([]model.Banks, error) {
	var banks []model.Banks

	err := r.db.Model().
		Table("banks").
		Column("id", "bank_name", "interest_rate", "maximum_loan", "minimum_down_payment", "loan_term").
		Select(&banks)
	if err != nil {
		log.Info(err)
		log.Error(err)
		return nil, err
	}
	//log.Info(banks)
	return banks, nil
}

func (r *BankRepository) CreateBankDB(c echo.Context, bank *model.Banks) (int, error) {

	//log.Info(bank)
	_, err := r.db.Model(bank).
		Insert()
	if err != nil {
		log.Info(err)
		log.Error(err)
		return 0, err
	}
	return bank.Id, nil
}
func (r *BankRepository) UpdateBankDB(c echo.Context, bank *model.Banks) (*model.Banks, error) {

	_, err := r.db.Model(bank).
		WherePK().
		Update()
	if err != nil {
		log.Info(err)
		log.Error(err)
		return nil, err
	}
	return bank, nil
}
func (r *BankRepository) DeleteBankDB(c echo.Context, bank *model.Banks) error {

	_, err := r.db.Model(bank).
		Where("id = ?", &bank.Id).
		Delete()
	if err != nil {
		log.Info(err)
		log.Error(err)
		return err
	}
	return nil
}
