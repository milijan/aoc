package main

import (
	"fmt"
)

type Node struct {
	x, y     int
	depth    byte
	shape    byte
	children []Node
}

type Puzzle struct {
	rows  int
	cols  int
	x     int
	y     int
	moves [][]int
	grid  [][]byte
	pipes [][]byte
	roots []Node
}

func ValidMove(puzzle *Puzzle, node Node, x int, y int) bool {
	if x > puzzle.cols || y > puzzle.rows {
		return false
	}
	if x == node.x && y > node.y {
		// west -> east
		switch puzzle.grid[x][y] {
		case '-', '7', 'J':
			switch node.shape {
			case '-', 'L', 'F':
				return true
			}
		}
	} else {
		// west <- east
		switch puzzle.grid[x][y] {
		case '-', 'L', 'F':
			switch node.shape {
			case '-', '7', 'J':
				return true
			}
		}
	}

	if y == node.y && x > node.x {
		// north -> south
		switch puzzle.grid[x][y] {
		case '|', 'L', 'J':
			switch node.shape {
			case '|', 'F', '7':
				return true
			}
		}
	} else {
		// south -> north
		switch puzzle.grid[x][y] {
		case '|', 'F', '7':
			switch node.shape {
			case '|', 'J', 'L':
				return true
			}
		}
	}

	return false
}

func NextNode(puzzle *Puzzle, node *Node, x int, y int) int {
	n := Node{
		x: node.x,
		y: node.y,
	}
	p := Node{
		x: node.x + x,
		y: node.y + y,
	}
	if ValidMove(puzzle, n, p.x, p.y) {
		n.shape = puzzle.grid[n.x][n.y]
		p.shape = puzzle.grid[p.x][p.y]
		n.children = append(n.children, p)
		puzzle.roots = append(puzzle.roots, n)
	}
	return 0
}

func Solver1(puzzle Puzzle) int {
	fmt.Println(puzzle)

	for _, pipe := range []byte{'-', '|', '7', 'J', 'L', 'F'} {
		for _, move := range puzzle.moves {
			fmt.Println(pipe, move,
				ValidMove(&puzzle, Node{x: puzzle.x, y: puzzle.y, shape: pipe, depth: 0},
					puzzle.x+move[0], puzzle.y+move[1]))
		}
	}
	return 0
}

func Solver2(puzzle Puzzle) int {
	return 0
}
