package main

import (
	"strconv"
	"strings"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func FirstNumeric(line string) (int, int) {
	digit := 0
	for i, c := range line {
		if c >= '1' && c <= '9' {
			return int(c - '0'), i
		}
	}
	return digit, -1
}

func FirstAlpha(line string, digitsStr []string) (int, int) {
	index := -1
	digitRank := -1
	for i, digit := range digitsStr {
		rank := strings.Index(line, digit)
		if rank >= 0 {
			if digitRank < 0 || rank < digitRank {
				digitRank = rank
				index = i
			}
		}
	}
	return index + 1, digitRank
}

func FirstDigit(line string, digitsStr []string) int {
	numDigit, numDigitPos := FirstNumeric(line)
	alphaDigit, alphaDigitPos := FirstAlpha(line, digitsStr)

	if numDigitPos < alphaDigitPos {
		if numDigitPos >= 0 {
			return numDigit
		} else {
			return alphaDigit
		}
	} else {
		if alphaDigitPos >= 0 {
			return alphaDigit
		} else {
			return numDigit
		}
	}
}

func Decode2(line string) int {
	digitsStr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digitsRevStr := []string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}

	f := FirstDigit(line, digitsStr)
	l := FirstDigit(Reverse(line), digitsRevStr)

	return 10*f + l
}

func Decode1(line string) int {
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

// combining the first digit and the last digit (in that order) to form a single two-digit number
// What is the sum of all of the calibration values?
func Solver1(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += Decode1(line)
	}

	return sum
}

// same thing but bow digits can be strings
func Solver2(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += Decode2(line)
	}

	return sum
}
