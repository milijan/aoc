package main

import (
	"testing"
)

func TestSolver1(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(puzzle Puzzle, expected int) {
		result := Solver1(puzzle)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test(Puzzle{}, 0)
	test(puzzle, 0)
}

func TestSolver2(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(puzzle Puzzle, expected int) {
		result := Solver2(puzzle)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test(Puzzle{}, 0)
	test(puzzle, 0)
}
