package main

import (
	"fmt"

	"github.com/roidaradal/fn/io"
)

// SolutionA: 7,5,6,10,11	1140
// SolutionB:

func input06(full bool) []string {
	lines, _ := io.ReadTextLines(getPath(6, full))
	return lines
}

func Day06A() {
	full := true
	for _, line := range input06(full) {
		fmt.Println("Marker:", findMarker(line))
	}
}

func findMarker(line string) int {
	count := len(line)
	for n := 4; n <= count; n++ {
		if allUnique(line[n-4 : n]) {
			return n
		}
	}
	return 0
}

func allUnique(text string) bool {
	set := make(map[rune]bool)
	for _, char := range text {
		set[char] = true
	}
	return len(text) == len(set)
}
