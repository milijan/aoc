package main

import (
	"fmt"
	"strings"
	"unicode"

	"gonum.org/v1/gonum/mat"
)

func Solver1(puzzle string) int {
	length := len(puzzle)
	if length == 0 {
		return 0
	}

	printMat := func(s string, m mat.Matrix) {
		// Print the result using the formatter.
		fa := mat.Formatted(m, mat.Prefix(strings.Repeat(" ", len(s)+3)), mat.Squeeze())
		fmt.Printf("\n%s = %v\n", s, fa)
	}

	get2DSlice := func(m mat.Matrix, i_r int, i_c int) (int, int, int, int) {
		n_r, n_c := m.Dims()
		// find next part
		for i_r != n_r && m.At(i_r, i_c%n_c) == 0 {
			i_c++
			if i_c%n_c == 0 {
				i_r++
			}
		}
		i, j, k, l := 0, 0, 0, 0
		if i_r != n_r {
			// find part size
			p_s := i_c
			for ; p_s < n_c && m.At(i_r, p_s%n_c) > 0; p_s++ {
			}

			i, k = max(0, i_r-1)%n_r, (min(n_r-1, i_r+1)%n_r)+1
			j = max(0, i_c-1) % n_c
			l = min(n_c, j+p_s)
		}
		return i, k, j, l
	}

	cols := strings.IndexRune(puzzle, '\n')
	rows := (length - cols + 1) / cols
	parts := mat.NewDense(rows, cols, nil)
	symc := mat.NewDense(rows, cols, nil)

	for k, c := range puzzle {
		j := k % (cols + 1)
		i := (k - j) / rows
		if unicode.IsNumber(c) {
			parts.Set(i, j, 1)
		} else if c != '.' && c != '\n' {
			symc.Set(i, j, 1)
		}
	}

	a, b, c, d := get2DSlice(parts, 2, 0)
	println(a, b, c, d)
	slice := symc.Slice(a, b, c, d)

	printMat("parts", parts)
	printMat("symbols", symc)
	printMat("slice", slice)
	fmt.Println(mat.Sum(slice))

	return 0
}

func Solver2(puzzle string) int {
	return 0
}
