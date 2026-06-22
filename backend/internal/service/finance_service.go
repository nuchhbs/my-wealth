package service

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/nuchhbs/my-wealth/backend/internal/model"
	"github.com/nuchhbs/my-wealth/backend/internal/sheets"
)

type FinanceService struct {
	sheets *sheets.Client
}

func NewFinanceService(s *sheets.Client) *FinanceService {
	return &FinanceService{sheets: s}
}

func (s *FinanceService) GetStatement() (*model.FinancialStatement, error) {
	rows, err := s.sheets.ReadAll()
	if err != nil {
		return nil, err
	}

	stmt := &model.FinancialStatement{}
	currentCategory := ""

	for _, row := range rows {
		// ต้องมีอย่างน้อย 3 คอลัม
		if len(row) < 3 {
			continue
		}

		colA := strings.TrimSpace(row[0]) // หมวด
		colB := strings.TrimSpace(row[1]) // ชื่อรายการ
		colC := parseAmount(safeGet(row, 2)) // รายจ่าย
		colD := parseAmount(safeGet(row, 3)) // รายรับ

		// อัปเดต category ถ้า colA มีค่า
		if colA != "" {
			currentCategory = colA
		}

		// ข้ามแถวที่ไม่มีชื่อรายการ หรือเป็นแถว title/total
		if colB == "" || isTotal(colB) {
			continue
		}

		// รายรับ: colD มีค่า, colC ว่าง
		if colD > 0 && colC == 0 && !isBalance(colB) {
			stmt.Income = append(stmt.Income, model.Item{
				Category: currentCategory,
				Name:     colB,
				Amount:   colD,
				Type:     "income",
			})
			continue
		}

		// รายจ่าย/ออม: colC มีค่า
		if colC > 0 && !isBalance(colB) {
			itemType := "expense"
			if isSaving(colB, currentCategory) {
				itemType = "saving"
			}
			item := model.Item{
				Category: currentCategory,
				Name:     colB,
				Amount:   colC,
				Type:     itemType,
			}
			if itemType == "saving" {
				stmt.Savings = append(stmt.Savings, item)
			} else {
				stmt.Expenses = append(stmt.Expenses, item)
			}
		}
	}

	// คำนวณ summary
	for _, i := range stmt.Income {
		stmt.Summary.TotalIncome += i.Amount
	}
	for _, i := range stmt.Expenses {
		stmt.Summary.TotalExpense += i.Amount
	}
	for _, i := range stmt.Savings {
		stmt.Summary.TotalSaving += i.Amount
	}
	stmt.Summary.Balance = stmt.Summary.TotalIncome - stmt.Summary.TotalExpense - stmt.Summary.TotalSaving

	return stmt, nil
}

// --- helpers ---

func safeGet(row []string, i int) string {
	if i < len(row) {
		return row[i]
	}
	return ""
}

func parseAmount(s string) float64 {
	s = strings.TrimSpace(s)
	// ลบ comma และ whitespace
	s = strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || r == '.' || r == '-' {
			return r
		}
		return -1
	}, s)
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func isTotal(s string) bool {
	keywords := []string{"รวม", "คงเหลือ", "total", "งบการเงิน"}
	s = strings.ToLower(strings.TrimSpace(s))
	for _, k := range keywords {
		if strings.Contains(s, k) {
			return true
		}
	}
	return false
}

func isBalance(s string) bool {
	return strings.Contains(s, "คงเหลือ")
}

func isSaving(name, category string) bool {
	keywords := []string{"ลงทุน", "ออม", "กองทุน", "หุ้น", "ดาวน์", "ประกัน"}
	combined := strings.ToLower(name + category)
	for _, k := range keywords {
		if strings.Contains(combined, k) {
			return true
		}
	}
	return false
}
