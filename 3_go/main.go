package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.SplitSeq(strings.TrimSpace(string(data)), "\n")

	sum1, sum2 := 0, 0

	for line := range lines {
		var as_chars = strings.Split(line, "")
		nums := make([]int, len(as_chars))
		for i, f := range as_chars {
			nums[i], _ = strconv.Atoi(f)
		}

		sum1 += get_max_n_jolts(nums, 2)
		sum2 += get_max_n_jolts(nums, 12)
	}

	fmt.Printf("part 1: %d\n", sum1)
	fmt.Printf("part 2: %d\n", sum2)
}

const line_length = 100

func get_max_n_jolts(bank []int, n int) int {
	var maxes = make([]int, n)
	start_index_offset := line_length - n

	for i, jolt := range bank {
		var start int
		if i <= start_index_offset {
			start = 0
		} else {
			start = i - start_index_offset
		}

		clear_remaining := false

		for j := start; j < n; j++ {
			if !clear_remaining {
				max := &maxes[j]
				if jolt > *max {
					*max = jolt
					clear_remaining = true
				}
			} else {
				maxes[j] = 0
			}
		}
	}

	return sum_digit_array(maxes)
}

func sum_digit_array(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum *= 10
		sum += n
	}
	return sum
}
