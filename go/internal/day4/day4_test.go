package day4_test

import (
	"testing"

	"github.com/Jimmeh/AOC2024/go/internal/day4"
	"github.com/Jimmeh/AOC2024/go/test/helpers"
)

func TestHorizontalMatches(t *testing.T) {
	reader := helpers.TestLineReader{Data: []string{"ABCDE"}}
	searcher, _ := day4.CreateWordSearcher(reader)
	actual := searcher.Find("XMAS")
	if actual != 0 {
		t.Errorf("%s: expected %d, got %d", "Horizontal: No matches", 0, actual)
	}
}
