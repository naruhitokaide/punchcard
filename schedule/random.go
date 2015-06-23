package schedule

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomSchedule creates random commits over the past 365/366 days.
// These commits will be created in the location specified in the command.
func RandomSchedule(min, max int) {
	// TODO add git init
	days := getDaysSinceDateMinusOneYear(time.Now())
	for day := range days {
		rnd := getRandomNumber(min, max)
		commits := RandomCommits(day, rnd)
		for commit := range commits {
			// TODO git add and commit file with given
			fmt.Println(commit.message)
		}
	}
}

// getRandomNumber returns a number in the range of min and max.
func getRandomNumber(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}
