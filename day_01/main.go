package main

import (
	"adventofcode2022/utils"
)

func run(input string) (interface{}, interface{}) {
	nums := utils.ParseIntList(input, "\n")
	solve1 := firstMesurement(nums)

	return solve1, nil
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

func main() {
	utils.Run(run, puzzle)
}
