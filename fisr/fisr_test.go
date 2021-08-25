package fisr

import (
	"math"
	"testing"
)

var (
	// All the test cases to test for
	TestCases = []float32{2, 3.33, 4, 5, 100, 1000, 1000003}
)

const (
	// Since this is an approximation we can only get an estimate
	MarginError = 0.01
)

// This test case checks against all the TestCases and asserts if
// the estimated inverse square root is within the margin of error.
func TestInverseSquareRoot(t *testing.T) {

	for _, testCase := range TestCases {
		if inverseSquareRoot, err := FastInverseSquareRoot(testCase); err == nil {

			referenceValue := float32(1 / math.Sqrt(float64(testCase)))
			if (inverseSquareRoot > referenceValue+MarginError) || (inverseSquareRoot < referenceValue-MarginError) {
				t.Fatalf("Estimated Square root for %.3f=%.3f not within margin of error %.3f", testCase, inverseSquareRoot, MarginError)
			} else {
				t.Logf("Estimated Square root for %.3f=%.3f with actual value %.3f", testCase, inverseSquareRoot, referenceValue)
			}

		} else {
			t.Fatalf("Error while computing Inverse Square root for %.3f: %s", testCase, err)
		}

	}

}
