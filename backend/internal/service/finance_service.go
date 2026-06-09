package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nuchhbs/my-wealth/backend/internal/model"
	"github.com/nuchhbs/my-wealth/backend/internal/sheets"
)

type FinanceService struct {
	sheets *sheets.Client
}

func NewFinanceService(s *sheets.Client) *FinanceService {
	return &FinanceService{sheets: s}
}

// GetTransactions reads from Sheet1 with columns: Date, Category, Amount, Note
func (s *FinanceService) GetTransactions() ([]model.Transaction, error) {
	rows, err := s.sheets.ReadRange("Sheet1!A2:D")
	if err != nil {
		return nil, err
	}

	var txns []model.Transaction
	for i, row := range rows {
		if len(row) < 3 {
			continue
		}
		date, err := time.Parse("2006-01-02", fmt.Sprint(row[0]))
		if err != nil {
			continue
		}
		amount, err := strconv.ParseFloat(fmt.Sprint(row[2]), 64)
		if err != nil {
			continue
		}
		note := ""
		if len(row) > 3 {
			note = fmt.Sprint(row[3])
		}
		txns = append(txns, model.Transaction{
			ID:       fmt.Sprintf("txn-%d", i+1),
			Date:     date,
			Category: fmt.Sprint(row[1]),
			Amount:   amount,
			Note:     note,
		})
	}
	return txns, nil
}

func (s *FinanceService) GetSummary() (*model.Summary, error) {
	txns, err := s.GetTransactions()
	if err != nil {
		return nil, err
	}

	var income, expense float64
	for _, t := range txns {
		if t.Amount >= 0 {
			income += t.Amount
		} else {
			expense += t.Amount
		}
	}
	return &model.Summary{
		TotalIncome:  income,
		TotalExpense: expense,
		Balance:      income + expense,
	}, nil
}
