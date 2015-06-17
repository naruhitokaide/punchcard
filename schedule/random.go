package schedule

import (
	"fmt"
	"math/rand"
)

func RandomSchedule(min, max) {
	fmt.Println("Start random scheduler")

	// representation of the past 365 days
	// for each day
	// get random value between --min and --max
	rnd := getRandomNumber(min, max)
	// save into structure representing the commits over the last year
	// start worker, which will execute all commits using some sort of
	// commit generator
}

// getRandomNumber returns a number in the range of min and max.
func getRandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}
