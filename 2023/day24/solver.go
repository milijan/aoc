package main

type Hail struct {
	position []int64
	velocity []int64
}

func IsCrossing(A Hail, B Hail) (bool, int64) {
	det := A.velocity[0]*B.velocity[1] - B.velocity[0]*A.velocity[1]
	return !(det == 0), det
}

func XIntercept(A Hail, B Hail) (bool, float64) {
	ok, det := IsCrossing(A, B)
	if ok {
		num := A.velocity[0]*B.velocity[1]*B.position[0] - B.velocity[0]*A.velocity[1]*A.position[0]
		num += A.velocity[0] * B.velocity[0] * (A.position[1] - B.position[1])
		return ok, float64(num) / float64(det)
	} else {
		return ok, 0
	}
}

func YIntercept(A Hail, B Hail, X float64) float64 {
	num := float64(A.velocity[1])*(X-float64(A.position[0])) +
		float64(A.velocity[0]*A.position[1])
	return num / float64(A.velocity[0])
}

func FutureIntercept(A Hail, X float64, Y float64) bool {
	if A.velocity[0] != 0 {
		return (X-float64(A.position[0]))/float64(A.velocity[0]) >= 0
	} else {
		return (Y-float64(A.position[1]))/float64(A.velocity[1]) >= 0
	}
}

func InAreaIntercept(area []uint64, X float64, Y float64) bool {
	return X >= float64(area[0]) &&
		X <= float64(area[1]) &&
		Y >= float64(area[0]) &&
		Y <= float64(area[1])
}

func Solver1(puzzle []Hail, area []uint64) int {
	count := 0
	n := len(puzzle)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			ok, X := XIntercept(puzzle[i], puzzle[j])
			if ok {
				Y := YIntercept(puzzle[i], puzzle[j], X)
				if FutureIntercept(puzzle[i], X, Y) &&
					FutureIntercept(puzzle[j], X, Y) &&
					InAreaIntercept(area, X, Y) {
					count++
				}
			}
		}
	}
	return count
}

func Solver2(puzzle []Hail) int {
	return 0
}
