package primer

const (
	Search  = "search"
	Create  = "create"
	Capture = "capture"
	Cancel  = "cancel"
	Refund  = "refund"
	Resume  = "resume"
	Get     = "get"
)

func (c *Client) Validate(idempotencyKey *string, apiType string, paymentId string) error {

	switch apiType {
	case Search:
		return nil
	case Create:
		if StrSafeDeref(idempotencyKey) == "" {
			return IdempotencyKeyMissing
		}
	case Capture, Cancel, Refund:
		if StrSafeDeref(idempotencyKey) == "" {
			return IdempotencyKeyMissing
		} else if paymentId == "" {
			return PaymentIdMissing
		}
	case Get, Resume:
		if paymentId == "" {
			return PaymentIdMissing
		}
	}
	return nil
}
