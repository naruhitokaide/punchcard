package schedule

import (
	"github.com/0xfoo/punchcard/git"
	"github.com/0xfoo/punchcard/utils"
	"math/rand"
	"strings"
	"time"
)

const (
	MIN_COMMIT_MESSAGE_LENGTH = 1
	MAX_COMMIT_MESSAGE_LENGTH = 8
)

// RandomSchedule creates random commits over the past 365/366 days.
// These commits will be created in the given git repo using the FileGenerator.
func RandomSchedule(min, max int, repo git.Git, filegen utils.FileGenerator) {
	messageBase := strings.Split(string(COMMIT_MESSAGE_BASE), BASE_SEPARATOR)
	days := GetDaysSinceDateMinusOneYear(time.Now())
	for day := range days {
		numCommits := getRandomNumber(min, max)
		commits := generateRandomCommits(day, numCommits, messageBase)
		for commit := range commits {
			repo.Add(filegen.CreateFile())
			repo.Commit(commit.message, commit.dateTime.String())
		}
	}
}

// generateRandomCommits returns a channel of random commits for a given day.
// These commits are a random selection of numCommits number of words from
// the given message base.
func generateRandomCommits(day time.Time, numCommits int, messageBase []string) <-chan Commit {
	commitChannel := make(chan Commit)
	go func() {
		for i := 0; i < numCommits; i++ {
			commitChannel <- Commit{
				dateTime: getRandomTime(day),
				message:  getRandomCommitMessage(messageBase, MAX_COMMIT_MESSAGE_LENGTH),
			}
		}
		close(commitChannel)
	}()
	return commitChannel
}

// getRandomTime sets a random time on the given date.
func getRandomTime(day time.Time) time.Time {
	hours := time.Duration(getRandomNumber(0, 23)) * time.Hour
	minutes := time.Duration(getRandomNumber(0, 59)) * time.Minute
	seconds := time.Duration(getRandomNumber(0, 59)) * time.Second
	return day.Add(hours + minutes + seconds)
}

// getRandomCommitMessage returns a commit message, no longer than length.
func getRandomCommitMessage(messageBase []string, length int) string {
	commitMessageLength := getRandomNumber(MIN_COMMIT_MESSAGE_LENGTH, length)
	return getRandomWords(messageBase, commitMessageLength)
}

// getRandomNumber returns a number in the range of min and max.
func getRandomNumber(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

// getRandomWords returns numWords random elements of the input.
func getRandomWords(inWords []string, numWords int) string {
	outWords := make([]string, numWords)
	for i := 0; i < numWords; i++ {
		randomIndex := getRandomNumber(0, len(inWords)-1)
		outWords = append(outWords, inWords[randomIndex])
	}
	return strings.TrimSpace(strings.Join(outWords, " "))
}
