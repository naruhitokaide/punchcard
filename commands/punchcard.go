package commands

import (
	"github.com/spf13/cobra"
)

const (
	VERSION = "0.1.0"
)

var Location string

var PunchCardCmd = &cobra.Command{
	Use:   "punchcard",
	Short: "Punchcard is a fun tool to create fake git commits.",
	Long: `Punchcard can create fake git commits in a repo.
The larger purposes is to have fun with contribution graphs, punchcards etc.`,
	Run: nil,
}

func init() {
	PunchCardCmd.PersistentFlags().StringVar(&Location, "location", ".",
		"location where the git repo will be initialized")
}