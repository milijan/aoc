package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Valid(line string) bool {

	parts := strings.Split(line, " ")
	springs := strings.Split(parts[0], ".")
	inventory := strings.Split(parts[1], ",")

	if len(springs) == len(inventory) {
		for i, spring := range springs {
			count, _ := strconv.Atoi(inventory[i])
			if len(spring) != count {
				return true
			}
		}
	}

	return false
}

func Swap(s string, i, j int) string {
	a := []rune(s)
	a[i], a[j] = a[j], a[i]
	return string(a)
}

func Permutation(line string, i int) {
	n := len(line)
	if i == (n - 1) {
		fmt.Println(line)
		return
	}

	for j := i; j < n; j++ {
		if j > i && line[i] == line[j] {
			continue
		}
		if j > i && line[j-1] == line[j] {
			continue
		}

		Swap(line, i, j)
		Permutation(line, i+1)
		Swap(line, i, j)
	}
}

func ListOptions(line string) []string {
	options := []string{}

	count := strings.Count(line, "?")
	options = append(options, strings.ReplaceAll(line, "?", "."))
	options = append(options, strings.ReplaceAll(line, "?", "#"))
	//o := []string{}
	if count > 1 {
		for i := 1; i < count-1; i++ {
			q := strings.Repeat("#", i) + "." + strings.Repeat(".", count-i)
			Permutation(q, 0)
		}
	}
	//fmt.Println(o)

	return options
}

func Solver1(lines []string) int {
	sum := 0

	for _, line := range lines {
		options := ListOptions(line)
		for _, option := range options {
			if Valid(option) {
				sum++
			}
		}
	}

	return sum
}

func Solver2(lines []string) int {
	return 0
}
