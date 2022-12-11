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
	inspector := registerMonkeys(input)
	inspector.InspectRounds(10000)
	// solve1 := inspector.ActivityCount()
	// fmt.Println(inspector)
	solve2 := inspector.ActivityCount()

	return solve2, nil
}

func registerMonkeys(input string) Inspector {
	inspector := Inspector{Monkeys: []*Monkey{}}

	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 7 {
		monkey := Monkey{}
		monkey.Items = extractItems(lines[i+1])
		monkey.Test = extractNum(lines[i+3])
		monkey.PositiveThrow = extractNum(lines[i+4])
		monkey.NegativeThrow = extractNum(lines[i+5])
		monkey.Operation = extractfunction(lines[i+2])
		monkey.ActivityLevel = 0

		inspector.Monkeys = append(inspector.Monkeys, &monkey)
	}

	return inspector
}

func extractItems(input string) []int {
	data := strings.Split(input, " ")
	items := []int{}
	for i, item := range data {
		if i < 4 {
			continue
		}
		if strings.HasSuffix(item, ",") {
			item = item[:len(item)-1]
		}
		num, _ := strconv.Atoi(item)
		items = append(items, num)
	}
	return items
}

func extractNum(input string) int {
	data := strings.Split(input, " ")
	num, _ := strconv.Atoi(data[len(data)-1])
	return num
}

func extractfunction(input string) func(int) int {
	data := strings.Split(input, " ")
	inputNum := data[len(data)-1]
	operator := data[len(data)-2]
	num := extractNum(inputNum)
	if operator == "+" {
		return func(old int) int {
			return old + num
		}
	} else if operator == "*" {
		if num == 0 {
			return func(old int) int {
				return old * old
			}
		} else {
			return func(old int) int {
				return old * num
			}
		}

	}
	return func(old int) int {
		fmt.Println("Inavalid operator: ", old)
		return 0
	}
}
