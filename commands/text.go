package commands

import (
	"log"

	"github.com/rtzll/punchcard/git"
	"github.com/rtzll/punchcard/schedule"
	"github.com/rtzll/punchcard/utils"
	"github.com/spf13/cobra"
)

const defaultText = "HELLO"

var text string

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Text will build the given text using commits.",
	Long: `Text will create commits to display the given text on the github
contribution graph. An error message will be shown in case the text will not
fit in the graphs width.`,
	Run: textRun,
}

// textRun creates repo and file generator based on the given location and
// starts the TextSchedule to create the text in commits.
func textRun(cmd *cobra.Command, args []string) {
	repo := git.Repo{Location}
	filegen := utils.RandomFileGenerator{Location}
	repo.Init()
	err := schedule.TextSchedule(text, repo, filegen)
	if err != nil {
		log.Fatal(err)
	}
}

// init initializes flags with default text and add textCmd to main cmd.
func init() {
	textCmd.Flags().StringVar(&text, "text", defaultText,
		"text which will build using commits")
	PunchCardCmd.AddCommand(textCmd)
}
