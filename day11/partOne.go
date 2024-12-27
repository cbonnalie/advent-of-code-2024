package main

import (
	"math"
)

var seenStones = map[int][]int{
	0: {1},
}

//func applyRules(stone int) []int {
//
//	if seenStones[stone] != nil {
//		return seenStones[stone]
//	}
//
//	if stone == 0 {
//		return []int{1}
//	} else if even, num := isEven(stone); even {
//		lenDiv2 := intPow(10, num)
//		left, right := divAndRem(stone, lenDiv2)
//		seenStones[stone] = []int{left, right}
//		return []int{left, right}
//	} else {
//		seenStones[stone] = []int{stone * 2024}
//		return []int{stone * 2024}
//	}
//}

func calculateNumStones(stone, blinks int) int {
	if blinks == 0 {
		return 1
	}

	transformedStones := applyRules(stone)
	sum := 0
	for _, num := range transformedStones {
		sum += calculateNumStones(num, blinks-1)
	}
	return sum
}

func isEven(number int) (bool, int) {

	digits := 0

	for number != 0 {
		number /= 10
		digits++
	}

	return digits%2 == 0, digits / 2
}

func divAndRem(x, y int) (int, int) {
	return x / y, x % y
}

func intPow[T int8 | int16 | int | int32 | int64](a, b T) T {

	return T(math.Pow(float64(a), float64(b)))

}

func partOne(stones []int) int {
	total := 0
	for _, stone := range stones {
		total += calculateNumStones(stone, 25)
	}
	return total
}

func partTwo(stones []int) int {
	total := 0
	for _, stone := range stones {
		total += calculateNumStones(stone, 75)
	}
	return total
}
