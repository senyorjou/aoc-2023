package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(path string) string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	strContent := string(data)
	return strings.TrimRight(strContent, "\n")
}

func main() {
	fmt.Println("AOC 2023, let's GO")
	day01()
	day02()
}
