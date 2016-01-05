package main

import (
	"bufio"
	"os"
)

func countLines(stdin *os.File) int {
	lines := 0

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		panic("Something went wrong")
	}

	return lines
}
