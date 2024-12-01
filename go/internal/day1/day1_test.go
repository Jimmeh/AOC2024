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

func TestMatchingList(t *testing.T) {
	reader := TestLineReader{[]string{"1   1"}}
	locationChecker := day1.CreateLocationChecker(reader)

	expected := 0
	actual := locationChecker.TotalDistance()

	if actual != expected {
		t.Errorf("Expected totalDistance of %d, got %d", expected, actual)
	}
}
