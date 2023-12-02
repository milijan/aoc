package main

import (
	"strconv"
	"strings"
)

func ReplaceFirst(line string, digitsStr []string) (string, int) {
	digitsNum := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	index := -1
	digitRank := -1
	count := 0
	for i, digit := range digitsStr {
		rank := strings.Index(line, digit)
		if rank >= 0 {
			count++
			if digitRank < 0 || rank < digitRank {
				digitRank = rank
				index = i
			}
		}
	}
	if index >= 0 {
		line = strings.Replace(line, digitsStr[index], digitsNum[index], 1)
	}
	return line, count
}

func ReplaceLast(line string, digitsStr []string) string {
	digitsNum := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	index := -1
	digitRank := -1
	for i, digit := range digitsStr {
		rank := strings.LastIndex(line, digit)
		if rank >= 0 {
			if digitRank < 0 || rank > digitRank {
				digitRank = rank
				index = i
			}
		}
	}
	if index >= 0 {
		line = strings.ReplaceAll(line, digitsStr[index], digitsNum[index])
	}
	return line
}

func Interpret(line string) string {
	//line = strings.ToLower(line)
	digitsStr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	line, count := ReplaceFirst(line, digitsStr)
	if count > 1 {
		line = ReplaceLast(line, digitsStr)
	}
	return line
}

func Decode(line string) int {
	code := 0
	first := true
	var f, l string
	if len(line) > 0 {
		for _, c := range line {
			if c >= '0' && c <= '9' {
				l = string(c)
				if first {
					f = l
					first = false
				}
			}
		}
		code, _ = strconv.Atoi(f + l)
	}
	return code
}

func Solver1(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += Decode(line)
	}

	return sum
}

func Solver2(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += Decode(Interpret(line))
	}

	return sum
}
