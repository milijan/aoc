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
	test([]string{puzzle[0]}, 18)
	test([]string{puzzle[1]}, 28)
	test([]string{puzzle[2]}, 68)
	test(puzzle, 114)
}

func TestSolver2(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(lines []string, expected int) {
		result := Solver2(lines)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test([]string{}, 0)
	test([]string{puzzle[0]}, -3)
	test([]string{puzzle[1]}, 0)
	test([]string{puzzle[2]}, 5)
	test(puzzle, 2)
}
