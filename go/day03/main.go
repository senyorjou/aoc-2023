package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/senyorjou/aoc-2023/go/utils"
)

func getAdjacents(matrix [][]string, row int, col int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			newRow := row + i
			newCol := col + j
			if newRow >= 0 && newRow < len(matrix) && newCol >= 0 && newCol < len(matrix[0]) {
				if matrix[newRow][newCol] != "." {
					count++
				}
			}
		}
	}
	return count
}

func extractIntegers(matrix [][]string, row int) []string {
	joinedString := strings.Join(matrix[row], "")
	integers := regexp.MustCompile("[0-9]+").FindAllString(joinedString, -1)

	return integers
}

func solve2(matrix [][]string) int {
	aPart := false
	currentNum := 0
	hasStar := false
	possibleStar := false
	starX := 0
	starY := 0
	posx := 0
	posy := 0
	starMap := createStarmap(matrix)

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if isDigit(matrix[row][col]) {
				currentDigit, _ := strconv.Atoi(matrix[row][col])
				currentNum = 10*currentNum + currentDigit
				possibleStar, posx, posy = hasAdjacentStar(matrix, row, col)
				if possibleStar {
					aPart = true
					starX = posx
					starY = posy
					hasStar = true
				}
			} else {
				if currentNum != 0 && hasStar && aPart {
					key := [2]int{starX, starY}
					starMap[key] = append(starMap[key], currentNum)

				}
				currentNum = 0
				hasStar = false
				aPart = false
				possibleStar = false
			}
		}
	}
	totalGears := calculateGears(starMap)

	return totalGears
}

func solve1(matrix [][]string) int {
	numbers := []int{}
	currentNum := 0
	aPart := false

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if isDigit(matrix[row][col]) {
				currentDigit, _ := strconv.Atoi(matrix[row][col])
				currentNum = 10*currentNum + currentDigit
				if hasAdjacentNonDigitNonDot(matrix, row, col) {
					aPart = true
				}
			} else {
				if currentNum != 0 && aPart {
					numbers = append(numbers, currentNum)
				}
				currentNum = 0
				aPart = false
			}
		}
	}

	return utils.SumIntList(numbers)
}

func main() {
	data := utils.ReadFile("./input.txt")
	matrix := parseMatrix(data)
	fmt.Println("Day 03")
	fmt.Printf("P1: %d\n", solve1(matrix))
	fmt.Printf("P2: %d\n", solve2(matrix))
}
