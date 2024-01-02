package main

import (
	"testing"
)

func TestSolver1(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(p Puzzle, expected int) {
		result := Solver1(p)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	//test(Puzzle{}, 0)
	test(puzzle, 136)
}

func TestSolver2(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(puzzle Puzzle, cycles int, expected int) {
		result := Solver2(puzzle, cycles)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test(Puzzle{}, 0, 0)
	test(puzzle, 1, 87)
	test(puzzle, 2, 69)
	test(puzzle, 3, 69)
	test(puzzle, 100, 68)
	test(puzzle, 1000, 64)
	test(puzzle, 10000, 69)
	test(puzzle, 100000, 65)
	test(puzzle, 1000000, 63)
	test(puzzle, 10000000, 69)
}
