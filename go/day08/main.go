package main

import (
	"fmt"
	"strings"

	"github.com/senyorjou/aoc-2023/go/utils"
)

type leftRight struct {
	left  string
	right string
}

func parseInstructions(data string) ([]string, map[string]leftRight) {
	nodes := map[string]leftRight{}
	lines := strings.Split(data, "\n")
	instructions := splitStingintoRunes(lines[0])

	for _, line := range lines[2:] {
		k, v := parseNodeData(line)
		nodes[k] = v
	}

	return instructions, nodes
}

func solve1(data string) int {
	instructions, nodes := parseInstructions(data)

	total := 0
	currentIndex := 0
	maxInstructions := len(instructions)
	currentNode := "AAA"
	for {
		if currentNode == "ZZZ" {
			break
		}
		total++
		if instructions[currentIndex] == "L" {
			currentNode = nodes[currentNode].left
		} else {
			currentNode = nodes[currentNode].right
		}
		currentIndex++
		if currentIndex == maxInstructions {
			currentIndex = 0
		}
	}

	return total
}

func getNodeFinalDestination(instructions []string, nodes map[string]leftRight, currentNode string) int {
	total := 0
	currentIndex := 0
	maxInstructions := len(instructions)
	for {
		if strings.HasSuffix(currentNode, "Z") {
			break
		}
		total++
		if instructions[currentIndex] == "L" {
			currentNode = nodes[currentNode].left
		} else {
			currentNode = nodes[currentNode].right
		}
		currentIndex++
		if currentIndex == maxInstructions {
			currentIndex = 0
		}
	}
	return total
}

func solve2(data string) int {
	instructions, nodes := parseInstructions(data)

	candidates := make([]int, 0)
	for _, candidate := range findNodesEndingWithA(nodes) {
		candidates = append(candidates, getNodeFinalDestination(instructions, nodes, candidate))
	}

	return lcmSlice(candidates)
}

func main() {
	data := string(utils.ReadFile("./input.txt"))
	fmt.Println("Day 08")
	fmt.Printf("P1: %d\n", solve1(data))
	fmt.Printf("P2: %d\n", solve2(data))
}
