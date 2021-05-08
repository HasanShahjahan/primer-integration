package primer

import (
	"context"
	"net/http"
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
