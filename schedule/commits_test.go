package schedule

import (
	"strings"
	"testing"
	"time"
)

func TestGetRandomCommit(t *testing.T) {
	var tests = []struct{ day time.Time }{
		{time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC)},
		{time.Date(2016, time.February, 28, 0, 0, 0, 0, time.UTC)},
	}
	messageBase := GetCommitMessageBase()
	for _, test := range tests {
		actual := GetRandomCommit(test.day, messageBase)
		if test.day.Day() != actual.DateTime.Day() {
			fmt := "GetRandomCommit should be on %v, but was %v"
			t.Errorf(fmt, test.day, actual.DateTime)
		}
		if len(actual.Message) < 0 {
			t.Error("GetRandomCommit should return a commit with a non empty message.")
		}
	}
}

func TestGetRandomNumber(t *testing.T) {
	var tests = []struct {
		min int
		max int
	}{
		{0, 0}, {0, 1}, {1, 1}, {0, 10}, {10, 20},
	}
	for _, test := range tests {
		actual := GetRandomNumber(test.min, test.max)
		if test.min > actual || actual > test.max {
			fmt := "GetRandomNumber(%d, %d) == %d; not min <= actual <= max"
			t.Errorf(fmt, test.min, test.max, actual)
		}
	}
}

func TestGetRandomCommitMessage(t *testing.T) {
	var tests = []struct{ length int }{{1}, {2}, {4}, {8}}
	messageBase := GetCommitMessageBase()
	for _, test := range tests {
		actual := getRandomCommitMessage(messageBase, test.length)
		actualLength := len(strings.Split(actual, " "))
		t.Logf("Message: %s (length: %d)", actual, actualLength)
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
		t.Logf("Date: %v and date with random time: %v", test.date, actual)
		dayBefore := test.date.AddDate(0, 0, -1)
		dayAfter := test.date.AddDate(0, 0, 1)
		if actual.After(dayAfter) || actual.Before(dayBefore) {
			fmt := "getRandomTime(%v) == %v; not dayBefore < actual < dayAfter"
			t.Errorf(fmt, test.date, actual)
		}
	}
}

func TestGenerateRandomCommits(t *testing.T) {
	var tests = []struct {
		date       time.Time
		numCommits int
	}{
		{time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC), 1},
		{time.Date(2016, time.February, 28, 0, 0, 0, 0, time.UTC), 2},
		{time.Date(2016, time.February, 29, 0, 0, 0, 0, time.UTC), 8},
	}
	messageBase := GetCommitMessageBase()
	for _, test := range tests {
		actual := GenerateRandomCommits(test.date, test.numCommits, messageBase)
		commitCount := 0
		for commit := range actual {
			if commit.DateTime.Day() != test.date.Day() {
				t.Error("Commit has a wrong date")
			}
			commitCount++
		}
		if commitCount != test.numCommits {
			t.Errorf("Expected %d commits, but got %d", test.numCommits, commitCount)
		}
	}
}
