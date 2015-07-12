package schedule

import (
	"testing"
	"time"
)

func TestBuildCommitSchedule(t *testing.T) {
	var tests = []struct {
		startDay            time.Time
		numNotAFieldEntries int
	}{
		{time.Date(2009, time.November, 9, 0, 0, 0, 0, time.UTC), 5},
		{time.Date(2016, time.February, 28, 0, 0, 0, 0, time.UTC), 12},
		{time.Date(2016, time.February, 29, 0, 0, 0, 0, time.UTC), 11},
		{time.Date(2013, time.February, 28, 0, 0, 0, 0, time.UTC), 4},
	}
	for _, test := range tests {
		days := getTestDays(test.startDay)
		schedule := BuildCommitSchedule(days)
		numNotAFieldEntries := 0
		for _, row := range schedule {
			for _, entry := range row {
				if entry != EMPTY && entry != NOT_A_FIELD {
					t.Errorf("Entry should be EMPTY or NOT_A_FIELD, but was %v", entry)
				}
				if entry == NOT_A_FIELD {
					numNotAFieldEntries++
				}
			}

		}
		if numNotAFieldEntries != test.numNotAFieldEntries {
			t.Errorf("Expected %d NOT_A_FIELD entries, but got %d",
				test.numNotAFieldEntries, numNotAFieldEntries)
		}

	}
}

func getTestDays(startDay time.Time) []time.Time {
	var resultingDays []time.Time
	days := GetDaysSinceDateMinusOneYear(startDay)
	for day := range days {
		resultingDays = append(resultingDays, day)
	}
	return resultingDays
}

func TestBuildFirstWeek(t *testing.T) {
	var tests = []struct {
		day              time.Weekday
		expectedSchedule []int
	}{
		{time.Sunday, []int{0, 0, 0, 0, 0, 0, 0}},
		{time.Monday, []int{-1, 0, 0, 0, 0, 0, 0}},
		{time.Tuesday, []int{-1, -1, 0, 0, 0, 0, 0}},
		{time.Wednesday, []int{-1, -1, -1, 0, 0, 0, 0}},
		{time.Thursday, []int{-1, -1, -1, -1, 0, 0, 0}},
		{time.Friday, []int{-1, -1, -1, -1, -1, 0, 0}},
		{time.Saturday, []int{-1, -1, -1, -1, -1, -1, 0}},
	}
	for _, test := range tests {
		actualSchedule := buildFirstWeek(test.day)
		if !sliceEqual(actualSchedule, test.expectedSchedule) {
			fmt := "Expected %v as schedule, but got %v"
			t.Errorf(fmt, test.expectedSchedule, actualSchedule)
		}
	}
}

func TestBuildLastWeek(t *testing.T) {
	var tests = []struct {
		day              time.Weekday
		expectedSchedule []int
	}{
		{time.Sunday, []int{0, -1, -1, -1, -1, -1, -1}},
		{time.Monday, []int{0, 0, -1, -1, -1, -1, -1}},
		{time.Tuesday, []int{0, 0, 0, -1, -1, -1, -1}},
		{time.Wednesday, []int{0, 0, 0, 0, -1, -1, -1}},
		{time.Thursday, []int{0, 0, 0, 0, 0, -1, -1}},
		{time.Friday, []int{0, 0, 0, 0, 0, 0, -1}},
		{time.Saturday, []int{0, 0, 0, 0, 0, 0, 0}},
	}
	for _, test := range tests {
		actualSchedule := buildLastWeek(test.day)
		if !sliceEqual(actualSchedule, test.expectedSchedule) {
			fmt := "Expected %v as schedule, but got %v"
			t.Errorf(fmt, test.expectedSchedule, actualSchedule)
		}
	}
}

func sliceEqual(sliceA, sliceB []int) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for i := 0; i < len(sliceA); i++ {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}
	return true
}

func TestConnectWeeksToSchedule(t *testing.T) {
	var tests = []struct {
		firstDay            time.Weekday
		lastDay             time.Weekday
		numEntries          int
		numNotAFieldEntries int
	}{
		{time.Sunday, time.Saturday, 371, 0},
		{time.Monday, time.Saturday, 371, 1},
		{time.Wednesday, time.Wednesday, 371, 6},
		{time.Saturday, time.Sunday, 371, 12},
	}
	for _, test := range tests {
		firstWeek := buildFirstWeek(test.firstDay)
		lastWeek := buildLastWeek(test.lastDay)
		schedule := connectWeeksToSchedule(firstWeek, lastWeek)
		length := 0
		numNotAFieldEntries := 0
		for _, row := range schedule {
			for _, entry := range row {
				if entry != EMPTY && entry != NOT_A_FIELD {
					t.Errorf("Entry should be EMPTY or NOT_A_FIELD, but was %v", entry)
				}
				if entry == NOT_A_FIELD {
					numNotAFieldEntries++
				}
				length++
			}
		}

		if length != test.numEntries {
			t.Errorf("Expected length was %d, but got %d", test.numEntries, length)
		}

		if numNotAFieldEntries != test.numNotAFieldEntries {
			t.Errorf("Expected %d NOT_A_FIELD entries, but got %d",
				test.numNotAFieldEntries, numNotAFieldEntries)
		}
	}
}
