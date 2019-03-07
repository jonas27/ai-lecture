package ticTacToe

import (
	"github.com/jonas27/a-star/astar"
)

var (
	name1 string
	name2 string
	m     *astar.Maze
)

func ini(maze *astar.Maze, n1 string, n2 string) {
	m = maze
	name1 = n1
	if name2 == "" {
		name2 = "pc"
	} else {
		name2 = n2
	}
}

func mark(r int, c int) bool {
	if m.Field(r, c).Value == 0 {
		m.Field(r, c).Value = 1
		m.Field(1, 1).Option = name1
		return true
	}
	return false
}

func solve() {
	if m.Field(1, 1).Value == 0 {
		m.Field(1, 1).Value = 1
		m.Field(1, 1).Option = name2
	}
}
