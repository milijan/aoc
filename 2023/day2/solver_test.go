package main

import (
	"testing"
)

func TestSolver1(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(games []string, bag map[string]int, expected int) {
		result := Solver1(games, bag)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}
	bag := map[string]int{"red": 12, "green": 13, "blue": 14}

	test([]string{""}, map[string]int{}, 0)
	test([]string{puzzle[0]}, bag, 1)
	test([]string{puzzle[1]}, bag, 1)
	test([]string{puzzle[2]}, bag, 0)
	test([]string{puzzle[3]}, bag, 0)
	test([]string{puzzle[4]}, bag, 1)
	test(puzzle, bag, 8)
}

func TestSolver2(t *testing.T) {

	puzzle := LoadPuzzle("test.txt")

	test := func(games []string, expected int) {
		result := Solver2(games)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test([]string{}, 0)
	test([]string{puzzle[0]}, 48)
	test([]string{puzzle[1]}, 12)
	test([]string{puzzle[2]}, 1560)
	test([]string{puzzle[3]}, 630)
	test([]string{puzzle[4]}, 36)
	test(puzzle, 2286)
}
