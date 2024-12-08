package day4

import (
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
			count += w.MatchDiagonally(row, col, required)
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

func (w wordSearcher) MatchDiagonally(row, col int, letters []byte) int {
	possible := 4
	for i, letter := range letters {
		if !w.LetterEquals(row+(i+1), col+(i+1), letter) {
			possible--
			break
		}
	}
	for i, letter := range letters {
		if !w.LetterEquals(row-(i+1), col-(i+1), letter) {
			possible--
			break
		}
	}
	for i, letter := range letters {
		if !w.LetterEquals(row-(i+1), col+(i+1), letter) {
			possible--
			break
		}
	}
	for i, letter := range letters {
		if !w.LetterEquals(row+(i+1), col-(i+1), letter) {
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
