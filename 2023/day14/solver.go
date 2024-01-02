package main

import (
	"fmt"
	"slices"
)

type Stones struct {
	x int
	y int
}

type Puzzle struct {
	rows       int
	cols       int
	load       int
	grid       [][]byte
	emptyGrid  [][]byte
	fixedNorth []Stones
	fixedWest  []Stones
	fixedSouth []Stones
	fixedEast  []Stones
}

func FilterStones(stones *[]Stones, sortbyX bool, i int) *[]Stones {
	filtered := []Stones{}
	for _, s := range *stones {
		if sortbyX && i == s.x {
			filtered = append(filtered, s)
		} else if i == s.y {
			filtered = append(filtered, s)
		}
	}
	return &filtered
}

func SortStones(stones []Stones, XFirst bool, ascend bool) []Stones {
	s := stones
	// first axis always asending
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if XFirst {
				if s[i].x > s[j].x {
					s[i], s[j] = s[j], s[i]
				}
			} else {
				if s[i].y > s[j].y {
					s[i], s[j] = s[j], s[i]
				}
			}

		}
	}
	// 2nd axis respecting order
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if XFirst { // X First
				if s[i].x == s[j].x {
					if ascend {
						if s[i].y > s[j].y {
							s[i], s[j] = s[j], s[i]
						}
					} else {
						if s[i].y < s[j].y {
							s[i], s[j] = s[j], s[i]
						}
					}
				} else {
					break
				}
			} else { // Y First
				if s[i].y == s[j].y {
					if ascend {
						if s[i].x > s[j].x {
							s[i], s[j] = s[j], s[i]
						}
					} else {
						if s[i].x < s[j].x {
							s[i], s[j] = s[j], s[i]
						}
					}
				} else {
					break
				}
			}
		}
	}
	return stones
}

func FindStones(grid [][]byte, stone byte) []Stones {
	stones := []Stones{}
	for i, row := range grid {
		for j, c := range row {
			if c == stone {
				stones = append(stones, Stones{x: i, y: j})
			}
		}
	}
	return stones
}

func RotateStones(stones []Stones, rows int, cols int, clockwise bool) {
	for i := 0; i < len(stones); i++ {
		temp := stones[i].x
		if clockwise {
			stones[i].x = stones[i].y
			stones[i].y = cols - temp - 1
		} else {
			stones[i].x = rows - stones[i].y - 1
			stones[i].y = temp
		}
	}
}

func RotateGrid(grid [][]byte, clockwise bool) [][]byte {
	rows := len(grid)
	rotatedGrid := [][]byte{}
	if rows > 0 {
		cols := len(grid[0])
		rotatedGrid = make([][]byte, rows)
		for i := range grid {
			rotatedGrid[i] = make([]byte, cols)
		}
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if clockwise {
					rotatedGrid[i][j] = grid[rows-j-1][i]
				} else {
					rotatedGrid[i][j] = grid[j][cols-i-1]
				}
			}
		}
	}
	return rotatedGrid
}

