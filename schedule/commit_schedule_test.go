package schedule

import (
	"testing"
	"time"
)

func TestBuildCommitScheduleFullWeeks(t *testing.T) {
	var tests = []struct {
		startDay time.Time
		numDays  int
	}{
		{time.Date(2009, time.November, 9, 0, 0, 0, 0, time.UTC), 7},
		{time.Date(2015, time.November, 9, 0, 0, 0, 0, time.UTC), 21},
		{time.Date(2014, time.July, 7, 0, 0, 0, 0, time.UTC), 365},
	}
	for _, test := range tests {
		days := getTestDays(test.startDay, test.numDays)
		schedule := BuildCommitSchedule(days)
		// TODO add assertions
	}
}

func TestBuildCommitScheduleWednesday(t *testing.T) {
	var tests = []struct {
		startDay time.Time
		numDays  int
	}{
		{}, // TODO add edge cases
	}
	for _, test := range tests {
		days := getTestDays(test.startDay, test.numDays)
		schedule := BuildCommitSchedule(days)
		// TODO add assertions
	}
}

func TestBuildCommitScheduleThrusdayEnd(t *testing.T) {
	var tests = []struct {
		startDay time.Time
		numDays  int
	}{
		{}, // TODO add edge cases
	}
	for _, test := range tests {
		days := getTestDays(test.startDay, test.numDays)
		schedule := BuildCommitSchedule(days)
		// TODO add assertions
	}
}

func getTestDays(startDay time.Time, numDays int) []time.Time {
	var resultingDays []time.Time
	for i := 0; i < numDays; i++ {
		resultingDays = append(resultingDays, startDay.AddDate(0, 0, i))
	}
	return resultingDays
}
