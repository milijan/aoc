package main

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
)

func PrintGraph(puzzle map[string][]string) {
	fmt.Println("Nodes: ", len(puzzle))
	for node, edges := range puzzle {
		fmt.Println(" ", node, edges)
	}
}

func BuildGraph(q map[string][]string, p map[string][]string, node string) map[string][]string {
	q[node] = slices.Clone(p[node])
	delete(p, node)
	for _, edge := range q[node] {
		n, m := EdgeNodes(edge)
		if n == node {
			n = m
		}
		_, ok := q[n]
		if !ok {
			q = BuildGraph(q, p, n)
		}
	}
	return q
}

func SplitGraph(puzzle map[string][]string, cuts []string) []map[string][]string {

	p := make(map[string][]string, len(puzzle))
	for node, edges := range puzzle {
		p[node] = slices.Clone(edges)
	}

	for _, cut := range cuts {
		n1, n2 := EdgeNodes(cut)
		i := slices.Index(p[n1], cut)
		p[n1] = slices.Delete(p[n1], i, i+1)

		j := slices.Index(p[n2], cut)
		p[n2] = slices.Delete(p[n2], j, j+1)
	}

	graphs := []map[string][]string{}
	for len(p) > 1 {
		q := map[string][]string{}
		for node, _ := range p {
			q = BuildGraph(q, p, node)
			break
		}
		graphs = append(graphs, q)
	}
	return graphs
}

func EdgeNodes(edge string) (string, string) {
	n1 := ""
	n2 := ""
	for i, c := range edge {
		if i < 3 {
			n1 += string(c)
		} else {
			n2 += string(c)
		}
	}
	return n1, n2
}

func ExtendNode(puzzle map[string][]string, node1 string, node2 string) map[string][]string {
	// add new edges to node1
	// remove circular edges
	for _, edge := range puzzle[node2] {
		j := slices.Index(puzzle[node1], edge)
		if j > -1 {
			puzzle[node1] = slices.Delete(puzzle[node1], j, j+1)
		} else {
			puzzle[node1] = append(puzzle[node1], edge)
		}
	}
	delete(puzzle, node2)

	return puzzle
}

func RadomEdge(puzzle map[string][]string) (string, string) {
	if len(puzzle) > 1 {
		n := rand.Intn(len(puzzle))
		i := 0
		for node1, edges1 := range puzzle {
			if i == n {
				j := rand.Intn(len(edges1))
				edge := edges1[j]
				for node2, edges2 := range puzzle {
					if node1 != node2 && slices.Contains(edges2, edge) {
						return node1, node2
					}
				}
			}
			i++
		}
	}
	return "", ""
}

// karger's algorithm
func MinCut(puzzle map[string][]string, interations int, expected int) (int, []string) {
	mincut := math.MaxInt32
	cuts := []string{}
	for i := 0; i < interations; i++ {
		p := make(map[string][]string, len(puzzle))
		for node, edges := range puzzle {
			p[node] = slices.Clone(edges)
		}
		for len(p) > 2 {
			node1, node2 := RadomEdge(p)
			ExtendNode(p, node1, node2)
		}

		// take first node (identical edges)
		for _, v := range p {
			if mincut > len(v) {
				cuts = v
				mincut = len(v)
			}
			if mincut == expected {
				return mincut, cuts
			}
			break
		}
	}
	return mincut, cuts
}

func Solver1(puzzle map[string][]string) int {

	_, cuts := MinCut(puzzle, 100, 3)

	graphs := SplitGraph(puzzle, cuts)

	if len(graphs) == 2 {
		return len(graphs[0]) * len(graphs[1])
	}
	return 0
}

func Solver2(puzzle map[string][]string) int {
	return 0
}
