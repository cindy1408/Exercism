package booking

import "time"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	if t.Before(time.Now()) {
		return true
	}
	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	if t.Hour() < 18 && t.Hour() >= 12{
		return true 
	}
	return false 
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return "You have an appointment on " + t.Format("Monday, January 2, 2006") + "," + " at " + t.Format("15:04") + "."
}


// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	current := time.Now()
	anniversary := "09-15"
	thisYear := current.Format("2006") + "-" + anniversary 
	thisYearAnniversary, _ := time.Parse("2006-1-2", thisYear)
	return thisYearAnniversary
}
