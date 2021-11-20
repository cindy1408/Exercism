package clock

import (
	"fmt"
	"time"
)

// Define the Clock type here.
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
	return c.Add(m)
}

func (c Clock) Subtract(m int) Clock {
	return c.Subtract(m)
}

func (c Clock) String() string {
	return c.String()
}

