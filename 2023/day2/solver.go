package main

import (
	"strconv"
	"strings"
)

// given a bag of red, blue, green cubes
// identify all the possible games and return the sum of all their ids
func Solver1(games []string, bag map[string]int) int {
	sum := 0
	for i, game := range games {
		possible := true
		_, sets, ok := strings.Cut(game, ": ")
		if !ok {
			continue
		}
		for _, set := range strings.Split(sets, "; ") {
			for _, draw := range strings.Split(set, ", ") {
				draw := strings.Split(draw, " ")
				cubes, err := strconv.Atoi(draw[0])
				if err != nil {
					panic(err)
				}
				color := draw[1]
				if bag[color] < cubes {
					possible = false
				}
			}
		}
		if possible {
			sum += i + 1
		}
	}

	return sum
}

func Solver2(lines []string) int {
	return 0
}
