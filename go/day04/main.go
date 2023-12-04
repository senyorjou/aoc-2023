package main

import (
	"fmt"
	"strings"

	"github.com/senyorjou/aoc-2023/go/utils"
)

func parseLine(line string) ([]int, []int) {
	parts := strings.Split(line, "|")
	left := strings.Split(parts[0], " ")
	right := strings.Split(parts[1], " ")

	return extractNumbers(left[1:]), extractNumbers(right)
}

func solve1(lines []string) int {
	points := 0
	for _, line := range lines {
		left, right := parseLine(line)
		common := getIntersection(left, right)
		points += twoToPower(len(common) - 1)
	}
	return points
}

func solve2(lines []string) int {
	stack := map[int]int{}
	for i, line := range lines {
		left, right := parseLine(line)
		numWins := getIntersection(left, right)
		stack[i] += 1
		for j := 1; j <= stack[i]; j++ {
			for k := range numWins {
				stack[i+k+1] += 1
			}
		}
	}
	total := 0
	for _, v := range stack {
		total += v
	}
	return total
}

func main() {
	data := utils.ReadFile("./input.txt")
	lines := strings.Split(data, "\n")
	fmt.Println("Day 04")
	fmt.Printf("P1: %d\n", solve1(lines))
	fmt.Printf("P2: %d\n", solve2(lines))
}
