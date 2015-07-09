package schedule

import (
	"errors"
	"github.com/0xfoo/punchcard/git"
	"github.com/0xfoo/punchcard/utils"
	"time"
)

const (
	SCHEDULE_WIDTH = 53
)

// TextSchedule creates commits over the past 365/366 days to build the given text.
// These commits will be created in the given git repo using the FileGenerator.
func TextSchedule(text string, repo git.Git, filegen utils.FileGenerator) error {
	messageBase := GetCommitMessageBase()
	days := GetDaysSinceNowMinusOneYear()
	commits, err := getTextCommitSchedule(text, days, messageBase)
	if err != nil {
		return err
	}
	for _, commit := range commits {
		repo.Add(filegen.CreateFile())
		repo.Commit(commit.Message, commit.DateTime.String())
	}
	return err
}

func getTextCommitSchedule(text string, days []time.Time, messageBase []string) ([]Commit, error) {
	if !textFits(text) {
		return nil, errors.New("Text does not fit.")
	}
	schedule := buildTextCommitSchedule(days, text) // BuildCommitSchedule(days)
	commits := convertScheduleToCommits(schedule, textField)
	return commits, nil
}

func textFits(text string) bool {
	textWidth := getTextWidth(text)
	textIsNotToWide := textWidth <= SCHEDULE_WIDTH-2 // adjust for margins
	textIsNotEmpty := textWidth > 0
	return textIsNotEmpty && textIsNotToWide
}

func getTextWidth(text string) int {
	// TODO get the width of each letter and add spacing
	return len(text)
}

func convertScheduleToCommits(schedule CommitSchedule, textField [][]int) []Commit {
	return nil
}

func buildTextCommitSchedule(days []time.Time, text string) CommitSchedule {
	schedule := BuildCommitSchedule(days)
	textFields := translateTextIntoXXX(text)
	addFieldsToSchedule(*schedule, textFields)
	return schedule
}

func translateTextIntoArray(text string) [][]int {
	// TODO concatenate letters with one column as space between letters
	return nil
}

func addFieldsToSchedule(schedule *CommitSchedule, fields [][]int) {
	for row_index, row := range fields {
		for column_index, field := range row {
			schedule[row_index][column_index] = field
		}
	}
}
