package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LoadPuzzle(filename string) []string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(filepath.Join(cwd, filename))
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	puzzle := strings.Split(string(content), "\n")

	return puzzle
}

func main() {
	puzzle := LoadPuzzle("puzzle.txt")

	sol1 := Solver1(puzzle)
	fmt.Println("-> puzzle 1:", sol1)

	sol2 := Solver2(puzzle)
	fmt.Println("-> puzzle 2:", sol2)
}
