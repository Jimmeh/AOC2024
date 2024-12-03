package day3_test

import (
	"testing"

	"github.com/Jimmeh/AOC2024/go/internal/day3"
	"github.com/Jimmeh/AOC2024/go/test/helpers"
)

func TestInstructions(t *testing.T) {
	Case(t, "Single instruction with 1x1=1", "mul(1,1)", 1)
	Case(t, "No instruction", "", 0)
	Case(t, "No instruction but a lot of chars", "sadfjhasdjfhakjbsdaihdf3829uohanvcval;dsfjw9u3ru09gtsv", 0)
	Case(t, "Single instruction with 1x2=2", "mul(1,2)", 2)
	Case(t, "Nearly an instruction (negative number)", "mul(-1,2)", 0)
	Case(t, "Nearly an instruction (spaces)", "mul(1 ,2)", 0)
	Case(t, "Instruction in the middle of stuff", "afhijiakfd adaslfmul(1,2)asdfijoas", 2)
	Case(t, "Two instructions 1x2 + 1x2 = 4", "mul(1,2)mul(1,2)", 4)
	Case(t, "Single instruction with 10x10=100", "mul(10,10)", 100)
	Case(t, "Single instruction with 100x100=10000", "mul(100,100)", 10000)
	Case(t, "Instruction can't have more than 3 digit numbers", "mul(1000,100)", 0)
	Case(t, "Example from problem", "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161)
}

func TestWithConditions(t *testing.T) {
	CaseWithConditions(t, "Example from problem", "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48)
}

func CaseWithConditions(t *testing.T, description string, input string, expected int) {
	reader := helpers.TestTextReader{Data: input}
	calculator, _ := day3.CreateInstructionCalculator(reader)
	actual := calculator.CalculateWithConditions()
	if expected != actual {
		t.Errorf("%s: expected %d, got %d", description, expected, actual)
	}
}

func Case(t *testing.T, description string, input string, expected int) {
	reader := helpers.TestTextReader{Data: input}
	calculator, _ := day3.CreateInstructionCalculator(reader)
	actual := calculator.Calculate()
	if expected != actual {
		t.Errorf("%s: expected %d, got %d", description, expected, actual)
	}
}
