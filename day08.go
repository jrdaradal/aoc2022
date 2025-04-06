package main

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/io"
)

// SolutionA:	21	1785
// SolutionB:	8	345168

type coords = [2]int

func input08(full bool) [][]int {
	lines, _ := io.ReadTextLines(getPath(8, full))
	return fn.Map(lines, func(line string) []int {
		return fn.Map(strings.Split(line, ""), fn.ParseInt)
	})
}

func Day08A() {
	full := true
	countVisible(input08(full))
}

func Day08B() {
	full := true
	findBestScore(input08(full))
}

func countVisible(grid [][]int) {
	nRows, nCols := len(grid), len(grid[0])

	visible := make(map[coords]bool)
	for row := 1; row < nRows-1; row++ {
		for col := 1; col < nCols-1; col++ {
			visible[coords{row, col}] = false
		}
	}

	for row := 1; row < nRows-1; row++ {
		for col := 1; col < nCols-1; col++ {
			checkRowVisible(grid, visible, row, col)
			checkColVisible(grid, visible, row, col)
		}
	}

	count := 0
	for _, ok := range visible {
		if ok {
			count += 1
		}
	}

	numEdges := (2 * nCols) + (2 * (nRows - 2))
	fmt.Println("Visible:", numEdges+count)
}

func checkRowVisible(grid [][]int, visible map[coords]bool, row, col int) {
	c := coords{row, col}
	if visible[c] {
		return
	}
	isValid := func(x int) bool {
		return x < grid[row][col]
	}
	ok := fn.All(grid[row][:col], isValid)
	if ok {
		visible[c] = true
		return
	}
	ok = fn.All(grid[row][col+1:], isValid)
	visible[c] = ok
}

func checkColVisible(grid [][]int, visible map[coords]bool, row, col int) {
	c := coords{row, col}
	if visible[c] {
		return
	}
	isValid := func(x int) bool {
		return x < grid[row][col]
	}
	above := make([]int, 0)
	for r := 0; r < row; r++ {
		above = append(above, grid[r][col])
	}
	ok := fn.All(above, isValid)
	if ok {
		visible[c] = true
		return
	}
	below := make([]int, 0)
	for r := row + 1; r < len(grid); r++ {
		below = append(below, grid[r][col])
	}
	ok = fn.All(below, isValid)
	visible[c] = ok
}

func findBestScore(grid [][]int) {
	nRows, nCols := len(grid), len(grid[0])
	best := 0
	for row := 1; row < nRows-1; row++ {
		for col := 1; col < nCols-1; col++ {
			best = max(best, computeScore(grid, row, col))
		}
	}
	fmt.Println("Best:", best)
}

func computeScore(grid [][]int, row, col int) int {
	nRows, nCols := len(grid), len(grid[0])
	value := grid[row][col]
	// Up
	n := 0
	for r := row - 1; r >= 0; r-- {
		n++
		if grid[r][col] >= value {
			break
		}
	}
	// Down
	s := 0
	for r := row + 1; r < nRows; r++ {
		s++
		if grid[r][col] >= value {
			break
		}
	}
	// Left
	w := 0
	for c := col - 1; c >= 0; c-- {
		w++
		if grid[row][c] >= value {
			break
		}
	}
	// Right
	e := 0
	for c := col + 1; c < nCols; c++ {
		e++
		if grid[row][c] >= value {
			break
		}
	}

	return n * e * w * s
}
