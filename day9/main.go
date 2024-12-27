package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var diskMap []string
var diskOne []string
var diskTwo []string

func main() {
	file, _ := os.Open("day9/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	diskMap = strings.Split(scanner.Text(), "")
	writeToDisk()
	defragOne()
	defragTwo()
	fmt.Println(calculateChecksum(diskOne))
	fmt.Println(calculateChecksum(diskTwo))
}

func writeToDisk() {
	diskOne = make([]string, 0)
	diskTwo = make([]string, 0)
	writing := true
	num := 0

	for i := 0; i < len(diskMap); i++ {

		// get the size from the diskOne map
		size, _ := strconv.Atoi(diskMap[i])

		var value string
		if writing {
			// if we're writing, we want to write the current num
			value = strconv.Itoa(num)
		} else {
			// otherwise, add free space with a "."
			value = "."
		}

		// write / allocate free space to the diskOne
		for j := 0; j < size; j++ {
			diskOne = append(diskOne, value)
			diskTwo = append(diskTwo, value)
		}

		// increment num if we wrote
		if writing {
			num++
		}

		writing = !writing
	}
}

func defragOne() {
	l := 0
	r := len(diskOne) - 1

	for {
		for diskOne[l] != "." {
			l++
		}

		for diskOne[r] == "." {
			r--
		}

		if l >= r {
			break
		}

		diskOne[l], diskOne[r] = diskOne[r], diskOne[l]
	}
}

func calculateChecksum(disk []string) int {
	checksum := 0
	for i := 0; i < len(disk); i++ {
		value, _ := strconv.Atoi(disk[i])
		product := value * i
		checksum += product
	}
	return checksum
}

func defragTwo() {
	// Start with the highest file ID and work down
	maxID := -1
	for _, cell := range diskTwo {
		if cell != "." {
			id, _ := strconv.Atoi(cell)
			if id > maxID {
				maxID = id
			}
		}
	}

	// Process files from the highest ID to lowest
	for fileID := maxID; fileID >= 0; fileID-- {
		// Find the file's current position and size
		fileStart := -1
		fileSize := 0

		// Find the start and size of current file
		for i := 0; i < len(diskTwo); i++ {
			if diskTwo[i] == strconv.Itoa(fileID) {
				if fileStart == -1 {
					fileStart = i
				}
				fileSize++
			}
		}

		if fileStart == -1 {
			continue // File not found, skip to next
		}

		// Find leftmost viable position
		currentPos := 0
		bestPos := -1

		for currentPos < fileStart {
			consecutiveSpace := 0

			// Count consecutive free space
			for i := currentPos; i < len(diskTwo) && diskTwo[i] == "."; i++ {
				consecutiveSpace++
				if consecutiveSpace == fileSize {
					bestPos = currentPos
					break
				}
			}

			if bestPos != -1 {
				break
			}

			// Move to next non-space character
			for currentPos < len(diskTwo) && diskTwo[currentPos] == "." {
				currentPos++
			}
			// Skip past non-space character
			if currentPos < len(diskTwo) {
				currentPos++
			}
		}

		// Move file if we found a valid position
		if bestPos != -1 {
			// Clear original position
			fileStr := strconv.Itoa(fileID)
			// Move file to new position
			for i := 0; i < fileSize; i++ {
				diskTwo[bestPos+i] = fileStr
				diskTwo[fileStart+i] = "."
			}
		}
	}
}
