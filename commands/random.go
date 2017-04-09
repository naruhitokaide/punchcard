package commands

import (
	"github.com/rtzll/punchcard/git"
	"github.com/rtzll/punchcard/schedule"
	"github.com/rtzll/punchcard/utils"
	"github.com/spf13/cobra"
)

const (
	minCommitsDefault = 1
	maxCommitsDefault = 10
)

var minCommits, maxCommits int

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Random will add commits throughout the past 365 days.",
	Long: `Random will create a git repo at the given location and create
random commits, random meaning the number of commits per day.
This will be done for the past 365 days and the commits are in the range of
--min and --max commits.`,
	Run: randomRun,
}

// randomRun creates repo and file generator based on the given location and
// starts the RandomSchedule using these and min, max number of commits.
func randomRun(cmd *cobra.Command, args []string) {
	repo := git.Repo{Location}
	filegen := utils.RandomFileGenerator{Location}
	repo.Init()
	schedule.RandomSchedule(minCommits, maxCommits, repo, filegen)
}

// init initializes flags with defaults and add randomCmd to main cmd.
func init() {
	randomCmd.Flags().IntVar(&minCommits, "min", minCommitsDefault,
		"minimal #commits on a given day.")
	randomCmd.Flags().IntVar(&maxCommits, "max", maxCommitsDefault,
		"maximal #commits on a given day.")
	PunchCardCmd.AddCommand(randomCmd)
}
