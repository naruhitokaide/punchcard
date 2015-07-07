package schedule

import (
	"testing"
	"time"
)

func XTestBuildCommitScheduleFullWeeks(t *testing.T) {
	var tests = []struct {
		startDay            time.Time
		numNotAFieldEntries int
	}{
		// TODO add edge cases, like leap years etc
		{time.Date(2009, time.November, 9, 0, 0, 0, 0, time.UTC), 0},
		{time.Date(2015, time.November, 9, 0, 0, 0, 0, time.UTC), 6},
		{time.Date(2014, time.July, 7, 0, 0, 0, 0, time.UTC), 12},
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
		expectedSchedule []ScheduleEntries
	}{
		{time.Sunday, []ScheduleEntries{0, 0, 0, 0, 0, 0, 0}},
		{time.Monday, []ScheduleEntries{-1, 0, 0, 0, 0, 0, 0}},
		{time.Tuesday, []ScheduleEntries{-1, -1, 0, 0, 0, 0, 0}},
		{time.Wednesday, []ScheduleEntries{-1, -1, -1, 0, 0, 0, 0}},
		{time.Thursday, []ScheduleEntries{-1, -1, -1, -1, 0, 0, 0}},
		{time.Friday, []ScheduleEntries{-1, -1, -1, -1, -1, 0, 0}},
		{time.Saturday, []ScheduleEntries{-1, -1, -1, -1, -1, -1, 0}},
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
		expectedSchedule []ScheduleEntries
	}{
		{time.Sunday, []ScheduleEntries{0, -1, -1, -1, -1, -1, -1}},
		{time.Monday, []ScheduleEntries{0, 0, -1, -1, -1, -1, -1}},
		{time.Tuesday, []ScheduleEntries{0, 0, 0, -1, -1, -1, -1}},
		{time.Wednesday, []ScheduleEntries{0, 0, 0, 0, -1, -1, -1}},
		{time.Thursday, []ScheduleEntries{0, 0, 0, 0, 0, -1, -1}},
		{time.Friday, []ScheduleEntries{0, 0, 0, 0, 0, 0, -1}},
		{time.Saturday, []ScheduleEntries{0, 0, 0, 0, 0, 0, 0}},
	}
	for _, test := range tests {
		actualSchedule := buildLastWeek(test.day)
		if !sliceEqual(actualSchedule, test.expectedSchedule) {
			fmt := "Expected %v as schedule, but got %v"
			t.Errorf(fmt, test.expectedSchedule, actualSchedule)
		}
	}
}

func sliceEqual(sliceA, sliceB []ScheduleEntries) bool {
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
