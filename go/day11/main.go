package main

import (
	"fmt"
	"strings"

	"github.com/senyorjou/aoc-2023/go/utils"
)

type galaxy struct {
	x int
	y int
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// calculate manhattan distance between two points of a matrix
func calcDistance(a, b galaxy) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func createExpandedMap(data string, extraRows int) [][]string {
	galaxyMap := make([][]string, 0)
	for y, line := range strings.Split(data, "\n") {
		galaxyMap = append(galaxyMap, make([]string, len(line)))
		for x, char := range line {
			galaxyMap[y][x] = string(char)
		}
	}
	newMap := rotateMatrixNeg(addHorizontalRows(rotateMatrix(addHorizontalRows(galaxyMap, extraRows)), extraRows))

	return newMap
}

// iterate over the [][]string and create a map of galaxy
func createGalaxyMap(matrix [][]string) []galaxy {
	galaxyMap := make([]galaxy, 0)
	for y, row := range matrix {
		for x, char := range row {
			if char == "#" {
				galaxyMap = append(galaxyMap, galaxy{x, y})
			}
		}
	}
	return galaxyMap
}
func solve1(data string) int {
	baseMap := createExpandedMap(data, 1)
	galaxies := createGalaxyMap(baseMap)
	total := 0
	for index, g1 := range galaxies {
		for _, g2 := range galaxies[index+1:] {
			total += calcDistance(g1, g2)
		}
	}

	return total
}

func solve2(data string) int {
	baseMap := createExpandedMap(data, 10000)
	galaxies := createGalaxyMap(baseMap)
	total := 0
	for index, g1 := range galaxies {
		for _, g2 := range galaxies[index+1:] {
			total += calcDistance(g1, g2)
		}
	}

	return total
}

func main() {
	data := string(utils.ReadFile("./input.txt"))
	fmt.Println("Day 11")
	fmt.Printf("P1: %d\n", solve1(data))
	fmt.Printf("P2: %d\n", solve2(data))
}
