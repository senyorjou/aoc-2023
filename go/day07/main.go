package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/senyorjou/aoc-2023/go/utils"
)

func solve1(data string) int {
	results := make(map[[2]int]int)
	results2 := make(map[int64]int)

	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bid, _ := strconv.Atoi(parts[1])
		handValue, decValue := calculateHighestCombination(hand, false)
		results[[2]int{handValue, decValue}] = bid
		results2[int64(handValue*10_000_000)+int64(decValue)] = bid
	}
	keys := make([]int64, len(results2))

	i := 0
	for k := range results2 {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	total := 0
	for index, key := range keys {
		index++
		total += results2[key] * index
	}

	return total
}

func solve2(data string) int {
	results := make(map[int64]int)

	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bid, _ := strconv.Atoi(parts[1])

		topVal, decValue := calculateHighestCombination(hand, true)
		for cardVal := range cardValues {
			newHand := strings.ReplaceAll(hand, "J", cardVal)
			handValue, _ := calculateHighestCombination(newHand, true)
			if handValue > topVal {
				topVal = handValue
			}
		}
		// handValue, decValue := calculateHighestCombination(hand)
		results[int64(topVal*100_000_000)+int64(decValue)] = bid
	}
	keys := make([]int64, len(results))

	i := 0
	for k := range results {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	total := 0
	for index, key := range keys {
		index++
		total += results[key] * index
	}

	return total
}

func main() {
	data := string(utils.ReadFile("./input.txt"))
	fmt.Println("Day 07")
	fmt.Printf("P1: %d\n", solve1(data))
	fmt.Printf("P2: %d\n", solve2(data))
}
