package schedule

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/rtzll/punchcard/git"
	"github.com/rtzll/punchcard/utils"
)

const (
	// ScheduleWidth is the width of the schedule
	ScheduleWidth = 53
	// TextRegex is the regular expression used to check the given text
	TextRegex = "[a-z ]{1,26}"
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
		// TODO error handling
		f, _ := filegen.CreateFile()
		repo.Add(f)
		repo.Commit(commit.Message, commit.DateTime.String())
	}
	return err
}

// getTextCommitSchedule returns a []Commit or an error if the given text will
// not fit onto the CommitSchedule.
func getTextCommitSchedule(text string, days []time.Time, messageBase []string) ([]Commit, error) {
	text = parseStringForTranslation(text)
	if err := checkText(text); err != nil {
		return nil, err
	}
	schedule := buildTextCommitSchedule(days, text)
	commits := convertScheduleToCommits(schedule)
	return commits, nil
}

// parseStringForTranslation returns a the text in lower case.
// also removing leading and trailing spaces.
func parseStringForTranslation(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}

// checkText checks wether or not the text will fit onto a CommitSchedule.
func checkText(text string) error {
	if matched, _ := regexp.MatchString(TextRegex, text); !matched {
		return errors.New("Text can only contain letters and spaces (not only spaces).")
	}
	textWidth := getTextWidth(text)
	textIsNotToWide := textWidth <= ScheduleWidth-2 // adjust for margins
	textIsNotEmpty := textWidth > 0
	if !(textIsNotEmpty && textIsNotToWide) {
		return errors.New("Text does not fit.")
	}
	return nil
}

// getTextWidth returns the width the text will need if put onto the CommitSchedule.
func getTextWidth(text string) int {
	width := 0
	for _, char := range strings.Split(text, "") {
		letter, _ := utils.TranslateLetter(char)
		width += len(letter[0]) + 1 // adjust for space between letters
	}
	return width - 1 // last letter does not need an extra space
}

// convertScheduleToCommits creates NumCommits commits for every entry.
func convertScheduleToCommits(schedule CommitSchedule) []Commit {
	var commits []Commit
	messageBase := GetCommitMessageBase()
	for _, row := range schedule {
		for _, entry := range row {
			for commit := range GenerateRandomCommits(entry.DateTime, entry.NumCommits, messageBase) {
				commits = append(commits, commit)

			}
		}
	}
	return commits
}

// buildTextCommitSchedule returns a CommitSchedule representing the given text.
func buildTextCommitSchedule(days []time.Time, text string) CommitSchedule {
	schedule := BuildCommitSchedule(days)
	mapTextOntoCommitSchedule(text, &schedule)
	return schedule
}

// mapTextOntoCommitSchedule will put text onto a CommitSchedule.
func mapTextOntoCommitSchedule(text string, schedule *CommitSchedule) {
	letters := buildTextFields(text)
	rightShift := 1 // adjust for left margin
	for _, fields := range letters {
		for rowIndex, row := range fields {
			for columnIndex, field := range row {
				schedule[rowIndex][columnIndex+rightShift].NumCommits = field
			}
		}
		rightShift += len(fields[0])
	}
}

// buildTextFields return [][][]int representation of the given text.
func buildTextFields(text string) [][][]int {
	var letters [][][]int
	if text == "" {
		return letters
	}
	space, _ := utils.TranslateLetter(" ")
	for _, char := range strings.Split(text, "") {
		letter, _ := utils.TranslateLetter(char)
		letters = append(letters, letter, space)
	}
	return letters[0 : len(letters)-1] // remove last extra space
}
