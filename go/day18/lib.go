package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func createdigPlan(input string) []DigStep {
	digPlan := make([]DigStep, 0)
	var direction string
	var steps int
	var color string
	for _, line := range strings.Split(input, "\n") {
		fmt.Sscanf(line, "%s %d (%s)", &direction, &steps, &color)
		digPlan = append(digPlan, DigStep{direction, steps, color[1 : len(color)-1]})

	}
	return digPlan
}

func polygonArea(points []Coor) int {
	area := 0.0
	for i := 0; i < len(points)-1; i++ {
		area += float64(points[i].x * points[i+1].y)
		area -= float64(points[i].y * points[i+1].x)
	}
	return int(math.Abs(area / 2))
}

func convertColor(hexStr string) (int, string) {
	dir := hexStr[len(hexStr)-1:]
	value, _ := strconv.ParseInt(hexStr[:5], 16, 64)

	return int(value), dir
}
