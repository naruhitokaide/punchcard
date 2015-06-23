package schedule

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomSchedule creates random commits over the past 365/366 days.
// These commits will be created in the location specified in the command.
func RandomSchedule(min, max int) {

	days := getDaysSinceDateMinusOneYear(time.Now())
	for day := range days {
		rnd := getRandomNumber(min, max)
		fmt.Println("%v - %d", day, rnd)
		// save into structure representing the commits over the last year
		// start worker, which will execute all commits using some sort of
		// commit generator
	}
}

// getRandomNumber returns a number in the range of min and max.
func getRandomNumber(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

// getDaysSinceDateMinusOneYear returns a slice of days since the given date
// minus one year. E.g. 01.01.2015 starts at the 01.01.2014.
func getDaysSinceDateMinusOneYear(givenDate time.Time) chan time.Time {
	dayChannel := make(chan time.Time)
	go func() {
		day := getDayMinusOneYear(givenDate)
		for givenDate.After(day) {
			dayChannel <- day
			day = day.AddDate(0, 0, 1)
		}
		// also add the givenDate, which will not be added using After()
		dayChannel <- givenDate
		close(dayChannel)
	}()
	return dayChannel
}

// getDayMinusOneYear returns the daya date minus one year, except the
// 29.02 will map to 28.02.
func getDayMinusOneYear(day time.Time) time.Time {
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
	return (day == 29 && month == time.February)
}
