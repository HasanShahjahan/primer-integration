package primer

import (
	"context"
	"fmt"
)

func (c *Client) Search(ctx context.Context) (*SearchPaymentResponse, error) {
	path := "/payments"
	if err := c.Validate(nil, Search, ""); err != nil {
		return nil, err
	}

	response := &SearchPaymentResponse{}
	if err := c.GetRequest(ctx, "", path, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Create(ctx context.Context, req *CreatePaymentRequest) (*PaymentResponse, error) {
	path := "/payments"
	if err := c.Validate(req.IdempotencyKey, Create, ""); err != nil {
		return nil, err
	}

	response := &PaymentResponse{}
	if err := c.PostRequest(ctx, StrSafeDeref(req.IdempotencyKey), path, req.CreatePayment, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Capture(ctx context.Context, req *CapturePaymentRequest, paymentId string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/payments/%s/capture", paymentId)
	if err := c.Validate(req.IdempotencyKey, Capture, paymentId); err != nil {
		return nil, err
	}

	response := &PaymentResponse{}
	if err := c.PostRequest(ctx, StrSafeDeref(req.IdempotencyKey), path, req.CapturePayment, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Cancel(ctx context.Context, req *CancelPaymentRequest, paymentId string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/payments/%s/cancel", paymentId)
	if err := c.Validate(req.IdempotencyKey, Cancel, paymentId); err != nil {
		return nil, err
	}

	response := &PaymentResponse{}
	if err := c.PostRequest(ctx, StrSafeDeref(req.IdempotencyKey), path, req.CancelPayment, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Refund(ctx context.Context, req *RefundPaymentRequest, paymentId string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/payments/%s/refund", paymentId)
	if err := c.Validate(req.IdempotencyKey, Refund, paymentId); err != nil {
		return nil, err
	}

	response := &PaymentResponse{}
	if err := c.PostRequest(ctx, StrSafeDeref(req.IdempotencyKey), path, req.RefundPayment, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Resume(ctx context.Context, req *ResumePaymentRequest, paymentId string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/payments/%s/resume", paymentId)
	if err := c.Validate(nil, Resume, paymentId); err != nil {
		return nil, err
	}

	response := &PaymentResponse{}
	if err := c.PostRequest(ctx, StrSafeDeref(req.IdempotencyKey), path, req.ResumePayment, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Get(ctx context.Context, req *GetPaymentRequest, paymentId string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/payments/%s", paymentId)
	if err := c.Validate(nil, Get, paymentId); err != nil {
		return nil, err
	}

	response := &PaymentResponse{}
	if err := c.GetRequest(ctx, StrSafeDeref(req.IdempotencyKey), path, response); err != nil {
		return nil, err
	}
	return response, nil
}
