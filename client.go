package bankcode

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

type option func(*Client) error

func NewClient(options ...option) (*Client, error) {
	c := &Client{
		base:       endpoint,
		httpClient: http.DefaultClient,
	}

	for _, o := range options {
		if err := o(c); err != nil {
			return nil, fmt.Errorf("initialize client: %w", err)
		}
	}
	return c, nil
}

// Client provides access method to BankCode API
type Client struct {
	keyToRequestHeader bool
	apiKey             string
	httpClient         *http.Client
	base               *url.URL
}

func (c *Client) call(ctx context.Context, req *http.Request, f func(resp io.ReadCloser) error) error {
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

func (c *Client) getRequest(ctx context.Context, u *url.URL, param *GetParameter) (*http.Request, error) {
	query := u.Query()
	if !c.keyToRequestHeader && param.APIKey != "" {
		query.Add("apikey", param.APIKey)
	}
	if len(param.Fields) > 0 {
		query.Add("fields", strings.Join(param.Fields, ","))
	}
	u.RawQuery = query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("generate http request: %w", err)
	}
	if c.keyToRequestHeader && param.APIKey != "" {
		req.Header.Add("apikey", param.APIKey)
	}
	return req, nil
}

func (c *Client) listRequest(ctx context.Context, u *url.URL, param *ListParameter) (*http.Request, error) {
	query := u.Query()
	if !c.keyToRequestHeader && param.APIKey != "" {
		query.Add("apikey", param.APIKey)
	}
	if param.Filter != "" {
		query.Add("filter", param.Filter)
	}
	if 0 < param.Limit && param.Limit <= 2000 {
		query.Add("limit", strconv.Itoa(param.Limit))
	}
	if param.Cursor != "" {
		query.Add("cursor", param.Cursor)
	}
	if len(param.Fields) > 0 {
		query.Add("fields", strings.Join(param.Fields, ","))
	}
	u.RawQuery = query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("generate http request: %w", err)
	}
	if c.keyToRequestHeader && param.APIKey != "" {
		req.Header.Add("apikey", param.APIKey)
	}
	return req, nil
}

type GetParameter struct {
	APIKey string
	Fields []string
}

type ListParameter struct {
	APIKey string
	Filter string
	Limit  int
	Cursor string
	Fields []string
}

func WithAPIKey(key string) option {
	return func(c *Client) error {
		if key == "" {
			return errors.New("API Key is empty")
		}
		c.apiKey = key
		return nil
	}
}

func WithEndpoint(endpoint string) option {
	return func(c *Client) error {
		if endpoint == "" {
			return errors.New("URL is empty")
		}
		u, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		c.base = u
		return nil
	}
}

func WithHeaderAPIKey(b bool) option {
	return func(c *Client) error {
		c.keyToRequestHeader = b
		return nil
	}
}
