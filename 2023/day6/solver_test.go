package main

import (
	"testing"
)

func TestSolver1(t *testing.T) {

	p := LoadPuzzle("test.txt")

	test := func(p puzzle, expected int) {
		result := Solver1(p)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test(puzzle{}, 0)
	test(p, 288)
}

func TestSolver2(t *testing.T) {

	p := LoadPuzzle("test.txt")

	test := func(p puzzle, expected int) {
		result := Solver2(p)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test(puzzle{}, 0)
	test(p, 71503)
}
