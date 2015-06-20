package schedule

import (
	"math/rand"
	"time"
)

// RandomSchedule creates random commits over the past 365 days.
// These commits will be created in the location specified in the command.
func RandomSchedule(min, max) {

	days := getDaysSinceThisDayLastYear()
	for day := range days {
		rnd := getRandomNumber(min, max)
		// save into structure representing the commits over the last year
		// start worker, which will execute all commits using some sort of
		// commit generator
	}
}

// getRandomNumber returns a number in the range of min and max.
func getRandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

// getDaysSinceThisDayLastYear returns a slice of days since todays date
// last year. E.g. 01.01.2015 starts at the 01.01.2014.
func getDaysSinceThisDayLastYear() []time.Time {
	daysSinceThisDayLastYear := make([]time.Time)
	now := time.Now()
	day := getDayLastYear(now)
	for day <= now {
		daysSinceThisDayLastYear.append(day)
		day = day.AddDate(0, 0, 1)
	}
	return daysSinceThisDayLastYear
}

// getDayLastYear returns the daya date minus one year, except the
// 29.02 will map to 28.02.
func getDayLastYear(day time.Time) time.Time {
	if isLeapDay(day) {
		// adjust for one year and one day
		return day.AddDate(-1, 0, -1)
	} else {
		return day.AddDate(-1, 0, 0)
	}
}

// isLeapDay checks if a given datetime is the 29.02 or not.
func isLeapDay(today time.Time) bool {
	_, month, day := today.Date()
	return (day == 29 && month == 2)
}
