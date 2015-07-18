package schedule

import (
	"time"
)

type CommitSchedule [7][53]ScheduleEntry

type ScheduleEntry struct {
	DateTime   time.Time
	NumCommits int
}

const (
	// NOT_A_FIELD   = -1
	EMPTY         = 0
	NUM_WEEK_DAYS = 7
)

var NOT_A_FIELD ScheduleEntry = ScheduleEntry{NumCommits: -1}

// BuildCommitSchedule returns an empty CommitSchedule, where all fiels are
// initialized with EMPTY except those which are not in the range of days.
// The CommitSchedule is a table of ints.
func BuildCommitSchedule(days []time.Time) CommitSchedule {
	firstWeek := buildFirstWeek(days[0])
	lastWeek := buildLastWeek(days[len(days)-1])
	return connectWeeksToSchedule(firstWeek, lastWeek)
}

// IsNotAField returns true if the given entry has the same datetime as NOT_A_FIELD.
func IsNotAField(entry ScheduleEntry) bool {
	return entry.DateTime.Equal(NOT_A_FIELD.DateTime)
}

// buildFirstWeek creates NUM_WEEK_DAYS schedule entries, where the entries
// before the given week day are NOT_A_FIELD and EMPTY afterwards (including given day)
func buildFirstWeek(day time.Time) []ScheduleEntry {
	var firstWeek []ScheduleEntry
	weekday := day.Weekday()
	for i := 0; i < NUM_WEEK_DAYS; i++ {
		if i < int(weekday) {
			firstWeek = append(firstWeek, NOT_A_FIELD)
		} else {
			firstWeek = append(firstWeek, ScheduleEntry{day, EMPTY})
			day = day.AddDate(0, 0, 1)
		}
	}
	return firstWeek
}

// buildLastWeek creates NUM_WEEK_DAYS schedule entries, where the entries
// after the given week day are NOT_A_FIELD and EMPTY before (including given day)
func buildLastWeek(day time.Time) []ScheduleEntry {
	var lastWeek []ScheduleEntry
	weekday := day.Weekday()
	day = getFirstDayOfWeek(day)
	for i := 0; i < NUM_WEEK_DAYS; i++ {
		if i > int(weekday) {
			lastWeek = append(lastWeek, NOT_A_FIELD)
		} else {
			lastWeek = append(lastWeek, ScheduleEntry{day, EMPTY})
			day = day.AddDate(0, 0, 1)
		}
	}
	return lastWeek
}

// getFirstDayOfWeek returns the first day of the given days week.
// i.e. we always return the last sunday before the given day.
func getFirstDayOfWeek(day time.Time) time.Time {
	return day.AddDate(0, 0, -int(day.Weekday()))
}

// connectWeeksToSchedule creates a CommitSchedule, by first and last week,
// filling in the weeks inbetween and initializing everything inbetween with EMPTY
func connectWeeksToSchedule(firstWeek, lastWeek []ScheduleEntry) CommitSchedule {
	schedule := new(CommitSchedule)
	var day = firstWeek[len(firstWeek)-1].DateTime
	for row_index, row := range schedule {
		for column_index, _ := range row {
			if column_index == 0 {
				schedule[row_index][column_index] = firstWeek[row_index]
			} else if column_index == 52 {
				schedule[row_index][column_index] = lastWeek[row_index]
			} else {
				day = day.AddDate(0, 0, 1)
				schedule[row_index][column_index] = ScheduleEntry{day, EMPTY}
			}
		}
	}
	return *schedule
}
