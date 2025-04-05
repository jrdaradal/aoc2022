package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/io"
)

// SolutionA: 	95437		1182909
// SolutionB:

const (
	cmdCD string = "$ cd"
	cmdLS string = "$ ls"
	dir   string = "dir"
	glue  string = "/"
)

var FS map[string]*Item

type Item struct {
	name     string
	path     string
	parent   *Item
	children []*Item
	size     uint
}

func (i *Item) AddChild(c *Item) {
	i.children = append(i.children, c)
}

func (i *Item) ComputeSize() uint {
	var size uint = 0
	for _, child := range i.children {
		size += child.size
	}
	i.size = size
	return size
}

func (i Item) Type() string {
	if i.children == nil {
		return "FILE"
	} else {
		return "DIR "
	}
}

func (i Item) IsDir() bool {
	return i.children != nil
}

func input07(full bool) []string {
	lines, _ := io.ReadTextLines(getPath(7, full))
	return lines
}

func Day07A() {
	full := true
	buildFS(input07(full))

	var limit uint = 100_000
	var total uint = 0
	for _, item := range FS {
		if item.IsDir() && item.size <= limit {
			total += item.size
		}
	}
	fmt.Println("Total:", total)
}

func buildFS(lines []string) {
	var cwd *Item = nil
	FS = make(map[string]*Item)
	for _, line := range lines {
		if strings.HasPrefix(line, cmdCD) {
			name := strings.Fields(line)[2]
			if name == ".." {
				cwd = cwd.parent
			} else {
				cwd, _ = getDir(name, cwd)
			}
		} else if line == cmdLS {
			continue
		} else {
			p := strings.Fields(line)
			if p[0] == dir {
				item, isNew := getDir(p[1], cwd)
				if isNew {
					cwd.AddChild(item)
				}
			} else {
				item, isNew := getFile(p[1], fn.ParseInt(p[0]), cwd)
				if isNew {
					cwd.AddChild(item)
				}
			}
		}
	}

	dirPaths := make([]string, 0)
	for path, item := range FS {
		if item.IsDir() {
			dirPaths = append(dirPaths, path)
		}
	}
	slices.SortFunc(dirPaths, func(p1, p2 string) int {
		// Sort by depth DESC (leaf nodes first going towards root)
		s1 := len(strings.Split(p1, glue))
		s2 := len(strings.Split(p2, glue))
		return cmp.Compare(s2, s1)
	})
	for _, path := range dirPaths {
		FS[path].ComputeSize()
	}
}

func getDir(name string, parent *Item) (*Item, bool) {
	var path string
	if parent == nil {
		path = name
	} else {
		path = parent.path + name + glue
	}
	if item, ok := FS[path]; ok {
		return item, false
	}
	item := &Item{
		name:     name,
		path:     path,
		parent:   parent,
		children: make([]*Item, 0),
		size:     0,
	}
	FS[path] = item
	return item, true
}

func getFile(name string, size int, parent *Item) (*Item, bool) {
	path := parent.path + name
	if item, ok := FS[path]; ok {
		return item, false
	}
	item := &Item{
		name:     name,
		path:     path,
		parent:   parent,
		children: nil,
		size:     uint(size),
	}
	FS[path] = item
	return item, true
}
