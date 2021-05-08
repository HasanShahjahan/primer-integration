package primer

import "time"

// TransactionStatus ...
type TransactionStatus string

// TransactionStatus enum
const (
	Pending          TransactionStatus = "PENDING"
	Failed           TransactionStatus = "FAILED"
	Authorized       TransactionStatus = "AUTHORIZED"
	Settling         TransactionStatus = "SETTLING"
	PartiallySettled TransactionStatus = "PARTIALLY_SETTLED"
	Settled          TransactionStatus = "SETTLED"
	Declined         TransactionStatus = "DECLINED"
	Cancelled        TransactionStatus = "CANCELLED"
)

// PaymentInstrumentType ...
type PaymentInstrumentType string

// PaymentInstrumentType enum
const (
	PaymentCard              PaymentInstrumentType = "PAYMENT_CARD"
	PaypalOrder              PaymentInstrumentType = "PAYPAL_ORDER"
	PaypalBillingAgreement   PaymentInstrumentType = "PAYPAL_BILLING_AGREEMENT"
	GooglePay                PaymentInstrumentType = "GOOGLE_PAY"
	GoCardlessMandate        PaymentInstrumentType = "GOCARDLESS_MANDATE"
	KlarnaAuthorizationToken PaymentInstrumentType = "KLARNA_AUTHORIZATION_TOKEN"
	KlarnaCustomerToken      PaymentInstrumentType = "KLARNA_CUSTOMER_TOKEN"
	ApplePay                 PaymentInstrumentType = "APPLE_PAY"
)

// TokenType ...
type TokenType string

// TokenType enum
const (
	MultiUse  TokenType = "MULTI_USE"
	SingleUse TokenType = "SINGLE_USE"
)

// BaseDto ...
type BaseDto struct {
	RequestId      *string `json:"requestId,omitempty"`
	IdempotencyKey *string `json:"xIdempotencyKey,omitempty"`
}

// PaymentInstrument ...
type PaymentInstrument struct {
	Token                      string                      `json:"token"`
	AnalyticsId                *string                     `json:"analyticsId,omitempty"`
	TokenType                  *string                     `json:"tokenType,omitempty"`
	PaymentInstrumentType      *string                     `json:"paymentInstrumentType,omitempty"`
	PaymentInstrumentData      *interface{}                `json:"paymentInstrumentData,omitempty"`
	ThreeDSecureAuthentication *ThreeDSecureAuthentication `json:"threeDSecureAuthentication,omitempty"`
}

// PaymentCardToken ...
type PaymentCardToken struct {
	Last4Digits        string   `json:"last4Digits"`
	ExpirationMonth    string   `json:"expirationMonth"`
	ExpirationYear     string   `json:"expirationYear"`
	CardHolderName     *string  `json:"cardholderName,omitempty"`
	Network            *string  `json:"network,omitempty"`
	IsNetworkTokenized *bool    `json:"isNetworkTokenized,omitempty"`
	BinData            *BinData `json:"binData,omitempty"`
}

// BinData ...
type BinData struct {
	Network                    string  `json:"network"`
	RegionalRestriction        string  `json:"regionalRestriction"`
	AccountNumberType          string  `json:"accountNumberType"`
	AccountFundingType         string  `json:"accountFundingType"`
	PrepaidReloadableIndicator string  `json:"prepaidReloadableIndicator"`
	ProductUsageType           string  `json:"productUsageType"`
	ProductCode                string  `json:"productCode"`
	ProductName                string  `json:"productName"`
	IssuerCountryCode          *string `json:"issuerCountryCode,omitempty"`
	IssuerName                 *string `json:"issuerName,omitempty"`
	IssuerCurrencyCode         *string `json:"issuerCurrencyCode,omitempty"`
}

// ThreeDSecureAuthentication ...
type ThreeDSecureAuthentication struct {
	ResponseCode    string  `json:"responseCode"`
	ReasonCode      *string `json:"reasonCode,omitempty"`
	ReasonText      *string `json:"reasonText,omitempty"`
	ProtocolVersion *string `json:"protocolVersion,omitempty"`
	ChallengeIssued *bool   `json:"challengeIssued,omitempty"`
}

// Customer ...
type Customer struct {
	Id              string           `json:"id"`
	Email           *string          `json:"email,omitempty"`
	BillingAddress  *BillingAddress  `json:"billingAddress,omitempty"`
	ShippingAddress *ShippingAddress `json:"shippingAddress,omitempty"`
}

// BillingAddress ...
type BillingAddress struct {
	Address
}

// ShippingAddress ...
type ShippingAddress struct {
	Address
}

