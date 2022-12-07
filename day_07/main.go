package main

import (
	"adventofcode2022/utils"
	"strconv"
	"strings"
)

func main() {
	utils.Run(run, puzzle)
}

func run(input string) (interface{}, interface{}) {
	root := registerDirectories(input)

	solve1 := findSum(root)
	solve2 := findDirToDelte(root)

	return solve1, solve2
}

func findSum(dir Directory) int {
	var treshold = 100000
	sum := dir.getSizeUnderTreshold(treshold)
	return sum
}

func findDirToDelte(dir Directory) int {
	const diskSpace = 70000000
	var freeSpace = diskSpace - dir.Size
	var target = 30000000 - freeSpace

	result := dir.getDirsForSpace(target)
	return result
}

func registerDirectories(input string) Directory {
	lines := strings.Split(input, "\n")
	root := Directory{Name: "/", ChildDir: []*Directory{}, Size: 0}

	currentDir := &root

	for i, line := range lines {
		data := strings.Split(line, " ")
		if i == 0 {
			continue
		}
		if data[0] == "$" { //command
			if data[1] == "cd" {
				if data[2] == ".." {
					currentDir = currentDir.ParentDir
				} else {
					for _, dir := range currentDir.ChildDir {
						if dir.Name == data[2] {
							currentDir = dir
							break
						}
					}
				}
			}
		} else if data[0] == "dir" {
			newDir := NewDir(currentDir, data[1])
			currentDir.ChildDir = append(currentDir.ChildDir, &newDir)
		} else {
			newFileSize, _ := strconv.Atoi(data[0])
			currentDir.Size += newFileSize
			currentDir.addSizeToParent(newFileSize)
		}
	}
	return root
}
