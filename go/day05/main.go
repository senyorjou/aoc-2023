package main

import (
	"fmt"
	"strings"

	"github.com/senyorjou/aoc-2023/go/utils"
)

// func createSeedMap(seedMap map[int]int, block string) map[int]int {
// 	lines := strings.Split(block, "\n")
// 	fmt.Println("line", lines[0])
// 	for _, line := range lines[1:] {

// 		parts := extractNums(line)
// 		dest := parts[0]
// 		source := parts[1]
// 		num := parts[2]

// 		fmt.Println("dest, source, num", dest, source, num)
// 		for i := source; num > 0; i++ {
// 			seedMap[i] = dest
// 			num -= 1
// 			dest += 1
// 			// fmt.Println("Dest", dest)
// 		}
// 	}
// 	return seedMap
// }

// func extractInstrucions(lines []string) map[int]int {
// 	// mapName := strings.Split(lines[0], " ")[0]
// 	// createSeedMap(lines[0])
// 	seeds := extractNums(lines[0])

// 	seedMap := make(map[int]int, 0)
// 	for _, instructions := range lines[1:] {
// 		seedMap = createSeedMap(seedMap, instructions)

// 		fmt.Println("seedMap", seedMap)
// 		fmt.Println(calcSeedsPos(seeds, seedMap))

// 	}

// 	return seedMap
// }

func extractConvs(lines []string) [][]int {
	convs := [][]int{}
	for _, line := range lines {
		nums := extractNums(line)

		if len(nums) == 3 {
			convs = append(convs, nums)
		}
	}
	return convs
}

func solve1(lines []string) int {
	seeds := extractNums(lines[0])
	convs := extractConvs(lines[1:])

	for index, seed := range seeds {
		for _, conv := range convs[1:] {
			if seed >= conv[1] && seed < (conv[1]+conv[2]) {
				seeds[index] = conv[0] + seed - conv[1]

			}
		}
		fmt.Println("Seeds", seeds)
	}

	min := seeds[0]
	for _, source := range seeds[1:] {
		if source < min {
			min = source
		}
	}
	return min
}

// func solve2(lines []string) int {
// 	return 0
// }

func main() {
	data := utils.ReadFile("./input-mini.txt")
	lines := strings.Split(data, "\n")
	fmt.Println("Day 05")
	fmt.Printf("P1: %d\n", solve1(lines))
	// fmt.Printf("P2: %d\n", solve2(lines))
}
