package robotname

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	randomName := r.RandomName()
	if randomName == "" {
		return "", fmt.Errorf("There's a problem")
	}
	listOfNames := []string{}
	if len(listOfNames) == 0 {
		r.name = randomName
		listOfNames = append(listOfNames, randomName)
	} else {
		for _, v := range listOfNames {
			if v == randomName {
				return "", fmt.Errorf("There's a problem")
			} else {
				r.name = randomName
				listOfNames = append(listOfNames, randomName)
			}
		}
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
	if r.name != "" {
		fmt.Errorf(`Robot name not cleared on reset.  Still %s.`, r.name)
	}

}

func (r *Robot) RandomName() string {
	if r.name != "" {
		return r.name
	}
	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	number := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomName := []string{}
	for i := 0; i < 2; i++ {
		randomIndexAlpha := r1.Intn(25)
		randomName = append(randomName, alphabet[randomIndexAlpha])
	}
	for j := 0; j < 3; j++ {
		randomIndexNum := r1.Intn(8)
		randomName = append(randomName, number[randomIndexNum])
	}
	fmt.Println(strings.Join(randomName, ""))
	return strings.Join(randomName, "")
}
