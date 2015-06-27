package schedule

import (
	"github.com/0xfoo/punchcard/git"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// RandomSchedule creates random commits over the past 365/366 days.
// These commits will be created in the location specified in the command.
func RandomSchedule(min, max int, location string) {
	git.Init(location)
	messageBase := getSplitFileContent(COMMIT_MESSAGE_BASE, BASE_SEPARATOR)
	days := GetDaysSinceDateMinusOneYear(time.Now())
	for day := range days {
		rnd := getRandomNumber(min, max)
		commits := RandomCommits(day, rnd, messageBase)
		for commit := range commits {
			filename := createFileInDir(location)
			git.Add(location, filename)
			git.Commit(location, commit.message, commit.dateTime.String())
		}
	}
}

// RandomCommits returns a channel of random commits for a given day.
func RandomCommits(day time.Time, numCommits int, messageBase []string) <-chan Commit {
	commitChannel := make(chan Commit)
	go func() {
		for i := 0; i < numCommits; i++ {
			commitChannel <- Commit{
				dateTime: getRandomTime(day),
				message:  getRandomCommitMessage(messageBase, 8),
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

// getRandomCommitMessage returns a commit message, no longer than length
func getRandomCommitMessage(messageBase []string, length int) string {
	return getRandomWords(messageBase, getRandomNumber(1, length))
}

// getSplitFileContent returns the content of a file (given by name) and
// split by a separator string
func getSplitFileContent(filename, sep string) []string {
	content, _ := ioutil.ReadFile(filename)
	return strings.Split(string(content), sep)
}

// getRandomNumber returns a number in the range of min and max.
func getRandomNumber(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

// getRandomWords returns numWords random elements of the input []string
func getRandomWords(inWords []string, numWords int) string {
	outWords := make([]string, numWords)
	for i := 0; i < numWords; i++ {
		outWords = append(outWords, inWords[getRandomNumber(0, len(inWords))])
	}
	return strings.TrimSpace(strings.Join(outWords, " "))
}

// createFileWithTimeStamp creates a file with the current nano seconds as the
// file name, and returns this time stamp (i.e. filename)
func createFileInDir(dir string) string {
	filename := strconv.Itoa(time.Now().Nanosecond())
	file, _ := os.Create(filepath.Join(dir, filename))
	file.Close()
	return filename
}
