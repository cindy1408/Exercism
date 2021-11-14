package main

import (
	"fmt"
	"time"
)

type Clock struct {
	hour int 
	minutes int
}

func New(h, m int) Clock {
	inputTime := time.Date(0, 0, 0, h, m, 0, 0, time.UTC)
	fmt.Println(inputTime)
	clockInput := Clock{
		hour: inputTime.Hour(), 
		minutes: inputTime.Minute(),
	}
	fmt.Println(clockInput)
	return clockInput
}

func (c Clock) Add(m int) Clock {
	fmt.Println(c.Add(m))
	return c.Add(m)
}

func (c Clock) Subtract(m int) Clock {
	fmt.Println(c.Subtract(m))
	return c.Subtract(m)
}

func (c Clock) String() string {
	fmt.Println(c.String())
	return c.String()
}

func main() {
	New(2, 30)
}
