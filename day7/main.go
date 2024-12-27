package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var isPartTwo bool

func getCalibrationResult(targets []string, values [][]string) {

	for _, v := range []bool{false, true} {
		var total int64 = 0
		isPartTwo = v
		for i := 0; i < len(targets); i++ {
			intTarget, _ := strconv.ParseInt(targets[i], 10, 64)
			if evaluate(values[i], intTarget, 0, 0) {
				total += intTarget
			}
		}
		fmt.Println(total)
	}
}

func evaluate(values []string, target, acc int64, index int) bool {

	// add first value to accumulator
	if index == 0 {
		intValue, _ := strconv.ParseInt(values[index], 10, 64)
		return evaluate(values, target, intValue, 1)
	}

	// check acc against the target when we've reached the last element of values
	if index == len(values)-1 {
		intValue, _ := strconv.ParseInt(values[index], 10, 64)

		if acc+intValue == target || acc*intValue == target {
			return true
		}

		if isPartTwo {
			accString := strconv.FormatInt(acc, 10)
			lastString := strconv.FormatInt(intValue, 10)
			concat, _ := strconv.ParseInt(accString+lastString, 10, 64)
			return concat == target
		}
		return false
	}

	nextValue, _ := strconv.ParseInt(values[index], 10, 64)

	if isPartTwo {
		accString := strconv.FormatInt(acc, 10)
		concat := accString + values[index]
		accInt, _ := strconv.ParseInt(concat, 10, 64)
		return evaluate(values, target, acc+nextValue, index+1) ||
			evaluate(values, target, acc*nextValue, index+1) ||
			evaluate(values, target, accInt, index+1)
	}
	return evaluate(values, target, acc+nextValue, index+1) ||
		evaluate(values, target, acc*nextValue, index+1)
}

func parseInput(filename string) ([]string, [][]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var targets []string
	var values [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		target := parts[0]
		targets = append(targets, target)

		valueStrings := strings.Fields(strings.TrimSpace(parts[1]))
		values = append(values, valueStrings)
	}

	return targets, values, nil
}

func main() {

	targets, values, err := parseInput("day7/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	getCalibrationResult(targets, values)
}
