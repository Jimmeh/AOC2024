package day4_test

import (
	"testing"

	"github.com/Jimmeh/AOC2024/go/internal/day4"
	"github.com/Jimmeh/AOC2024/go/test/helpers"
)

func TestHorizontalMatches(t *testing.T) {
	XmasMatchesCase(t, "Horizontal: No matches", []string{"ABCDE"}, 0)
	XmasMatchesCase(t, "Horizontal: Exact match", []string{"XMAS"}, 1)
	XmasMatchesCase(t, "Horizontal: Exact match twice", []string{"XMASXMAS"}, 2)
	XmasMatchesCase(t, "Horizontal: Match amongst non/near matches", []string{"XMAHFKAJFXMASXMSSUHGFS"}, 1)
	XmasMatchesCase(t, "Horizontal: Exact match (backwards)", []string{"SAMX"}, 1)
	XmasMatchesCase(t, "Horizontal: Forward and backward match", []string{"SAMXXMAS"}, 2)
}

func XmasMatchesCase(t *testing.T, descrption string, input []string, expected int) {
	reader := helpers.TestLineReader{Data: input}
	searcher, _ := day4.CreateWordSearcher(reader)
	actual := searcher.Find("XMAS")
	if actual != expected {
		t.Errorf("%s: expected %d, got %d", descrption, expected, actual)
	}
}
