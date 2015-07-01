package schedule

type ScheduleEntries int

type CommitSchedule [][]int

const (
	NOT_A_FIELD ScheduleEntries = -1
	EMPTY       ScheduleEntries = 0
	ONE         ScheduleEntries = 1
	TWO         ScheduleEntries = 2
	THREE       ScheduleEntries = 3
	FOUR        ScheduleEntries = 4
)

func BuildCommitSchedule(days []time.Time) CommitSchedule {
	// get weeks, which determine width and height is seven
	// fill entries with EMPTY or NOT_A_FIELD
	return nil
}
