package main

import (
	"adventofcode2022/utils"
	"fmt"
)

func main() {
	nums := utils.ParseIntList(puzzle, "\n")
	firstCount := firstMesurement(nums)
	fmt.Println("Result -> ", firstCount)
}

func firstMesurement(data []int) int {
	count := 0
	for i := 0; i < len(data)-1; i++ {
		if data[i+1] > data[i] {
			count++
		}
	}
	return count
}
