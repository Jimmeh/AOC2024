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

func (c instructionCalculator) Calculate() int {
	regex, _ := regexp.Compile(`mul\(\d,\d\)`)
	matches := regex.FindAllString(c.data, -1)
	for _, match := range matches {
		startTrimmed := strings.Trim(match, "mul(")
		commaSeparatedOperands := strings.Trim(startTrimmed, ")")
		numbers := strings.Split(commaSeparatedOperands, ",")
		result := 1
		for _, num := range numbers {
			i, _ := strconv.Atoi(num)
			result *= i
		}
		return result
	}
	if len(matches) > 0 {
		return 1
	}
	return 0
}
