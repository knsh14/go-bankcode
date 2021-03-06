package bankcode

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

// Bank describes bank information
type Bank struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	HalfWidthKana string `json:"halfWidthKana"`
	FullWidthKana string `json:"fullWidthKana"`
	Hiragana      string `json:"hiragana"`
}

// Banks is result of ListBanks API
type Banks struct {
	Data       []*Bank `json:"data"`
	Size       int     `json:"size"`
	Limit      int     `json:"limit"`
	HasNext    bool    `json:"hasNext"`
	NextCursor string  `json:"nextCursor"`
	HasPrev    bool    `json:"hasPrev"`
	PrevCursor string  `json:"prevCursor"`
	Version    string  `json:"version"`
}

// GetBan returns bank info of input bank code
func (c *Client) GetBank(ctx context.Context, code string, param *GetParameter) (*Bank, error) {
	u, err := url.Parse(c.base.String() + "/banks/" + code)
	if err != nil {
		return nil, fmt.Errorf("generate URL: %w", err)
	}
	req, err := c.getRequest(ctx, u, param)
	if err != nil {
		return nil, fmt.Errorf("generate request: %w", err)
	}

	var res Bank
	err = c.call(ctx, req, func(resp io.ReadCloser) error {
		if err := json.NewDecoder(resp).Decode(&res); err != nil {
			return fmt.Errorf("decode to response: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// ListBanks returns list of banks and cursor information
func (c *Client) ListBanks(ctx context.Context, param *ListParameter) (*Banks, error) {
	u, err := url.Parse(c.base.String() + "/banks")
	if err != nil {
		return nil, err
	}
	req, err := c.listRequest(ctx, u, param)
	if err != nil {
		return nil, fmt.Errorf("generate request: %w", err)
	}

	var res Banks
	err = c.call(ctx, req, func(resp io.ReadCloser) error {
		if err := json.NewDecoder(resp).Decode(&res); err != nil {
			return fmt.Errorf("decode to response: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &res, nil
}
