package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func LoadPuzzle(filename string) []Hail {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(filepath.Join(cwd, filename))
	if err != nil {
		fmt.Println(err)
		return []Hail{}
	}

	puzzle := []Hail{}
	for _, line := range strings.Split(string(content), "\n") {
		state := strings.Split(line, "@")
		h := Hail{}
		for _, p := range strings.Split(state[0], ",") {
			d, _ := strconv.Atoi(strings.Trim(p, " "))
			h.position = append(h.position, int64(d))
		}
		for _, v := range strings.Split(state[1], ",") {
			d, _ := strconv.Atoi(strings.Trim(v, " "))
			h.velocity = append(h.velocity, int64(d))
		}
		puzzle = append(puzzle, h)
	}

	return puzzle
}

func main() {
	puzzle := LoadPuzzle("puzzle.txt")

	sol1 := Solver1(puzzle, []uint64{200000000000000, 400000000000000})
	fmt.Println("-> puzzle 1:", sol1)

	sol2 := Solver2(puzzle)
	fmt.Println("-> puzzle 2:", sol2)
}
