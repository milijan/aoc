package main

import (
	"fmt"
)

type Stones struct {
	x int
	y int
}

type Puzzle struct {
	rows int
	cols int
	load int
	grid [][]byte
}

func FilterStones(stones []Stones, sortbyX bool, i int) []Stones {
	filtered := []Stones{}
	for _, s := range stones {
		if sortbyX && i == s.x {
			filtered = append(filtered, s)
		} else if i == s.y {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func SortStones(stones []Stones, sortbyX bool) []Stones {
	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if sortbyX && stones[i].x > stones[j].x {
				stones[i], stones[j] = stones[j], stones[i]
			} else if stones[i].y > stones[j].y {
				stones[i], stones[j] = stones[j], stones[i]
			}
		}
	}
	return stones
}

func FindStones(puzzle Puzzle, stone byte) []Stones {
	stones := []Stones{}
	for i, row := range puzzle.grid {
		for j, c := range row {
			if c == stone {
				stones = append(stones, Stones{x: i, y: j})
			}
		}
	}
	return SortStones(stones, true)
}

func Rotate(puzzle Puzzle, clockwise bool) Puzzle {
	newPuzzle := Puzzle{
		rows: puzzle.cols,
		cols: puzzle.rows,
		load: puzzle.load,
	}
	newPuzzle.grid = make([][]byte, newPuzzle.rows)
	for i := range newPuzzle.grid {
		newPuzzle.grid[i] = make([]byte, newPuzzle.cols)
	}
	for i := 0; i < newPuzzle.rows; i++ {
		for j := 0; j < newPuzzle.cols; j++ {
			if clockwise {
				newPuzzle.grid[i][j] = puzzle.grid[puzzle.rows-j-1][i]
			} else {
				newPuzzle.grid[i][j] = puzzle.grid[j][puzzle.cols-i-1]
			}
		}
	}
	return newPuzzle
}

func Transpose(puzzle Puzzle) Puzzle {
	newPuzzle := Puzzle{
		rows: puzzle.cols,
		cols: puzzle.rows,
		load: puzzle.load,
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

func MoveStones(tiltedStones []Stones, rollingStones []Stones, fixedStones []Stones) []Stones {
	wall := 0
	for _, s := range rollingStones {
		for _, f := range fixedStones {
			if s.x > f.x {
				wall = max(wall, f.x+1)
			}
		}
		tiltedStones = append(tiltedStones, Stones{x: wall, y: s.y})
		wall++
	}
	return tiltedStones
}

func TiltTable(puzzle Puzzle) Puzzle {
	fixedStones := FindStones(puzzle, '#')
	rollingStones := FindStones(puzzle, 'O')

	tiltedStones := []Stones{}

	for j := 0; j < puzzle.cols; j++ {
		fixedStonesCol := FilterStones(fixedStones, false, j)
		rollingStonesCol := FilterStones(rollingStones, false, j)
		tiltedStones = MoveStones(tiltedStones, rollingStonesCol, fixedStonesCol)
	}

	newPuzzle := Puzzle{
		rows: puzzle.rows,
		cols: puzzle.cols,
		load: 0,
		grid: make([][]byte, puzzle.cols),
	}
	for i := range newPuzzle.grid {
		newPuzzle.grid[i] = make([]byte, puzzle.rows)
	}

	for _, s := range fixedStones {
		newPuzzle.grid[s.x][s.y] = '#'
	}

	for _, s := range tiltedStones {
		newPuzzle.load += puzzle.rows - s.x
		newPuzzle.grid[s.x][s.y] = 'O'
	}

	return newPuzzle
}

func PrintTable(puzzle Puzzle) {
	for _, r := range puzzle.grid {
		for _, c := range r {
			if c == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func Solver1(puzzle Puzzle) int {

	p := TiltTable(puzzle)

	return p.load
}

func Solver2(puzzle Puzzle, cycles int) int {

	for i := 0; i < cycles; i++ {
		// cycle
		for j := 0; j < 4; j++ {
			puzzle = TiltTable(puzzle)
			puzzle = Rotate(puzzle, true)
		}
	}

	return puzzle.load
}
