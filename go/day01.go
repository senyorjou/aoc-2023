package main

import (
	"fmt"
	"regexp"
	"strings"
)

type extract func(string) int

var values = map[string]int{
	"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4,
	"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

func convert(val string) int {
	return values[val]
}

func extractAny(line string, pattern *regexp.Regexp) int {
	matches := pattern.FindAllStringSubmatch(line, -1)
	first := convert(matches[0][0])
	last := convert(matches[len(matches)-1][0])

	return first*10 + last
}

func extractNums(line string) int {
	pattern := regexp.MustCompile(`(\d)`)
	return extractAny(line, pattern)
}

func extractComplex(line string) int {
	pattern := regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine|\d)`)
	return extractAny(line, pattern)
}

func solve(data string, exctractfn extract) int {
	lines := strings.Split(data, "\n")
	result := 0
	for _, line := range lines {
		result += exctractfn(line)
	}
	return result
}

func day01() {
	data := readFile("./data/day01-input.txt")
	fmt.Println("Day 01")
	fmt.Printf("P1: %d\n", solve(data, extractNums))
	fmt.Printf("P2: %d ", solve(data, extractComplex))
}
