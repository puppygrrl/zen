package game

import (
	"fmt"
	"testing"
)

func TestGuess(t *testing.T) {
	var tests = []struct {
		word, secret string
		want string
	}{
		// Character position match == 1
		// Else, character in string: 2
		// Else, 0
		{"ABCDE", "ABCDE", ""},
		{"ABCDE", "ABCDF", "11110"},
		{"BRAIN", "BOATS", "10100"},
		{"BLOAT", "BOATS", "10222"},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%s,%s", tt.word, tt.secret)
		t.Run(testname, func(t *testing.T) {
			ans := Guess(tt.word, tt.secret)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
