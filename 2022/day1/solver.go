package main

import (
	"fmt"
	"strconv"
)

func Solver1(lines []string) int {
	sum := 0
	val := 0
	maxsum := 0
	var err error

	for _, line := range lines {
		if len(line) > 0 {
			val, err = strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			} else {
				sum += val
			}
		} else {
			if sum > maxsum {
				maxsum = sum
			}
			sum = 0
		}
	}

	return maxsum
}

func Solver2(lines []string) int {
	return 0
}
