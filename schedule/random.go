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
// Every day maps to itself minus one year except the 29.02 will map to 28.02.
func getDaysSinceThisDayLastYear() []time.Date {
	return
}
