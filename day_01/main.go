package main

import (
	"adventofcode2022/utils"
	"sort"
	"strconv"
	"strings"
)

// --- Day 1: Calorie Counting ---

func main() {
	utils.Run(run, puzzle)
}

func run(input string) (interface{}, interface{}) {
	caloriesTotal := strings.Split(input, "\n")
	sortedCalorieList := readAndCount(caloriesTotal)

	solve1 := calcMaxCalories(sortedCalorieList)
	solve2 := calcTopThreeCalories(sortedCalorieList)

	return solve1, solve2
}

func readAndCount(caloriesTotal []string) []int {
	caloriesPerElf := []int{}
	accumulator := 0
	for _, line := range caloriesTotal {
		number, err := strconv.Atoi(line)
		if err != nil {
			caloriesPerElf = append(caloriesPerElf, accumulator)
			accumulator = 0
		} else {
			accumulator += number
		}
	}
	sort.Ints(caloriesPerElf)
	return caloriesPerElf
}

func calcTopThreeCalories(list []int) int {
	index := len(list) - 1
	total := 0
	for i := index; i > index-3; i-- {
		total += list[i]
	}
	return total
}

func calcMaxCalories(list []int) int {
	index := len(list) - 1
	return list[index]
}
