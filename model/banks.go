package model

type Banks struct {
	tableName          struct{} `pg:"banks,alias:banks"`
	Id                 int      `json:"id,omitempty" pg:"id"`
	BankName           string   `json:"bank_name,omitempty" pg:"bank_name"`
	InterestRate       float64  `json:"interest_rate,omitempty" pg:"interest_rate"`
	MaximumLoan        int      `json:"maximum_loan,omitempty" pg:"maximum_loan"`
	MinimumDownPayment int      `json:"minimum_down_payment,omitempty" pg:"minimum_down_payment"`
	LoanTerm           int      `json:"loan_term,omitempty" pg:"loan_term"`
}

func (b *Banks) TableName() string {
	return "banks"
}
