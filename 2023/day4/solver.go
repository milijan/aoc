package main

import (
	"math"
	"regexp"
)

func Solver1(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	re := regexp.MustCompile(`[0-9]+`)
	points := 0
	for _, line := range lines {
		matches := 0
		card := re.FindAllString(line, -1)
		// find duplicate in list
		for j, c1 := range card[1:] {
			for _, c2 := range card[j+2:] {
				if c1 == c2 {
					matches++
				}
			}
		}
		points += int(math.Floor(math.Pow(2, float64(matches-1))))
	}

	return points
}

func Solver2(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	re := regexp.MustCompile(`[0-9]+`)
	cardsTab := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		cardsTab[i] += 1
		matches := 0
		card := re.FindAllString(lines[i], -1)
		// find duplicate in list
		for j, c1 := range card[1:] {
			for _, c2 := range card[j+2:] {
				if c1 == c2 {
					matches++
				}
			}
		}
		if matches > 0 {
			for j := i + 1; j < min(i+1+matches, len(lines)); j++ {
				cardsTab[j] += 1 * cardsTab[i]
			}
		}
	}
	// sum of array elements
	sum := func(a []int) int {
		s := 0
		for _, v := range a {
			s += v
		}
		return s
	}
	return sum(cardsTab)
}
