package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

var dirs = [][][2]int{
	{{-1, -1}, {1, -1}}, //up
	{{-1, 1}, {1, 1}},   //down
	{{-1, 1}, {-1, -1}}, //left
	{{1, 1}, {1, -1}},   //right
}

var opposites = [][][2]int{
	{{-1, 1}, {1, 1}},   //down
	{{-1, -1}, {1, -1}}, //up
	{{1, 1}, {1, -1}},   //right
	{{-1, 1}, {-1, -1}}, //left
}

var target = []rune{'X', 'M', 'A', 'S'}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func partOne() int {

	var grid [][]rune
	file, _ := os.Open("day4/input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	rows, cols := len(grid), len(grid[0])
	total := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			total += countFromCell(grid, r, c)
		}
	}

	return total
}

func countFromCell(grid [][]rune, r, c int) int {
	count := 0

	for _, dir := range directions {
		count += searchDirection(grid, r, c, dir[0], dir[1])
	}

	return count
}

func searchDirection(grid [][]rune, startR, startC, dR, dC int) int {

	rows, cols := len(grid), len(grid[0])
	r, c := startR, startC

	for i := 0; i < len(target); i++ {

		if r < 0 || r >= rows || c < 0 || c >= cols {
			return 0
		}
		if grid[r][c] != target[i] {
			return 0
		}
		r += dR
		c += dC
	}
	return 1
}

func partTwo() int {
	var grid [][]rune
	file, _ := os.Open("day4/input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	rows, cols := len(grid), len(grid[0])
	total := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'A' {
				total += countMAS(grid, r, c)
			}
		}
	}

	return total
}

func countMAS(grid [][]rune, r, c int) int {

	count := 0

	for i := 0; i < len(dirs); i++ {

		dir := dirs[i]
		opposite := opposites[i]

		r1, c1 := r+dir[0][0], c+dir[0][1]
		r2, c2 := r+dir[1][0], c+dir[1][1]
		or1, oc1 := r+opposite[0][0], c+opposite[0][1]
		or2, oc2 := r+opposite[1][0], c+opposite[1][1]

		if r1 < 0 || r1 >= len(grid) || c1 < 0 || c1 >= len(grid[0]) {
			continue
		}

		if r2 < 0 || r2 >= len(grid) || c2 < 0 || c2 >= len(grid[0]) {
			continue
		}

		if or1 < 0 || or1 >= len(grid) || oc1 < 0 || oc1 >= len(grid[0]) {
			continue
		}

		if or2 < 0 || or2 >= len(grid) || oc2 < 0 || oc2 >= len(grid[0]) {
			continue
		}

		if grid[r1][c1] == grid[r2][c2] && grid[r1][c1] == 'M' {
			if grid[or1][oc1] == grid[or2][oc2] && grid[or1][oc1] == 'S' {
				count = 1
				break
			}
		}

		if grid[r1][c1] == grid[r2][c2] && grid[r1][c1] == 'S' {
			if grid[or1][oc1] == grid[or2][oc2] && grid[or1][oc1] == 'M' {
				count = 1
				break
			}
		}

	}

	return count
}
