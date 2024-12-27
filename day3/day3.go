package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const EMPTY string = ""
const DO string = "do()"

func part1(data string) int {

	mulRegex := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	mulMatches := mulRegex.FindAllStringSubmatch(data, -1)

	total := 0

	for _, match := range mulMatches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		total += num1 * num2
	}

	return total
}

func part2(data string) int {

	/*
		0 = mul(~)
		1 = mul(~)
		2 = first number
		3 = second number
		4 = do() / dont()
	*/
	regex := regexp.MustCompile(`(mul\(([0-9]{1,3}),([0-9]{1,3})\))|(do\(\)|don't\(\))`)
	matches := regex.FindAllStringSubmatch(data, -1)

	doDontMap := map[string][]string{
		"do()":    {},
		"don't()": {},
	}

	doOrDont := DO

	for _, match := range matches {
		mulMatch := match[1]
		doOrDontMatch := match[4]

		if doOrDontMatch != EMPTY {
			doOrDont = doOrDontMatch
		}

		if mulMatch != EMPTY {
			doDontMap[doOrDont] = append(doDontMap[doOrDont], mulMatch)
		}
	}

	var total int
	for _, entry := range doDontMap[DO] {
		nums := strings.Split(entry, ",")
		one, _ := strconv.Atoi(nums[0][4:])
		two, _ := strconv.Atoi(nums[1][:len(nums[1])-1])
		total += one * two
	}

	return total
}

func main() {

	data := func() string { d, _ := os.ReadFile("day3/input3.txt"); return string(d) }()

	partOneResult := part1(data)
	partTwoResult := part2(data)

	fmt.Println("--- PART 1 ---")
	fmt.Println(partOneResult)
	fmt.Println("--- PART 2 ---")
	fmt.Println(partTwoResult)
}
