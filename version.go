package bankcode

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

func (c *Client) GetVersion(ctx context.Context, apiKey string) (string, error) {
	p, err := url.Parse(c.base.String() + "/version")
	if err != nil {
		return "", err
	}
	var res struct {
		Version string `json:"version"`
	}
	err = c.call(ctx, p, func(resp io.ReadCloser) error {
		if err := json.NewDecoder(resp).Decode(&res); err != nil {
			return fmt.Errorf("decode to response: %w", err)
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return res.Version, nil
}
