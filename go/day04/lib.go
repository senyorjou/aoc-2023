package main

import "strconv"

func twoToPower(n int) int {
	return 1 << uint(n)
}

func getIntersection(a, b []int) []int {
	intersection := []int{}
	for _, v := range a {
		for _, vv := range b {
			if v == vv {
				intersection = append(intersection, v)
			}
		}
	}
	return intersection
}

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
