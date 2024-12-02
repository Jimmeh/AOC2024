package day2

import (
	"strconv"
	"strings"

	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func CreateReportSafetyChecker(reader file_reading.LineReader) reportSafetyChecker {
	return reportSafetyChecker{reader}
}

type reportSafetyChecker struct {
	reader file_reading.LineReader
}

func (c reportSafetyChecker) TotalSafeReports() (int, error) {
	lines, err := c.reader.Lines()
	if err != nil {
		return -1, err
	}
	report := []int{}
	for _, entry := range strings.Split(lines[0], " ") {
		num, _ := strconv.Atoi(entry)
		report = append(report, num)
	}

	for i, v := range report {
		if i == 0 {
			continue
		}
		if v <= report[i-1] {
			return 0, nil
		}

	}
	return 1, nil
}
