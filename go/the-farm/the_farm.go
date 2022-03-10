package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction && amount > 0 {
		amount *= 2
		return amount / float64(cows), nil
	} else if amount < 0 {
		if err == ErrScaleMalfunction || err == nil {
			return 0, errors.New("negative fodder")
		}
		return 0, err
	} else if float64(cows) == 0 {
		return 0, errors.New("division by zero")
	} else if float64(cows) < 0 {
		return 0, &SillyNephewError{cows: cows}
	} else if err != nil {
		return 0, err
	} else {
		return amount / float64(cows), nil
	}

}
