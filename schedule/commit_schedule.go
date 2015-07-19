package schedule

import (
	"bytes"
	"fmt"
	"time"
)

type CommitSchedule [7][53]ScheduleEntry

type ScheduleEntry struct {
	DateTime   time.Time
	NumCommits int
}

const (
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
	// var adjustedDay time.Time
	for rowIndex, row := range schedule {
		for columnIndex, _ := range row {
			if columnIndex == 0 {
				schedule[rowIndex][columnIndex] = firstWeek[rowIndex]
			} else if columnIndex == 52 {
				schedule[rowIndex][columnIndex] = lastWeek[rowIndex]
			} else {
				adjustedDay := day.AddDate(0, 0, getDeltaDays(rowIndex, columnIndex))
				schedule[rowIndex][columnIndex] = ScheduleEntry{adjustedDay, EMPTY}
			}
		}
	}
	return *schedule
}

// getDeltaDays returns the day which need to be added to the last day of the
// first week to get the date for the position at (rowIndex, columnIndes).
func getDeltaDays(rowIndex, columnIndex int) int {
	return columnIndex*7 - (6 - rowIndex)
}

// String returns a string representing the CommitSchedule.
func (schedule CommitSchedule) String() string {
	var buffer bytes.Buffer
	for _, row := range schedule {
		for _, entry := range row {
			// entryString := fmt.Sprintf("(%s,%d) ", entry.DateTime.String(), entry.NumCommits)
			entryString := fmt.Sprintf("(%s,%d) ", entry.DateTime.Weekday().String(), entry.NumCommits)
			buffer.WriteString(entryString)
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
