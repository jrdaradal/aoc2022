package main

import (
	"fmt"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/io"
)

// SolutionA:	2	485
// SolutionB:	4	857

type setrange = [2]int
type rangepair = [2]setrange

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

func parseRange(text string) setrange {
	p := fn.Map(fn.CleanSplit(text, "-"), fn.ParseInt)
	return setrange{p[0], p[1]}
}

func isSupersetPair(p rangepair) bool {
	return isSuperset(p[0], p[1]) || isSuperset(p[1], p[0])
}

func isSuperset(r1, r2 [2]int) bool {
	return r1[0] <= r2[0] && r2[1] <= r1[1]
}

func isOverlappingPair(p rangepair) bool {
	s1, e1 := p[0][0], p[0][1]
	s2, e2 := p[1][0], p[1][1]
	if s1 < s2 {
		return s2 <= e1
	} else {
		return s1 <= e2
	}
}
