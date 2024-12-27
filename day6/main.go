package main

//
//import (
//	"bufio"
//	"io"
//	"os"
//)
//
//var Grid [][]rune
//
//type Coodinate struct {
//	x      int
//	y      int
//	facing Direction
//}
//
//type Direction struct {
//	dirX int
//	dirY int
//}
//
//type Input struct {
//	mat    [][]int
//	width  int
//	height int
//	start  Coodinate
//}
//
//func parseInput(src io.Reader) *Input {
//	i := Input{
//		mat: make([][]int, 0),
//	}
//
//	scanner := bufio.NewScanner(src)
//	vi := 0
//
//	for scanner.Scan() {
//		line := scanner.Bytes()
//		if len(line) == 0 {
//			break
//		}
//
//		if i.width == 0 {
//			i.width = len(string(line))
//		}
//
//		i.mat = append(i.mat, make([]int, i.width))
//
//	}
//}
//
//func main() {
//	file, _ := os.Open("day6/input.txt")
//	defer file.Close()
//	scanner := bufio.NewScanner(file)
//}
