package main

import (
	"fmt"
	"time"
)

func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

	if t.Before(time.Now()) {
		return true
	}
	return false
}

func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	if t.Hour() < 18 && t.Hour() >= 12 {
		return true
	}
	return false
}

func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return "You have an appointment on " + t.Format("Monday, January 2, 2006") + "," + " at " + t.Format("15:04") + "."
}

func AnniversaryDate() time.Time {
	current := time.Now()
	anniversary := "09-15"
	thisYear := current.Format("2006") + "-" + anniversary 
	thisYearAnniversary, _ := time.Parse("2006-1-2", thisYear)
	return thisYearAnniversary
}

func main(){
	fmt.Println(Schedule("7/25/2019 13:45:00")) //fine
	fmt.Println(Schedule("11/28/1984 2:02:02"))
	fmt.Println(Schedule("2/29/2112 11:59:59"))
	fmt.Println(Schedule("7/13/2020 20:32:00")) //fine 
	fmt.Println(Schedule("2/29/2112 11:59:59"))

	fmt.Println(HasPassed("July 25, 2019 13:45:00"))
	fmt.Println(IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00"))
	fmt.Println(IsAfternoonAppointment("Friday, March 8, 1974 12:02:02"))
	fmt.Println(IsAfternoonAppointment("Friday, September 9, 2112 11:59:59"))

	fmt.Println(Description("7/25/2019 13:45:00"))
	fmt.Println(AnniversaryDate())
}
