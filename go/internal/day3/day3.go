package day3

import "github.com/Jimmeh/AOC2024/go/internal/file_reading"

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
	return 1
}
