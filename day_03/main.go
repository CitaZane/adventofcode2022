package main

import (
	"adventofcode2022/utils"
	"bytes"
	"strings"
)

// --- Day 3: Rucksack Reorganization ---

func main() {
	utils.Run(run, puzzle)
}

func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")

	solve1 := sumOfMissingItems(lines)
	solve2 := sumOfBadgeValues(lines)

	return solve1, solve2
}

func sumOfMissingItems(itemList []string) int {
	sum := 0

	for _, items := range itemList {
		var half = len(items) / 2
		var split = []byte(items)
		var itemsSplit = bytes.Runes(split)

		pocket1 := putItemsInPocket(itemsSplit[:half])
		pocket2 := putItemsInPocket(itemsSplit[half:])

		dublicatedItem := findDublicate(pocket1, pocket2)
		sum += itemValue(dublicatedItem)
	}
	return sum
}

func sumOfBadgeValues(itemList []string) int {
	sum := 0

	for i := 0; i < len(itemList); i += 3 {
		itemsRunes1 := bytes.Runes([]byte(itemList[i]))
		pocket1 := putItemsInPocket(itemsRunes1)

		itemsRunes2 := bytes.Runes([]byte(itemList[i+1]))
		pocket2 := putMoreItemsIn(pocket1, itemsRunes2, 2)

		itemsRunes3 := bytes.Runes([]byte(itemList[i+2]))
		pocket3 := putMoreItemsIn(pocket2, itemsRunes3, 3)

		badge := findBadge(pocket3)
		sum += itemValue(badge)
	}
	return sum
}

func putItemsInPocket(items []rune) map[rune]int {
	var itemsCounted = map[rune]int{}

	for _, item := range items {
		itemsCounted[item] = 1
	}
	return itemsCounted
}

func putMoreItemsIn(target map[rune]int, source []rune, value int) map[rune]int {
	for _, item := range source {
		_, ok := target[item]
		if ok {
			target[item] = value
		}
	}
	for item, count := range target {
		if count != value {
			delete(target, item)
		}
	}
	return target
}
func findBadge(target map[rune]int) rune {
	for item, value := range target {
		if value == 3 {
			return item
		}
	}
	return 'a'
}

func findDublicate(pocket1, pocket2 map[rune]int) rune {
	for item := range pocket2 {
		_, ok := pocket1[item]
		if ok {
			return item
		}
	}
	return 'a'
}

func itemValue(item rune) int {
	if item >= 97 {
		return int(item) - 96
	} else {
		return int(item) - 38
	}
}
