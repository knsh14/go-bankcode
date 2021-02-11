package bankcode

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const baseURL = "https://apis.bankcode-jp.com/v1"

var (
	endpoint *url.URL
)

func init() {
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	endpoint = u
}

type option interface {
	Apply(*Client)
}

func NewClient(...option) *Client {

	return &Client{
		base:       endpoint,
		httpClient: http.DefaultClient,
	}
}

type Client struct {
	keyToRequestHeader bool
	httpClient         *http.Client
	base               *url.URL
}

func (c *Client) call(ctx context.Context, u *url.URL, f func(resp io.ReadCloser) error) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("generate http request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request to bank code: %w", err)
	}
	defer func() {
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
	}()

	if resp.StatusCode != 200 {
		return fmt.Errorf("http status error: %s", resp.Status)
	}

	return f(resp.Body)
}
