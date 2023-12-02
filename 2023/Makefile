SOURCES = $(wildcard */main.go)
SUBDIRS = $(dir $(SOURCES))
PUZZLES = $(SUBDIRS:/=)
TESTS = $(subst day,,$(PUZZLES))

define MAIN_TEMPLATE
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// only consider numbers
	puzzle1, err := os.ReadFile(filepath.Join(cwd, "puzzle.txt"))
	if err != nil {
		fmt.Println(err)
	}
	puzzle := strings.Split(string(puzzle1), "\n")
	sol1 := Solver1(puzzle)
	fmt.Println("-> puzzle 1:", sol1)

	// consider numbers and literals
	puzzle2, err := os.ReadFile(filepath.Join(cwd, "puzzle.txt"))
	if err != nil {
		fmt.Println(err)
	}
	puzzle = strings.Split(string(puzzle2), "\n")
	sol2 := Solver2(puzzle)
	fmt.Println("-> puzzle 2:", sol2)
}
endef
export MAIN_TEMPLATE

define SOLVER_TEMPLATE
package main

func Solver1(lines []string) int {
	return 0
}

func Solver2(lines []string) int {
	return 0
}
endef
export SOLVER_TEMPLATE

define TEST_TEMPLATE
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

endef
export TEST_TEMPLATE

# generate test file if missing
.PRECIOUS: %/solver_test.go
%/solver_test.go:
	@mkdir -p $(@D)
	@echo generating test template
	@echo "$$TEST_TEMPLATE" > $@

# generate puzzle solver if missing
.PRECIOUS: %/solver.go
%/solver.go: 
	@mkdir -p $(@D)
	@echo generating puzzle solver template
	@echo "$$SOLVER_TEMPLATE" > $@
	@echo generating main program 
	@echo "$$MAIN_TEMPLATE" > $(@D)/main.go
	@echo create puzzle file
	@touch $(@D)/puzzle.txt

# solve puzzle
$(PUZZLES): %: %/solver.go %/main.go
	@echo Solving $@
	@cd $(<D); go run .

# test solver
test%: day%/solver_test.go day%/solver.go
	@cd $(<D); go test

# initialise puzzle
init%: day%/solver.go day%/solver_test.go
	@echo $(<D) initialised

.PHONY: $(PUZZLES) 
all: $(PUZZLES)

