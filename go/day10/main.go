package main

import (
	"fmt"

	"github.com/senyorjou/aoc-2023/go/utils"
)

func findMovements(matrix [][]string, startPos [2]int, forbidden [2]int) [2]int {
	movs := map[[2]int][]string{
		{0, -1}: {"|", "F", "7"},
		{0, 1}:  {"|", "L", "J"},
		{1, 0}:  {"-", "J", "7"},
		{-1, 0}: {"-", "F", "L"},
	}

	nextPos := startPos
	for dir := range movs {
		nextPos = [2]int{startPos[0] + dir[0], startPos[1] + dir[1]}

		if nextPos == forbidden {
			continue
		}
		if nextPos[0] < 0 || nextPos[1] < 0 {
			continue
		}
		if nextPos[0] >= len(matrix[0]) || nextPos[1] >= len(matrix) {
			continue
		}

		if matrix[nextPos[0]][nextPos[1]] == "S" {
			break
		}
		if dirInAllowedDirs(matrix[nextPos[0]][nextPos[1]], movs[dir]) {
			break
		}
	}
	return nextPos
}

func calcNextPos(startPos [2]int, forbidden [2]int, possibleDirs [2][2]int) [2]int {
	for _, possibleDir := range possibleDirs {
		nextPos := [2]int{startPos[0] + possibleDir[0], startPos[1] + possibleDir[1]}
		if nextPos[0] != forbidden[0] || nextPos[1] != forbidden[1] {
			return nextPos
		}
	}
	return [2]int{0, 0}
}

func solve1(data string) int {
	matrix := transposeAndCovert(parseData(data))
	startPos := findInMatrix(matrix, "S")

	forbiddenMov := [2]int{0, 0}
	nextPos := findMovements(matrix, startPos, forbiddenMov)

	forbiddenMov = startPos

	movs := 1
	possibleDirs := [2][2]int{}
	for {
		movs++
		direction := matrix[nextPos[0]][nextPos[1]]
		switch direction {
		case "|":
			possibleDirs = [2][2]int{{0, -1}, {0, 1}}
		case "-":
			possibleDirs = [2][2]int{{1, 0}, {-1, 0}}
		case "F":
			possibleDirs = [2][2]int{{1, 0}, {0, 1}}
		case "L":
			possibleDirs = [2][2]int{{1, 0}, {0, -1}}
		case "J":
			possibleDirs = [2][2]int{{0, -1}, {-1, 0}}
		case "7":
			possibleDirs = [2][2]int{{0, 1}, {-1, 0}}
		}
		oldPos := nextPos
		nextPos = calcNextPos(nextPos, forbiddenMov, possibleDirs)
		forbiddenMov = oldPos

		if matrix[nextPos[0]][nextPos[1]] == "S" {
			break
		}
	}

	return movs / 2
}

func main() {
	data := string(utils.ReadFile("./input.txt"))
	fmt.Println("Day 10")
	fmt.Printf("P1: %d\n", solve1(data))
	// fmt.Printf("P2: %d\n", solve2(data))
}
