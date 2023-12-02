package main

import (
	"testing"
)

func TestSolver1(t *testing.T) {

	test := func(lines []string, expected int) {
		result := Solver1(lines)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test([]string{}, 0)
	test([]string{"asdsd"}, 0)
	test([]string{"asd4asd5asd"}, 45)
	test([]string{"123"}, 13)
	test([]string{"13"}, 13)
	test([]string{"a12aaa3"}, 13)
	test([]string{"12aa3aa"}, 13)
	test([]string{"5"}, 55)
	test([]string{"a5"}, 55)
	test([]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}, 142)
}

func TestSolver2(t *testing.T) {

	test := func(lines []string, expected int) {
		result := Solver2(lines)
		if result != expected {
			t.Errorf("Expected %d, result %d", expected, result)
		}
	}

	test([]string{}, 0)
	test([]string{"13"}, 13)
	test([]string{"asdsd"}, 0)
	test([]string{"asd4asd5asd"}, 45)
	test([]string{"123"}, 13)
	test([]string{"a12aaa3"}, 13)
	test([]string{"12aa3aa"}, 13)
	test([]string{"5"}, 55)
	test([]string{"a5"}, 55)
	test([]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}, 142)
	test([]string{"one"}, 11)
	test([]string{"oneone"}, 11)
	test([]string{"two"}, 22)
	test([]string{"three"}, 33)
	test([]string{"4four"}, 44)
	test([]string{"five5"}, 55)
	test([]string{"six"}, 66)
	test([]string{"seven"}, 77)
	test([]string{"eight"}, 88)
	test([]string{"nine"}, 99)
	test([]string{"two1nine"}, 29)
	test([]string{"twone"}, 21)
	test([]string{"eightwothree"}, 83)
	test([]string{"eightwo"}, 82)
	test([]string{"oneightwoo"}, 12)
	test([]string{"toneightwoneightwo"}, 12)
	test([]string{"oneightwoneightwoneightwoneightwo"}, 12)
	test([]string{"sevenine"}, 79)
	test([]string{"nine6nine"}, 99)
	test([]string{"none6nine"}, 19)
	test([]string{"sevenineighthree"}, 73)
	test([]string{"sevenineighthreeight"}, 78)
	test([]string{"sevenineighthreeightwone"}, 71)
	test([]string{"sevenineighthreeightwoneoneightwo"}, 72)
	test([]string{"oneightwoneightwone"}, 11)
	test([]string{"abcone2threexyz"}, 13)
	test([]string{"xtwone3four"}, 24)
	test([]string{"4nineeightseven2"}, 42)
	test([]string{"nineeightseven2"}, 92)
	test([]string{"zoneight234"}, 14)
	test([]string{"7pqrstsixteen"}, 76)
	test([]string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}, 281)
}
