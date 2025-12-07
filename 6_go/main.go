package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := readFileToStr("input.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	lines, ops := parseInput(input)
	results := parseInts(strings.Fields(lines[0]))
	lines = lines[1:]

	for _, line := range lines {
		nums := parseInts(strings.Fields(line))

		for i, num := range nums {
			results[i] = applyOp(
				ops[i],
				results[i],
				num)
		}
	}

	sum := sum(results)
	fmt.Printf("part 1: %d\n", sum)
}

func part2(input string) {
	vertical_nums := getVerticalNums(input)
	_, ops := parseInput(input)

	group_count := len(ops)

	subtotals := make([]int, group_count)
	group_idx := 0
	for _, num := range vertical_nums {
		if num == 0 {
			group_idx += 1
			continue
		}

		subtotals[group_idx] = applyOp(
			ops[group_idx],
			subtotals[group_idx],
			num)
	}

	sum := sum(subtotals)
	fmt.Printf("part 2: %d", sum)
}

// returns array of ints that represent the vertical nums
// in the input, and where 0s represent empty columns
func getVerticalNums(table string) []int {
	width := strings.Index(table, "\n")
	height := strings.Count(table, "\n") - 1

	nums := make([]int, width)

	for col := range width {
		sum := 0
		for row := range height {
			idx := row*(width+1) + col
			ch := table[idx]

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

//////////
// helpers

func parseInput(input string) (lines []string, ops []string) {
	lines = slices.Collect(strings.Lines(input))
	ops = strings.Fields(lines[len(lines)-1])
	lines = lines[:len(lines)-1]
	return lines, ops
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func applyOp(op string, acc int, val int) int {
	if op == "*" {
		if acc == 0 {
			acc = 1
		}
		return acc * val
	}

	// otherwise '+'
	return acc + val
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
