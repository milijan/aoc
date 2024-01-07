package main

type Hail struct {
	position []float64
	velocity []float64
}

func IsCrossing(A Hail, B Hail) (bool, float64) {
	det := A.velocity[0]*B.velocity[1] - B.velocity[0]*A.velocity[1]
	return !(det == 0.0), det
}

func XIntercept(A Hail, B Hail) (bool, float64) {
	ok, det := IsCrossing(A, B)
	if ok {
		num := A.velocity[0]*B.velocity[1]*B.position[0] - B.velocity[0]*A.velocity[1]*A.position[0]
		num += A.velocity[0] * B.velocity[0] * (A.position[1] - B.position[1])
		return ok, num / det
	} else {
		return ok, 0
	}
}

func YIntercept(A Hail, B Hail, X float64) float64 {
	num := A.velocity[1]*X - A.position[0] +
		A.velocity[0]*A.position[1]
	return num / A.velocity[0]
}

func FutureIntercept(A Hail, X float64, Y float64) bool {
	if A.velocity[0] != 0 {
		return (X-A.position[0])/A.velocity[0] >= 0
	} else {
		return (Y-A.position[1])/A.velocity[1] >= 0
	}
}

func InAreaIntercept(area []float64, X float64, Y float64) bool {
	return X >= area[0] && X <= area[1] && Y >= area[0] && Y <= area[1]
}

func Solver1(puzzle []Hail, area []float64) int {
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
