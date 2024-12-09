package day5_test

import (
	"testing"

	"github.com/Jimmeh/AOC2024/go/internal/day5"
	"github.com/Jimmeh/AOC2024/go/test/helpers"
)

func TestIndividualRule(t *testing.T) {
	IndividualRuleCase(t, "Empty update passes rule", "1|2", []int{}, true)
	IndividualRuleCase(t, "In order update passes rule", "1|2", []int{1, 2}, true)
	IndividualRuleCase(t, "Out of order update fails rule", "1|2", []int{2, 1}, false)
	IndividualRuleCase(t, "Only left number in update passes rule", "1|2", []int{1}, true)
	IndividualRuleCase(t, "Only right number in update passes rule", "1|2", []int{2}, true)
	IndividualRuleCase(t, "Many other numbers with correct order", "1|2", []int{9, 1, 4, 5, 2, 3, 7}, true)
	IndividualRuleCase(t, "Many other numbers with incorrect order", "1|2", []int{9, 2, 4, 5, 1, 3, 7}, false)
}
func TestMultipleRules(t *testing.T) {
	MultiRuleCase(t, "Empty update passes rule", []string{"1|2", "", "3"}, 3)
	MultiRuleCase(t, "Empty update passes rule", []string{"1|2", "", "1,2,3"}, 2)
	MultiRuleCase(t, "Empty update passes rule", []string{"1|2", "3|4", "", "1,2,3,4,5,6,7"}, 4)
}

func IndividualRuleCase(t *testing.T, desc, input string, update []int, expected bool) {
	rule := day5.CreateRule(input)
	actual := rule.Passes(update)
	if actual != expected {
		t.Errorf("%s: expected: %t, got %t", desc, expected, actual)
	}
}

func MultiRuleCase(t *testing.T, desc string, lines []string, expected int) {
	reader := helpers.TestLineReader{Data: lines}
	checker := day5.CreateRuleChecker(reader)
	actual := checker.MiddleNumberSum()
	if actual != expected {
		t.Errorf("%s: expected: %d, got %d", desc, expected, actual)
	}
}