func Transpose(puzzle *Puzzle) Puzzle {
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

func MoveNorth(stones []Stones, walls []Stones) {
	stop := 0
	y := 0
	SortStones(stones, false, true) // Y then X ascending
	for i, stone := range stones {
		if stone.y != y { // change of column
			stop = 0
			y = stone.y
		}
		for _, wall := range walls {
			if stone.y < wall.y {
				break
			} else if stone.y == wall.y {
				if stone.x > wall.x {
					stop = max(stop, wall.x+1)
				}
			}
		}
		stones[i].x = stop
		stop++
	}
}

func MoveSouth(stones []Stones, walls []Stones, rows int) {
	stop := rows - 1
	y := 0
	SortStones(stones, false, false) // Y then X descending
	for i, stone := range stones {
		if stone.y != y { // change of column
			stop = rows - 1
			y = stone.y
		}
		for _, wall := range walls {
			if stone.y < wall.y { // no more walls
				break
			} else if stone.y == wall.y {
				if stone.x < wall.x {
					stop = min(stop, wall.x-1)
				}
			}
		}
		stones[i].x = stop
		stop--
	}
}

func MoveWest(stones []Stones, walls []Stones) {
	stop := 0
	x := 0
	SortStones(stones, true, true) // X then Y ascending
	for i, stone := range stones {
		if stone.x != x { // change of row
			stop = 0
			x = stone.x
		}
		for _, wall := range walls {
			if stone.x < wall.x { // no more walls
				break
			} else if stone.x == wall.x {
				if stone.y > wall.y {
					stop = max(stop, wall.y+1)
				}
			}
		}
		stones[i].y = stop
		stop++
	}
}

func MoveEast(stones []Stones, walls []Stones, cols int) {
	stop := cols - 1
	x := 0
	SortStones(stones, true, false) // X then Y descending
	for i, stone := range stones {
		if stone.x != x { // change of row
			stop = cols - 1
			x = stone.x
		}
		for _, wall := range walls {
			if stone.x < wall.x { // no more walls
				break
			} else if stone.x == wall.x {
				if stone.y < wall.y {
					stop = min(stop, wall.y-1)
				}
			}
		}
		stones[i].y = stop
		stop--
	}
}

func CleanGrid(grid [][]byte) {
	for _, r := range grid {
		for i := 0; i < len(r); i++ {
			if r[i] == 'O' {
				r[i] = '.'
			}
		}
	}
}

func PlaceStones(grid [][]byte, stones []Stones) [][]byte {
	rows := len(grid)
	newGrid := [][]byte{}
	if rows > 0 {
		newGrid = make([][]byte, rows)
		for i, r := range grid {
			newGrid[i] = slices.Clone(r)
		}

		for _, s := range stones {
			newGrid[s.x][s.y] = 'O'
		}
	}
	return newGrid
}

func PrintTable(grid [][]byte) {
	for _, r := range grid {
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

func GetLoadNorth(stones *[]Stones, puzzle *Puzzle) int {
	load := 0
	for _, s := range *stones {
		load += puzzle.rows - s.x
	}
	return load
}

func AutoCorrelation(timeseries []int) (int, float64) {
	n := len(timeseries)
	k := n / 2
	max := 0.0
	delay := 0
	mean := 0.0

	for i := 0; i < n; i++ {
		mean += float64(timeseries[i])
	}
	mean /= float64(n)

	for j := 0; j < k; j++ {
		a := 0.0
		for i := j; i < n-j; i++ {
			a += (float64(timeseries[i]) - mean) * (float64(timeseries[i+j]) - mean)
		}
		if j > 5 && a > max {
			max = a
			delay = j
		}
	}
	return delay, max
}

func InitPuzzle(puzzle *Puzzle) {
	puzzle.emptyGrid = make([][]byte, len(puzzle.grid))
	for i, r := range puzzle.grid {
		puzzle.emptyGrid[i] = slices.Clone(r)
	}
	CleanGrid(puzzle.emptyGrid)

	puzzle.fixedNorth = FindStones(puzzle.emptyGrid, '#')
	SortStones(puzzle.fixedNorth, false, true) // Y then X ascending
	puzzle.fixedWest = slices.Clone(puzzle.fixedNorth)
	SortStones(puzzle.fixedWest, true, true) // X then Y ascending
	puzzle.fixedSouth = slices.Clone(puzzle.fixedNorth)
	SortStones(puzzle.fixedSouth, false, false) // Y then X descending
	puzzle.fixedEast = slices.Clone(puzzle.fixedNorth)
	SortStones(puzzle.fixedEast, true, false) // X then Y descending
}

func Solver1(puzzle Puzzle) int {
	InitPuzzle(&puzzle)
	stones := FindStones(puzzle.grid, 'O')

	MoveNorth(stones, puzzle.fixedNorth)

	return GetLoadNorth(&stones, &puzzle)
}

func Solver2(puzzle Puzzle, cycles int) int {
	InitPuzzle(&puzzle)
	stones := FindStones(puzzle.grid, 'O')

	load := 0
	period := 0
	n := 1000
	history := make([]int, n)
	prediction := make([]bool, n)

	for i := 0; i < cycles; i++ {

		// check our predictions
		ok := true
		for _, p := range prediction {
			ok = ok && p
		}

		if ok {
			j := (i - period) % n
			k := (cycles - i - 1) % period
			if j+period > n {
				periodicSlice := history[j-period : j]
				load = periodicSlice[k]
			} else {
				periodicSlice := history[j : j+period]
				load = periodicSlice[k]
			}
			break
		} else {
			MoveNorth(stones, puzzle.fixedNorth)
			MoveWest(stones, puzzle.fixedWest)
			MoveSouth(stones, puzzle.fixedSouth, puzzle.rows)
			MoveEast(stones, puzzle.fixedEast, puzzle.cols)

			// calculate load and store in history array
			load = GetLoadNorth(&stones, &puzzle)
			history[i%n] = load

			if i > n {
				// look for repeating pattern
				period, _ = AutoCorrelation(history)
				if period > 0 {
					prediction[i%n] = load == history[(i-period)%n]
				}
			}
		}
	}

	return load
}
