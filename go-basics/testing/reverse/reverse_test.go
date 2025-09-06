package reverse

import "testing"

func TestReverseToReturnReversedInputString(t *testing.T) {
	actualResult := reverse("Hello")
	var expectedResult = "olleH"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
