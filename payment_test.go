package primer

import (
	"context"
	"encoding/json"
	"testing"
)

type Config struct {
	baseUrl string `json:"baseUrl"`
	apiKey  string `json:"apiKey"`
}

type configurationWrapper struct {
	Config `json:"primer"`
}

var confStr = `{
  "primer": {
    "baseUrl": "https://api.sandbox.primer.io",
    "apiKey": "dc0567e4-4c86-4bc2-b4d1-387ee9dfc1c5",
  }
}`

var conf configurationWrapper

func setup(t *testing.T) {
	if err := json.Unmarshal([]byte(confStr), &conf); err != nil {
		t.Errorf("unmarshal primer config failed, err=%#v", err)
	}
}

var idempotencyKey string
var requestId string
var apiKey string
var baseUrl string
var paymentId string

func TestGetPayment(t *testing.T) {
	apiKey = "dc0567e4-4c86-4bc2-b4d1-387ee9dfc1c5"
	baseUrl = "https://api.sandbox.primer.io"
	//setup(t)
	c, _ := NewClient(WithApiKey(apiKey), WithBaseUrl(baseUrl))
	idempotencyKey = ""
	requestId = ""
	paymentId = "123456"
	getPaymentRequest := &GetPaymentRequest{
		BaseDto: BaseDto{
			RequestId:      &requestId,
			IdempotencyKey: &idempotencyKey,
		},
	}

	if got, err := c.Get(context.Background(), getPaymentRequest, paymentId); err != nil {
		t.Errorf("get payment is failed, err=%v", err)
	} else {
		t.Logf("successfully get payment request, resp=%v", got)
	}
}

func TestSearchPayment(t *testing.T) {
	apiKey = "dc0567e4-4c86-4bc2-b4d1-387ee9dfc1c5"
	baseUrl = "https://api.sandbox.primer.io"
	//setup(t)
	c, _ := NewClient(WithApiKey(apiKey), WithBaseUrl(baseUrl))

	if got, err := c.Search(context.Background()); err != nil {
		t.Errorf("search payment is failed, err=%v", err)
	} else {
		t.Logf("successfully search payment request, resp=%v", got)
	}
}
