package main

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/io"
)

// SolutionA: 13	6339

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

func moveRope(head, tail, delta coords, visited map[coords]bool) (coords, coords) {
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
	for range steps {
		head[idx] += 1 * factor
		if !isAdjacent(head, tail) {
			tail = moveTail(head, tail)
			visited[tail] = true
		}
	}
	return head, tail
}

func moveTail(head, tail coords) coords {
	dy := head[0] - tail[0]
	dx := head[1] - tail[1]
	if dx > 0 {
		tail[1] += 1
	} else if dx < 0 {
		tail[1] -= 1
	}
	if dy > 0 {
		tail[0] += 1
	} else if dy < 0 {
		tail[0] -= 1
	}
	return tail
}

func isAdjacent(head, tail coords) bool {
	dy := fn.Abs(head[0] - tail[0])
	dx := fn.Abs(head[1] - tail[1])
	return dy <= 1 && dx <= 1
}
