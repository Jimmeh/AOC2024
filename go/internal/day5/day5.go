package day5

import (
	"strconv"
	"strings"
)

func CreateRule(input string) rule {
	split := strings.Split(input, "|")
	left, _ := strconv.Atoi(split[0])
	right, _ := strconv.Atoi(split[1])
	return rule{left, right}
}

type rule struct {
	first  int
	second int
}

func (r rule) Passes(update []int) bool {
	firstIndex := -1
	secondIndex := -1
	for i, page := range update {
		if page == r.first {
			firstIndex = i
		}
		if page == r.second {
			secondIndex = i
		}
		if firstIndex != -1 && secondIndex != -1 {
			break
		}
	}
	return firstIndex == -1 || secondIndex == -1 || firstIndex < secondIndex
}
