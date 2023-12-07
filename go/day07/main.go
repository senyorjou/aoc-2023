package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/senyorjou/aoc-2023/go/utils"
)

func solve(data string, joker bool) int {
	results := make(map[int64]int)

	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bid, _ := strconv.Atoi(parts[1])

		handValue, decValue := calculateHighestCombination(hand, joker)
		if joker {
			currVal := handValue
			for cardVal := range cardValues {
				newHand := strings.ReplaceAll(hand, "J", cardVal)
				currVal, _ = calculateHighestCombination(newHand, joker)
				if currVal > handValue {
					handValue = currVal
				}
			}
		}
		results[int64(handValue*1_000_000)+int64(decValue)] = bid
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
	fmt.Printf("P1: %d\n", solve(data, false))
	fmt.Printf("P2: %d\n", solve(data, true))
}
