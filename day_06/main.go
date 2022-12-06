package main

import "adventofcode2022/utils"

// --- Day 6: Tuning Trouble ---
func main() {
	utils.Run(run, puzzle)
}

func run(input string) (interface{}, interface{}) {

	solve1 := findStartMarker(input)
	solve2 := findMessageMarker(input)

	return solve1, solve2
}

func findStartMarker(stream string) int {
	streamAsBytes := []byte(stream)

	for i := 0; i < len(streamAsBytes); i++ {
		var mapedValues = map[byte]bool{}
		for j := i; j < i+4; j++ {
			mapedValues[streamAsBytes[j]] = true
		}
		if len(mapedValues) == 4 {
			return i + 4
		}

	}
	return 0
}

func findMessageMarker(stream string) int {
	streamAsBytes := []byte(stream)

	for i := 0; i < len(streamAsBytes); i++ {
		var mapedValues = map[byte]bool{}
		for j := i; j < i+14; j++ {
			mapedValues[streamAsBytes[j]] = true
		}
		if len(mapedValues) == 14 {
			return i + 14
		}

	}
	return 0
}
