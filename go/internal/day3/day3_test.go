package day3_test

import (
	"testing"

	"github.com/Jimmeh/AOC2024/go/internal/day3"
	"github.com/Jimmeh/AOC2024/go/test/helpers"
)

func TestSingleInstruction(t *testing.T) {
	Case(t, "Single instruction with 1x1=1", "mul(1,1)", 1)
	Case(t, "No instruction", "", 0)
	Case(t, "No instruction but a lot of chars", "sadfjhasdjfhakjbsdaihdf3829uohanvcval;dsfjw9u3ru09gtsv", 0)
	Case(t, "Single isntruction with 1x2=2", "mul(1,2)", 2)
	Case(t, "Nearly an instruction", "mul(-1,2)", 0)
}

func Case(t *testing.T, description string, input string, expected int) {
	reader := helpers.TestTextReader{Data: input}
	calculator, _ := day3.CreateInstructionCalculator(reader)
	actual := calculator.Calculate()
	if expected != actual {
		t.Errorf("%s: expected %d, got %d", description, expected, actual)
	}
}
