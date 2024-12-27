package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rules = make(map[string][]string)
var updates [][]string
var valid [][]string
var invalid [][]string

func parseRules(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		source := parts[0]
		dest := parts[1]
		rules[source] = append(rules[source], dest)
	}
}

func parsePages(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		var currentPage []string
		for _, part := range parts {
			currentPage = append(currentPage, part)
		}
		updates = append(updates, currentPage)
	}
}

func determineIfValid(update []string) {
	for i, val := range update {
		for _, rule := range rules[val] {
			for j := i - 1; j >= 0; j-- {
				if update[j] == rule {
					invalid = append(invalid, update)
					return
				}
			}
		}
	}
	valid = append(valid, update)
}

func getSumOfMiddles(updates [][]string) int {
	sum := 0
	for _, update := range updates {
		mid := len(update) / 2
		value, _ := strconv.Atoi(update[mid])
		sum += value
	}
	return sum
}

func sortInvalid() {
	for idx := range invalid {
		swapped := true
		for swapped {
			swapped = false
			for i := 0; i < len(invalid[idx])-1; i++ {
				currentVal := invalid[idx][i]
				for _, rule := range rules[currentVal] {
					for j := i + 1; j < len(invalid[idx]); j++ {
						if invalid[idx][j] == rule {
							invalid[idx][i], invalid[idx][j] = invalid[idx][j], invalid[idx][i]
							swapped = true
							break
						}
					}
					if swapped {
						break
					}
				}
				if swapped {
					break
				}
			}
		}
	}
}

func main() {
	file, _ := os.Open("day5/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	parseRules(scanner)
	parsePages(scanner)

	for _, update := range updates {
		determineIfValid(update)
	}

	sum := getSumOfMiddles(valid)
	fmt.Println("PART 1")
	fmt.Println(sum)

	sortInvalid()
	sum = getSumOfMiddles(invalid)
	fmt.Println("PART 2")
	fmt.Println(sum)
}
