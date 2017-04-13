package schedule

import (
	"time"

	"github.com/rtzll/punchcard/git"
	"github.com/rtzll/punchcard/utils"
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
			// TODO handle error
			f, _ := filegen.CreateFile()
			repo.Add(f)
			repo.Commit(commit.Message, commit.DateTime.String())
		}
	}
}
