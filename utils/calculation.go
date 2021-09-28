package utils

import (
	"github.com/labstack/gommon/log"
	"math"
)

func Calculation(InitialLoan, InterestRate, LoanTerm int) float64 {
	log.Info(InitialLoan)
	log.Info(InterestRate)
	log.Info(LoanTerm)

	ir := float64(InterestRate) / 100.00

	first := math.Pow((1 + ir/12), float64(LoanTerm))

	second := (float64(InitialLoan) * (ir / 12))

	m := second * math.Round(first*100) / 100

	third := math.Pow((1 + ir/12), float64(LoanTerm))

	fourth := third - 1

	calc := m / fourth

	return math.Round(calc*100) / 100
}
