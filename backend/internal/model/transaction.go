package model

import "time"

type Transaction struct {
	ID       string    `json:"id"`
	Date     time.Time `json:"date"`
	Category string    `json:"category"`
	Amount   float64   `json:"amount"`
	Note     string    `json:"note"`
}

type Summary struct {
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
	Balance      float64 `json:"balance"`
}
