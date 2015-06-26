package schedule

import (
	"bufio"
	"os"
	"path/filepath"
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

func TestCreateFileInDir(t *testing.T) {
	testDir := "testDir"
	os.MkdirAll(testDir, 0755)
	filename := createFileInDir(testDir)
	if _, err := os.Stat(filepath.Join(testDir, filename)); os.IsNotExist(err) {
		t.Errorf("Expected file (%s) to be created", filename)
	}
	os.RemoveAll(testDir)
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
	for _, test := range tests {
		actual := RandomCommits(test.date, test.numCommits)
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
		min      int
		max      int
		location string
	}{
		{1, 1, "testWithOneCommit"},
		{2, 8, "testWithUpToEightCommits"},
		{10, 100, "testWithUpToOneHundredCommits"},
	}
	for _, test := range tests {
		currentDir, _ := os.Getwd()
		testDir := filepath.Join(currentDir, test.location)
		// t.Logf("test")
		RandomSchedule(test.min, test.max, test.location)
		numCommits := getNumberOfGitLogLines(testDir)
		if numCommits < test.min || test.max < test.max {
			fmt := "Number of commits: %d; should be in the range of %d to %d"
			t.Errorf(fmt, numCommits, test.min, test.max)
		}
		// os.RemoveAll(testDir)
	}
}

func getNumberOfGitLogLines(gitLocation string) int {
	log := filepath.Join(gitLocation, ".git", "logs", "refs", "heads", "master")
	logFile, _ := os.Open(log)
	defer logFile.Close()
	logScanner := bufio.NewScanner(logFile)
	lineCount := 0
	for logScanner.Scan() {
		lineCount++
	}
	return lineCount
}
