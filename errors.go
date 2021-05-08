package primer

import (
	"errors"
	"net/http"
)

var (
	ApiKeyMissing         = errors.New("ApiKey is missing")
	BaseUrlMissing        = errors.New("base Url is missing")
	IdempotencyKeyMissing = errors.New("Idempotency-Key is missing")
	PaymentIdMissing      = errors.New("unique payment id is required")
)

type Error struct {
	RequestId string `json:"requestID,omitempty"`
	Status    int    `json:"status,omitempty"`
	Message   string `json:"message,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

func wrapError(err error) *Error {
	return &Error{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
	}
}
