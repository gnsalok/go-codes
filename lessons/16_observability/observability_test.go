package observability

import "testing"

func TestIsSuccessOrRedirect(t *testing.T) {
	for _, status := range []int{200, 201, 302} {
		if !IsSuccessOrRedirect(status) {
			t.Fatalf("IsSuccessOrRedirect(%d) = false; want true", status)
		}
	}
	if IsSuccessOrRedirect(500) {
		t.Fatal("IsSuccessOrRedirect(500) = true; want false")
	}
}
