package main

import (
	"fmt"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/io"
)

// SolutionA:	157		7553
// SolutionB:	70		2758

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

func Day03B() {
	full := true
	lines := input03(full)
	numLines := len(lines)
	total := 0
	for i := 0; i < numLines; i += 3 {
		total += getPriority(findBadge(lines[i : i+3]))
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

func findBadge(lines []string) rune {
	common := make(map[rune]bool)
	for _, char := range lines[0] {
		common[char] = true
	}
	for i := 1; i < len(lines); i++ {
		uncommon := make(map[rune]bool)
		for char := range common {
			uncommon[char] = true
		}
		for _, char := range lines[i] {
			uncommon[char] = false
		}
		for char := range uncommon {
			if uncommon[char] {
				delete(common, char)
			}
		}
	}
	return fn.MapKeys(common)[0]
}
