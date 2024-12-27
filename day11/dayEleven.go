package main

func applyRules(stone int) []int {

	if seenStones[stone] != nil {
		return seenStones[stone]
	}

	if stone == 0 {
		return []int{1}
	} else if even, num := isEven(stone); even {
		lenDiv2 := intPow(10, num)
		left, right := divAndRem(stone, lenDiv2)
		seenStones[stone] = []int{left, right}
		return []int{left, right}
	} else {
		seenStones[stone] = []int{stone * 2024}
		return []int{stone * 2024}
	}
}
