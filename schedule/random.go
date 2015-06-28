package schedule

import (
	"github.com/0xfoo/punchcard/git"
	"github.com/0xfoo/punchcard/utils"
	"time"
)

// RandomSchedule creates random commits over the past 365/366 days.
// These commits will be created in the given git repo using the FileGenerator.
func RandomSchedule(min, max int, repo git.Git, filegen utils.FileGenerator) {
	messageBase := GetCommitMessageBase()
	days := GetDaysSinceDateMinusOneYear(time.Now())
	for day := range days {
		numCommits := GetRandomNumber(min, max)
		commits := GenerateRandomCommits(day, numCommits, messageBase)
		for commit := range commits {
			repo.Add(filegen.CreateFile())
			repo.Commit(commit.Message, commit.DateTime.String())
		}
	}
}
