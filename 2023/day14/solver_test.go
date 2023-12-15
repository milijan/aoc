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
	test(puzzle, 136)
}

func TestSolver2(t *testing.T) {

	puzzle := LoadPuzzle("test2.txt")

	test := func(puzzle Puzzle, cycles int, expected int) {
		result := Solver2(puzzle, cycles)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test(Puzzle{}, 0, 0)
	test(puzzle, 1, 115)
	test(puzzle, 2, 121)
	test(puzzle, 3, 127)
	test(puzzle, 1000000, 118)
}
