package main

import (
	"fmt"

	"github.com/senyorjou/aoc-2023/go/utils"
)

type Direction int

type DigStep struct {
	dir   string
	steps int
	color string
}

type Coor struct {
	x int
	y int
}

func createPerimeter(digplan []DigStep, part string) ([]Coor, int) {
	perimeter := make([]Coor, 0)
	perimeter = append(perimeter, Coor{0, 0})
	currentPos := Coor{0, 0}
	rope := 0
	for _, dig := range digplan {
		steps := dig.steps
		dir := dig.dir
		if part == "2" {
			steps, dir = convertColor(dig.color)
		}
		switch dir {
		case "R", "0":
			currentPos = Coor{currentPos.x + steps, currentPos.y}
		case "L", "2":
			currentPos = Coor{currentPos.x - steps, currentPos.y}
		case "U", "3":
			currentPos = Coor{currentPos.x, currentPos.y - steps}
		case "D", "1":
			currentPos = Coor{currentPos.x, currentPos.y + steps}
		}
		perimeter = append(perimeter, currentPos)
		rope += steps
	}

	return perimeter, rope
}

func solve(input string, part string) int {
	digplan := createdigPlan(input)
	perimeter, rope := createPerimeter(digplan, part)
	// fmt.Println("Ropea", rope)
	// fmt.Println("Area", polygonArea(perimeter))
	// fmt.Println("Total Area:", polygonArea(perimeter)+(rope/2)+1)
	return polygonArea(perimeter) + (rope / 2) + 1
}

func main() {
	data := (utils.ReadFile("./input.txt"))
	fmt.Println("Day 18")
	fmt.Printf("P1: %d\n", solve(data, "1"))
	fmt.Printf("P2: %d\n", solve(data, "2"))
}
