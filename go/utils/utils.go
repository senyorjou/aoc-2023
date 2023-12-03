package utils

import (
	"os"
	"strings"
)

func ReadFile(path string) string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	strContent := string(data)
	return strings.TrimRight(strContent, "\n")
}

func SumIntList(list []int) int {
	total := 0
	for _, n := range list {
		total += n
	}
	return total
}

func ProdIntList(list []int) int {
	total := 1
	for _, n := range list {
		total *= n
	}
	return total
}
