package day2_test

import (
	"testing"

	"github.com/Jimmeh/AOC2024/go/internal/day2"
	"github.com/Jimmeh/AOC2024/go/test/helpers"
)

func TestSingleLine(t *testing.T) {
	Case(t, "Single safe report, always increasing by 1", []string{"1 2 3 4 5"}, 1)
	Case(t, "Single safe report, always decreasing by 1", []string{"5 4 3 2 1"}, 1)
	Case(t, "Single unsafe report, always increasing with leap", []string{"1 2 10 11 12"}, 0)
	Case(t, "Single unsafe report, always decreasing with leap", []string{"12 11 10 2 1"}, 0)
	Case(t, "Single unsafe report, changes direction increase to decrease", []string{"1 2 3 2 1"}, 0)
	Case(t, "Single unsafe report, changes direction decrease to increase", []string{"3 2 1 2 3"}, 0)
	Case(t, "Single unsafe report, always 1", []string{"1 1 1 1 1"}, 0)
}

func TestMultipleLines(t *testing.T) {
	Case(t, "Multiple safe lines, always increasing", []string{"1 2 3 4 5", "6 7 8 9 10"}, 2)
	Case(t, "Multiple safe lines, always decreasing", []string{"10 9 8 7 6", "5 4 3 2 1"}, 2)
	Case(t, "Multiple lines, one safe, one not (increasing)", []string{"2 4 6 8 10", "1 5 7 9 10"}, 1)
	Case(t, "Multiple lines, one safe, one not (decreasing)", []string{"10 9 3 2 1", "10 8 6 4 2"}, 1)
	Case(t, "Multiple unsafe lines", []string{"10 9 3 2 1", "1 2 6 3 1"}, 0)
	Case(t, "Example from problem", []string{"7 6 4 2 1", "1 2 7 8 9", "9 7 6 2 1", "1 3 2 4 5", "8 6 4 4 1", "1 3 6 7 9"}, 2)
}

func Case(t *testing.T, description string, input []string, expected int) {
	reader := helpers.TestLineReader{Data: input}
	checker := day2.CreateReportSafetyChecker(reader)
	actual, _ := checker.TotalSafeReports()
	if actual != expected {
		t.Errorf("%s: expected %d, got %d", description, expected, actual)
	}
}
