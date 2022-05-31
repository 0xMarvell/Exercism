package allyourbase

import (
	"errors"
	"math"
)

var inputBaseErr = errors.New("input base must be >= 2")
var outputBaseErr = errors.New("output base must be >= 2")
var inputDigitsErr = errors.New("all digits must satisfy 0 <= d < input base")

// ConvertToBase converts from one number base to another.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {

	err := validateInput(inputBase, inputDigits, outputBase)
	if err != nil {
		return nil, err
	}

	if inputBase == 10 {
		return fromDecimal(inputDigits, outputBase), nil
	} else if outputBase == 10 {
		return toDecimal(inputBase, inputDigits), nil
	}

	decimal := toDecimal(inputBase, inputDigits)
	return fromDecimal(decimal, outputBase), nil

}

// validateInput checks for errors in the provided input.
func validateInput(inputBase int, inputDigits []int, outputBase int) error {
	if inputBase < 2 {
		return inputBaseErr
	}

	for _, digit := range inputDigits {
		if digit < 0 || inputBase <= digit {
			return inputDigitsErr
		}
	}

	if outputBase < 2 {
		return outputBaseErr
	}

	return nil
}

// sliceToInt converts input from slice of integers to a sinle integer value.
func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}

	return res
}

// intToSlice converts integer to slice of integers
func intToSlice(i int) []int {
	digits := make([]int, 0)
	for i != 0 {
		digits = append(digits, i%10)
		i = i / 10
	}

	return reverseSlice(digits)
}

// reverseSlice reverses the elements in a slice.
func reverseSlice(digits []int) []int {
	reversed := make([]int, 0, len(digits))
	for i := len(digits) - 1; i >= 0; i-- {
		reversed = append(reversed, digits[i])
	}

	return reversed
}

// fromDecimal converts input from decimal (base 10) to the specified base.
func fromDecimal(inputDigits []int, outputBase int) []int {
	digit := sliceToInt(inputDigits)

	if digit == 0 {
		return []int{0}
	}

	result := make([]int, 0)
	for digit != 0 {
		result = append(result, digit%outputBase)
		digit = digit / outputBase
	}

	return reverseSlice(result)
}

// toDecimal converts input to decimal (base 10).
func toDecimal(inputBase int, inputDigits []int) []int {
	var sum int
	for i, digit := range inputDigits {
		q := len(inputDigits) - 1 - i // 2, 1, 0 for 3-digit number
		result := digit * int(math.Pow(float64(inputBase), float64(q)))
		sum += result
	}

	if len(inputDigits) == 0 {
		return []int{0}
	}

	return intToSlice(sum)
}
