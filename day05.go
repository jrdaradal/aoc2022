package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/io"
)

// SolutionA: 	CMZ		HNSNMTLHQ
// SolutionB: 	MCD		RNLFDJMCT

type config05 struct {
	stacks [][]rune
	moves  [][3]int // count, idx1, idx2
}

type TransferFn = func(int, []rune, []rune) ([]rune, []rune)

func input05(full bool) *config05 {
	cfg := &config05{
		stacks: nil,
		moves:  make([][3]int, 0),
	}
	stackMode := true
	lines, _ := io.ReadTextFile(getPath(5, full))
	for _, line := range strings.Split(lines, "\n") {
		cleanLine := strings.TrimSpace(line)
		if cleanLine == "" {
			stackMode = false
			continue
		}
		if stackMode {
			if !strings.HasPrefix(cleanLine, "[") {
				continue
			}
			if cfg.stacks == nil {
				count := len(line) / 4
				cfg.stacks = make([][]rune, count)
				for i := range count {
					cfg.stacks[i] = make([]rune, 0)
				}
			}
			for i, char := range line {
				if i%4 != 1 || char == ' ' {
					continue
				}
				idx := i / 4
				cfg.stacks[idx] = append(cfg.stacks[idx], char)
			}
		} else {
			cfg.moves = append(cfg.moves, parseMove(line))
		}
	}
	return cfg
}

func Day05A() {
	full := true
	processMoves(input05(full), transferReverse)
}

func Day05B() {
	full := true
	processMoves(input05(full), transferAsIs)
}

func parseMove(line string) [3]int {
	p := fn.CleanSplit(line, "from")
	count := fn.ParseInt(fn.SpaceSplit(p[0])[1])
	i := fn.Map(fn.CleanSplit(p[1], "to"), fn.ParseInt)
	return [3]int{count, i[0] - 1, i[1] - 1}
}

func processMoves(cfg *config05, transferFn TransferFn) {
	s := cfg.stacks
	for _, move := range cfg.moves {
		count, idx1, idx2 := move[0], move[1], move[2]
		s[idx1], s[idx2] = transferFn(count, s[idx1], s[idx2])
	}
	top := make([]rune, len(s))
	for i, stack := range s {
		top[i] = stack[0]
	}
	fmt.Println("Top:", string(top))
}

func transferReverse(count int, s1, s2 []rune) ([]rune, []rune) {
	move := fn.CopySlice(s1[:count])
	slices.Reverse(move)
	n1 := s1[count:]
	n2 := append(move, s2...)
	return n1, n2
}

func transferAsIs(count int, s1, s2 []rune) ([]rune, []rune) {
	move := fn.CopySlice(s1[:count])
	n1 := s1[count:]
	n2 := append(move, s2...)
	return n1, n2
}
