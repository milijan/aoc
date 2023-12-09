package main

import (
	"strconv"
	"strings"
)

func diff(a []int, b []int) []int {
	r := []int{}
	for i, v := range a {
		r = append(r, v-b[i])
	}
	return r
}

func allZeros(a []int) bool {
	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}

func extrapolateRight(data []int) int {
	n := len(data)
	diff := diff(data[1:n], data[0:n-1])
	if allZeros(diff) {
		return data[n-1]
	}
	return extrapolateRight(diff) + data[n-1]
}

func extrapolateLeft(data []int) int {
	n := len(data)
	diff := diff(data[1:n], data[0:n-1])
	if allZeros(diff) {
		return data[0]
	}
	return data[0] - extrapolateLeft(diff)
}

func parseValues(a []string) []int {
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

func Solver1(lines []string) int {
	sum := 0
	for _, line := range lines {
		data := parseValues(strings.Split(line, " "))
		sum += extrapolateRight(data)
	}
	return sum
}

func Solver2(lines []string) int {
	sum := 0
	for _, line := range lines {
		data := parseValues(strings.Split(line, " "))
		sum += extrapolateLeft(data)
	}
	return sum
}
