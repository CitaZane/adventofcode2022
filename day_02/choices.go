package main

type Choice int

const (
	Rock     Choice = 1
	Paper    Choice = 2
	Scissors Choice = 3
)

func translateChoice(input string) Choice {
	switch input {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	default:
		return Scissors
	}
}

func (c Choice) calcRoundResult(against Choice) Result {
	if c == against {
		return Draw
	}

	switch c {
	case Rock:
		if against == Paper {
			return Lose
		}
		if against == Scissors {
			return Win
		}
	case Paper:
		if against == Rock {
			return Win
		}
		if against == Scissors {
			return Lose
		}
	case Scissors:
		if against == Paper {
			return Win
		}
		if against == Rock {
			return Lose
		}
	}
	return 0
}

func (c Choice) calcBasedOnResult(result Result) Choice {
	if result == Draw {
		return c
	}
	switch c {
	case Rock:
		if result == Win {
			return Paper
		}
		if result == Lose {
			return Scissors
		}
	case Paper:
		if result == Win {
			return Scissors
		}
		if result == Lose {
			return Rock
		}
	case Scissors:
		if result == Win {
			return Rock
		}
		if result == Lose {
			return Paper
		}
	}
	return 0
}
