package main

import (
	"adventofcode2022/utils"
	"container/list"
	"strings"
)

func main() {
	utils.Run(run, puzzle)
}

// 2789 too high
func run(input string) (interface{}, interface{}) {
	diagram, startPos := registerDiagram(input)

	// solve1 := findShortestPath(diagram, startPos[0])
	solve2 := findBestStartingPoint(diagram, startPos)

	return nil, solve2
}

type Node struct {
	Position Position
	Value    rune
	Friends  map[Position]*Node
	Parent   Position
}

type Position struct {
	X int
	Y int
}

func findBestStartingPoint(diagram map[Position]*Node, start []Position) int {
	bestSoFar := 999999

	for _, pos := range start {
		length := findShortestPath(diagram, pos)
		if length < bestSoFar {
			bestSoFar = length
		}

	}
	return bestSoFar
}

func findShortestPath(diagram map[Position]*Node, start Position) int {

	visited := make(map[Position]*Node)
	queue := list.New()
	queue.PushBack(diagram[start])
	visited[start] = diagram[start]
	end := Position{-1, -1}

	for queue.Len() > 0 {
		qnode := queue.Front()
		if qnode.Value.(*Node).Value == 123 {
			end = qnode.Value.(*Node).Position
			break
		}
		// iterate through all of its friends
		// mark the visited nodes; enqueue the non-visted
		for pos, node := range qnode.Value.(*Node).Friends {
			if _, ok := visited[pos]; !ok {
				if diagram[pos].Value <= qnode.Value.(*Node).Value+1 {
					visited[pos] = node
					diagram[pos].Parent = qnode.Value.(*Node).Position
					queue.PushBack(node)
				}
			}
		}
		queue.Remove(qnode)
	}
	if end.X == -1 {
		return 9999999
	}
	steps := 0
	look := end
	for {
		if look != start {
			steps += 1
			look = diagram[look].Parent
		} else {
			break
		}
	}
	return steps
}

func registerDiagram(input string) (map[Position]*Node, []Position) {
	lines := strings.Split(input, "\n")
	diagram := map[Position]*Node{}
	var start []Position
	for i, line := range lines {
		for j, char := range line {
			pos := Position{X: j, Y: i}
			if char == 'S' {
				start = append(start, pos)
				char = 96
			}
			if char == 'a' {
				start = append(start, pos)
			}
			if char == 'E' {
				char = 123
			}
			node := Node{Position: pos, Value: char, Friends: map[Position]*Node{}}
			diagram[pos] = &node
		}
	}

	// add friends
	for pos, node := range diagram {
		up := Position{X: pos.X, Y: pos.Y - 1}
		if _, ok := diagram[up]; ok {
			node.Friends[up] = diagram[up]
		}
		down := Position{X: pos.X, Y: pos.Y + 1}
		if _, ok := diagram[down]; ok {
			node.Friends[down] = diagram[down]
		}
		left := Position{X: pos.X - 1, Y: pos.Y}
		if _, ok := diagram[left]; ok {
			node.Friends[left] = diagram[left]
		}
		right := Position{X: pos.X + 1, Y: pos.Y}
		if _, ok := diagram[right]; ok {
			node.Friends[right] = diagram[right]
		}
	}
	return diagram, start
}
