package main

import (
	"fmt"

	"github.com/roidaradal/fn/io"
)

// SolutionA: 7,5,6,10,11		1140
// SolutionB: 19,23,23,29,26	3495

func input06(full bool) []string {
	lines, _ := io.ReadTextLines(getPath(6, full))
	return lines
}

func Day06A() {
	full := true
	for _, line := range input06(full) {
		fmt.Println("Marker:", findMarker(line, 4))
	}
}

func Day06B() {
	full := true
	for _, line := range input06(full) {
		fmt.Println("Marker:", findMarker(line, 14))
	}
}

func findMarker(line string, numUnique int) int {
	count := len(line)
	for n := numUnique; n <= count; n++ {
		if allUnique(line[n-numUnique : n]) {
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
