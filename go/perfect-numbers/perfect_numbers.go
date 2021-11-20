package perfect

import (
	"errors"
)

type Classification string

const ClassificationPerfect Classification = "ClassificationPerfect"
const ClassificationDeficient Classification = "ClassificationDeficient"
const ClassificationAbundant Classification = "ClassificationAbundant"
var ErrOnlyPositive error = errors.New("only positive numbers")

func Classify(n int64) (Classification, error) {
	if n < 0 {
		return "only errors", ErrOnlyPositive
	}
	factors := []int{}
	for i:=0; i < int(n); i++ {
		if int(n) % i == 0 {
			factors = append(factors, i)
		}
	}
	total := 0 
	for i, value := range factors {
		if i == 1 {
			total = value 
		}
		total = total + value
	}
	if int(n) == total {
		return ClassificationPerfect, nil 
	} else if int(n) < total {
		return ClassificationDeficient, nil 
	} else if int(n) > total {
		return ClassificationAbundant, nil 
	}
	return "", nil
}
