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
}

func Case(t *testing.T, description string, input []string, expected int) {
	reader := TestLineReader{input}
	locationChecker := day1.CreateLocationChecker(reader)
	actual, _ := locationChecker.TotalDistance()
	if actual != expected {
		t.Errorf("%s: expected %d, got %d", description, expected, actual)
	}
}
