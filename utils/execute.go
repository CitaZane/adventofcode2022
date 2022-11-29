package utils

import (
	"fmt"
	"time"
)

func Run(run func(string) (interface{}, interface{}), puzzle string) {
	start := time.Now()
	solve1, solve2 := run(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("SOLVE1: %v\nSOLVE2: %v\n", solve1, solve2)
	fmt.Printf("Program took %s", elapsed)
}
