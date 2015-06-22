package schedule

import (
	"testing"
	"time"
)

func TestIsLeapYear(t *testing.T) {
	var tests = []struct {
		date     time.Time
		expected bool
	}{
		{time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC), false},
		{time.Date(2016, time.February, 28, 0, 0, 0, 0, time.UTC), false},
		{time.Date(2016, time.February, 29, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.UTC), true},
	}
	for _, test := range tests {
		actual := isLeapDay(test.date)
		if actual != test.expected {
			fmt := "isLeapYear(%v) == %b; but wanted %b"
			t.Errorf(fmt, test.date, actual, test.expected)
		}
	}
}

func TestGetDateLastYear(t *testing.T) {
	var tests = []struct {
		date     time.Time
		expected time.Time
	}{
		{time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC),
			time.Date(2008, time.November, 10, 0, 0, 0, 0, time.UTC)},
		{time.Date(2016, time.February, 28, 0, 0, 0, 0, time.UTC),
			time.Date(2015, time.February, 28, 0, 0, 0, 0, time.UTC)},
		{time.Date(2016, time.February, 29, 0, 0, 0, 0, time.UTC),
			time.Date(2015, time.February, 28, 0, 0, 0, 0, time.UTC)},
		{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.UTC),
			time.Date(2011, time.February, 28, 0, 0, 0, 0, time.UTC)},
	}
	for _, test := range tests {
		actual := getDayMinusOneYear(test.date)
		if actual != test.expected {
			fmt := "getDayLastYear(%v) == %v; but wanted %v"
			t.Errorf(fmt, test.date, actual, test.expected)
		}
	}
}

func TestGetRandomNumber(t *testing.T) {
	var tests = []struct {
		min int
		max int
	}{
		{0, 0},
		{0, 1},
		{1, 1},
		{0, 10},
		{10, 20},
	}
	for _, test := range tests {
		actual := getRandomNumber(test.min, test.max)
		if test.min > actual || actual > test.max {
			fmt := "getRandomNumber(%d, %d) == %d; not min <= actual <= max"
			t.Errorf(fmt, test.min, test.max, actual)
		}
	}
}

func TestGetDaysSinceDateMinusOneYear(t *testing.T) {
	var tests = []struct {
		date           time.Time
		expectedLength int
	}{
		{time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC), 365},
		{time.Date(2016, time.February, 28, 0, 0, 0, 0, time.UTC), 365},
		{time.Date(2016, time.February, 29, 0, 0, 0, 0, time.UTC), 366},
		{time.Date(2012, time.February, 29, 0, 0, 0, 0, time.UTC), 366},
	}
	for _, test := range tests {
		// TODO how to get the size of a channel
		actual := getDaysSinceDateMinusOneYear(test.date)
		if len(actual) == test.expectedLength {
			fmt := "getDaysSinceDateMinusOneYear(%v) == %d; but wanted %d"
			t.Errorf(fmt, test.date, len(actual), test.expectedLength)
		}
	}
}
