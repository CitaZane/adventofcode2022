package main

import (
	"adventofcode2022/utils"
	"strconv"
	"strings"
)

func main() {
	utils.Run(run, puzzle)
}

func run(input string) (interface{}, interface{}) {
	forest := registerForest(input)

	solve1 := countVisibleTrees(forest)
	solve2 := findBestScenicScore(forest)

	return solve1, solve2
}

func countVisibleTrees(forest Forest) int {
	forest.CountVisibleTrees()

	return forest.VisibleTrees
}

func findBestScenicScore(forest Forest) int {
	forest.findTopScenicScore()

	return forest.ScenicScore
}

func registerForest(input string) Forest {
	lines := strings.Split(input, "\n")
	forest := newForest()

	for _, line := range lines {
		var row []int
		for _, tree := range line {
			num, _ := strconv.Atoi(string(tree))
			row = append(row, num)
		}
		forest.Map = append(forest.Map, row)
	}

	forest.Height = len(forest.Map)
	forest.Width = len(forest.Map[0])
	return forest
}
