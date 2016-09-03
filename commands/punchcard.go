package commands

import (
	"github.com/spf13/cobra"
)

const (
	// VERSION of this tool
	VERSION = "0.1.0"
	// DefaultLocation is the current working directory
	DefaultLocation = "."
)

// Location where the git repo will be created
var Location string

// PunchCardCmd is the command line tool setup
var PunchCardCmd = &cobra.Command{
	Use:   "punchcard",
	Short: "Punchcard is a fun tool to create fake git commits.",
	Long: `Punchcard can create fake git commits in a repo.
The larger purpose is to have fun with contribution graphs, punchcards etc.`,
	Run: nil,
}

func init() {
	PunchCardCmd.PersistentFlags().StringVar(&Location, "location", DefaultLocation,
		"location where the git repo will be initialized")
}
