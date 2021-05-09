package primer

import (
	"context"
	"encoding/json"
	"testing"
)

type Config struct {
	BaseUrl string `json:"BaseUrl"`
	ApiKey  string `json:"ApiKey"`
}

type configurationWrapper struct {
	Config `json:"primer"`
}

var configurationStr = `{
"primer": {
		"BaseUrl": "https://api.sandbox.primer.io",
		"ApiKey": "dc0567e4-4c86-4bc2-b4d1-387ee9dfc1c5"
	}
}`

var configuration configurationWrapper

func setup(t *testing.T) {
	if err := json.Unmarshal([]byte(configurationStr), &configuration); err != nil {
		t.Errorf("unmarshal adyen config failed, err=%#v", err)
	}
}

var idempotencyKey string
var requestId string
var paymentId string
var orderId string
var amount int64
var CurrencyCode string
var token string
var reason string

func TestSearchPayment(testing *testing.T) {
	setup(testing)
	httpClient, _ := NewClient(WithApiKey(configuration.Config.ApiKey), WithBaseUrl(configuration.Config.BaseUrl))

	if result, err := httpClient.Search(context.Background()); err != nil {
		testing.Errorf("search payment is failed, err=%v", err)
	} else {
		testing.Logf("successfully search payment request, resp=%v", result)
	}
}

func TestCreatePayment(testing *testing.T) {
	setup(testing)
	httpClient, _ := NewClient(WithApiKey(configuration.Config.ApiKey), WithBaseUrl(configuration.Config.BaseUrl))

	idempotencyKey = "Hasan"
	requestId = "Hasan-123"

	orderId = "eatigo"
	amount = 100
	CurrencyCode = "USD"

	//TODO:The payment method token used to auth the transaction. Is it the authentication token or from checkout SDK?
	token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2MjA2NzI1ODIsImFjY2Vzc1Rva2VuIjoiZmZiZWI4MGItYWMzYS00YzNiLTg0NjAtYWE1MGM4ZGVkNmFlIiwiYW5hbHl0aWNzVXJsIjoiaHR0cHM6Ly9hbmFseXRpY3MuYXBpLnNhbmRib3guY29yZS5wcmltZXIuaW8vbWl4cGFuZWwiLCJpbnRlbnQiOiJDSEVDS09VVCIsImNvbmZpZ3VyYXRpb25VcmwiOiJodHRwczovL2FwaS5zYW5kYm94LnByaW1lci5pby9jbGllbnQtc2RrL2NvbmZpZ3VyYXRpb24iLCJjb3JlVXJsIjoiaHR0cHM6Ly9hcGkuc2FuZGJveC5wcmltZXIuaW8iLCJwY2lVcmwiOiJodHRwczovL3Nkay5hcGkuc2FuZGJveC5wcmltZXIuaW8iLCJlbnYiOiJTQU5EQk9YIiwidGhyZWVEU2VjdXJlSW5pdFVybCI6Imh0dHBzOi8vc29uZ2JpcmRzdGFnLmNhcmRpbmFsY29tbWVyY2UuY29tL2NhcmRpbmFsY3J1aXNlL3YxL3NvbmdiaXJkLmpzIiwidGhyZWVEU2VjdXJlVG9rZW4iOiJleUowZVhBaU9pSktWMVFpTENKaGJHY2lPaUpJVXpJMU5pSjkuZXlKcWRHa2lPaUpqWkRjd05HTm1OQzFpWVRZNUxUUXpaR0l0T0dJMU55MHhZV0UzWVdNMU1HVTBNR01pTENKcFlYUWlPakUyTWpBMU9EWXhPRElzSW1semN5STZJalZsWWpWaVlXVmpaVFpsWXpjeU5tVmhOV1ppWVRkbE5TSXNJazl5WjFWdWFYUkpaQ0k2SWpWbFlqVmlZVFF4WkRRNFptSmtOakE0T0RoaU9HVTBOQ0o5Lk1nN1hCYVphRldGWGQ5b2VjVnpKeDBJX2Y5OFRDOTBQa0NHbEhMQ0NoMjgiLCJwYXltZW50RmxvdyI6IkRFRkFVTFQifQ.ItKmPP85kWeStka0lUQobuUD9cQqrUC6XHNXcGZhJVU"
	createPayment := &CreatePaymentRequest{
		BaseDto: BaseDto{
			RequestId:      &requestId,
			IdempotencyKey: &idempotencyKey,
		},
		CreatePayment: CreatePayment{
			OrderId:      orderId,
			Amount:       amount,
			CurrencyCode: CurrencyCode,
			PaymentInstrument: PaymentInstrument{
				Token: token,
			},
		},
	}

	if result, err := httpClient.Create(context.Background(), createPayment); err != nil {
		testing.Errorf("create payment is failed, err=%v", err)
	} else {
		testing.Logf("successfully create payment request, resp=%v", result)
	}
}

