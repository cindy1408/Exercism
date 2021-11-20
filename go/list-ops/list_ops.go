package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

type binFunc func(x, y int) int

type predFunc func(n int) bool 

type unaryFunc func(x int) int 


func (s IntList) Foldl(fn binFunc, initial int) int {
	for i, _ := range s {
		initial = fn(initial, s[i])
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i:= len(s)-1; i >= 0 ; i -- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	matchList := []int{}
	for _, value := range s {
		match := fn(value)
		if match {
			matchList = append(matchList, value)
		}
	}
	return matchList
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	m := []int{}
	for _, value := range s {
		m = append(m, fn(value))
	}
	return m
}

func (s IntList) Reverse() IntList {
	reverse := IntList{}
	for i:= len(s)-1; i>=0; i-- {
		reverse = append(reverse, s[i])
	}
	return reverse
}

func (s IntList) Append(lst IntList) IntList {
	s = append(s, lst ...)
	return s
}

func (s IntList) Concat(lists []IntList) IntList {
	flattenList := []int{}
	for _, list := range lists {
		flattenList = append(flattenList, list ...)
	}
	s = append(s, flattenList...)
	return s
}
