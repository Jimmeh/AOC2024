package day4

import (
	"slices"

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

func (w wordSearcher) FindWords() int {
	count := 0
	for row := range w.grid {
		for col, letter := range w.grid[row] {
			if letter != 'X' {
				continue
			}
			required := []byte{'M', 'A', 'S'}
			count += w.MatchHorizontally(row, col, required)
			count += w.MatchVertically(row, col, required)
		}
	}
	return count
}

func (w wordSearcher) MatchHorizontally(row, col int, letters []byte) int {
	possible := 2
	for i, letter := range letters {
		if !w.LetterEquals(row, col+(i+1), letter) {
			possible--
			break
		}
	}
	for i, letter := range letters {
		if !w.LetterEquals(row, col-(i+1), letter) {
			possible--
			break
		}
	}
	return possible
}

func (w wordSearcher) MatchVertically(row, col int, letters []byte) int {
	possible := 2
	for i, letter := range letters {
		if !w.LetterEquals(row+(i+1), col, letter) {
			possible--
			break
		}
	}
	for i, letter := range letters {
		if !w.LetterEquals(row-(i+1), col, letter) {
			possible--
			break
		}
	}
	return possible
}

func (w wordSearcher) LetterEquals(i, j int, letter byte) bool {
	return i >= 0 &&
		i < len(w.grid) &&
		j >= 0 &&
		j < len(w.grid[i]) &&
		w.grid[i][j] == letter
}

func (w wordSearcher) Find(word string) int {
	result := 0
	for _, sequence := range getSequences(w.grid) {
		result += forwardMatches(sequence, word)
		result += backwardMatches(sequence, word)
	}
	return result
}

func backwardMatches(sequence []byte, word string) int {
	return forwardMatches(sequence, reverseString(word))
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
	cols := getColumns(grid)
	rightDiags := getDiagonals(grid)
	reversedGrid := [][]byte{}
	for _, row := range grid {
		reversedGrid = append(reversedGrid, reverseBytes(row))
	}
	leftDiags := getDiagonals(reversedGrid)
	return slices.Concat(rows, cols, rightDiags, leftDiags)
}

func getColumns(grid [][]byte) [][]byte {
	cols := [][]byte{}
	for i := range grid[0] {
		col := []byte{}
		for _, row := range grid {
			col = append(col, row[i])
		}
		cols = append(cols, col)
	}
	return cols
}

func getDiagonals(grid [][]byte) [][]byte {
	diags := [][]byte{}
	for i := range grid {
		diag := []byte{}
		for j := 0; j < len(grid)-i; j++ {
			if j >= len(grid[i+j]) {
				break
			}
			diag = append(diag, grid[i+j][j])
		}
		diags = append(diags, diag)
	}
	for i := range grid[0] {
		if i == 0 {
			continue
		}
		diag := []byte{}
		for j := 0; j < len(grid[0])-i; j++ {
			if j >= len(grid) {
				break
			}
			diag = append(diag, grid[j][i+j])
		}
		diags = append(diags, diag)
	}
	return diags
}

func reverseBytes(b []byte) []byte {
	result := []byte(b)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
