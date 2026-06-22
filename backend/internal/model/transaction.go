package model

type Item struct {
	Category string  `json:"category"` // หัก, ลงทุน, ออม ฯลฯ
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
	Type     string  `json:"type"` // "income" | "expense" | "saving"
}

type Summary struct {
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
	TotalSaving  float64 `json:"total_saving"`
	Balance      float64 `json:"balance"`
}

type FinancialStatement struct {
	Income   []Item  `json:"income"`
	Expenses []Item  `json:"expenses"`
	Savings  []Item  `json:"savings"`
	Summary  Summary `json:"summary"`
}
