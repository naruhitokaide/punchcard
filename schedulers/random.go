package schedulers

import (
	"fmt"
)

func RandomSchedule() {
	fmt.Println("Start random scheduler")
	// representation of the past 365 days
	// for each day
	// get random value between --min and --max
	// save into structure representing the commits over the last year
	// start worker, which will execute all commits using some sort of
	// commit generator
}
