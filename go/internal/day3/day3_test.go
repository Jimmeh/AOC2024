package day3_test

import (
	"testing"

	"github.com/Jimmeh/AOC2024/go/internal/day3"
	"github.com/Jimmeh/AOC2024/go/test/helpers"
)

func TestSingleInstruction(t *testing.T) {
	Case(t, "Single instruction with single digit operands", "mul(1,1)", 1)
}

func Case(t *testing.T, description string, input string, expected int) {
	reader := helpers.TestTextReader{Data: input}
	calculator, _ := day3.CreateInstructionCalculator(reader)
	actual := calculator.Calculate()
	if expected != actual {
		t.Errorf("%s: expected %d, got %d", description, expected, actual)
	}
}
