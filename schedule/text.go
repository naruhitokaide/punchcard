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
	// TODO
	// check if text can be put into a commit schedule (width < 52)
	if !textFits(text) {
		return nil, errors.New("Text does not fit.")
	}
	// concatenate letters with one column as space between letters
	// put the result into a commit schedule
	// translate the commit schedule to a []Commit
	return nil, nil
}

func textFits(text string) bool {
	textWidth := getTextWidth(text)
	textIsNotToWide := textWidth <= SCHEDULE_WIDTH-2 // adjust for margins
	textIsNotEmpty := textWidth > 0
	return textIsNotEmpty && textIsNotToWide
}

func getTextWidth(text string) int {
	// TODO translate into commit letters or some how get width
	return len(text)
}
