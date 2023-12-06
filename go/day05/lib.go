package main

import (
	"regexp"
	"strconv"
)

func calcSeedsPos(seeds []int, instructionsMap map[int]int) map[int]int {
	seedMap := map[int]int{}
	for _, seed := range seeds {
		if instructionsMap[seed] == 0 {
			seedMap[seed] = seed
			continue
		} else {
			seedMap[seed] = instructionsMap[seed]
		}
	}
	return seedMap
}

func extractNums(line string) []int {
	asNum := []int{}
	pattern := regexp.MustCompile(`(\d+)`)

	for _, num := range pattern.FindAllString(line, -1) {
		possibleNum, _ := strconv.Atoi(num)
		asNum = append(asNum, possibleNum)
	}
	return asNum
}
