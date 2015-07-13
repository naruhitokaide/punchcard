package schedule

import (
	"errors"
	"github.com/0xfoo/punchcard/git"
	"github.com/0xfoo/punchcard/utils"
	"strings"
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
	schedule := buildTextCommitSchedule(days, text)
	commits := convertScheduleToCommits(schedule)
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

func convertScheduleToCommits(schedule CommitSchedule) []Commit {
	return nil
}

func buildTextCommitSchedule(days []time.Time, text string) CommitSchedule {
	schedule := BuildCommitSchedule(days)
	mapTextOntoCommitSchedule(text, &schedule)
	return schedule
}

func mapTextOntoCommitSchedule(text string, schedule *CommitSchedule) {
	letters := buildTextFields(text)
	rightShift := 0
	for _, fields := range letters {
		for rowIndex, row := range fields {
			for columnIndex, field := range row {
				schedule[rowIndex][columnIndex+rightShift] = field
			}
			rightShift += len(row)
		}
	}
}

func buildTextFields(text string) [][][]int {
	var letters [][][]int
	space, _ := utils.TranslateLetter(" ")
	for _, char := range strings.Split(text, "") {
		letter, _ := utils.TranslateLetter(char)
		letters = append(letters, letter, space)
	}
	return letters
}
