package model

type Banks struct {
	Id                 int    `json:"id,omitempty" pg:"id,pk" example:"2"`
	BankName           string `json:"bank_name,omitempty" pg:"bank_name" example:"Privat"`
	InterestRate       int    `json:"interest_rate,omitempty" pg:"interest_rate" example:"35"`
	MaximumLoan        int    `json:"maximum_loan,omitempty" pg:"maximum_loan" example:"10000"`
	MinimumDownPayment int    `json:"minimum_down_payment,omitempty" pg:"minimum_down_payment" example:"1000"`
	LoanTerm           int    `json:"loan_term,omitempty" pg:"loan_term" example:"20"`
}

type CalculationBorrowed struct {
	InitialLoan int    `json:"initial_loan,omitempty"`
	DownPayment int    `json:"down_payment,omitempty"`
	BankName    string `json:"bank_name,omitempty"`
}
