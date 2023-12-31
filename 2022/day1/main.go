package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	puzzle, err := os.ReadFile(filepath.Join(cwd, "puzzle.txt"))
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(puzzle), "\n")

	sol1 := Solver1(lines)
	fmt.Println("-> puzzle 1:", sol1)

	sol2 := Solver2(lines)
	fmt.Println("-> puzzle 2:", sol2)
}
