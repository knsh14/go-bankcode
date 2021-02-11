package bankcode

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
)

type GetBankRequest struct {
	APIKey string
	Fields []string
}

type Bank struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	HalfWidthKana string `json:"halfWidthKana"`
	FullWidthKana string `json:"fullWidthKana"`
	Hiragana      string `json:"hiragana"`
}

func (c *Client) GetBank(ctx context.Context, code string, param *GetBankRequest) (*Bank, error) {
	p, err := url.Parse(c.base.String() + "/banks/" + code)
	if err != nil {
		return nil, err
	}
	query := p.Query()
	p.RawQuery = query.Encode()

	var res Bank
	err = c.call(ctx, p, func(resp io.ReadCloser) error {
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

type ListBankRequest struct {
	APIKey string
	Filter string
	Limit  int
	Cursor string
	Fields []string
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

func (c *Client) ListBanks(ctx context.Context, param *ListBankRequest) (*Banks, error) {
	p, err := url.Parse(c.base.String() + "/banks")
	if err != nil {
		return nil, err
	}
	query := p.Query()
	query.Add("limit", strconv.Itoa(param.Limit))
	p.RawQuery = query.Encode()
	var res Banks
	err = c.call(ctx, p, func(resp io.ReadCloser) error {
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
