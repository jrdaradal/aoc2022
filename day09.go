package main

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/io"
)

// SolutionA: 13	6339
// SolutionB: 36	2541

func input09(full bool) []coords {
	lines, _ := io.ReadTextLines(getPath(9, full))
	return fn.Map(lines, func(line string) coords {
		p := strings.Fields(line)
		d := fn.ParseInt(p[1])
		switch p[0] {
		case "U":
			return coords{d, 0}
		case "D":
			return coords{-d, 0}
		case "L":
			return coords{0, -d}
		case "R":
			return coords{0, d}
		default:
			return coords{0, 0}
		}
	})
}

func Day09A() {
	full := true
	head, tail := coords{0, 0}, coords{0, 0}
	visited := make(map[coords]bool)
	visited[tail] = true
	for _, delta := range input09(full) {
		head, tail = moveRope(head, tail, delta, visited)
	}
	fmt.Println("Visited:", len(visited))
}

func Day09B() {
	full := true
	tail := 9
	pos := make([]coords, tail+1)
	for i := range tail + 1 {
		pos[i] = coords{0, 0}
	}
	visited := make(map[coords]bool)
	visited[pos[tail]] = true
	for _, delta := range input09(full) {
		pos = moveChain(pos, delta, visited)
	}
	fmt.Println("Visited:", len(visited))

}

func moveRope(head, tail, delta coords, visited map[coords]bool) (coords, coords) {
	steps, idx, factor := unpackDelta(delta)
	for range steps {
		head[idx] += 1 * factor
		if !isAdjacent(head, tail) {
			tail = follow(head, tail)
			visited[tail] = true
		}
	}
	return head, tail
}

func moveChain(pos []coords, delta coords, visited map[coords]bool) []coords {
	tail := len(pos) - 1
	steps, idx, factor := unpackDelta(delta)
	for range steps {
		pos[0][idx] += 1 * factor
		for i := 1; i <= tail; i++ {
			if !isAdjacent(pos[i-1], pos[i]) {
				pos[i] = follow(pos[i-1], pos[i])
			}
		}
		visited[pos[tail]] = true
	}
	return pos
}

func follow(c1, c2 coords) coords {
	dy := c1[0] - c2[0]
	dx := c1[1] - c2[1]
	if dx > 0 {
		c2[1] += 1
	} else if dx < 0 {
		c2[1] -= 1
	}
	if dy > 0 {
		c2[0] += 1
	} else if dy < 0 {
		c2[0] -= 1
	}
	return c2
}

func isAdjacent(c1, c2 coords) bool {
	dy := fn.Abs(c1[0] - c2[0])
	dx := fn.Abs(c1[1] - c2[1])
	return dy <= 1 && dx <= 1
}

func unpackDelta(delta coords) (int, int, int) {
	dy, dx := delta[0], delta[1]
	var steps, idx, factor int
	if dy == 0 && dx > 0 { // right
		steps, idx, factor = dx, 1, 1
	} else if dy == 0 && dx < 0 { // left
		steps, idx, factor = -dx, 1, -1
	} else if dx == 0 && dy > 0 { // up
		steps, idx, factor = dy, 0, 1
	} else if dx == 0 && dy < 0 { // down
		steps, idx, factor = -dy, 0, -1
	}
	return steps, idx, factor
}
