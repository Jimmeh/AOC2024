package day4

import "github.com/Jimmeh/AOC2024/go/internal/file_reading"

func CreateWordSearcher(reader file_reading.LineReader) (wordSearcher, error) {
	lines, err := reader.Lines()
	if err != nil {
		return wordSearcher{}, nil
	}
	asByteArrays := [][]byte{}
	for _, line := range lines {
		asByteArrays = append(asByteArrays, []byte(line))
	}
	return wordSearcher{asByteArrays}, nil
}

type wordSearcher struct {
	grid [][]byte
}

func (w wordSearcher) Find(word string) int {
	result := 0
	for _, sequence := range getSequences(w.grid) {
		matched := 0
		for i, char := range sequence {
			if matched == 0 && len(word) > len(sequence)-i {
				break
			}
			if char == word[matched] {
				matched++
			}
			if matched == len(word) {
				matched = 0
				result++
			}
		}
	}
	return result
}

func getSequences(grid [][]byte) [][]byte {
	return grid
}
