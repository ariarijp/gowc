package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func countWordsInLine(line string) int {
	re := regexp.MustCompile("\\s+")

	trimmed := strings.TrimSpace(line)
	normalized := re.ReplaceAllString(trimmed, " ")

	if len(normalized) == 0 {
		return 0
	}

	return strings.Count(normalized, " ") + 1
}

func countWordsInLines(stdin *os.File) int {
	words := 0

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		words += countWordsInLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic("Something went wrong")
	}

	return words
}
