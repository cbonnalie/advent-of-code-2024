package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Schematic struct {
	heights []int
	isLock  bool
}

const (
	schematicWidth    = 5
	maxCombinedHeight = 5
	lockPattern       = "#####"
)

func ParseInput(filename string) ([]Schematic, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("failed to close input file: %v", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	var schematics []Schematic

	for scanner.Scan() {
		schematic, err := parseSchematic(scanner)
		if err != nil {
			return nil, err
		}
		schematics = append(schematics, schematic)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan input file: %w", err)
	}

	return schematics, nil
}

func parseSchematic(scanner *bufio.Scanner) (Schematic, error) {
	var lines []string
	line := scanner.Text()

	for line != "" {
		lines = append(lines, line)
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
	}

	heights := calculateHeights(lines)
	isLock := lines[0] == lockPattern

	return Schematic{
		heights,
		isLock,
	}, nil
}

func calculateHeights(lines []string) []int {
	heights := make([]int, schematicWidth)
	for i := range heights {
		heights[i] = -1
	}

	for _, row := range lines {
		for i, char := range row {
			if char == '#' {
				heights[i]++
			}
		}
	}

	return heights
}

func MatchSchematics(schematics []Schematic) int {
	var locks, keys [][]int

	for _, s := range schematics {
		if s.isLock {
			locks = append(locks, s.heights)
		} else {
			keys = append(keys, s.heights)
		}
	}

	return countValidPairs(locks, keys)
}

func countValidPairs(locks, keys [][]int) int {
	total := 0

	for _, lock := range locks {
		for _, key := range keys {
			if isValidPair(lock, key) {
				total++
			}
		}
	}

	return total
}

func isValidPair(lock, key []int) bool {
	for i := 0; i < schematicWidth; i++ {
		if lock[i]+key[i] > maxCombinedHeight {
			return false
		}
	}
	return true
}

func main() {
	schematics, err := ParseInput("day25/input.txt")
	if err != nil {
		log.Fatalf("Failed to parse input: %v", err)
	}

	total := MatchSchematics(schematics)
	fmt.Printf("Part 1: %d\n", total)
}
