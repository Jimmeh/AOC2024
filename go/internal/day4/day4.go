package day4

import (
	"fmt"

	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

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
		result += forwardMatches(sequence, word)
		result += backwardMatches(sequence, word)
		fmt.Printf(reverse(word))
		fmt.Printf(word)
		fmt.Println()
	}
	return result
}

func backwardMatches(sequence []byte, word string) int {
	return forwardMatches(sequence, reverse(word))
}

func forwardMatches(sequence []byte, word string) int {
	result := 0
	matched := 0
	for i, char := range sequence {
		if matched == 0 && len(word) > len(sequence)-i {
			break
		}
		if char == word[matched] {
			matched++
			if matched == len(word) {
				matched = 0
				result++
			}
		} else if char == word[0] {
			matched = 1
		} else {
			matched = 0
		}
	}
	return result
}

func getSequences(grid [][]byte) [][]byte {
	rows := grid
	return rows
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
