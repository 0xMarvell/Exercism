package grains

import "errors"

// Square calculates the amount of grains on a given square (on a chessboard)
// Note that the chessboard has 64 sqaures
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("square value should be greater than 0 and less or equal 64")
	}
	return 1 << (number - 1), nil
}

// Total calculates the total number of grains on the chessboard
func Total() uint64 {
	return 1<<64 - 1
}
