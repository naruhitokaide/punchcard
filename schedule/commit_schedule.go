package schedule

import (
	"time"
)

type ScheduleEntries int

type CommitSchedule [7][53]ScheduleEntries

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
	firstWeek := buildFirstWeek(days[0].Weekday())
	lastWeek := buildLastWeek(days[len(days)-1].Weekday())
	return connectWeeksToSchedule(firstWeek, lastWeek)
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

// connectWeeksToSchedule creates a CommitSchedule, by first and last week,
// filling in the weeks inbetween and initializing everything inbetween with EMPTY
func connectWeeksToSchedule(firstWeek, lastWeek []ScheduleEntries) CommitSchedule {
	schedule := new(CommitSchedule)
	for row_index, row := range schedule {
		for column_index, _ := range row {
			if column_index == 0 {
				schedule[row_index][column_index] = firstWeek[row_index]
			} else if column_index == 52 {
				schedule[row_index][column_index] = lastWeek[row_index]
			} else {
				schedule[row_index][column_index] = EMPTY
			}
		}
	}
	return *schedule
}
