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

type directionModifier struct {
	vertical, horizontal int
}

type wordSearcher struct {
	grid [][]byte
}

func (w wordSearcher) FindXmas() int {
	count := 0
	for row := range w.grid {
		for col, letter := range w.grid[row] {
			if letter != 'X' {
				continue
			}
			required := []byte{'M', 'A', 'S'}
			directions := []directionModifier{
				{0, 1},
				{0, -1},
				{1, 0},
				{-1, 0},
				{1, 1},
				{-1, -1},
				{1, -1},
				{-1, 1},
			}
			count += w.matchDirections(row, col, required, directions)
		}
	}
	return count
}

func (w wordSearcher) FindCrosses() int {
	count := 0
	for row := range w.grid {
		for col, letter := range w.grid[row] {
			if letter != 'A' {
				continue
			}
			if w.crossDiagonalsCorrect(row, col) {
				count++
			}
		}
	}
	return count
}

func (w wordSearcher) crossDiagonalsCorrect(i, j int) bool {
	leftDiag := (w.letterEquals(i+1, j+1, 'S') && w.letterEquals(i-1, j-1, 'M')) || (w.letterEquals(i+1, j+1, 'M') && w.letterEquals(i-1, j-1, 'S'))
	rightDiag := (w.letterEquals(i-1, j+1, 'S') && w.letterEquals(i+1, j-1, 'M')) || (w.letterEquals(i-1, j+1, 'M') && w.letterEquals(i+1, j-1, 'S'))
	return leftDiag && rightDiag
}

func (w wordSearcher) matchDirections(row, col int, letters []byte, directions []directionModifier) int {
	possible := len(directions)
	for _, dir := range directions {
		for i, letter := range letters {
			if !w.letterEquals(row+((i+1)*dir.vertical), col+((i+1)*dir.horizontal), letter) {
				possible--
				break
			}
		}
	}
	return possible
}

func (w wordSearcher) letterEquals(i, j int, letter byte) bool {
	return i >= 0 &&
		i < len(w.grid) &&
		j >= 0 &&
		j < len(w.grid[i]) &&
		w.grid[i][j] == letter
}
