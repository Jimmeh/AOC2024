package file_reading_test

import (
	"path/filepath"
	"testing"

	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func TestReadingLines(t *testing.T) {
	path, _ := filepath.Abs("../../test/data/file_with_5_lines.txt")
	reader := file_reading.NewReader(path)
	lines, err := reader.Lines()
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}
	if len(lines) != 5 {
		t.Errorf("Expected 5 lines, got %d", len(lines))
	}
}

func TestReadingFullText(t *testing.T) {
	path, _ := filepath.Abs("../../test/data/file_with_5_lines.txt")
	reader := file_reading.NewReader(path)
	text, err := reader.Text()
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}
	if text != "0\n1\n2\n3\n4" {
		t.Errorf("expected 5 lines of indices")
	}
}
