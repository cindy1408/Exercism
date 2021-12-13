package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	errorMessage string
	numOfCows int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %v cows", s.numOfCows)
}

func someError(num int) error {
	return &SillyNephewError{
		numOfCows: num,
	}
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	if cows == 0 {
		return 0, errors.New("Division by zero")
	}
	if cows < 0 {
		return 0, someError(cows)
	}
	if err == ErrScaleMalfunction {
		fmt.Println("this is scale Malfunction", err)
		fodder*= 2
	} else if err != nil {
		return 0, err
	}
	if fodder < 0 {
		return 0, errors.New("Negative fodder")
	}
	fodderEachCow := fodder / float64(cows)
	fmt.Println(fodderEachCow, nil)
	return fodderEachCow, nil
}
