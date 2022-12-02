package main

import (
	"adventofcode2022/utils"
	"strings"
)

func main() {
	utils.Run(run, puzzle)
}

// 11449
// 13187
func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")

	solve1 := calcGameScore(lines, false)
	solve2 := calcGameScore(lines, true)

	return solve1, solve2
}

func calcGameScore(input []string, smart bool) int {
	totalScore := 0

	for _, round := range input {
		var roundScore int
		if smart{
			roundScore = calcSmartScorePerRound(round)
		}else{
			roundScore = calcScorePerRound(round)
		}
		totalScore += roundScore
	}
	return totalScore
}

func calcScorePerRound(input string) int {
	choices := strings.Split(input, " ")

	roundScore := 0
	competitor := translateChoice(choices[0])
	yourChoice := translateChoice(choices[1])

	roundScore += int(yourChoice)
	roundScore += int(yourChoice.calcRoundResult(competitor))
	return roundScore
}


func calcSmartScorePerRound(input string) int {
	choices := strings.Split(input, " ")

	roundScore := 0
	competitor := translateChoice(choices[0])
	result := translateResult(choices[1])

	roundScore += int(result)
	roundScore += int(competitor.calcBasedOnResult(result))
	return roundScore
}
