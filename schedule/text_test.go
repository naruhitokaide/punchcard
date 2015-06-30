package schedule

import (
	"testing"
)

func XTestTextSchedule(t *testing.T) {
	var tests = []struct {
		text        string
		numCommits  int
		isPrintable bool
	}{
		{"hello", 123, true}, // TODO replace with actual number of commits
		{"", 0, false},
		{"this is to long to print", 0, false},
	}
	for _, test := range tests {
		git := &MockGit{}
		git.Init()
		filegen := MockFileGenerator{}
		err := TextSchedule(test.text, git, filegen)

		if (err == nil) != test.isPrintable {
			fmt := "The message %s should have been isPrintable==%b"
			t.Errorf(fmt, test.text, test.isPrintable)
		}

		if git.numCommitCalls != test.numCommits {
			fmt := "The message %s should have had %d commits, but got %d."
			t.Errorf(fmt, test.text, test.numCommits, git.numCommitCalls)
		}
	}
}

func XTestGetTextCommitSchedule(t *testing.T) {
	var tests = []struct {
		text        string
		numCommits  int
		isPrintable bool
	}{
		{"hello", 123, true}, // TODO replace with actual number of commits
		{"", 0, false},
		{"this is to long to print", 0, false},
	}
	for _, test := range tests {
		days := GetDaysSinceNowMinusOneYear()
		messageBase := GetCommitMessageBase()
		commits, err := getTextCommitSchedule(test.text, days, messageBase)

		if (err == nil) != test.isPrintable {
			fmt := "The message %s should have been isPrintable==%b, but was %b"
			t.Errorf(fmt, test.text, test.isPrintable, (err == nil))
		}

		if len(commits) != test.numCommits {
			fmt := "The message %s should have had %d commits, but got %d."
			t.Errorf(fmt, test.text, test.numCommits, len(commits))
		}
	}
}
