package commands

import (
	"fmt"
	"github.com/0xfoo/punchcard/schedule"
	"github.com/spf13/cobra"
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

func randomRun(cmd *cobra.Command, args []string) {
	schedule.RandomSchedule(minCommits, maxCommits)
}

func init() {
	randomCmd.Flags().IntVar(&minCommits, "min", 1,
		"minimal #commits on a given day.")
	randomCmd.Flags().IntVar(&maxCommits, "max", 10,
		"maximal #commits on a given day.")
	PunchCardCmd.AddCommand(randomCmd)
}
