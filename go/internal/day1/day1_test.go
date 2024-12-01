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
}

func Case(t *testing.T, description string, input []string, expected int) {
	reader := TestLineReader{[]string{"1   1"}}
	locationChecker := day1.CreateLocationChecker(reader)
	actual := locationChecker.TotalDistance()
	if actual != expected {
		t.Errorf("%s: expected %d, got %d", description, expected, actual)
	}
}
