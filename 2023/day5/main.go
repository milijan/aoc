package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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
	sections := strings.Split(string(content), ":")

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
	parsePairs := func(a []string) [][]int {
		r := parseValues(a)
		m := [][]int{}
		for i := 0; i < len(r)/2; i++ {
			m = append(m, r[2*i:2*i+2])
		}
		return m
	}
	parseMaps := func(a []string) [][]int {
		r := parseValues(a)
		m := [][]int{}
		for i := 0; i < len(r)/3; i++ {
			m = append(m, r[3*i:3*i+3])
		}
		return m
	}

	p := puzzle{}

	re := regexp.MustCompile(`[0-9]+`)
	for i, section := range sections {
		param := re.FindAllString(section, -1)
		switch i {
		case 1:
			p.seeds = parseValues(param)
			p.seedPairs = parsePairs(param)
		case 2:
			p.seedToSoil = parseMaps(param)
		case 3:
			p.soilToFert = parseMaps(param)
		case 4:
			p.fertToWater = parseMaps(param)
		case 5:
			p.waterToLight = parseMaps(param)
		case 6:
			p.lightToTemp = parseMaps(param)
		case 7:
			p.tempToHum = parseMaps(param)
		case 8:
			p.humToLoc = parseMaps(param)
		}
	}

	return p
}

func main() {
	puzzle := LoadPuzzle("puzzle.txt")

	sol1 := Solver1(puzzle)
	fmt.Println("-> puzzle 1:", sol1)

	sol2 := Solver2(puzzle)
	fmt.Println("-> puzzle 2:", sol2)
}
