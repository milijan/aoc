package main

import (
	"testing"
)

func TestGetDestination(t *testing.T) {

	test := func(s int, m [][]int, expected int) {
		result := GetDestination(s, m)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test(0, [][]int{}, 0)
	test(20, [][]int{{10, 20, 0}}, 20)
	test(21, [][]int{{10, 20, 1}}, 21)
	test(20, [][]int{{10, 20, 1}, {40, 10, 5}}, 10)
	test(16, [][]int{{10, 20, 1}, {40, 10, 5}}, 16)
	test(13, [][]int{{10, 20, 1}, {40, 10, 5}}, 43)
}

func TestSolver1(t *testing.T) {

	puzz := LoadPuzzle("test.txt")

	test := func(p puzzle, expected int) {
		result := Solver1(p)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test(puzz, 35)
}

func TestSolver2(t *testing.T) {

	puzz := LoadPuzzle("test.txt")

	test := func(p puzzle, expected int) {
		result := Solver2(p)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test(puzz, 46)
}
