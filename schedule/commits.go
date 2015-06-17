package schedule

import (
	"time"
)

type Commit struct {
	dateTime time.Date
	message  string
}

type Commits []Commit
