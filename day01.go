package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/io"
)

// SolutionA: 	24000	70698
// SolutionB: 	45000	206643

func input01(full bool) []string {
	lines, _ := io.ReadTextFile(getPath(1, full))
	return fn.Map(strings.Split(lines, "\n"), strings.TrimSpace)
}

func Day01A() {
	full := true
	maxCalories, current := 0, 0
	for _, line := range input01(full) {
		if line == "" {
			maxCalories = max(maxCalories, current)
			current = 0
		} else {
			current += fn.ParseInt(line)
		}
	}
	maxCalories = max(maxCalories, current)
	fmt.Println("MaxCalories:", maxCalories)
}

func Day01B() {
	full := true
	top3, current := []int{0, 0, 0}, 0
	for _, line := range input01(full) {
		if line == "" {
			top3 = adjustTop3(top3, current)
			current = 0
		} else {
			current += fn.ParseInt(line)
		}
	}
	top3 = adjustTop3(top3, current)
	fmt.Println("Total:", top3[0]+top3[1]+top3[2])
}

func adjustTop3(top3 []int, current int) []int {
	top3 = append(top3, current)
	sort.Ints(top3)
	return top3[1:4]
}
