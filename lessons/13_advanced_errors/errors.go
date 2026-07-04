package advancederrors

import (
	"errors"
	"fmt"
)

var ErrPaymentDeclined = errors.New("payment declined")

type ValidationError struct {
	Field string
	Rule  string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s failed %s validation", e.Field, e.Rule)
}

func Charge(cardID string, amountCents int) error {
	if cardID == "" {
		return ValidationError{Field: "card_id", Rule: "required"}
	}
	if amountCents <= 0 {
		return ValidationError{Field: "amount_cents", Rule: "positive"}
	}
	if cardID == "declined" {
		return fmt.Errorf("charge card %q: %w", cardID, ErrPaymentDeclined)
	}
	return nil
}
