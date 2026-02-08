package errors

import (
	"testing"
)

func Test_Divide(t *testing.T) {
	testCases := []struct {
		testName       string
		a              float64
		b              float64
		expectedResult float64
		expectedError  error
	}{
		{testName: "Pass", a: 10, b: 2, expectedResult: 5, expectedError: nil},
		{testName: "Fail", a: 10, b: 0, expectedResult: 0, expectedError: ErrDivideByZero},
	}

	for _, tc := range testCases {
		result, err := Divide(tc.a, tc.b)

		if err != nil && tc.expectedError == nil {
			t.Errorf("Test (%f, %f): Expected no error, got %v", tc.a, tc.b, err)
		} else if err == nil && tc.expectedError != nil {
			t.Errorf("Test (%f, %f): Expected error '%v', got nil", tc.a, tc.b, tc.expectedError)
		} else if err != nil && tc.expectedError != nil && err != tc.expectedError {
			t.Errorf("Test (%f, %f): Expected error '%v', got '%v'", tc.a, tc.b, tc.expectedError, err)
		}

		if result != tc.expectedResult {
			t.Errorf("Test (%f, %f): Expected result %f, got %f", tc.a, tc.b, tc.expectedResult, result)
		}

	}

}
