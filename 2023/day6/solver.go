package main

import (
	"runtime"
	"strconv"
	"sync"
)

type puzzle struct {
	time     []int
	distance []int
}

func Reduce(record int, newdistance <-chan int, result chan<- int) {
	count := 0
	for d := range newdistance {
		if d > record {
			count++
		}
	}
	result <- count
}

func Map(minHold int, raceDuration int, newDistance chan<- int) {
	var wg sync.WaitGroup
	for hold := minHold; hold < raceDuration; hold++ {
		wg.Add(1)
		go func(i int, r int) {
			defer wg.Done()
			newDistance <- i * (r - i)
		}(hold, raceDuration)
	}
	wg.Wait()
	close(newDistance)
}

func CountWays(raceDuration int, record int, minhold int) int {
	n := 3
	runtime.GOMAXPROCS(n)

	count := make(chan int, n)
	distance := make(chan int, n)

	go Reduce(record, distance, count)

	Map(minhold, raceDuration, distance)

	return <-count
}

func Solver1(p puzzle) int {
	if len(p.time) == 0 {
		return 0
	}
	result := 1
	for i := range p.time {
		result *= CountWays(p.time[i], p.distance[i], 1)
	}
	return result
}

func Solver2(p puzzle) int {
	n := len(p.time)
	if n == 0 {
		return 0
	}
	duration := ""
	distance := ""
	// merge races
	for i := range p.time {
		duration += strconv.Itoa(p.time[i])
		distance += strconv.Itoa(p.distance[i])
	}
	t, _ := strconv.Atoi(duration)
	d, _ := strconv.Atoi(distance)

	return CountWays(t, d, 14)
}
