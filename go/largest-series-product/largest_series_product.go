package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}

	if span > len(digits) {
		return -1, errors.New("span must be smaller than string length")
	}

	maxProduct := -1
	for i := 0; i < len(digits)-span+1; i++ {
		internalSlice := digits[i : i+span]

		prod := 1
		for _, ch := range internalSlice {
			digit := int(ch) - int('0')
			if digit < 0 || digit > 9 {
				return -1, errors.New("digits input must only contain digits")
			}
			prod *= digit
		}
		if prod > maxProduct {
			maxProduct = prod
		}
	}

	return int64(maxProduct), nil
}
