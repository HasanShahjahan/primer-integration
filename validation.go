package primer

func (c *Client) Validate(idempotencyKey *string) error {
	if StrSafeDeref(idempotencyKey) == "" {
		return IdempotencyKeyMissing
	}
	return nil
}
