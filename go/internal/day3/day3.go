package day3

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func CreateInstructionCalculator(reader file_reading.TextReader) (instructionCalculator, error) {
	data, err := reader.Text()
	if err != nil {
		return instructionCalculator{}, nil
	}
	return instructionCalculator{data}, nil
}

type instructionCalculator struct {
	data string
}

func (c instructionCalculator) CalculateWithConditions() int {
	splitByDont := strings.Split("do()"+c.data, "don't()")
	result := 0
	for _, section := range splitByDont {
		splitByDo := strings.SplitN(section, "do()", 2)
		if len(splitByDo) == 1 {
			continue
		}
		enabledPart := splitByDo[1]
		result += sumOfSection(enabledPart)
	}
	return result
}

func (c instructionCalculator) Calculate() int {
	return sumOfSection(c.data)
}

func sumOfSection(section string) int {
	regex, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := regex.FindAllString(section, -1)
	result := 0
	for _, match := range matches {
		a, b := getNumbers(match)
		result += a * b
	}
	return result
}

func getNumbers(instruction string) (int, int) {
	startTrimmed := strings.Trim(instruction, "mul(")
	commaSeparatedOperands := strings.Trim(startTrimmed, ")")
	numbers := strings.Split(commaSeparatedOperands, ",")
	a, _ := strconv.Atoi(numbers[0])
	b, _ := strconv.Atoi(numbers[1])
	return a, b
}
