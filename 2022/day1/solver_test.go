package main

import (
	"testing"
)

func TestSolver1(t *testing.T) {

	test := func(lines []string, expected int) {
		result := Solver1(lines)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test([]string{}, 0)
}

func TestSolver2(t *testing.T) {

	test := func(lines []string, expected int) {
		result := Solver2(lines)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test([]string{}, 0)
}

