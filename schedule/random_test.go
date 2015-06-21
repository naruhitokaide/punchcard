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

func TestGetDayLastYear(t *testing.T) {
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
		actual := getDayLastYear(test.date)
		if actual != test.expected {
			fmt := "getDayLastYear(%v) == %v; but wanted %v"
			t.Errorf(fmt, test.date, actual, test.expected)
		}
	}
}

// TODO add test cases for getDaysSinceThisDayLastYear and getRandomNumber
