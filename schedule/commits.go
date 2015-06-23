package schedule

import (
	"time"
)

type Commit struct {
	dateTime time.Time
	message  string
}

// RandomCommits returns a channel of random commits for a given day.
func RandomCommits(day time.Time, rnd int) chan Commit {
	commitChannel := make(chan Commit)
	go func() {
		for i := 0; i < rnd; i++ {
			commitChannel <- Commit{
				dateTime: getRandomTime(), message: getRandomCommitMessage(),
			}
		}
		close(commitChannel)
	}()
	return commitChannel
}

func getRandomTime() time.Time {
	return time.Now()
}

func getRandomCommitMessage() string {
	return "not so random string"
}
