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
		for r, row := range schedule {
			for c, col := range row {
				if col != EMPTY {
					fmt := "Expected only EMPTY values, but got %d at (%d,%d)"
					t.Errorf(fmt, col, r, c)
				}
			}
		}
	}
}

func TestBuildCommitScheduleWednesdayStart(t *testing.T) {
	var tests = []struct {
		startDay time.Time
		numDays  int
	}{
		{time.Date(2009, time.November, 11, 0, 0, 0, 0, time.UTC), 9},
		{time.Date(2015, time.November, 11, 0, 0, 0, 0, time.UTC), 23},
		{time.Date(2014, time.July, 9, 0, 0, 0, 0, time.UTC), 367},
	}
	for _, test := range tests {
		days := getTestDays(test.startDay, test.numDays)
		schedule := BuildCommitSchedule(days)
		for r, row := range schedule {
			for c, col := range row {
				firstWeekMondayOrTuesDay := (c == 0 && r < 2)
				if firstWeekMondayOrTuesDay {
					if col != NOT_A_FIELD {
						fmt := "Expected NOT_A_FIELD values, but got %d at (%d,%d)"
						t.Errorf(fmt, col, r, c)
					}
				} else if col != EMPTY {
					fmt := "Expected only EMPTY values, but got %d at (%d,%d)"
					t.Errorf(fmt, col, r, c)
				}
			}
		}
	}
}

func TestBuildCommitScheduleThrusdayEnd(t *testing.T) {
	var tests = []struct {
		startDay time.Time
		numDays  int
	}{
		{time.Date(2009, time.November, 9, 0, 0, 0, 0, time.UTC), 4},
		{time.Date(2015, time.November, 9, 0, 0, 0, 0, time.UTC), 19},
		{time.Date(2014, time.July, 7, 0, 0, 0, 0, time.UTC), 362},
	}
	for _, test := range tests {
		days := getTestDays(test.startDay, test.numDays)
		schedule := BuildCommitSchedule(days)
		for r, row := range schedule {
			for c, col := range row {
				lastWeekFridayOrSaturdayOrSunday := (c == len(row) && r > 4)
				if lastWeekFridayOrSaturdayOrSunday {
					if col != NOT_A_FIELD {
						fmt := "Expected NOT_A_FIELD values, but got %d at (%d,%d)"
						t.Errorf(fmt, col, r, c)
					}
				} else if col != EMPTY {
					fmt := "Expected only EMPTY values, but got %d at (%d,%d)"
					t.Errorf(fmt, col, r, c)
				}
			}
		}
	}
}

func getTestDays(startDay time.Time, numDays int) []time.Time {
	var resultingDays []time.Time
	for i := 0; i < numDays; i++ {
		resultingDays = append(resultingDays, startDay.AddDate(0, 0, i))
	}
	return resultingDays
}
