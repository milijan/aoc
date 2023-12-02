// function that reads a text file that contains
// a number per line and calculates the number of non-empty lines
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    puzzle, err := os.ReadFile("./puzzle.txt")
    if err != nil {
        fmt.Println(err)
    }
    lines := strings.Split(string(puzzle), "\n")
    sum := 0
    val := 0
    maxsum := 0
    for _, line := range lines {
        if len(line) > 0 {
            val, err = strconv.Atoi(line)
            if err != nil {
                fmt.Println(err)
            } else {
                sum += val
            }
        } else {
            if sum > maxsum {
                maxsum = sum
            }
            sum = 0
        }
    }
    fmt.Println("Number of non-empty lines:", maxsum)
}