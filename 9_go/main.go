package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Vec2 struct {
	x, y int
}

func createVec2(line string) Vec2 {
	xy := strings.Split(strings.TrimSpace(line), ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	return Vec2{x, y}
}

func main() {
	input := readFileToStr("input.txt")
	lines := slices.Collect(strings.Lines(input))

	data := make([]Vec2, len(lines))
	for i, line := range lines {
		data[i] = createVec2(line)
	}

	part1(&data)
}

func area(a, b Vec2) int {
	res := (a.x - b.x + 1) * (a.y - b.y + 1)
	if res < 0 {
		res *= -1
	}
	return res
}

func part1(data *[]Vec2) {
	largest := math.MinInt
	for i, a := range *data {
		for _, b := range (*data)[i:] {
			area := area(a, b)
			if area > largest {
				largest = area
			}
		}
	}
	fmt.Printf("part 1: %d\n", largest)
}

func readFileToStr(fname string) string {
	data, _ := os.ReadFile(fname)
	return string(data)
}
