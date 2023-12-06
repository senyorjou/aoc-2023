package main

import (
	"regexp"
	"strconv"
	"strings"
)

func extractNumbers(sNums []string) []int {
	nums := []int{}
	for _, v := range sNums {
		num, _ := strconv.Atoi(v)
		if num == 0 {
			continue
		}
		nums = append(nums, num)
	}
	return nums
}

func parseGoodData(input string) (int, int) {
	lines := strings.Split(input, "\n")
	pattern := regexp.MustCompile(`(\d+)`)

	matches := pattern.FindAllString(lines[0], -1)
	time, _ := strconv.Atoi(strings.Join(matches, ""))

	matches = pattern.FindAllString(lines[1], -1)
	distance, _ := strconv.Atoi(strings.Join(matches, ""))

	return time, distance
}

func parseBadData(input string) [][]int {
	lines := strings.Split(input, "\n")
	convs := make([][]int, 0)
	for _, line := range lines {
		nums := extractNumbers(strings.Split(line, " "))
		convs = append(convs, nums)
	}
	entries := transpose(convs)

	return entries
}

// made by copilot!
func transpose(matrix [][]int) [][]int {
	out := make([][]int, len(matrix[0]))
	for i := range out {
		out[i] = make([]int, len(matrix))
	}
	for i, row := range matrix {
		for j, val := range row {
			out[j][i] = val
		}
	}
	return out
}
