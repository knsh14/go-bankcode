package bankcode

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

// GetVersion returns version of bankcode API
func (c *Client) GetVersion(ctx context.Context, apiKey string) (string, error) {
	u, err := url.Parse(c.base.String() + "/version")
	if err != nil {
		return "", fmt.Errorf("generate URL: %w", err)
	}
	req, err := c.getRequest(ctx, u, &GetParameter{APIKey: apiKey})
	if err != nil {
		return "", fmt.Errorf("generate request: %w", err)
	}

	var res struct {
		Version string `json:"version"`
	}
	err = c.call(ctx, req, func(resp io.ReadCloser) error {
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
