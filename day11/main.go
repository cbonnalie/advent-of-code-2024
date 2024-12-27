package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("day11/input.txt")

	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	initialStones := make([]int, len(line))
	for i := range initialStones {
		num, _ := strconv.Atoi(line[i])
		initialStones[i] = num
	}

	//fmt.Println(partOne(initialStones))
	fmt.Println(partTwo(initialStones))
}
