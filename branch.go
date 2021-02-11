package bankcode

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
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
	u, err := url.Parse(c.base.String() + "/banks/" + bankCode + "/branches/" + branchCode)
	if err != nil {
		return nil, err
	}
	req, err := c.getRequest(ctx, u, param)
	if err != nil {
		return nil, fmt.Errorf("generate request: %w", err)
	}

	var res []*Branch
	err = c.call(ctx, req, func(resp io.ReadCloser) error {
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
	u, err := url.Parse(c.base.String() + "/banks/" + bankCode + "/branches")
	if err != nil {
		return nil, err
	}
	req, err := c.listRequest(ctx, u, param)
	if err != nil {
		return nil, fmt.Errorf("generate request: %w", err)
	}

	var res Branches
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
