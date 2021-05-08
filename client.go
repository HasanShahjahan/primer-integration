package primer

import (
	"context"
	"net/http"
	"strings"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
	baseUrl    string
}

func (c *Client) GetRequest(ctx context.Context, idempotencyKey, path string, apiReq interface{}, apiResp Dto) error {
	req, err := c.createRequest(http.MethodGet, idempotencyKey, path, apiReq)
	if err != nil {
		return wrapError(err)
	}
	return c.do(ctx, req, apiResp)
}

func (c *Client) PostRequest(ctx context.Context, idempotencyKey, path string, apiReq interface{}, apiResp Dto) error {
	req, err := c.createRequest(http.MethodPost, idempotencyKey, path, apiReq)
	if err != nil {
		return wrapError(err)
	}

	return c.do(ctx, req, apiResp)
}

func (c *Client) createRequest(method, idempotencyKey, path string, request interface{}) (*http.Request, error) {
	body, err := MarshalRequest(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.baseUrl+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("X-Idempotency-Key", idempotencyKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, apiResp Dto) error {
	client := c.httpClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return wrapError(err)
	}
	defer resp.Body.Close()
	return DecodeResponse(resp, apiResp)
}

// ClientOption constructor parameter for NewClient(...)
type ClientOption func(*Client) error

// NewClient constructs a new Client which can make requests to the Primer APIs.
func NewClient(options ...ClientOption) (*Client, error) {
	c := &Client{}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	if strings.TrimSpace(c.apiKey) == "" {
		return nil, ApiKeyMissing
	}
	if strings.TrimSpace(c.baseUrl) == "" {
		return nil, BaseUrlMissing
	}
	return c, nil
}

// WithHttpClient configures a Primer API client with a http.Client to make requests over.
func WithHttpClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		if c.Transport == nil {
			c.Transport = http.DefaultTransport
		}
		client.httpClient = c
		return nil
	}
}

// WithApiKey configures a Primer API client with an API Key
func WithApiKey(apiKey string) ClientOption {
	return func(c *Client) error {
		c.apiKey = apiKey
		return nil
	}
}

// WithBaseUrl configures a Primer API client with a custom base url
func WithBaseUrl(baseUrl string) ClientOption {
	return func(c *Client) error {
		c.baseUrl = baseUrl
		return nil
	}
}
