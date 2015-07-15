package schedule

import (
	"testing"
	"time"
)

var testSunday = time.Date(2015, time.February, 1, 0, 0, 0, 0, time.UTC)
var testMonday = time.Date(2015, time.February, 2, 0, 0, 0, 0, time.UTC)
var testTuesday = time.Date(2015, time.February, 3, 0, 0, 0, 0, time.UTC)
var testWednesday = time.Date(2015, time.February, 4, 0, 0, 0, 0, time.UTC)
var testThursday = time.Date(2015, time.February, 5, 0, 0, 0, 0, time.UTC)
var testFriday = time.Date(2015, time.February, 6, 0, 0, 0, 0, time.UTC)
var testSaturday = time.Date(2015, time.February, 7, 0, 0, 0, 0, time.UTC)

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
				if entry.NumCommits != EMPTY && entry != NOT_A_FIELD {
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

func TestIsNotAField(t *testing.T) {
	var tests = []struct {
		entry       ScheduleEntry
		isNotAField bool
	}{
		{NOT_A_FIELD, true}, {ScheduleEntry{testMonday, 1}, false},
	}
	for _, test := range tests {
		actual := IsNotAField(test.entry)
		if actual != test.isNotAField {
			t.Errorf("Expected IsNotAField to be %b, but was %b", test.isNotAField, actual)
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
		day             time.Time
		expectedCommits []int
	}{
		{testSunday, []int{0, 0, 0, 0, 0, 0, 0}},
		{testMonday, []int{-1, 0, 0, 0, 0, 0, 0}},
		{testTuesday, []int{-1, -1, 0, 0, 0, 0, 0}},
		{testWednesday, []int{-1, -1, -1, 0, 0, 0, 0}},
		{testThursday, []int{-1, -1, -1, -1, 0, 0, 0}},
		{testFriday, []int{-1, -1, -1, -1, -1, 0, 0}},
		{testSaturday, []int{-1, -1, -1, -1, -1, -1, 0}},
	}
	for _, test := range tests {
		actualSchedule := buildFirstWeek(test.day)
		if !sliceEqual(actualSchedule, test.expectedCommits) {
			fmt := "Expected %v as schedule, but got %v"
			t.Errorf(fmt, test.expectedCommits, actualSchedule)
		}
	}
}

func TestBuildLastWeek(t *testing.T) {
	var tests = []struct {
		day             time.Time
		expectedCommits []int
	}{
		{testSunday, []int{0, -1, -1, -1, -1, -1, -1}},
		{testMonday, []int{0, 0, -1, -1, -1, -1, -1}},
		{testTuesday, []int{0, 0, 0, -1, -1, -1, -1}},
		{testWednesday, []int{0, 0, 0, 0, -1, -1, -1}},
		{testThursday, []int{0, 0, 0, 0, 0, -1, -1}},
		{testFriday, []int{0, 0, 0, 0, 0, 0, -1}},
		{testSaturday, []int{0, 0, 0, 0, 0, 0, 0}},
	}
	for _, test := range tests {
		actualSchedule := buildLastWeek(test.day)
		if !sliceEqual(actualSchedule, test.expectedCommits) {
			fmt := "Expected %v as schedule, but got %v"
			t.Errorf(fmt, test.expectedCommits, actualSchedule)
		}
	}
}

func sliceEqual(scheduleEntries []ScheduleEntry, numCommits []int) bool {
	if len(scheduleEntries) != len(numCommits) {
		return false
	}
	for i := 0; i < len(scheduleEntries); i++ {
		if scheduleEntries[i].NumCommits != numCommits[i] {
			return false
		}
	}
	return true
}

func TestGetFirstDayOfWeek(t *testing.T) {
	var tests = []struct {
		day                  time.Time
		expectedFirstWeekDay time.Time
	}{
		{time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC), time.Date(2009, time.November, 8, 0, 0, 0, 0, time.UTC)},
		{time.Date(2016, time.February, 28, 0, 0, 0, 0, time.UTC), time.Date(2016, time.February, 28, 0, 0, 0, 0, time.UTC)},
		{time.Date(2013, time.February, 28, 0, 0, 0, 0, time.UTC), time.Date(2013, time.February, 24, 0, 0, 0, 0, time.UTC)},
	}
	for _, test := range tests {
		actual := getFirstDayOfWeek(test.day)
		if !actual.Equal(test.expectedFirstWeekDay) {
			t.Errorf("Expected first weekday to be %v, but got %v", actual, test.expectedFirstWeekDay)
		}
	}
}

func TestConnectWeeksToSchedule(t *testing.T) {
	var tests = []struct {
		firstDay            time.Time
		lastDay             time.Time
		numEntries          int
		numNotAFieldEntries int
	}{
		{testSunday, testSaturday, 371, 0},
		{testMonday, testSaturday, 371, 1},
		{testWednesday, testWednesday, 371, 6},
		{testSaturday, testSunday, 371, 12},
	}
	for _, test := range tests {
		firstWeek := buildFirstWeek(test.firstDay)
		lastWeek := buildLastWeek(test.lastDay)
		schedule := connectWeeksToSchedule(firstWeek, lastWeek)
		length := 0
		numNotAFieldEntries := 0
		for _, row := range schedule {
			for _, entry := range row {
				if entry.NumCommits != EMPTY && entry != NOT_A_FIELD {
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
