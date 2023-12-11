package main

import (
	"strings"
)

// rotate a matrix -90 degrees
func rotateMatrixNeg(matrix [][]string) [][]string {
	rotatedMatrix := make([][]string, len(matrix[0]))
	for i := range rotatedMatrix {
		rotatedMatrix[i] = make([]string, len(matrix))
		for j := range rotatedMatrix[i] {
			rotatedMatrix[i][j] = matrix[j][len(matrix[0])-i-1]
		}
	}
	return rotatedMatrix
}

// rotate a matrix 90 degrees
func rotateMatrix(matrix [][]string) [][]string {
	rotatedMatrix := make([][]string, len(matrix[0]))
	for i := range rotatedMatrix {
		rotatedMatrix[i] = make([]string, len(matrix))
		for j := range rotatedMatrix[i] {
			rotatedMatrix[i][j] = matrix[len(matrix)-j-1][i]
		}
	}
	return rotatedMatrix
}

func addHorizontalRows(matrix [][]string, numRows int) [][]string {
	newMatrix := make([][]string, 0)
	for _, row := range matrix {
		newMatrix = append(newMatrix, row)
		if !strings.Contains(strings.Join(row, ""), "#") {
			for i := 0; i < numRows; i++ {
				newMatrix = append(newMatrix, row)
			}
		}
	}
	return newMatrix
}
