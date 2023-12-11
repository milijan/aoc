package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LoadPuzzle(filename string) Puzzle {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(filepath.Join(cwd, filename))
	if err != nil {
		fmt.Println(err)
		return Puzzle{}
	}

	c := string(content)

	cols := strings.IndexRune(c, '\n')
	rows := (len(c) - cols + 1) / cols

	puzzle := Puzzle{
		rows: rows,
		cols: cols,
		grid: make([][]byte, 0),
	}

	for i := 0; i < rows; i++ {
		row := content[i*(cols+1) : i*(cols+1)+cols]
		puzzle.grid = append(puzzle.grid, row)
	}

	return puzzle
}

func main() {
	puzzle := LoadPuzzle("puzzle.txt")

	sol1 := Solver1(puzzle)
	fmt.Println("-> puzzle 1:", sol1)

	sol2 := Solver2(puzzle, 1000000)
	fmt.Println("-> puzzle 2:", sol2)
}
