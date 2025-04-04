package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Day01A()
	// Day01B()
	Day02A()

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
