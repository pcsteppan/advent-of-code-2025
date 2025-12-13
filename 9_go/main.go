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
	part2(&data)
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

// this approach won't work, because some edges may lie in the interior of shapes
// which might mean we have to go back to actually modeling something and creating a set of inner tiles?
// I hope not
func part2(data *[]Vec2) {
	tiles := make(map[Vec2]bool)

	origin := (*data)[0]
	minX, maxX, minY, maxY := origin.x, origin.x, origin.y, origin.y

	// doing 2 things here: writing edges into map, and finding bounding box
	// i still starts at 0 when we range over a slice like so
	for i, a := range *data {
		var b Vec2
		if i == 0 {
			b = (*data)[len(*data)-1]
		} else {
			b = (*data)[i-1]
		}

		minX = min(minX, a.x)
		maxX = max(maxX, a.x)
		minY = min(minY, a.y)
		maxY = max(maxY, a.y)

		if a.x == b.x {
			// iterate through y diff
			for y := min(a.y, b.y); y < max(a.y, b.y); y++ {
				tiles[Vec2{x: a.x, y: y}] = true
			}
		} else {
			// iterate through x diff
			for x := min(a.x, b.x); x < max(a.x, b.x); x++ {
				tiles[Vec2{x: x, y: a.y}] = true
			}
		}
	}

	fmt.Println(minX, minY, maxX, maxY)

	for _, v := range *data {
		tiles[v] = true
	}

	// once we have edges set up, we can fill in the shapes by
	// using edge detection

	fmt.Println("filling")

	for y := minY; y <= maxY; y++ {
		prev := false
		inside := false

		fmt.Printf("%d/%d\r", y, maxY)

		for x := minX; x <= maxX; x++ {
			if (tiles[Vec2{x, y}]) {
				// fmt.Print("#")
				if !prev {
					inside = !inside
				}
				prev = true
			} else {
				if inside {
					tiles[Vec2{x, y}] = true
					// fmt.Print("X")
				} else {
					// fmt.Print(".")
				}
				prev = false
			}
		}
		// fmt.Println()
	}
	// do the same thing as part 1, but verify everything
	// inside is in the tiles set/map
	fmt.Println("")
	fmt.Println("finding")

	largest := 1000
	for i, a := range *data {
		fmt.Printf("%d/%d\r", i, len(*data))
		for _, b := range (*data)[i:] {
			area := area(a, b)
			if area > largest && isValidRect(&a, &b, &tiles) {
				largest = area
			}
		}
	}

	fmt.Printf("part 2: %d\n", largest)
}

func isValidRect(a *Vec2, b *Vec2, lookup *map[Vec2]bool) bool {
	for y := min(a.y, b.y); y < max(a.y, b.y); y++ {
		for x := min(a.x, b.x); x < max(a.x, b.x); x++ {
			if !(*lookup)[Vec2{x: x, y: y}] {
				return false
			}
		}
	}
	return true
}

func readFileToStr(fname string) string {
	data, _ := os.ReadFile(fname)
	return string(data)
}
