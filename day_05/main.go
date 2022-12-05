package main

import (
	"adventofcode2022/utils"
	"regexp"
	"strconv"
	"strings"
)

// --- Day 5: Supply Stacks ---
func main() {
	utils.Run(run, puzzle)
}

func run(input string) (interface{}, interface{}) {
	crates, instructions := readInput(input)

	// solve1 := reorderCrates(crates, instructions)
	solve2 := reorderCratesWithNewCrane(crates, instructions)

	return nil, solve2
}

func reorderCrates(crates [][]string, instructions [][]int) string {
	for _, instruction := range instructions {
		moveCrates(crates, instruction)
	}
	result := getTopCrates(crates)
	return result
}

func reorderCratesWithNewCrane(crates [][]string, instructions [][]int) string {
	for _, instruction := range instructions {
		moveCratesWithNewCrane(crates, instruction)
	}
	result := getTopCrates(crates)
	return result
}

func getTopCrates(crates [][]string) string {
	var result = ""
	for _, crate := range crates {
		result += crate[0]
	}
	return result
}

func moveCrates(crates [][]string, instruction []int) {
	from := instruction[1] - 1
	to := instruction[2] - 1
	for i := 0; i < instruction[0]; i++ {
		var move = crates[from][0]
		target := append([]string{move}, crates[to]...)
		source := crates[from][1:]
		crates[to] = target
		crates[from] = source
	}
}

func moveCratesWithNewCrane(crates [][]string, instruction []int) {
	from := instruction[1] - 1
	to := instruction[2] - 1
	crateCount := instruction[0]
	var move = crates[from][:crateCount]

	target := append([]string{}, move...)
	target = append(target, crates[to]...)
	source := crates[from][crateCount:]
	crates[to] = target
	crates[from] = source

}

func readInput(input string) ([][]string, [][]int) {
	lines := strings.Split(input, "\n")
	var stackCount = (len(lines[1]) + 1) / 4
	var crates = make([][]string, stackCount)
	var instructions = make([][]int, 0)
	var crateRegex = regexp.MustCompile(`^((\[([A-Z])\]\s?)|(\s{3})\s?)+$`)

	for _, line := range lines {
		if crateRegex.MatchString(line) { //crates
			readCrates(line, crates)
		} else if strings.Contains(line, "move") {
			instruction := readInstructions(line)
			instructions = append(instructions, instruction)
		}
	}
	return crates, instructions
}

func readCrates(input string, target [][]string) {
	reg := regexp.MustCompile(`(([A-Z])|(\s{4}))`)
	vals := reg.FindAllString(input, -1)
	for i, value := range vals {
		if len(value) != 4 {
			target[i] = append(target[i], value)
		}
	}
}

func readInstructions(input string) []int {
	var instructionRegex = regexp.MustCompile(`\d+`)
	vals := instructionRegex.FindAllString(input, -1)
	var instruction = []int{}

	for _, value := range vals {
		num, _ := strconv.Atoi(value)
		instruction = append(instruction, num)
	}
	return instruction
}
