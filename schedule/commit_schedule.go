package schedule

import (
	"time"
)

type ScheduleEntries int

type CommitSchedule [][]ScheduleEntries

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

// buildFirstWeek creates NUM_WEEK_DAYS schedule entries, where the entries
// before the given week day are NOT_A_FIELD and EMPTY afterwards (including given day)
func buildFirstWeek(day time.Weekday) []ScheduleEntries {
	var firstWeek []ScheduleEntries
	for i := 0; i < NUM_WEEK_DAYS; i++ {
		if i < int(day) {
			firstWeek = append(firstWeek, NOT_A_FIELD)
		} else {
			firstWeek = append(firstWeek, EMPTY)
		}
	}
	return firstWeek
}

// buildLastWeek creates NUM_WEEK_DAYS schedule entries, where the entries
// after the given week day are NOT_A_FIELD and EMPTY before (including given day)
func buildLastWeek(day time.Weekday) []ScheduleEntries {
	var lastWeek []ScheduleEntries
	for i := 0; i < NUM_WEEK_DAYS; i++ {
		if i > int(day) {
			lastWeek = append(lastWeek, NOT_A_FIELD)
		} else {
			lastWeek = append(lastWeek, EMPTY)
		}
	}
	return lastWeek
}

func connectWeeksToSchedule(firstWeek, lastWeek []ScheduleEntries) [][]ScheduleEntries {
	return nil
}
