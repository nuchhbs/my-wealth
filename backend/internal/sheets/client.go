package sheets

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

// Public Google Sheets CSV export URL
// Sheet must be shared as "Anyone with the link can view"
const exportURL = "https://docs.google.com/spreadsheets/d/%s/export?format=csv&gid=%s"

type Client struct {
	sheetID string
	gid     string // tab/sheet gid (default "0" = first sheet)
}

func NewClient(sheetID, gid string) *Client {
	if gid == "" {
		gid = "0"
	}
	return &Client{sheetID: sheetID, gid: gid}
}

// ReadAll fetches the sheet and returns rows as [][]string (skips header row)
func (c *Client) ReadAll() ([][]string, error) {
	url := fmt.Sprintf(exportURL, c.sheetID, c.gid)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch sheet: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sheet returned status %d — make sure the sheet is public", resp.StatusCode)
	}

	rows, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parse csv: %w", err)
	}

	if len(rows) <= 1 {
		return nil, nil // empty or header only
	}
	return rows[1:], nil // skip header
}
