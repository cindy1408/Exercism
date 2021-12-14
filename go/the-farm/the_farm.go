package thefarm

import (
	"errors"
	"fmt"
)

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

func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	if cows == 0 {
		return 0, errors.New("Division by zero")
	}
	if cows < 0 {
		return 0, someError(cows)
	}
	if err == ErrScaleMalfunction {
		fodder*= 2
	} else if err != nil {
		return 0, err
	}
	if fodder < 0 {
		return 0, errors.New("Negative fodder")
	}
	return fodder / float64(cows), nil
}
