package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/senyorjou/aoc-2023/go/utils"
)

func parseData(data string) []string {
	hands := []string{}
	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, " ")
		hands = append(hands, parts[0])
	}
	return hands
}

func solve1(data string) int {
	hands := parseData(data)
	results := make(map[[2]int]int)
	results2 := make(map[int64]int)

	fmt.Println("hands", hands)
	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bid, _ := strconv.Atoi(parts[1])
		handValue, decValue := calculateHighestCombination(hand)
		results[[2]int{handValue, decValue}] = bid
		results2[int64(handValue*10_000_000)+int64(decValue)] = bid
	}
	// fmt.Println("results", results)
	// fmt.Println("results2", results2)

	keys := make([]int64, len(results2))

	i := 0
	for k := range results2 {
		keys[i] = k
		i++
	}
	// fmt.Println("keys", keys)
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	// fmt.Println("keys", keys)

	total := 0
	for index, key := range keys {
		index++
		// fmt.Println("key", key, results2[key], results2[key]*index)
		total += results2[key] * index
	}

	return total
}

func main() {
	data := string(utils.ReadFile("./input.txt"))
	// data := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"
	fmt.Println("Day 07")
	fmt.Printf("P1: %d\n", solve1(data))
	// fmt.Printf("P2: %d\n", solve2(data))
}
