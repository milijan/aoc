package main

import (
	"testing"
)

func TestSolver1(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(lines []string, expected int) {
		result := Solver1(lines)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test([]string{}, 0)
	test(puzzle, 13)
}

func TestSolver2(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(lines []string, expected int) {
		result := Solver2(lines)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	//test([]string{}, 0)
	//test([]string{puzzle[0]}, 1)
	test([]string{puzzle[0], puzzle[1]}, 3)
	test(puzzle, 30)
}
