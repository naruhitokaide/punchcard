package utils

import (
	"testing"
)

func TestTranslateLetter(t *testing.T) {
	var tests = []struct {
		letter string
		sum    int
	}{
		{"a", 14}, {"i", 6}, {" ", 0},
	}
	for _, test := range tests {
		actual := 0
		fields, err := TranslateLetter(test.letter)
		if err != nil {
			t.Error("Got err != nil, but should have been nil.")
		}
		for _, row := range fields {
			for _, entry := range row {
				actual += entry
			}
		}
		if actual != test.sum {
			t.Errorf("Expected %d, but got %d", test.sum, actual)
		}
	}
}
