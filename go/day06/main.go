package main

import (
	"fmt"

	"github.com/senyorjou/aoc-2023/go/utils"
)

func raceResults(race []int) int {
	time := race[0]
	maxDistance := race[1]
	totalWins := 0
	for speed := 1; speed <= time; speed++ {
		raceDistance := speed * (time - speed)
		if raceDistance > maxDistance {
			totalWins += 1
		}
	}

	return totalWins
}

func solve1(data string) int {
	races := parseBadData(data)
	winnings := 1
	for _, race := range races {
		winnings *= raceResults(race)
	}

	return winnings
}

func solve2(data string) int {
	time, distance := parseGoodData(data)
	input := []int{time, distance}
	return raceResults(input)
}

func main() {
	data := string(utils.ReadFile("./input.txt"))
	fmt.Println("Day 06")
	fmt.Printf("P1: %d\n", solve1(data))
	fmt.Printf("P2: %d\n", solve2(data))
}
