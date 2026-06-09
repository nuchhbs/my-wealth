package sheets

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Client struct {
	svc     *sheets.Service
	sheetID string
}

func NewClient(credentialsFile, sheetID string) (*Client, error) {
	ctx := context.Background()
	svc, err := sheets.NewService(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, err
	}
	return &Client{svc: svc, sheetID: sheetID}, nil
}

func (c *Client) ReadRange(rangeName string) ([][]interface{}, error) {
	resp, err := c.svc.Spreadsheets.Values.Get(c.sheetID, rangeName).Do()
	if err != nil {
		return nil, err
	}
	return resp.Values, nil
}
