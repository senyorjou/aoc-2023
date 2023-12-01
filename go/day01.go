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

func convert(val string) int {
	values := map[string]int{
		"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4,
		"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	}

	return values[val]
}

func extractComplex(line string) int {
	pattern := regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine|\d)`)
	matches := pattern.FindAllStringSubmatch(line, -1)
	first := convert(matches[0][0])
	last := convert(matches[len(matches)-1][0])

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

func day01P2(data string) int {
	lines := strings.Split(data, "\n")
	result := 0
	for _, line := range lines {
		result += extractComplex(line)
	}
	return result
}

func day01() {
	data := readFile("./data/day01-input.txt")
	fmt.Println("Day 01")
	fmt.Printf("P1: %d\n", day01P1(data))
	fmt.Printf("P2: %d ", day01P2(data))
}
