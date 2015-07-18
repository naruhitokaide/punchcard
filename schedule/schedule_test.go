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

func TestGetDaysSinceDateMinusOneYear(t *testing.T) {
	var tests = []struct {
		date           time.Time
		expectedLength int
	}{
		{time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC), 366},
		{time.Date(2016, time.February, 28, 0, 0, 0, 0, time.UTC), 366},
		{time.Date(2016, time.February, 29, 0, 0, 0, 0, time.UTC), 367},
		{time.Date(2013, time.February, 28, 0, 0, 0, 0, time.UTC), 367},
	}
	for _, test := range tests {
		actual := GetDaysSinceDateMinusOneYear(test.date)
		length := 0
		actualLastDate := test.date
		for day := range actual {
			actualLastDate = day
			length++
		}
		if length != test.expectedLength {
			fmt := "len(GetDaysSinceDateMinusOneYear(%v)) == %d; but wanted %d"
			t.Errorf(fmt, test.date, length, test.expectedLength)
		}
		if !test.date.Equal(actualLastDate) {
			fmt := "Last day should be equal to %v but was %v"
			t.Errorf(fmt, test.date, actualLastDate)
		}
	}
}

func TestGetDaysSinceNowMinusOneYear(t *testing.T) {

	actual := GetDaysSinceNowMinusOneYear()
	length := 0
	var actualLastDate time.Time
	for _, day := range actual {
		actualLastDate = day
		length++
	}
	if length != 366 && length != 367 {
		fmt := "len(GetDaysSinceNowMinusOneYear()) should be 366 or 367; but was %d"
		t.Errorf(fmt, length)
	}
	today := time.Now()
	if !assertEqualDates(today, actualLastDate) {
		fmt := "Last day should be equal to %v (now) but was %v"
		t.Errorf(fmt, today, actualLastDate)

	}
}

func assertEqualDates(dayOne, dayTwo time.Time) bool {
	return dayOne.Day() == dayTwo.Day() && dayOne.Month() == dayTwo.Month() &&
		dayOne.Year() == dayTwo.Year()
}
