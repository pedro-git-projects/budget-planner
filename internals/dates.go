package internals

import "time"

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func currentMonthNumberOfDays() float64 {
	now := time.Now()
	cYear, cMonth, _ := now.Date()
	cLocation := now.Location()
	firstDay := time.Date(cYear, cMonth, 1, 0, 0, 0, 0, cLocation)
	lastDay := firstDay.AddDate(0, 1, -1)

	number := float64(lastDay.Day() - firstDay.Day() + 1)
	return number
}
