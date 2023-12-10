package main

import (
	"strings"
)

func dirInAllowedDirs(s string, list []string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	return false
}

func findInMatrix(matrix [][]string, target string) [2]int {
	for y, row := range matrix {
		for x, char := range row {
			if char == target {
				return [2]int{y, x}
			}
		}
	}
	return [2]int{0, 0}
}

func transposeAndCovert(matrix [][]rune) [][]string {
	newMatrix := make([][]string, len(matrix[0]))
	for i := range newMatrix {
		newMatrix[i] = make([]string, len(matrix))
	}
	for i, row := range matrix {
		for j, char := range row {
			newMatrix[j][i] = string(char)
		}
	}
	return newMatrix
}

func parseData(data string) [][]rune {
	rows := strings.Split(data, "\n")
	matrix := make([][]rune, len(rows))
	for i, row := range rows {
		matrix[i] = []rune(row)
	}
	return matrix
}
