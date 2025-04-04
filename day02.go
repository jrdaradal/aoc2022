package main

import (
	"fmt"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/io"
)

// SolutionA:	15	12740
// SolutionB:	12 	11980

type rps = [2]int

const (
	R int = 1
	P int = 2
	S int = 3
	L int = 0
	D int = 3
	W int = 6
)

var (
	winsOver = map[int]int{R: S, P: R, S: P}
	losesTo  = map[int]int{S: R, R: P, P: S}
)

func input02(full bool, mask map[string]int) []rps {
	lines, _ := io.ReadTextLines(getPath(2, full))
	return fn.Map(lines, func(line string) rps {
		p := fn.SpaceSplit(line)
		return rps{mask[p[0]], mask[p[1]]}
	})
}

func Day02A() {
	full := true
	mask := map[string]int{
		"A": R,
		"B": P,
		"C": S,
		"X": R,
		"Y": P,
		"Z": S,
	}
	total := 0
	for _, game := range input02(full, mask) {
		total += computeGameScore(game)
	}
	fmt.Println("Total:", total)
}

func Day02B() {
	full := true
	mask := map[string]int{
		"A": R,
		"B": P,
		"C": S,
		"X": L,
		"Y": D,
		"Z": W,
	}
	total := 0
	for _, game := range input02(full, mask) {
		total += coerceGameScore(game)
	}
	fmt.Println("Total:", total)
}

func computeGameScore(game rps) int {
	opp, you := game[0], game[1]
	score := you
	if opp == you {
		score += D
	} else if winsOver[you] == opp {
		score += W
	}
	return score
}

func coerceGameScore(cfg rps) int {
	opp, out := cfg[0], cfg[1]
	var you int
	if out == D {
		you = opp
	} else if out == W {
		you = losesTo[opp]
	} else if out == L {
		you = winsOver[opp]
	}
	return computeGameScore(rps{opp, you})
}
