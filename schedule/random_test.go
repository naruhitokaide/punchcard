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

func TestGetSplitFileContent(t *testing.T) {
	actual := getSplitFileContent(COMMIT_MESSAGE_BASE, " ")
	t.Log(actual)
	if len(actual) != 100 {
		t.Errorf("File has %d words, but got %d", 100, len(actual))
	}
}

func TestGetRandomCommitMessage(t *testing.T) {
	var tests = []struct {
		length int
	}{
		{1}, {2}, {4}, {8},
	}
	messageBase := getMessageBase()
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

func getMessageBase() []string {
	return getSplitFileContent(COMMIT_MESSAGE_BASE, BASE_SEPARATOR)
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

func TestRandomCommits(t *testing.T) {
	var tests = []struct {
		date       time.Time
		numCommits int
	}{
		{time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC), 1},
		{time.Date(2016, time.February, 28, 0, 0, 0, 0, time.UTC), 2},
		{time.Date(2016, time.February, 29, 0, 0, 0, 0, time.UTC), 8},
	}
	messageBase := getMessageBase()
	for _, test := range tests {
		actual := generateRandomCommits(test.date, test.numCommits, messageBase)
		commitCount := 0
		for commit := range actual {
			if commit.dateTime.Day() != test.date.Day() {
				t.Error("Commit has a wrong date")
			}
			commitCount++
		}
		if commitCount != test.numCommits {
			t.Errorf("Expected %d commits, but got %d", test.numCommits, commitCount)
		}
	}
}

func TestRandomSchedule(t *testing.T) {
	var tests = []struct {
		min int
		max int
	}{
		{1, 1}, {2, 8}, {10, 100},
	}
	for _, test := range tests {
		git := &MockGit{}
		git.Init()
		filegen := MockFileGenerator{}
		RandomSchedule(test.min, test.max, git, filegen)
		if git.numInitCalls != 1 {
			t.Errorf("Expected one init call, but got %d", git.numInitCalls)
		}
		if git.numAddCalls != git.numCommitCalls {
			t.Error("Add calls should happen as often as commit calls.")
		}
		if test.min*366 > git.numCommitCalls || test.max*366 < git.numCommitCalls {
			fmt := "Total commits should be between %d and %d, but was %d"
			t.Errorf(fmt, test.min*366, test.max*366, git.numCommitCalls)
		}
	}
}

type MockFileGenerator struct{}

func (m MockFileGenerator) CreateFile() string {
	return ""
}

type MockGit struct {
	numInitCalls   int
	numAddCalls    int
	numCommitCalls int
}

func (m *MockGit) Init() {
	m.numInitCalls++
}

func (m *MockGit) Add(filename string) {
	m.numAddCalls++
}

func (m *MockGit) Commit(message, date string) {
	m.numCommitCalls++
}
