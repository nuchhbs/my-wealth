# API Documentation

Base URL: `http://localhost:8080`

## Endpoints

### GET /health
Returns server status.

**Response**
```json
{ "status": "ok" }
```

---

### GET /api/summary
Returns financial summary aggregated from Google Sheets.

**Response**
```json
{
  "total_income": 50000.00,
  "total_expense": -20000.00,
  "balance": 30000.00
}
```

---

### GET /api/transactions
Returns all transactions from Google Sheets.

**Response**
```json
[
  {
    "id": "txn-1",
    "date": "2026-06-01T00:00:00Z",
    "category": "Salary",
    "amount": 50000.00,
    "note": "June salary"
  }
]
```

## Google Sheet Format

Sheet1 columns (row 1 = header, data starts row 2):

| A (Date)   | B (Category) | C (Amount) | D (Note)    |
|------------|--------------|------------|-------------|
| 2026-06-01 | Salary       | 50000      | June salary |
| 2026-06-05 | Food         | -500       | Dinner      |

- Amount positive = income
- Amount negative = expense
- Date format: `YYYY-MM-DD`
