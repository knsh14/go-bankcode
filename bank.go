package bankcode

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type Bank struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	HalfWidthKana string `json:"halfWidthKana"`
	FullWidthKana string `json:"fullWidthKana"`
	Hiragana      string `json:"hiragana"`
}

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

func (c *Client) GetBank(ctx context.Context, code string, param *GetParameter) (*Bank, error) {
	p, err := url.Parse(c.base.String() + "/banks/" + code)
	if err != nil {
		return nil, fmt.Errorf("generate URL: %w", err)
	}
	req, err := c.getRequest(ctx, p, param)
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

func (c *Client) ListBanks(ctx context.Context, param *ListParameter) (*Banks, error) {
	p, err := url.Parse(c.base.String() + "/banks")
	if err != nil {
		return nil, err
	}
	req, err := c.listRequest(ctx, p, param)
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
