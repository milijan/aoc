package main

import (
	"testing"
)

func TestSolver1(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(p string, expected int) {
		result := Solver1(p)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	//test([]string{}, 0)
	test(puzzle, 0)
	//test([]string{puzzle[0], puzzle[1]}, 467)
	//test([]string{puzzle[0], puzzle[1], puzzle[2]}, 467+35)
	// test([]string{puzzle[1], puzzle[2], puzzle[3]}, 0)
	// test([]string{puzzle[2], puzzle[3], puzzle[4]}, 35)
	// test([]string{puzzle[3], puzzle[4], puzzle[5]}, 58)
	// test([]string{puzzle[4], puzzle[5], puzzle[6]}, 58)
	// test([]string{puzzle[5], puzzle[6], puzzle[7]}, 58+755)
	// test(puzzle, 4361)
}

func TestSolver2(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(p string, expected int) {
		result := Solver2(p)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test(puzzle, 0)
}
