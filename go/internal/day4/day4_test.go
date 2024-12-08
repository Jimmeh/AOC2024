package day4_test

import (
	"testing"

	"github.com/Jimmeh/AOC2024/go/internal/day4"
	"github.com/Jimmeh/AOC2024/go/test/helpers"
)

func TestHorizontalMatches(t *testing.T) {
	MatchXmasCase(t, "Horizontal: No matches", []string{"ABCDE"}, 0)
	MatchXmasCase(t, "Horizontal: Exact match", []string{"XMAS"}, 1)
	MatchXmasCase(t, "Horizontal: Exact match, preceding X", []string{"XXMAS"}, 1)
	MatchXmasCase(t, "Horizontal: Exact match twice", []string{"XMASXMAS"}, 2)
	MatchXmasCase(t, "Horizontal: Match amongst non/near matches", []string{"XMAHFKAJFXMASXMSSUHGFS"}, 1)
	MatchXmasCase(t, "Horizontal: Exact match (backwards)", []string{"SAMX"}, 1)
	MatchXmasCase(t, "Horizontal: Forward and backward match", []string{"SAMXXMAS"}, 2)
}

func TestColumnMatches(t *testing.T) {
	MatchXmasCase(t, "Vertical Exact match", []string{"X", "M", "A", "S"}, 1)
	MatchXmasCase(t, "Vertical matches", []string{"XXX", "MMM", "AAA", "SSS"}, 3)
}

func TestDiagMatches(t *testing.T) {
	XmasMatchesCase(t, "Square grid exact match", []string{"XABC", "AMBC", "ABAC", "ABCS"}, 1)
	XmasMatchesCase(t, "Assymetric grid exact match right", []string{"AXABC", "AAMBC", "AABAC", "AABCS"}, 1)
	XmasMatchesCase(t, "Assymetric grid exact match left", []string{"AAAAA", "XABCD", "AMBCD", "ABACD", "ABCSD"}, 1)
}

func MatchXmasCase(t *testing.T, descrption string, input []string, expected int) {
	reader := helpers.TestLineReader{Data: input}
	searcher, _ := day4.CreateWordSearcher(reader)
	actual := searcher.FindWords()
	if actual != expected {
		t.Errorf("%s: expected %d, got %d", descrption, expected, actual)
	}
}

func XmasMatchesCase(t *testing.T, descrption string, input []string, expected int) {
	reader := helpers.TestLineReader{Data: input}
	searcher, _ := day4.CreateWordSearcher(reader)
	actual := searcher.Find("XMAS")
	if actual != expected {
		t.Errorf("%s: expected %d, got %d", descrption, expected, actual)
	}
}
