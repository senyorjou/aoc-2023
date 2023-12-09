package main

import (
	"fmt"
	"strings"

	"github.com/senyorjou/aoc-2023/go/utils"
)

func createNextIteration(iteration []int) []int {
	nextIteration := make([]int, 0)
	for i := 0; i < len(iteration)-1; i++ {
		nextIteration = append(nextIteration, iteration[i+1]-iteration[i])
	}
	return nextIteration
}

func sumAllLastNumberOfIterations(iterations [][]int) int {
	total := 0
	for _, iteration := range iterations {
		total += iteration[len(iteration)-1]
	}
	return total
}

func calculatePreviousNumbers(iterations [][]int) int {
	firsts := make([]int, 0)
	for _, iteration := range iterations {
		firsts = append(firsts, iteration[0])
	}
	pointer := 0
	for i := len(firsts) - 1; i > 0; i-- {
		pointer = firsts[i-1] - pointer
	}

	return pointer
}

func extractIters(nums string) [][]int {
	iterations := make([][]int, 0)
	iter := extractNums(nums)
	iterations = append(iterations, iter)
	for {
		iter = createNextIteration(iter)
		iterations = append(iterations, iter)
		if allZeroes(iter) {
			break
		}
	}
	return iterations
}

func solve1(data string) int {
	lines := strings.Split(data, "\n")
	total := 0
	for _, line := range lines {
		total += sumAllLastNumberOfIterations(extractIters(line))
	}
	return total
}

func solve2(data string) int {
	lines := strings.Split(data, "\n")

	total := 0
	for _, line := range lines {
		total += calculatePreviousNumbers(extractIters(line))
	}
	return total
}

func main() {
	data := string(utils.ReadFile("./input.txt"))
	fmt.Println("Day 09")
	fmt.Printf("P1: %d\n", solve1(data))
	fmt.Printf("P2: %d\n", solve2(data))
}
