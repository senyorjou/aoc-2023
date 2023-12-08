package main

import (
	"regexp"
	"strings"
)

func splitStingintoRunes(s string) []string {
	runes := []string{}
	for _, r := range s {
		runes = append(runes, string(r))
	}
	return runes
}

func parseNodeData(line string) (string, leftRight) {
	pattern := regexp.MustCompile(`(\w+)`)
	matches := pattern.FindAllString(line, -1)

	return matches[0], leftRight{left: matches[1], right: matches[2]}
}

func findNodesEndingWithA(nodes map[string]leftRight) []string {
	candidates := []string{}
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			candidates = append(candidates, k)
		}
	}
	return candidates
}

// function to calculate GCD (Greatest Common Divisor)
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Function to calculate LCM (Least Common Multiple)
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Function to calculate LCM of a slice of ints
func lcmSlice(nums []int) int {
	current := 1
	for _, num := range nums {
		current = lcm(current, num)
	}
	return current
}
