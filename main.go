package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWord(line string) int {
	trimmed := strings.TrimSpace(line)

	if len(trimmed) == 0 {
		return 0
	}

	return strings.Count(trimmed, " ") + 1
}

func main() {
	words := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words += countWord(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic("Something went wrong")
	}

	fmt.Println(words)
}
