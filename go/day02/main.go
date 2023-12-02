package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parsePlay(play string) (int, int, int) {
	R := 0
	G := 0
	B := 0

	for _, part := range strings.Split(play, ";") {
		for _, block := range strings.Split(strings.Trim(part, " "), ",") {
			play := strings.Split(strings.Trim(block, " "), " ")
			num, _ := strconv.Atoi(play[0])
			color := play[1]

			if color == "red" {
				R += num
			}
			if color == "green" {
				G += num
			}
			if color == "blue" {
				B += num
			}

		}
	}
	return R, G, B
}

func parseLine(line string) (int, int, int, int) {
	parts := strings.Split(line, ":")
	gameId, _ := strconv.Atoi((strings.Split(parts[0], " "))[1])

	plays := strings.Split(parts[1], ";")
	RMax := 0
	GMax := 0
	BMax := 0
	for _, play := range plays {
		R, G, B := parsePlay(play)
		if R > RMax {
			RMax = R
		}
		if G > GMax {
			GMax = G
		}
		if B > BMax {
			BMax = B
		}
	}
	return gameId, RMax, GMax, BMax
}

func solvePossible(input string) int {
	lines := strings.Split(input, "\n")

	total := 0
	RMax := 12
	GMax := 13
	BMax := 14

	for _, line := range lines {
		game, R, G, B := parseLine(line)
		if R <= RMax && G <= GMax && B <= BMax {
			total += game
		}
	}

	return total
}

func solveMinimum(input string) int {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		_, R, G, B := parseLine(line)
		total += R * G * B
	}

	return total
}

func day02() {
	data := utils.readFile("./data/day02-input.txt")
	fmt.Println("Day 02")
	fmt.Printf("P1: %d\n", solvePossible(data))
	fmt.Printf("P2: %d\n", solveMinimum(data))
}
