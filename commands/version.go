package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Punchcard",
	Long:  `Print the version number of Punchcard`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Punchcard Version " + VERSION)
	},
}

func init() {
	PunchCardCmd.AddCommand(versionCmd)
}
