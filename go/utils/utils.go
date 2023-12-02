package utils

import (
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