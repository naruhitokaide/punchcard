package schedule

import (
	"strings"
	"testing"
	"time"
)

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

func TestGetRandomCommitMessage(t *testing.T) {
	var tests = []struct {
		length int
	}{
		{1}, {2}, {4}, {8},
	}
	for _, test := range tests {
		actual := getRandomCommitMessage(test.length)
		actualLength := len(strings.Split(actual, " "))
		t.Log("Message: %s (length: %d)", actual, actualLength)
		if actualLength < 0 || test.length < actualLength {
			fmt := "getRandomCommitMessage(%d) == %s (length: %d)"
			t.Errorf(fmt, test.length, actual, actualLength)
		}
	}
}

func TestGetRandomTime(t *testing.T) {
	var tests = []struct {
		date time.Time
	}{
		{time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC)},
		{time.Date(2016, time.February, 28, 0, 0, 0, 0, time.UTC)},
	}
	for _, test := range tests {
		actual := getRandomTime(test.date)
		t.Log("Date: %v and date with random time: %v", test.date, actual)
		dayBefore := test.date.AddDate(0, 0, -1)
		dayAfter := test.date.AddDate(0, 0, 1)
		if actual.After(dayAfter) || actual.Before(dayBefore) {
			fmt := "getRandomTime(%v) == %v; not dayBefore < actual < dayAfter"
			t.Errorf(fmt, test.date, actual)
		}
	}
}
