package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func LoadPuzzle(filename string) puzzle {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(filepath.Join(cwd, filename))
	if err != nil {
		fmt.Println(err)
		return puzzle{}
	}

	re := regexp.MustCompile(`[0-9]+`)
	param := re.FindAllString(string(content), -1)

	parseValues := func(a []string) []int {
		r := []int{}
		for _, p := range a {
			v, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			r = append(r, v)
		}
		return r
	}
	v := parseValues(param)

	n := len(v) / 2
	p := puzzle{}
	p.time = append(p.time, v[:n]...)
	p.distance = append(p.distance, v[n:2*n]...)

	return p
}

func main() {
	puzzle := LoadPuzzle("puzzle.txt")

	sol1 := Solver1(puzzle)
	fmt.Println("-> puzzle 1:", sol1)

	sol2 := Solver2(puzzle)
	fmt.Println("-> puzzle 2:", sol2)
}