func TestCapturePayment(testing *testing.T) {
	setup(testing)
	httpClient, _ := NewClient(WithApiKey(configuration.Config.ApiKey), WithBaseUrl(configuration.Config.BaseUrl))

	idempotencyKey = "Hasan1"
	requestId = "234567"

	paymentId = "2345678"
	amount = 100

	capturePaymentRequest := &CapturePaymentRequest{
		BaseDto: BaseDto{
			RequestId:      &requestId,
			IdempotencyKey: &idempotencyKey,
		},
		CapturePayment: CapturePayment{
			Amount: amount,
			Final:  true,
		},
	}

	if result, err := httpClient.Capture(context.Background(), capturePaymentRequest, paymentId); err != nil {
		testing.Errorf("capture payment is failed, err=%v", err)
	} else {
		testing.Logf("successfully cappture payment request, resp=%v", result)
	}
}

func TestCancelPayment(testing *testing.T) {
	setup(testing)
	httpClient, _ := NewClient(WithApiKey(configuration.Config.ApiKey), WithBaseUrl(configuration.Config.BaseUrl))

	idempotencyKey = "Hasan2"
	requestId = "23456789"

	paymentId = "2345678"
	reason = "Insufficient Balance"

	cancelPaymentRequest := &CancelPaymentRequest{
		BaseDto: BaseDto{
			RequestId:      &requestId,
			IdempotencyKey: &idempotencyKey,
		},
		CancelPayment: CancelPayment{
			Reason: reason,
		},
	}

	if result, err := httpClient.Cancel(context.Background(), cancelPaymentRequest, paymentId); err != nil {
		testing.Errorf("cancel payment is failed, err=%v", err)
	} else {
		testing.Logf("successfully cancel payment request, resp=%v", result)
	}
}

func TestRefundPayment(testing *testing.T) {
	setup(testing)
	httpClient, _ := NewClient(WithApiKey(configuration.Config.ApiKey), WithBaseUrl(configuration.Config.BaseUrl))

	idempotencyKey = "Hasan3"
	requestId = "2345678"

	paymentId = "2345678"
	amount = 100

	refundPaymentRequest := &RefundPaymentRequest{
		BaseDto: BaseDto{
			RequestId:      &requestId,
			IdempotencyKey: &idempotencyKey,
		},
		RefundPayment: RefundPayment{
			Amount: amount,
		},
	}

	if result, err := httpClient.Refund(context.Background(), refundPaymentRequest, paymentId); err != nil {
		testing.Errorf("refund payment is failed, err=%v", err)
	} else {
		testing.Logf("successfully refund payment request, resp=%v", result)
	}
}

func TestResumePayment(testing *testing.T) {
	setup(testing)
	httpClient, _ := NewClient(WithApiKey(configuration.Config.ApiKey), WithBaseUrl(configuration.Config.BaseUrl))

	idempotencyKey = ""
	requestId = ""

	paymentId = "2345678"
	amount = 100

	//TODO:A token that is sent back from the checkout to complete a blocked payment flow?
	token = "6742b173-62f0-48bf-a210-60865437254a"

	resumePaymentRequest := &ResumePaymentRequest{
		BaseDto: BaseDto{
			RequestId:      &requestId,
			IdempotencyKey: &idempotencyKey,
		},
		ResumePayment: ResumePayment{
			ResumeToken: token,
		},
	}

	if result, err := httpClient.Resume(context.Background(), resumePaymentRequest, paymentId); err != nil {
		testing.Errorf("resume payment is failed, err=%v", err)
	} else {
		testing.Logf("successfully resume payment request, resp=%v", result)
	}
}

func TestGetPayment(testing *testing.T) {
	setup(testing)
	httpClient, _ := NewClient(WithApiKey(configuration.Config.ApiKey), WithBaseUrl(configuration.Config.BaseUrl))

	idempotencyKey = ""
	requestId = ""
	paymentId = "123456"

	getPaymentRequest := &GetPaymentRequest{
		BaseDto: BaseDto{
			RequestId:      &requestId,
			IdempotencyKey: &idempotencyKey,
		},
	}

	if result, err := httpClient.Get(context.Background(), getPaymentRequest, paymentId); err != nil {
		testing.Errorf("get payment is failed, err=%v", err)
	} else {
		testing.Logf("successfully get payment request, resp=%v", result)
	}
}
