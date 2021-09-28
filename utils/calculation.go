package utils

import (
	"github.com/labstack/gommon/log"
	"math"
)

func Calculation(initialLoan, interestRate, loanTerm int) float64 {
	log.Info(initialLoan)
	log.Info(interestRate)
	log.Info(loanTerm)

	ir := float64(interestRate) / 100.00
	first := math.Pow((1 + ir/12), float64(loanTerm))
	second := (float64(initialLoan) * (ir / 12))
	m := second * math.Round(first*100) / 100
	third := math.Pow((1 + ir/12), float64(loanTerm))
	fourth := third - 1
	calc := m / fourth

	return math.Round(calc*100) / 100
}
