package bankcode

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
)

type Branch struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	HalfWidthKana string `json:"halfWidthKana"`
	FullWidthKana string `json:"fullWidthKana"`
	Hiragana      string `json:"hiragana"`
}

type Branches struct {
	Data       []*Branch `json:"data"`
	Size       int       `json:"size"`
	Limit      int       `json:"limit"`
	HasNext    bool      `json:"hasNext"`
	NextCursor string    `json:"nextCursor"`
	HasPrev    bool      `json:"hasPrev"`
	PrevCursor string    `json:"prevCursor"`
	Version    string    `json:"version"`
}

func (c *Client) GetBranch(ctx context.Context, bankCode, branchCode string, param *GetParameter) ([]*Branch, error) {
	p, err := url.Parse(c.base.String() + "/banks/" + bankCode + "/branches/" + branchCode)
	if err != nil {
		return nil, err
	}
	query := p.Query()
	p.RawQuery = query.Encode()

	var res []*Branch
	err = c.call(ctx, p, func(resp io.ReadCloser) error {
		if err := json.NewDecoder(resp).Decode(&res); err != nil {
			return fmt.Errorf("decode to response: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) ListBranch(ctx context.Context, bankCode string, param *ListParameter) (*Branches, error) {
	p, err := url.Parse(c.base.String() + "/banks/" + bankCode + "/branches")
	if err != nil {
		return nil, err
	}
	query := p.Query()
	query.Add("limit", strconv.Itoa(param.Limit))
	p.RawQuery = query.Encode()

	var res Branches
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
