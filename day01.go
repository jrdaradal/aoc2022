package main

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/io"
)

// SolutionA: 	24000	70698

func input01(full bool) []string {
	lines, _ := io.ReadTextFile(getPath(1, full))
	return fn.Map(strings.Split(lines, "\n"), strings.TrimSpace)
}

func Day01A() {
	lines := input01(true)
	maxCalories := 0
	current := 0
	for _, line := range lines {
		if line == "" {
			if current > maxCalories {
				maxCalories = current
			}
			current = 0
		} else {
			current += fn.ParseInt(line)
		}
	}
	if current > maxCalories {
		maxCalories = current
	}
	fmt.Println("MaxCalories:", maxCalories)
}
