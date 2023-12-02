// combining the first digit and the last digit (in that order) to form a single two-digit number
// What is the sum of all of the calibration values?
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

	// only consider numbers
	puzzle1, err := os.ReadFile(filepath.Join(cwd, "puzzle.txt"))
	if err != nil {
		fmt.Println(err)
	}
	puzzle := strings.Split(string(puzzle1), "\n")
	sol1 := Solver1(puzzle)
	fmt.Println("-> puzzle 1:", sol1)

	// consider numbers and literals
	puzzle2, err := os.ReadFile(filepath.Join(cwd, "puzzle.txt"))
	if err != nil {
		fmt.Println(err)
	}
	puzzle = strings.Split(string(puzzle2), "\n")
	sol2 := Solver2(puzzle)
	fmt.Println("-> puzzle 2:", sol2)
}
