package main

import (
	"adventofcode2022/utils"
	"strconv"
	"strings"
)

// --- Day 4: Camp Cleanup ---
func main() {
	utils.Run(run, puzzle)
}

func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")

	solve1 := findFullyContainedRange(lines)
	solve2 := findOverlppingRanges(lines)

	return solve1, solve2
}

func findFullyContainedRange(input []string) int {
	pairCount := 0

	for _, pair := range input {
		pair1, pair2 := parsePair(pair)
		if pairContained := pairContained(pair1, pair2); pairContained {
			pairCount += 1
		}
	}
	return pairCount
}
func findOverlppingRanges(input []string) int {
	overlappCount := 0

	for _, pair := range input {
		pair1, pair2 := parsePair(pair)
		if pairOverlap := pairOverlap(pair1, pair2); pairOverlap {
			overlappCount += 1
		}
	}
	return overlappCount
}

func parsePair(pair string) ([]int, []int) {
	elements := strings.Split(pair, ",")
	pair1 := parseRange(elements[0])
	pair2 := parseRange(elements[1])
	return pair1, pair2
}

func pairOverlap(pair1, pair2 []int) bool {
	if pairContained(pair1, pair2){
		return true
	}
	if pair1[0] <=  pair2[0] && pair1[1] >= pair2[0]{
		return true
	}
	if pair1[0] <=  pair2[1] && pair1[1] >= pair2[1]{
		return true
	}
	return false
}

func pairContained(pair1, pair2 []int) bool {
	if pair1[0] <= pair2[0] && pair1[1] >= pair2[1] {
		return true
	} else if pair1[0] >= pair2[0] && pair1[1] <= pair2[1] {
		return true
	}
	return false
}

func parseRange(rawRange string) []int {
	result := []int{}
	rangeSplit := strings.Split(rawRange, "-")

	rangeStart, _ := strconv.Atoi(rangeSplit[0])
	rangeEnd, _ := strconv.Atoi(rangeSplit[1])
	result = append(result, rangeStart, rangeEnd)
	return result
}
