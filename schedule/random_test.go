package schedule

import (
	"testing"
)

func TestGetRandomNumber(t *testing.T) {
	var tests = []struct {
		min int
		max int
	}{
		{0, 0},
		{0, 1},
		{1, 1},
		{0, 10},
		{10, 20},
	}
	for _, test := range tests {
		actual := getRandomNumber(test.min, test.max)
		if test.min > actual || actual > test.max {
			fmt := "getRandomNumber(%d, %d) == %d; not min <= actual <= max"
			t.Errorf(fmt, test.min, test.max, actual)
		}
	}
}

// TODO test random commit message and random time
