package main

type Result int

const (
	Win     Result = 6
	Lose     Result = 0
	Draw     Result = 3
)

func translateResult(input string) Result {
	switch input {
	case "X":
		return Lose
	case "Y":
		return Draw
	default:
		return Win
	}
}