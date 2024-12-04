package day4

import "github.com/Jimmeh/AOC2024/go/internal/file_reading"

func CreateWordSearcher(reader file_reading.LineReader) (wordSearcher, error) {
	return wordSearcher{[][]byte{}}, nil
}

type wordSearcher struct {
	grid [][]byte
}

func (w wordSearcher) Find(word string) int {
	return 0
}
