package day2

import "github.com/Jimmeh/AOC2024/go/internal/file_reading"

func CreateReportSafetyChecker(reader file_reading.LineReader) reportSafetyChecker {
	return reportSafetyChecker{reader}
}

type reportSafetyChecker struct {
	reader file_reading.LineReader
}

func (c reportSafetyChecker) TotalSafeReports() int {
	return 1
}
