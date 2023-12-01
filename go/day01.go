package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func extractNums(line string) int {
	pattern := regexp.MustCompile(`(\d)`)
	matches := pattern.FindAllStringSubmatch(line, -1)
	first, _ := strconv.Atoi(matches[0][0])
	last, _ := strconv.Atoi(matches[len(matches)-1][0])

	return first*10 + last
}

func day01P1(data string) int {
	lines := strings.Split(data, "\n")
	result := 0
	for _, line := range lines {
		result += extractNums(line)
	}
	return result
}

func day01() {
	data := readFile("./data/day01-input.txt")
	fmt.Println("Day 01")
	fmt.Printf("P1: %d ", day01P1(data))
}
