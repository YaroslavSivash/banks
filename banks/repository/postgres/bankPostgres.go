package postgres

import (
	"bank/model"
	"context"
	"github.com/go-pg/pg/v10"
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

func (r *BankRepository) AllBanksDB(ctx context.Context) ([]model.Banks, error) {
	var banks []model.Banks
	//SELECT EXISTS(SELECT 1 FROM banks
	//WHERE id = ? AND interest_rate = ?);
	//err := r.db.Model(&banks).
	//	Column("id", "bank_name", "interest_rate", "maximum_loan", "minimum_down_payment", "loan_term").
	//
	//	//Where("interest_rate = ?", 35).
	//	WhereOr("interest_rate = ? OR interest_rate = ?", 100, 35).
	//	Where("maximum_loan = 1000000").
	//	Select()
	err := r.db.Model(&banks).
		//Table("banks").
		ColumnExpr("interest_rate, loan_term").
		//Where("id = ?", 5).
		//ColumnExpr("DISTINCT loan_term").
		Select(&banks)

	if err != nil {
		log.Info()
		log.Error(err)
		return nil, err
	}
	return banks, nil
}

func (r *BankRepository) CreateBankDB(ctx context.Context, bank *model.Banks) (int, error) {

	var bankId int
	_, err := r.db.Model(bank).
		Insert(&bankId)
	log.Info(bankId)
	if err != nil {
		log.Info(bankId)
		log.Info(err)
		log.Error(err)
		return 0, err
	}
	return bank.Id, nil
}
func (r *BankRepository) UpdateBankDB(ctx context.Context, bank *model.Banks) (*model.Banks, error) {
	//openItemsRez, err := r.db.Model().
	//	Table("items").
	//	Set("is_open = true").
	//	Where("tender_id IN(SELECT tender_id FROM tenders WHERE start_at <= ? AND finalized_at > ?",
	//		timeNow, timeNow).
	//	Where("is_open = false").
	//	Update()
	//_, err := r.db.Model().
	//	Table("banks").
	//	//Set("bank_name = ?Liverpool").
	//	//Set("interest_rate = 50").
	//	Set("interest_rate = ?", bank.LoanTerm).
	//	Set("bank_name = ?", bank.BankName).
	//	Where("minimum_down_payment = ?", bank.MinimumDownPayment).
	//	Where("loan_term < ?", bank.LoanTerm).
	//	Update()
	var b *model.Banks
	_, err := r.db.Model().
		Table("banks").
		//Set("bank_name = ?Liverpool").
		//Set("interest_rate = 50").
		Set("interest_rate = 100").
		Update(&b)
	log.Info(&b)
	if err != nil {

		log.Error(err)
		return nil, err
	}
	log.Info(bank)
	return bank, nil
}
func (r *BankRepository) DeleteBankDB(ctx context.Context, bank *model.Banks) error {

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
