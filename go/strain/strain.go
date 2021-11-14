package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var keep Ints
	for _, individualInt := range(i) {
		if filter(individualInt) {
			keep = append(keep, individualInt)
		}
	}
	return keep
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var discard Ints
	for _, individualInt := range(i) {
		if !filter(individualInt) {
			discard = append(discard, individualInt)
		}
	}
	return discard
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var keep Lists
	for _, individualList := range(l) {
		if filter(individualList) {
			keep = append(keep, individualList)
		}
	}
	return keep
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var keep Strings
	for _, individualString := range(s) {
		if filter(individualString) {
			keep = append(keep, individualString)
		}
	}
	return keep
}
