package utils

import (
	"strconv"
	"strings"
)

func ParseIntList(input, seperator string) []int {
	lines := strings.Split(input, seperator)
	list := make([]int, len(lines))

	for index, line := range lines {
		number, err := strconv.Atoi(line)
		Check(err)
		list[index] = number
	}
	return list
}
