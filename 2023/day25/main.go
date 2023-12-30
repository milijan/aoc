package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func LoadPuzzle(filename string) map[string][]string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(filepath.Join(cwd, filename))
	if err != nil {
		fmt.Println(err)
		return map[string][]string{}
	}
	lines := strings.Split(string(content), "\n")

	puzzle := make(map[string][]string)

	for _, line := range lines {
		lst := strings.Split(line, ": ")
		node1 := lst[0]
		nodeList := lst[1]
		edge := ""
		for _, node2 := range strings.Split(nodeList, " ") {
			if node1 < node2 {
				edge = node1 + node2
			} else {
				edge = node2 + node1
			}
			if !slices.Contains(puzzle[node1], edge) {
				puzzle[node1] = append(puzzle[node1], edge)
			}
			if !slices.Contains(puzzle[node2], edge) {
				puzzle[node2] = append(puzzle[node2], edge)
			}
		}
	}

	return puzzle
}

func main() {
	puzzle := LoadPuzzle("puzzle.txt")

	sol1 := Solver1(puzzle)
	fmt.Println("-> puzzle 1:", sol1)

	sol2 := Solver2(puzzle)
	fmt.Println("-> puzzle 2:", sol2)
}
