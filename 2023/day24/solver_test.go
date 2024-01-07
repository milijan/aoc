package main

import (
	"testing"
)

func TestSolver1(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(puzzle []Hail, area []float64, expected int) {
		result := Solver1(puzzle, area)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test([]Hail{}, []float64{}, 0)
	test(puzzle[0:1], []float64{7, 27}, 0)
	test(puzzle[0:2], []float64{7, 27}, 1)
	test(puzzle[0:3], []float64{7, 27}, 2)
	test(puzzle, []float64{7, 27}, 2)
}

func TestSolver2(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(puzzle []Hail, area []float64, expected int) {
		result := Solver2(puzzle)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test([]Hail{}, []float64{}, 0)
	test(puzzle, []float64{7, 27}, 0)
}
