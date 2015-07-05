package schedule

import (
	"time"
)

type ScheduleEntries int

type CommitSchedule [][]int

const (
	NOT_A_FIELD ScheduleEntries = -1
	EMPTY       ScheduleEntries = 0
	ONE         ScheduleEntries = 1
	TWO         ScheduleEntries = 2
	THREE       ScheduleEntries = 3
	FOUR        ScheduleEntries = 4

	NUM_WEEK_DAYS = 7
)

// BuildCommitSchedule returns an empty CommitSchedule, where all fiels are
// initialized with EMPTY except those which are not in the range of days.
// The CommitSchedule is a table of ints.
func BuildCommitSchedule(days []time.Time) CommitSchedule {
	// get weeks, which determine width and height is seven
	// fill entries with EMPTY or NOT_A_FIELD
	schedule := make(CommitSchedule, 0) // TODO figure out num weeks
	// firstWeek := buildFirstWeek(days[0].Weekday())
	// lastWeek := buildLastWeek(days[len(days)-1].Weekday())
	// TODO get days inbetween first and last week and join them
	return schedule
}

func buildFirstWeek(day time.Weekday) []int {
	var firstWeek []int
	for i := 0; i < NUM_WEEK_DAYS; i++ {
		firstWeek = append(firstWeek, i)
	}
	return firstWeek
}

func buildLastWeek(day time.Weekday) []int {
	var lastWeek []int
	for i := 0; i < NUM_WEEK_DAYS; i++ {
		lastWeek = append(lastWeek, i)
	}
	return lastWeek
}
