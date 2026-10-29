package advancederrors

import (
	"errors"
	"testing"
)

func TestChargePreservesStableDomainError(t *testing.T) {
	err := Charge("declined", 100)
	if !errors.Is(err, ErrPaymentDeclined) {
		t.Fatalf("Charge() error = %v; want ErrPaymentDeclined in chain", err)
	}
}

func TestChargeReturnsTypedValidationError(t *testing.T) {
	err := Charge("", 100)
	var validation ValidationError
	if !errors.As(err, &validation) || validation.Field != "card_id" {
		t.Fatalf("Charge() error = %v; want card_id ValidationError", err)
	}
}

func TestChargeRejectsNonPositiveAmount(t *testing.T) {
	err := Charge("card", 0)
	var validation ValidationError
	if !errors.As(err, &validation) || validation.Field != "amount_cents" {
		t.Fatalf("Charge() error = %v; want amount_cents ValidationError", err)
	}
}
