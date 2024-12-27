package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Grid [][]rune

type ParseError struct {
	Path string
	Err  error
}

type Point struct {
	Row, Col int
}

type Region struct {
	Char    rune
	Points  []Point
	Visited bool
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("failed to parse %s: %v", e.Path, e.Err)
}

func ReadGrid(path string) (Grid, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, &ParseError{path, err}
	}
	defer file.Close()

	return ParseGrid(file)
}

func ParseGrid(r io.Reader) (Grid, error) {
	scanner := bufio.NewScanner(r)
	var lines [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, 0, len(line)+2)
		row = append(row, '.')
		row = append(row, []rune(line)...)
		row = append(row, '.')
		lines = append(lines, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("input file is empty")
	}

	width := len(lines[0])
	padding := make([]rune, width)
	for i := range padding {
		padding[i] = '.'
	}

	grid := make(Grid, 0, len(lines)+2)
	grid = append(grid, append([]rune{}, padding...))
	grid = append(grid, lines...)
	grid = append(grid, append([]rune{}, padding...))

	return grid, nil
}

func (g Grid) Print() {
	for _, row := range g {
		fmt.Println(string(row))
	}
}

func (g Grid) FindRegions() []Region {
	visited := make(map[Point]bool)
	var regions []Region

	for row := 1; row < len(g)-1; row++ {
		for col := 1; col < len(g[row])-1; col++ {
			point := Point{row, col}

			if visited[point] || g[row][col] == '.' {
				continue
			}

			region := g.exploreRegion(point, visited)
			if len(region.Points) > 0 {
				regions = append(regions, region)
			}
		}
	}
	return regions
}

func (g Grid) exploreRegion(start Point, visited map[Point]bool) Region {
	char := g[start.Row][start.Col]
	region := Region{
		Char:   char,
		Points: make([]Point, 0),
	}

	queue := []Point{start}
	visited[start] = true

	dirs := []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		region.Points = append(region.Points, current)

		for _, dir := range dirs {
			next := Point{current.Row + dir.Row, current.Col + dir.Col}
			if visited[next] || g[next.Row][next.Col] != char {
				continue
			}

			visited[next] = true
			queue = append(queue, next)
		}
	}
	return region
}

func (g Grid) PrintRegion(region Region) {
	// Create a copy of the grid with only the region visible
	display := make([][]rune, len(g))
	for i := range g {
		display[i] = make([]rune, len(g[i]))
		for j := range g[i] {
			display[i][j] = '.'
		}
	}

	// Fill in the region
	for _, p := range region.Points {
		display[p.Row][p.Col] = region.Char
	}

	// Print the result
	for _, row := range display {
		fmt.Println(string(row))
	}
}

func (region Region) getPerimeter(g Grid) int {
	perimeter := 0
	for _, point := range region.Points {
		sides := 4
		if g[point.Row-1][point.Col] == region.Char {
			sides--
		}
		if g[point.Row+1][point.Col] == region.Char {
			sides--
		}
		if g[point.Row][point.Col-1] == region.Char {
			sides--
		}
		if g[point.Row][point.Col+1] == region.Char {
			sides--
		}
		perimeter += sides
	}
	return perimeter
}

func main() {
	grid, err := ReadGrid(filepath.Join("day12", "input.txt"))
	if err != nil {
		log.Fatalf("Failed to read grid: %v", err)
	}

	grid.Print()

	regions := grid.FindRegions()

	totalCost := 0
	fmt.Printf("Found %d regions:\n", len(regions))
	for i, region := range regions {
		perimeter := region.getPerimeter(grid)
		area := len(region.Points)
		totalCost += perimeter * area
		fmt.Printf("\nRegion %d (character '%c', size %d):\n",
			i+1, region.Char, len(region.Points))
		grid.PrintRegion(region)
		fmt.Println("Perimeter:", perimeter)
		fmt.Println("Area:", area)
	}

	fmt.Printf("Total cost:\n%d\n", totalCost)
}
