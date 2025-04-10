package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Day01A()
	// Day01B()
	// Day02A()
	// Day02B()
	// Day03A()
	// Day03B()
	// Day04A()
	// Day04B()
	// Day05A()
	// Day05B()
	// Day06A()
	// Day06B()
	// Day07A()
	// Day07B()]
	// Day08A()
	// Day08B()
	// Day09A()
	Day09B()

	fmt.Printf("\nTime: %v\n", time.Since(now))
}

func getPath(day int, full bool) string {
	var suffix string
	if full {
		suffix = "test"
	} else {
		suffix = "sample"
	}
	return fmt.Sprintf("data/%.2d_%s.txt", day, suffix)
}
