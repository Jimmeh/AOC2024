package day3

import (
	"regexp"

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
	regex, _ := regexp.Compile(`mul\(1,1\)`)
	if regex.Match([]byte(c.data)) {
		return 1
	}
	return 0
}
