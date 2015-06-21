package schedule

import (
	"time"
)

type Commit struct {
	dateTime time.Time
	message  string
}

type Commits []Commit
