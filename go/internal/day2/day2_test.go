package day2_test

import (
	"testing"

	"github.com/Jimmeh/AOC2024/go/internal/day2"
	"github.com/Jimmeh/AOC2024/go/test/helpers"
)

func TestSingleLine(t *testing.T) {
	Case(t, "Single safe report, always incrementing by 1", []string{"1 2 3 4 5"}, 1)
	Case(t, "Single unsafe report, always 1", []string{"1 1 1 1 1"}, 0)
}

func Case(t *testing.T, description string, input []string, expected int) {
	reader := helpers.TestLineReader{Data: input}
	checker := day2.CreateReportSafetyChecker(reader)
	actual, _ := checker.TotalSafeReports()
	if actual != expected {
		t.Errorf("%s: expected %d, got %d", description, expected, actual)
	}
}
