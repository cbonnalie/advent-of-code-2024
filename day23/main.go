package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Connection struct {
	nodes []string
}

func NewConnection(nodes []string) Connection {
	sortedNodes := make([]string, len(nodes))
	copy(sortedNodes, nodes)
	sort.Strings(sortedNodes)
	return Connection{nodes: sortedNodes}
}

func (c Connection) String() string {
	return strings.Join(c.nodes, ",")
}

func prep() map[string][]string {
	conns := make(map[string][]string)

	file, err := os.Open("day23/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		a, b := parts[0], parts[1]
		conns[a] = append(conns[a], b)
		conns[b] = append(conns[b], a)
	}

	return conns
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func startsWithT(s string) bool {
	return len(s) > 0 && s[0] == 't'
}

func part1(conns map[string][]string) int {
	seen := make(map[string]bool)

	for a, bs := range conns {
		for _, b := range bs {
			for _, c := range conns[b] {
				if contains(conns[c], a) {
					nodes := []string{a, b, c}
					hasT := false
					for _, node := range nodes {
						if startsWithT(node) {
							hasT = true
							break
						}
					}
					if hasT {
						conn := NewConnection(nodes)
						seen[conn.String()] = true
					}
				}
			}
		}
	}

	return len(seen)
}

func setDifference(a, b []string) []string {
	bMap := make(map[string]bool)
	for _, item := range b {
		bMap[item] = true
	}

	var diff []string
	for _, item := range a {
		if !bMap[item] {
			diff = append(diff, item)
		}
	}
	return diff
}

func setIntersection(a, b []string) []string {
	bMap := make(map[string]bool)
	for _, item := range b {
		bMap[item] = true
	}

	var inter []string
	for _, item := range a {
		if bMap[item] {
			inter = append(inter, item)
		}
	}
	return inter
}

func setUnion(a, b []string) []string {
	unionMap := make(map[string]bool)
	for _, item := range a {
		unionMap[item] = true
	}

	for _, item := range b {
		unionMap[item] = true
	}

	union := make([]string, 0, len(unionMap))
	for item := range unionMap {
		union = append(union, item)
	}
	return union
}

func neighbors(v string, graph map[string][]string) []string {
	return graph[v]
}

func bronKerbosch(r, p, x []string, graph map[string][]string, maxCliques *[][]string) []string {
	if len(p) == 0 && len(x) == 0 {
		clique := make([]string, len(r))
		copy(clique, r)
		*maxCliques = append(*maxCliques, clique)
	}
	for _, v := range p {
		bronKerbosch(
			setUnion(r, []string{v}),
			setIntersection(p, neighbors(v, graph)),
			setIntersection(x, neighbors(v, graph)),
			graph,
			maxCliques,
		)
		p = setDifference(p, []string{v})
		x = setUnion(x, []string{v})
	}
	return nil
}

func part2(conns map[string][]string) string {
	vertices := make([]string, 0)
	for v := range conns {
		vertices = append(vertices, v)
	}

	var maxCliques [][]string
	bronKerbosch([]string{}, vertices, []string{}, conns, &maxCliques)

	var largestClique []string
	maxSize := 0
	for _, clique := range maxCliques {
		if len(clique) > maxSize {
			maxSize = len(clique)
			largestClique = clique
		}
	}

	sort.Strings(largestClique)
	return strings.Join(largestClique, ",")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func main() {
	start := time.Now()
	conns := prep()
	timeTrack(start, "Prep")

	start = time.Now()
	fmt.Printf("Part 1: %d\n", part1(conns))
	timeTrack(start, "Part 1")

	start = time.Now()
	fmt.Printf("Part 2: %s\n", part2(conns))
	timeTrack(start, "Part 2")
}
