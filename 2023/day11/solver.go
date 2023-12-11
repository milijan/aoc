package main

import (
	"math"
)

type Galaxy struct {
	x int
	y int
}

type Puzzle struct {
	rows int
	cols int
	grid [][]byte
}

func Distance(puzzle Puzzle, scale int, a, b Galaxy) int {
	k := 0
	if a.x < b.x {
		for i := a.x; i < b.x; i++ {
			if puzzle.grid[i][0] == '*' {
				k++
			}
		}
	} else {
		for i := b.x; i < a.x; i++ {
			if puzzle.grid[i][0] == '*' {
				k++
			}
		}
	}
	if a.y < b.y {
		for j := a.y; j < b.y; j++ {
			if puzzle.grid[0][j] == '*' {
				k++
			}
		}
	} else {
		for j := b.y; j < a.y; j++ {
			if puzzle.grid[0][j] == '*' {
				k++
			}
		}
	}
	d := int(math.Abs(float64(a.x-b.x))+math.Abs(float64(a.y-b.y))) + k*(scale-1)
	return d
}

func FindGalaxies(puzzle Puzzle) []Galaxy {
	galaxies := []Galaxy{}
	for i, row := range puzzle.grid {
		for j, c := range row {
			if c == '#' {
				galaxies = append(galaxies, Galaxy{x: i, y: j})
			}
		}
	}
	return galaxies
}

func Transpose(puzzle Puzzle) Puzzle {
	newPuzzle := Puzzle{
		rows: puzzle.cols,
		cols: puzzle.rows,
	}
	newPuzzle.grid = make([][]byte, newPuzzle.rows)
	for i := range newPuzzle.grid {
		newPuzzle.grid[i] = make([]byte, newPuzzle.cols)
	}
	for i := 0; i < newPuzzle.rows; i++ {
		for j := 0; j < newPuzzle.cols; j++ {
			newPuzzle.grid[i][j] = puzzle.grid[j][i]
		}
	}
	return newPuzzle
}

func Expand(puzzle Puzzle) Puzzle {
	newPuzzle := Puzzle{
		rows: puzzle.rows,
		cols: puzzle.cols,
	}
	r := make([]byte, puzzle.cols)
	for i := range r {
		r[i] = '*'
	}
	for _, row := range puzzle.grid {
		empty := true
		for _, c := range row {
			if c == '#' {
				empty = false
				break
			}
		}
		if empty {
			newPuzzle.grid = append(newPuzzle.grid, r)
		} else {
			newPuzzle.grid = append(newPuzzle.grid, row)
		}
	}
	return newPuzzle
}

func Solver1(puzzle Puzzle) int {
	newPuzzle := Expand(puzzle)
	newPuzzle = Transpose(newPuzzle)
	newPuzzle = Expand(newPuzzle)
	newPuzzle = Transpose(newPuzzle)

	galaxies := FindGalaxies(newPuzzle)

	sum := 0
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			if g1 != g2 {
				sum += Distance(newPuzzle, 2, g1, g2)
			}
		}

	}
	return sum
}

func Solver2(puzzle Puzzle, scale int) int {
	newPuzzle := Expand(puzzle)
	newPuzzle = Transpose(newPuzzle)
	newPuzzle = Expand(newPuzzle)
	newPuzzle = Transpose(newPuzzle)

	galaxies := FindGalaxies(newPuzzle)

	sum := 0
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			if g1 != g2 {
				sum += Distance(newPuzzle, scale, g1, g2)
			}
		}

	}
	return sum
}
