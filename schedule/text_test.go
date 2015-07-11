package schedule

import (
	"math/rand"
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

func TestTextFits(t *testing.T) {
	var tests = []struct {
		text        string
		isPrintable bool
	}{
		{"hello", true},
		{"", false},
		{"this is to long to print", false},
	}
	for _, test := range tests {
		actual := textFits(test.text)
		if actual != test.isPrintable {
			t.Errorf("Expected check to be %v, but was %v", test.isPrintable, actual)
		}
	}
}

func TestGetTextWidth(t *testing.T) {
	var tests = []struct {
		text      string
		textWidth int
	}{
		{"hello", 24},
		{"", 0},
		{"this is to long to print", 96},
	}
	for _, test := range tests {
		actual := getTextWidth(test.text)
		if actual != test.textWidth {
			t.Errorf("Expected width to be %d, but was %d", test.textWidth, actual)
		}
	}
}

func TestConvertScheduleToCommits(t *testing.T) {
	var tests = []struct {
		numCommits int
	}{
		{0}, {1}, {365},
	}
	for _, test := range tests {
		days := GetDaysSinceNowMinusOneYear()
		schedule := BuildCommitSchedule(days)
		addCommitsToSchedule(&schedule, test.numCommits)
		commits := convertScheduleToCommits(schedule)
		actual := len(commits)
		if actual != test.numCommits {
			t.Errorf("Expected %d commits, but got %d", test.numCommits, actual)
		}
	}
}

func addCommitsToSchedule(schedule *CommitSchedule, numCommits int) {
	for i := 0; i < numCommits; i++ {
		randRow := rand.Intn(7)
		randCol := rand.Intn(53)
		schedule[randCol][randRow] += 1
	}
}

func TestBuildTextCommitSchedule(t *testing.T) {
	var tests = []struct {
		text       string
		numCommits int
	}{
		{"hello", 62},
		{"i", 6},
	}
	for _, test := range tests {
		days := GetDaysSinceNowMinusOneYear()
		commits := buildTextCommitSchedule(days, test.text)
		actual := len(commits)
		if actual != test.numCommits {
			t.Errorf("Expected width to be %d, but was %d", test.numCommits, actual)
		}
	}
}

func TestTranslateTextIntoArray(t *testing.T) {
	var tests = []struct {
		letter    string
		numPixels int
	}{
		{"a", 14},
		{"i", 6},
	}
	for _, test := range tests {
		letter := translateTextIntoArray(test.letter)
		sumEntries := 0
		for _, entry := range letter {
			sumEntries += entry
		}
		actual := sumEntries
		if actual != test.numPixels {
			t.Errorf("Expected width to be %d, but was %d", test.numPixels, actual)
		}
	}
}

func TestAddFieldsToSchedule(t *testing.T) {
	// TODO
}
