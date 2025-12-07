package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input := readFileToStr("input.txt")
	lines := slices.Collect(strings.Lines(input))
	p1 := solve(lines)
	fmt.Println(p1)
}

func solve(lines []string) uint {
	rays := make(map[int]bool)
	start_ray_pos := strings.Index(lines[0], "S")
	rays[start_ray_pos] = true

	split_count := 0
	for _, line := range lines {
		// iterate over positions of ^
		for i, c := range line {
			if c == rune('^') && rays[i] {
				rays[i] = false
				split_count += 1
				rays[i-1] = true
				rays[i+1] = true
			}
		}
		// for each intersection, sum += 1
		// and we update the rays set accordingly
	}

	return uint(split_count)
}

// type Point struct {
// 	x int
// 	y int
// }

// type Tree struct {
// 	data map[Point]Point
// }

// func createTree(lines []string) Tree {

// }

func readFileToStr(fname string) string {
	data, _ := os.ReadFile(fname)
	return string(data)
}
