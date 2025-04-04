package main

import (
	"fmt"

	"github.com/roidaradal/fn/io"
)

// SolutionA:	157		7553
// SolutionB:

func input03(full bool) []string {
	lines, _ := io.ReadTextLines(getPath(3, full))
	return lines
}

func Day03A() {
	full := true
	total := 0
	for _, line := range input03(full) {
		total += getPriority(findCommon(line))
	}
	fmt.Println("Total:", total)
}

func getPriority(char rune) int {
	if 97 <= char && char <= 122 {
		return int(char) - 96
	} else if 65 <= char && char <= 90 {
		return int(char) - 38
	}
	return 0
}

func findCommon(line string) rune {
	mid := len(line) / 2
	set := make(map[rune]bool)
	for i, char := range line {
		if i < mid {
			set[char] = true
		} else {
			if _, ok := set[char]; ok {
				return char
			}
		}
	}
	return '?'
}
