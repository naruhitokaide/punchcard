package schedule

import (
	"math/rand"
	"strings"
	"time"
)

const (
	COMMIT_MESSAGE_BASE       = `Non eram nescius, Brute, cum, quae summis ingeniis exquisitaque doctrina philosophi Graeco sermone tractavissent, ea Latinis litteris mandaremus, fore ut hic noster labor in varias reprehensiones incurreret. nam quibusdam, et iis quidem non admodum indoctis, totum hoc displicet philosophari. quidam autem non tam id reprehendunt, si remissius agatur, sed tantum studium tamque multam operam ponendam in eo non arbitrantur. erunt etiam, et ii quidem eruditi Graecis litteris, contemnentes Latinas, qui se dicant in Graecis legendis operam malle consumere. postremo aliquos futuros suspicor, qui me ad alias litteras vocent, genus hoc scribendi, etsi sit elegans, personae tamen et dignitatis esse negent.`
	BASE_SEPARATOR            = " "
	MIN_COMMIT_MESSAGE_LENGTH = 1
	MAX_COMMIT_MESSAGE_LENGTH = 8
)

type Commit struct {
	DateTime time.Time
	Message  string
}

// GenerateRandomCommits returns a channel of random commits for a given day.
// These commits are a random selection of numCommits number of words from
// the given message base.
func GenerateRandomCommits(day time.Time, numCommits int, messageBase []string) <-chan Commit {
	commitChannel := make(chan Commit)
	go func() {
		for i := 0; i < numCommits; i++ {
			commitChannel <- Commit{
				DateTime: getRandomTime(day),
				Message:  getRandomCommitMessage(messageBase, MAX_COMMIT_MESSAGE_LENGTH),
			}
		}
		close(commitChannel)
	}()
	return commitChannel
}

// GetRandomNumber returns a number in the range of min and max.
func GetRandomNumber(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

func GetCommitMessageBase() []string {
	return strings.Split(string(COMMIT_MESSAGE_BASE), BASE_SEPARATOR)
}

// getRandomTime sets a random time on the given date.
func getRandomTime(day time.Time) time.Time {
	hours := time.Duration(GetRandomNumber(0, 23)) * time.Hour
	minutes := time.Duration(GetRandomNumber(0, 59)) * time.Minute
	seconds := time.Duration(GetRandomNumber(0, 59)) * time.Second
	return day.Add(hours + minutes + seconds)
}

// getRandomCommitMessage returns a commit message, no longer than length.
func getRandomCommitMessage(messageBase []string, length int) string {
	commitMessageLength := GetRandomNumber(MIN_COMMIT_MESSAGE_LENGTH, length)
	return getRandomWords(messageBase, commitMessageLength)
}

// getRandomWords returns numWords random elements of the input.
func getRandomWords(inWords []string, numWords int) string {
	outWords := make([]string, numWords)
	for i := 0; i < numWords; i++ {
		randomIndex := GetRandomNumber(0, len(inWords)-1)
		outWords = append(outWords, inWords[randomIndex])
	}
	return strings.TrimSpace(strings.Join(outWords, " "))
}
