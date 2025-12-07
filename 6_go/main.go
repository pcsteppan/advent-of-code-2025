package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const _example = "123 328  51 64 \r\n 45 64  387 23 \r\n  6 98  215 314\r\n*   +   *   +  \r\n"

func main() {
	input := readFileToStr("input.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := slices.Collect(strings.Lines(input))
	results := parseInts(strings.Fields(lines[0]))
	op_line := strings.Fields(lines[len(lines)-1])
	lines = lines[1 : len(lines)-1]

	for _, line := range lines {
		nums := parseInts(strings.Fields(line))
		for i, num := range nums {
			if op_line[i] == "*" {
				results[i] *= num
			} else {
				results[i] += num
			}
		}
	}

	sum := 0
	for _, n := range results {
		sum += n
	}
	fmt.Printf("part 1: %d\n", sum)
}

func part2(input string) {
	vertical_nums := getVerticalNums(input)
	lines := slices.Collect(strings.Lines(input))
	op_line := strings.Fields(lines[len(lines)-1])

	group_count := 1
	for _, n := range vertical_nums {
		if n == 0 {
			group_count += 1
		}
	}

	subtotals := make([]int, group_count)
	group_idx := 0
	for _, num := range vertical_nums {
		if num == 0 {
			group_idx += 1
			continue
		}

		if op_line[group_idx] == "*" {
			if subtotals[group_idx] == 0 {
				subtotals[group_idx] = 1
			}

			subtotals[group_idx] *= num
		} else {
			subtotals[group_idx] += num
		}
	}

	sum := 0
	for _, num := range subtotals {
		sum += num
	}

	fmt.Printf("part 2: %d", sum)
}

func readFileToStr(fname string) string {
	data, _ := os.ReadFile(fname)
	return string(data)
}

func parseInts(num_strs []string) []int {
	nums := make([]int, len(num_strs))
	for i, str := range num_strs {
		nums[i], _ = strconv.Atoi(str)
	}
	return nums
}

// returns array of ints that represent the vertical nums
// in the input, and where 0s represent empty columns
func getVerticalNums(table string) []int {
	data := strings.ReplaceAll(table, "\r\n", "\n")

	width := strings.Index(data, "\n")
	height := strings.Count(data, "\n") - 1

	nums := make([]int, width)

	for col := range width {
		sum := 0
		for row := range height {
			idx := row*(width+1) + col
			ch := data[idx]

			var n int
			if ch == byte(' ') {
				continue
			} else {
				n = int(ch - '0')
			}

			sum *= 10
			sum += n
		}
		nums[col] = sum
	}

	return nums
}
