package file_reading

import (
	"bufio"
	"os"
)

type LineReader interface {
	Lines() ([]string, error)
}

type reader struct {
	path string
}

func NewReader(path string) reader {
	return reader{path}
}

func (r reader) Lines() ([]string, error) {
	file, err := os.Open(r.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
