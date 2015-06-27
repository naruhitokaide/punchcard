package schedule

import (
	"time"
)

const (
	COMMIT_MESSAGE_BASE = "commit_message_base.txt"
	BASE_SEPARATOR      = " "
)

type Commit struct {
	dateTime time.Time
	message  string
}
