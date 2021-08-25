package fisr

import (
	"errors"
	"math"
)

const (
	// The Magic nummber is the first approximation
	// of the log term
	Magic32 uint32 = 0x5F375A86
)

// This function is responsible for calculating
// the inverse square root of the algorithm using
// binary bit manipulation and newton's method.
// The funciton accepts positive 32 bit floating
// point numbers and returns a 32 bit floating point
// number. If the number happens to be negative, Nan
// is returned with an error
func FastInverseSquareRoot(number float32) (inverseSqrt float32, err error) {

	// Check for negative numbers
	if number < 0 {
		inverseSqrt = float32(math.NaN())
		err = errors.New("The number passed was negative")
		return
	}

	// Create a binary representation of the number
	// but with decimal support as per IEEE 754
	numberBits := math.Float32bits(number)

	// Shift it by the magic number and divide by 2
	numberBits = Magic32 - (numberBits >> 1)

	// Convert the number bits back to float32
	inverseSqrt = math.Float32frombits(numberBits)

	// Calculate one half of the number
	oneHalf := 0.5 * number

	// Perform one iteration of Newton's method
	inverseSqrt *= float32(1.5) - (oneHalf * inverseSqrt * inverseSqrt)

	// Since we have the return type names explicitly specified
	// we can just return
	return

}
