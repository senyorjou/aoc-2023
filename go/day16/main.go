package main

import (
	"fmt"
	"strings"

	"github.com/senyorjou/aoc-2023/go/utils"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type checkpoint struct {
	x   int
	y   int
	dir Direction
}

func createGrid(input string) [][]rune {
	grid := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	return grid
}

// create a function to calculate if a position is within the grid
func isWithinGrid(grid [][]rune, x, y int) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0])
}

func moveOneStep(x, y int, direction Direction) (int, int) {
	switch direction {
	case North:
		y--
	case East:
		x++
	case South:
		y++
	case West:
		x--
	}
	return x, y
}

func walk(grid [][]rune, walked map[checkpoint]bool, x, y int, direction Direction) map[checkpoint]bool {
	for {
		if !isWithinGrid(grid, x, y) {
			break
		}
		if _, exists := walked[checkpoint{x, y, direction}]; exists {
			break
		}
		walked[checkpoint{x, y, direction}] = true

		switch grid[y][x] {
		case
			'.':
			x, y = moveOneStep(x, y, direction)
		case '|':
			if direction == East || direction == West {
				walked = walk(grid, walked, x, y-1, North)
				walked = walk(grid, walked, x, y+1, South)
			} else {
				x, y = moveOneStep(x, y, direction)
			}
		case '-':
			if direction == North || direction == South {
				walked = walk(grid, walked, x-1, y, West)
				walked = walk(grid, walked, x+1, y, East)
			} else {
				x, y = moveOneStep(x, y, direction)
			}
		case '/':
			switch direction {
			case North:
				walked = walk(grid, walked, x+1, y, East)
			case East:
				walked = walk(grid, walked, x, y-1, North)
			case South:
				walked = walk(grid, walked, x-1, y, West)
			case West:
				walked = walk(grid, walked, x, y+1, South)
			}
		case '\\':
			switch direction {
			case North:
				walked = walk(grid, walked, x-1, y, West)
			case East:
				walked = walk(grid, walked, x, y+1, South)
			case South:
				walked = walk(grid, walked, x+1, y, East)
			case West:
				walked = walk(grid, walked, x, y-1, North)
			}
		}
	}

	return walked
}

func totalChecks(grid [][]rune, walked map[checkpoint]bool) int {
	total := 0
	for y, row := range grid {
		for x := range row {
			if walked[checkpoint{x, y, East}] ||
				walked[checkpoint{x, y, West}] ||
				walked[checkpoint{x, y, North}] ||
				walked[checkpoint{x, y, South}] {
				total++
			}
		}
	}
	return total
}

func solveGrid(grid [][]rune, x, y int, direction Direction) int {
	walked := make(map[checkpoint]bool)
	walk(grid, walked, x, y, direction)
	total := totalChecks(grid, walked)

	return total
}

func solve1(input string) int {
	grid := createGrid(input)
	return solveGrid(grid, 0, 0, West)
}

func solve2(input string) int {
	grid := createGrid(input)
	results := make([]int, 0)
	for y, row := range grid {
		for x := range row {
			if y == 0 ||
				y == len(grid)-1 ||
				x == 0 ||
				x == len(grid[0])-1 {
				for dir := range []Direction{North, South, East, West} {
					results = append(results, solveGrid(grid, x, y, Direction(dir)))
				}
			}
		}
	}
	return utils.MaxIntList(results)
}

func main() {
	data := (utils.ReadFile("./input.txt"))
	fmt.Println("Day 16")
	fmt.Printf("P1: %d\n", solve1(data))
	fmt.Printf("P2: %d\n", solve2(data))
}
