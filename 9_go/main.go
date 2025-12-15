package main

import (
	. "daynine/lib"
	"fmt"
	"maps"
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
	part2(data)
}

func area(a, b Vec2) int {
	minx := min(a.x, b.x)
	miny := min(a.y, b.y)
	maxx := max(a.x, b.x)
	maxy := max(a.y, b.y)
	return (maxx - minx + 1) * (maxy - miny + 1)
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

type Grid struct {
	data [][]int
}

func (g *Grid) isInside(v Vec2) bool {
	for i, x := range g.data[v.y] {
		if x > v.x {
			return i%2 == 1
		}
	}
	return false
}

func part2(data []Vec2) {
	tiles := make(map[Vec2]struct{})

	origin := data[0]
	minX, maxX, minY, maxY := origin.x, origin.x, origin.y, origin.y
	tiles[origin] = struct{}{}

	for i, a := range data {
		var b Vec2
		if i == 0 {
			b = data[len(data)-1]
		} else {
			b = data[i-1]
		}

		minX = min(minX, a.x)
		maxX = max(maxX, a.x)
		minY = min(minY, a.y)
		maxY = max(maxY, a.y)

		tiles[a] = struct{}{}
		tiles[b] = struct{}{}
		if a.x == b.x {
			for y := min(a.y, b.y); y <= max(a.y, b.y); y++ {
				tiles[Vec2{x: a.x, y: y}] = struct{}{}
			}
		}
	}

	xValuesOfInterest := make(map[int]struct{})
	for _, v := range data {
		tiles[v] = struct{}{}
		xValuesOfInterest[v.x] = struct{}{}
	}

	xValues := slices.Collect(maps.Keys(xValuesOfInterest))
	slices.Sort(xValues)

	grid := Grid{data: make([][]int, maxY+1)}
	for y := minY; y <= maxY; y++ {
		// convert to format for helper
		edges := make([]Edge, 0)
		for _, x := range xValues {
			edgeType := getEdgeType(Vec2{x, y}, tiles)
			if edgeType == None {
				continue
			}
			edges = append(edges, Edge{edgeType, x})
		}

		if len(edges) == 0 {
			continue
		}

		// reduce row to only outer edges
		outerEdges, err := GetEdges(edges)
		if err != nil {
			fmt.Printf("=== ERROR AT %d ===\n", y)
			fmt.Println(edges)
			panic(err)
		}

		// get indices
		row := make([]int, len(outerEdges))
		for i, e := range outerEdges {
			row[i] = e.Id
		}

		grid.data[y] = row
	}

	largest := 0
	for i, a := range data {
		for _, b := range data[i:] {
			area := area(a, b)
			if area > largest && isValidRect(&a, &b, &grid) {
				largest = area
			}
		}
	}

	fmt.Printf("part 2: %d\n", largest)
}

func getEdgeType(v Vec2, tiles map[Vec2]struct{}) EdgeType {
	_, self := tiles[v]
	if !self {
		return None
	}

	up := Vec2{x: v.x, y: v.y - 1}
	down := Vec2{x: v.x, y: v.y + 1}
	_, isUp := tiles[up]
	_, isDown := tiles[down]

	if isUp && isDown {
		return Both
	}
	if isUp && !isDown {
		return Up
	}
	if !isUp && isDown {
		return Down
	}

	return None
}

func isValidRect(a *Vec2, b *Vec2, grid *Grid) bool {
	start := min(a.x, b.x)
	end := max(a.x, b.x)
	for y := min(a.y, b.y); y <= max(a.y, b.y); y++ {
		if !grid.isValidRow(y, start, end) {
			return false
		}
	}
	return true
}

func (grid *Grid) isValidRow(y, start, end int) bool {
	row := (*grid).data[y]
	if len(row) == 0 {
		return false
	}
	i := findInsertionPoint(row, start)

	if i >= len(row) {
		return false
	}

	return i%2 == 1 && row[i] >= end
}

func findInsertionPoint(arr []int, target int) int {
	lo, hi := 0, len(arr)

	for lo < hi {
		mid := (lo + hi) / 2

		if arr[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}

func readFileToStr(fname string) string {
	data, _ := os.ReadFile(fname)
	return string(data)
}
