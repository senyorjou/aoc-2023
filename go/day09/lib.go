package main

import (
	"regexp"
	"strconv"
)

func extractNums(line string) []int {
	pattern := regexp.MustCompile(`(-?\d+)`)
	matches := pattern.FindAllString(line, -1)
	nums := make([]int, 0)

	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		nums = append(nums, num)
	}
	return nums
}

func allZeroes(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}
