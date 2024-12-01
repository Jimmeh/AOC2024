package day1_test

import (
	"testing"

	"github.com/Jimmeh/AOC2024/go/internal/day1"
)

type TestLineReader struct {
	lines []string
}

func (r TestLineReader) Lines() ([]string, error) {
	return r.lines, nil
}

func TestSingleEntry(t *testing.T) {
	Case(t, "Single matching entry", []string{"1   1"}, 0)
	Case(t, "Single entry, first smaller than second", []string{"1   2"}, 1)
	Case(t, "Single entry, first larger than second", []string{"2   1"}, 1)
	Case(t, "Single entry, first much smaller than second", []string{"1   50"}, 49)
	Case(t, "Single entry, first much larger than second", []string{"50   1"}, 49)
}

func TestMultipleEntriesInOrder(t *testing.T) {
	Case(t, "Multiple matching entries", []string{"1   1", "2   2"}, 0)
	Case(t, "Multiple entries, all 1 apart, first smaller than second", []string{"1   2", "2   3"}, 2)
	Case(t, "Multiple entries, all 1 apart, first larger than second", []string{"2   1", "3   2"}, 2)
	Case(t, "Multiple entries, all far apart, first smaller than second", []string{"1   41", "5   50"}, 85)
	Case(t, "Multiple entries, all far apart, first larger than second", []string{"50   11", "54   45"}, 48)
}

func TestMultipleEntriesOutOfOrder(t *testing.T) {
	Case(t, "Multiple matching enties out of order", []string{"2   2", "1   1", "3   3"}, 0)
	Case(t, "Multiple matching enties out of order but paired correctly", []string{"2   3", "1   2", "3   4"}, 3)
	Case(t, "Multiple matching enties out of order not paired correctly", []string{"2   6", "1   4", "3   2"}, 6)
	Case(t, "Example from problem", []string{"3   4", "4   3", "2   5", "1   3", "3   9", "3   3"}, 11)
}

func Case(t *testing.T, description string, input []string, expected int) {
	reader := TestLineReader{input}
	locationChecker, _ := day1.CreateLocationChecker(reader)
	actual := locationChecker.TotalDistance()
	if actual != expected {
		t.Errorf("%s: expected %d, got %d", description, expected, actual)
	}
}
