package main

type puzzle struct {
	seeds        []int
	seedPairs    [][]int
	seedToSoil   [][]int
	soilToFert   [][]int
	fertToWater  [][]int
	waterToLight [][]int
	lightToTemp  [][]int
	tempToHum    [][]int
	humToLoc     [][]int
}

func GetDestination(s int, m [][]int) int {
	for _, r := range m {
		if s >= r[1] && s < r[1]+r[2] {
			return r[0] + s - r[1]
		}
	}
	return s
}

func GetSeedLocation(s int, p puzzle) int {
	loc := GetDestination(s, p.seedToSoil)
	loc = GetDestination(loc, p.soilToFert)
	loc = GetDestination(loc, p.fertToWater)
	loc = GetDestination(loc, p.waterToLight)
	loc = GetDestination(loc, p.lightToTemp)
	loc = GetDestination(loc, p.tempToHum)
	loc = GetDestination(loc, p.humToLoc)

	return loc
}

func Solver1(puz puzzle) int {

	m := -1
	for _, s := range puz.seeds {
		l := GetSeedLocation(s, puz)
		if m < 0 || l < m {
			m = l
		}
	}
	return m
}

func Solver2(puz puzzle) int {
	m := -1
	for _, s := range puz.seedPairs {
		for i := s[0]; i < s[0]+s[1]; i++ {
			l := GetSeedLocation(i, puz)
			if m < 0 || l < m {
				m = l
			}
		}
	}
	return m
}
