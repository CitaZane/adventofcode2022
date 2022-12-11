package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	utils.Run(run, puzzle)
}

func run(input string) (interface{}, interface{}) {
	signals := registerSignals(input)

	// solve1 := calcSignalStrenght(signals)
	solve2 := drawSprite(signals)

	return nil, solve2
}

func registerSignals(input string) []int {
	lines := strings.Split(input, "\n")
	signals := []int{}

	for _, line := range lines {
		data := strings.Split(line, " ")
		if len(data) == 1 {
			signals = append(signals, 0)
		} else {
			num, _ := strconv.Atoi(data[1])
			signals = append(signals, num)
		}
	}
	return signals
}

func calcSignalStrenght(signals []int) int {
	sum := 1
	signalStrength := 0
	cycle := 0
	signalIndex := 0
	firstIter := true

	for {
		if signalIndex == len(signals)-1 {
			break
		}
		cycle += 1
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			signalStrength += (cycle * sum)
		}
		if signals[signalIndex] != 0 {
			if firstIter {
				firstIter = false
			} else {
				sum += signals[signalIndex]
				firstIter = true
				signalIndex++
			}
		} else {
			signalIndex++
		}
	}
	return signalStrength
}

func drawSprite(signals []int) int {
	var screen = []string{}

	sum := 1
	cycle := 0
	signalIndex := 0
	firstIter := true

	for {
		if signalIndex == len(signals)-1 {
			break
		}
		cycle += 1

		if signals[signalIndex] != 0 {
			if firstIter {
				firstIter = false
			} else {
				sum += signals[signalIndex]
				firstIter = true
				signalIndex++
			}
		} else {
			signalIndex++
		}

		// draw
		if cycle%40 < sum+2 && cycle%40 > sum-2 {
			screen = append(screen, "#")
		} else {
			screen = append(screen, ".")
		}
	}

	for i, pixel := range screen {

		fmt.Print(pixel)
		if i%40 == 0 && i != 0 {
			fmt.Println()
		}
	}
	return 0
}
