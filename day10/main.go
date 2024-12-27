package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type coordinate struct {
	x int
	y int
}

// Store complete paths for each trailhead
var paths = map[coordinate][][]coordinate{}

func main() {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := parseInput(file)
	trailheads := findTrailheads(grid)
	findPaths(grid, trailheads)
	fmt.Println(getPathSum())
	fmt.Println(getPathSumTwo())
}

func parseInput(file *os.File) [][]int {
	grid := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, v := range line {
			row[i], _ = strconv.Atoi(string(v))
		}
		grid = append(grid, row)
	}
	return grid
}

func findTrailheads(grid [][]int) []coordinate {
	coords := make([]coordinate, 0)
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 0 {
				coords = append(coords, coordinate{r, c})
			}
		}
	}
	return coords
}

func findPaths(grid [][]int, trailheads []coordinate) {
	for _, trailhead := range trailheads {
		currentPath := []coordinate{trailhead}
		traverse(grid, trailhead, currentPath, trailhead.x, trailhead.y)
	}
}

func traverse(grid [][]int, trailhead coordinate, currentPath []coordinate, r, c int) {
	// If we reached a 9, we found a valid path
	if grid[r][c] == 9 {
		// Make a copy of the current path
		pathCopy := make([]coordinate, len(currentPath))
		copy(pathCopy, currentPath)

		// Add this path to the trailhead's paths
		paths[trailhead] = append(paths[trailhead], pathCopy)
		return
	}

	nextValue := grid[r][c] + 1
	// Check all four directions
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right

	for _, dir := range directions {
		newR, newC := r+dir[0], c+dir[1]

		// Check bounds
		if newR < 0 || newR >= len(grid) || newC < 0 || newC >= len(grid[0]) {
			continue
		}

		// Check if next value matches and position hasn't been visited
		if grid[newR][newC] == nextValue && !containsCoordinate(currentPath, coordinate{newR, newC}) {
			// Add new position to path
			newPath := append(currentPath, coordinate{newR, newC})
			traverse(grid, trailhead, newPath, newR, newC)
		}
	}
}

func containsCoordinate(path []coordinate, coord coordinate) bool {
	for _, c := range path {
		if c == coord {
			return true
		}
	}
	return false
}

func pathToString(path []coordinate) string {
	result := ""
	for _, coord := range path {
		result += fmt.Sprintf("(%d,%d)", coord.x, coord.y)
	}
	return result
}

func getPathSumTwo() int {
	sum := 0
	// For each trailhead, count unique complete paths
	for _, pathList := range paths {
		// Use a map to track unique complete paths
		uniquePaths := make(map[string]bool)
		for _, path := range pathList {
			if len(path) > 0 {
				uniquePaths[pathToString(path)] = true
			}
		}
		sum += len(uniquePaths)
	}
	return sum
}

func getPathSum() int {
	sum := 0
	// For each trailhead, count unique endpoints
	for _, pathList := range paths {
		// Use a map to track unique endpoints
		uniqueEnds := make(map[coordinate]bool)
		for _, path := range pathList {
			if len(path) > 0 {
				uniqueEnds[path[len(path)-1]] = true
			}
		}
		sum += len(uniqueEnds)
	}
	return sum
}
