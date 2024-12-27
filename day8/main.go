package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var grid [][]string

var antinodes map[coordinate]bool

type coordinate struct {
	x int
	y int
}

var partOne bool

func main() {
	parseGrid()

	for _, v := range []bool{true, false} {
		if v {
			fmt.Println("Part 1:")
		} else {
			fmt.Println("Part 2:")
		}

		partOne = v
		antinodes = make(map[coordinate]bool)
		findAntinodes()
		fmt.Println(len(antinodes))
	}
}

func parseGrid() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []string{})
		for _, ch := range line {
			grid[len(grid)-1] = append(grid[len(grid)-1], string(ch))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func printGrid() {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func findAntinodes() {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] != "." {
				searchForPair(grid[r][c], r, c)
			}
		}
	}
}

func searchForPair(char string, r, c int) {

	startR := r
	startC := c

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == char && r != startR && r != startC {

				if partOne {
					calculateAntinodePositionOne(startR, startC, r, c)
				} else {
					calculateAntinodePositionTwo(startR, startC, r, c)
				}

			}
		}
	}

}

func calculateAntinodePositionOne(startR, startC, r, c int) {
	rDiff := r - startR
	cDiff := c - startC

	if startR-rDiff >= 0 && startR-rDiff < len(grid) {
		if startC-cDiff >= 0 && startC-cDiff < len(grid[0]) {
			antinodes[coordinate{startR - rDiff, startC - cDiff}] = true
		}
	}

	if r+rDiff >= 0 && r+rDiff < len(grid) {
		if c+cDiff >= 0 && c+cDiff < len(grid[0]) {
			antinodes[coordinate{r + rDiff, c + cDiff}] = true
		}
	}
}

func calculateAntinodePositionTwo(startR, startC, r, c int) {

	antinodes[coordinate{startR, startC}] = true
	antinodes[coordinate{r, c}] = true

	rDiff := r - startR
	cDiff := c - startC

	for {
		if startR-rDiff >= 0 && startR-rDiff < len(grid) {
			if startC-cDiff >= 0 && startC-cDiff < len(grid[0]) {
				antinodes[coordinate{startR - rDiff, startC - cDiff}] = true
				rDiff += r - startR
				cDiff += c - startC
				continue
			}
		}
		break
	}

	rDiff = r - startR
	cDiff = c - startC

	for {
		if r+rDiff >= 0 && r+rDiff < len(grid) {
			if c+cDiff >= 0 && c+cDiff < len(grid[0]) {
				antinodes[coordinate{r + rDiff, c + cDiff}] = true
				rDiff += r - startR
				cDiff += c - startC
				continue
			}
		}
		break
	}
}
