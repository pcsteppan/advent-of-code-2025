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
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
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
	}

	fmt.Printf("part 1: %d\n", split_count)
}

func part2(lines []string) {
	rays := make(map[int]int)
	start_ray_pos := strings.Index(lines[0], "S")
	rays[start_ray_pos] = 1

	for _, line := range lines {
		// iterate over positions of ^
		for i, c := range line {
			if c == rune('^') && rays[i] > 0 {
				paths := rays[i]
				rays[i] = 0
				rays[i-1] += paths
				rays[i+1] += paths
			}
		}
	}

	// sum final path counts
	sum := 0
	for _, n := range rays {
		sum += n
	}

	fmt.Printf("part 1: %d\n", sum)
}

func readFileToStr(fname string) string {
	data, _ := os.ReadFile(fname)
	return string(data)
}
