package schedule

import (
	"github.com/0xfoo/punchcard/git"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// RandomSchedule creates random commits over the past 365/366 days.
// These commits will be created in the location specified in the command.
func RandomSchedule(min, max int, location string) {
	git.Init(location)
	days := getDaysSinceDateMinusOneYear(time.Now())
	for day := range days {
		rnd := getRandomNumber(min, max)
		commits := RandomCommits(day, rnd)
		for commit := range commits {
			// TODO git add and commit file with given
			filename := "filename"
			git.Add(location, filename)
			git.Commit(location, commit.message, commit.dateTime.String())
		}
	}
}

// RandomCommits returns a channel of random commits for a given day.
func RandomCommits(day time.Time, rnd int) chan Commit {
	commitChannel := make(chan Commit)
	go func() {
		for i := 0; i < rnd; i++ {
			commitChannel <- Commit{
				dateTime: getRandomTime(day),
				message:  getRandomCommitMessage(),
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

func getRandomCommitMessage() string {
	content, _ := ioutil.ReadFile(COMMIT_MESSAGE_BASE)
	words := strings.Split(string(content), " ")
	return getRandomWords(words, getRandomNumber(4, 8))
}

// getRandomNumber returns a number in the range of min and max.
func getRandomNumber(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

func getRandomWords(words []string, numberOfWords int) string {
	rndWords := make([]string, numberOfWords)
	for i := 0; i < numberOfWords; i++ {
		rndWords = append(rndWords, words(getRandomNumber(0, len(words))))
	}
	return strings.Join(rndWords, " ")
}
