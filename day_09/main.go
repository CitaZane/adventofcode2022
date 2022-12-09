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
	movements := registerMovements(input)

	// solve1 := countVisitedPlaces(movements)
	solve2 := countVisitedPlaces(movements, 10)

	return nil, solve2
}

func registerMovements(input string) []Movement {
	lines := strings.Split(input, "\n")
	movements := []Movement{}
	for _, line := range lines {
		data := strings.Split(line, " ")
		num, _ := strconv.Atoi(data[1])
		move := Movement{Direction: data[0], Count: num}
		movements = append(movements, move)
	}
	return movements
}

func countVisitedPlaces(movements []Movement, ropeLength int) int {
	simulator := NewSimulator(movements, ropeLength)
	simulator.ExecuteMovements()

	return len(simulator.PlacesVisited)
}