// Address ...
type Address struct {
	AddressLine1 string  `json:"addressLine1"`
	City         string  `json:"city"`
	CountryCode  string  `json:"countryCode"`
	PostalCode   string  `json:"postalCode"`
	FirstName    *string `json:"firstName,omitempty"`
	LastName     *string `json:"lastName,omitempty"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	State        *string `json:"state,omitempty"`
}

// Transaction ...
type Transaction struct {
	Id                     string        `json:"id"`
	Processor              string        `json:"processor"`
	TransactionType        string        `json:"type"`
	Status                 string        `json:"status"`
	ProcessorTransactionId *string       `json:"processorTransactionId,omitempty"`
	PaymentError           *PaymentError `json:"paymentError,omitempty"`
}

// PaymentError ...
type PaymentError struct {
	Date             time.Time `json:"date"`
	PaymentErrorType string    `json:"type"`
	DeclineCode      *string   `json:"declineCode,omitempty"`
	DeclineType      *string   `json:"declineType,omitempty"`
	ProcessorMessage *string   `json:"processorMessage,omitempty"`
}

// Payment ...
type Payment struct {
	Id                       string                    `json:"id"`
	Date                     time.Time                 `json:"date"`
	Status                   string                    `json:"status"`
	OrderId                  string                    `json:"orderId"`
	CurrencyCode             string                    `json:"currencyCode"`
	Amount                   int64                     `json:"amount"`
	AmountAuthorized         int64                     `json:"amountAuthorized"`
	AmountCapture            int64                     `json:"amountCaptured"`
	AmountRefunded           int64                     `json:"amountRefunded"`
	PaymentInstrument        PaymentInstrument         `json:"paymentInstrument"`
	Transactions             []Transaction             `json:"transactions"`
	Processor                *string                   `json:"processor,omitempty"`
	RequiredAction           *RequiredAction           `json:"requiredAction,omitempty"`
	StatementDescriptor      *string                   `json:"statementDescriptor,omitempty"`
	VaultedPaymentInstrument *VaultedPaymentInstrument `json:"vaultedPaymentInstrument,omitempty"`
	Customer                 *Customer                 `json:"customer,omitempty"`
	LastPaymentError         *PaymentError             `json:"lastPaymentError,omitempty"`
	Metadata                 *map[string]string        `json:"metadata"`
	WorkflowExecutionError   *WorkflowExecutionError   `json:"workflowExecutionError,omitempty"`
}

// RequiredAction ...
type RequiredAction struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ClientToken *string `json:"clientToken,omitempty"`
}

// VaultedPaymentInstrument ...
type VaultedPaymentInstrument struct {
	Token                           string                      `json:"token"`
	AnalyticsId                     string                      `json:"analyticsId,omitempty"`
	TokenType                       string                      `json:"tokenType,omitempty"`
	PaymentInstrumentType           string                      `json:"paymentInstrumentType,omitempty"`
	PaymentInstrumentData           interface{}                 `json:"paymentInstrumentData,omitempty"`
	ThreeDomainSecureAuthentication *ThreeDSecureAuthentication `json:"threeDSecureAuthentication,omitempty"`
}

// WorkflowExecutionError ...
type WorkflowExecutionError struct {
	Reason string  `json:"reason"`
	StepId *string `json:"stepId,omitempty"`
}

// PaymentResponse ...
type PaymentResponse struct {
	BaseDto
	Payment
}

func (b *BaseDto) SetRequestId(id string) {
	b.RequestId = &id
}

// CreatePayment ...
type CreatePayment struct {
	OrderId             string             `json:"orderId"`
	CurrencyCode        string             `json:"currencyCode"`
	Amount              int64              `json:"amount"`
	PaymentInstrument   PaymentInstrument  `json:"paymentInstrument"`
	StatementDescriptor *string            `json:"statementDescriptor,omitempty"`
	Customer            *Customer          `json:"customer,omitempty"`
	Metadata            *map[string]string `json:"metadata,omitempty"`
}

// CreatePaymentRequest ...
type CreatePaymentRequest struct {
	BaseDto
	CreatePayment
}

// CapturePayment ...
type CapturePayment struct {
	Amount int64 `json:"amount"`
	Final  bool  `json:"final"`
}

// CapturePaymentRequest ...
type CapturePaymentRequest struct {
	BaseDto
	CapturePayment
}

// CancelPayment ...
type CancelPayment struct {
	Reason string `json:"reason"`
}

// CancelPaymentRequest ...
type CancelPaymentRequest struct {
	BaseDto
	CancelPayment
}

// RefundPayment ...
type RefundPayment struct {
	Amount  int64   `json:"amount"`
	OrderId *string `json:"orderId,omitempty"`
	Reason  *string `json:"reason,omitempty"`
}

// RefundPaymentRequest ...
type RefundPaymentRequest struct {
	BaseDto
	RefundPayment
}

type ResumePayment struct {
	ResumeToken string `json:"resumeToken"`
}

// ResumePaymentRequest ...
type ResumePaymentRequest struct {
	BaseDto
	ResumePayment
}

// SearchPayment ...
type SearchPayment struct {
	Data       []Data  `json:"data"`
	NextCursor *string `json:"nextCursor,omitempty"`
	PrevCursor *string `json:"prevCursor,omitempty"`
}

// Data ...
type Data struct {
	Id           string    `json:"id"`
	Date         time.Time `json:"date"`
	Status       string    `json:"status"`
	OrderId      string    `json:"orderId"`
	CurrencyCode string    `json:"currencyCode"`
	Amount       int64     `json:"amount"`
	Processor    *string   `json:"processor,omitempty"`
}

// SearchPaymentResponse ...
type SearchPaymentResponse struct {
	BaseDto
	SearchPayment
}

// GetPaymentRequest ...
type GetPaymentRequest struct {
	BaseDto
}
