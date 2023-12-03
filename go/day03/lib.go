package main

import (
	"strings"

	"github.com/senyorjou/aoc-2023/go/utils"
)

func parseMatrix(data string) [][]string {
	matrix := [][]string{}
	for _, row := range strings.Split(data, "\n") {
		stringRow := []string{}
		for _, char := range row {
			stringRow = append(stringRow, string(char))
		}
		matrix = append(matrix, stringRow)
	}
	return matrix
}

func hasAdjacentNonDigitNonDot(matrix [][]string, row, col int) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			newCol := col + j
			newRow := row + i

			if newCol >= 0 && newCol < len(matrix[0]) && newRow >= 0 && newRow < len(matrix) {
				if matrix[newRow][newCol] != "." && !isDigit(matrix[newRow][newCol]) {
					return true
				}
			}
		}
	}

	return false
}
func hasAdjacentStar(matrix [][]string, row, col int) (bool, int, int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			newCol := col + j
			newRow := row + i

			if newCol >= 0 && newCol < len(matrix[0]) && newRow >= 0 && newRow < len(matrix) {
				if matrix[newRow][newCol] == "*" {
					return true, newRow, newCol
				}
			}
		}
	}

	return false, 0, 0
}

func isDigit(char string) bool {
	for _, digit := range "0123456789" {
		if char == string(digit) {
			return true
		}
	}

	return false
}

func createStarmap(matrix [][]string) map[[2]int][]int {
	starMap := make(map[[2]int][]int)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == "*" {
				key := [2]int{row, col}
				starMap[key] = []int{}
			}
		}
	}
	return starMap
}

func calculateGears(starMap map[[2]int][]int) int {
	total := 0
	for _, values := range starMap {
		if len(values) > 1 {
			total += utils.ProdIntList(values)
		}
	}
	return total
}
