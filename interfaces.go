package primer

import "context"

type ePayment interface {
	Search(ctx context.Context) (*SearchPaymentResponse, error)
	Create(ctx context.Context, req *CreatePaymentRequest) (*PaymentResponse, error)
	Capture(ctx context.Context, req *CapturePaymentRequest, paymentId string) (*PaymentResponse, error)
	Cancel(ctx context.Context, req *CancelPaymentRequest, paymentId string) (*PaymentResponse, error)
	Refund(ctx context.Context, req *RefundPaymentRequest, paymentId string) (*PaymentResponse, error)
	Resume(ctx context.Context, req *ResumePaymentRequest, paymentId string) (*PaymentResponse, error)
	Get(ctx context.Context, req *GetPaymentRequest, paymentId string) (*PaymentResponse, error)
}
