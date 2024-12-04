package day4_test

import (
	"testing"

	"github.com/Jimmeh/AOC2024/go/internal/day4"
	"github.com/Jimmeh/AOC2024/go/test/helpers"
)

func TestHorizontalMatches(t *testing.T) {
	XmasMatchesCase(t, "Horizontal: No matches", []string{"ABCDE"}, 0)
}

func XmasMatchesCase(t *testing.T, descrption string, input []string, expected int) {
	reader := helpers.TestLineReader{Data: input}
	searcher, _ := day4.CreateWordSearcher(reader)
	actual := searcher.Find("XMAS")
	if actual != 0 {
		t.Errorf("%s: expected %d, got %d", descrption, expected, actual)
	}
}
