package main

import (
	"fmt"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/io"
)

// SolutionA:	2	485
// SolutionB:	4	857

type rangepair = [2][2]int

func input04(full bool) []rangepair {
	lines, _ := io.ReadTextLines(getPath(4, full))
	return fn.Map(lines, func(line string) rangepair {
		p := fn.CleanSplit(line, ",")
		return rangepair{parseRange(p[0]), parseRange(p[1])}
	})
}

func Day04A() {
	full := true
	count := 0
	for _, pair := range input04(full) {
		if isSupersetPair(pair) {
			count += 1
		}
	}
	fmt.Println("Count:", count)
}

func Day04B() {
	full := true
	count := 0
	for _, pair := range input04(full) {
		if isOverlappingPair(pair) {
			count += 1
		}
	}
	fmt.Println("Count:", count)
}

func parseRange(text string) [2]int {
	p := fn.Map(fn.CleanSplit(text, "-"), fn.ParseInt)
	return [2]int{p[0], p[1]}
}

func isSupersetPair(pair rangepair) bool {
	s1, s2 := pair[0], pair[1]
	return isSuperset(s1, s2) || isSuperset(s2, s1)
}

func isSuperset(r1, r2 [2]int) bool {
	s1, e1 := r1[0], r1[1]
	s2, e2 := r2[0], r2[1]
	return s1 <= s2 && e2 <= e1
}

func isOverlappingPair(pair rangepair) bool {
	r1, r2 := pair[0], pair[1]
	s1, e1 := r1[0], r1[1]
	s2, e2 := r2[0], r2[1]
	if s1 < s2 {
		return s2 <= e1
	} else {
		return s1 <= e2
	}
}
