package main

import (
	"testing"
)

func TestSolver1(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(p map[string][]string, expected int) {
		result := Solver1(p)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test(map[string][]string{}, 0)
	test(puzzle, 54)
}

func TestSolver2(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(p map[string][]string, expected int) {
		result := Solver2(p)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test(map[string][]string{}, 0)
	test(puzzle, 0)
}
