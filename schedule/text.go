package schedule

import (
	"github.com/0xfoo/punchcard/git"
	"github.com/0xfoo/punchcard/utils"
	"time"
)

// TextSchedule creates commits over the past 365/366 days to build the given text.
// These commits will be created in the given git repo using the FileGenerator.
func TextSchedule(text string, repo git.Git, filegen utils.FileGenerator) error {
	messageBase := GetCommitMessageBase()
	days := GetDaysSinceNowMinusOneYear()
	commits, err := GetTextCommitSchedule(text, days, messageBase)
	if err != nil {
		return err
	}
	for _, commit := range commits {
		repo.Add(filegen.CreateFile())
		repo.Commit(commit.Message, commit.DateTime.String())
	}
	return err
}

func GetTextCommitSchedule(text string, days []time.Time, messageBase []string) ([]Commit, error) {
	// TODO
	// check if text can be put into commits
	// concatenate letters with one column as space between letters
	// map the result onto commits using the given days
	return nil, nil
}
