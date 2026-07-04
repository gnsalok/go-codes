package observability

import "testing"

func TestIsSuccessful(t *testing.T) {
	for _, status := range []int{200, 201, 302} {
		if !IsSuccessful(status) {
			t.Fatalf("IsSuccessful(%d) = false; want true", status)
		}
	}
	if IsSuccessful(500) {
		t.Fatal("IsSuccessful(500) = true; want false")
	}
}
