package main

import (
	"cmp"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := readFileToStr("input.txt")
	lines := slices.Collect(strings.Lines(input))
	solve(lines)
}

type Vec3 struct {
	x, y, z int
}

// a "graph" that is useful for this problem
// maintains a map of what nodes belong to what groups
// and some different sizes, which help with knowing when we're done connecting nodes for part 2
// notably doesn't maintain any data about edges
type Graph struct {
	I                int // incrementer for labeling groups of nodes
	Groups           map[Vec3]int
	NodeCount        int
	GroupCount       int
	DesiredNodeCount int
}

func (g *Graph) Connect(a, b Vec3) int {
	a_val, a_ok := g.Groups[a]
	b_val, b_ok := g.Groups[b]
	if a_ok && b_ok {
		if a_val == b_val {
			return 0
		}
		// need to combine the groups
		// make all b nodes go to a's group
		for n, x := range g.Groups {
			if x == b_val {
				g.Groups[n] = a_val
			}
		}

		// tracking we collapsed one region into another
		g.GroupCount -= 1

	} else if a_ok {
		// need to assign b to a's group
		g.Groups[b] = a_val
		g.NodeCount += 1
	} else if b_ok {
		// need to assign a to b's group
		g.Groups[a] = b_val
		g.NodeCount += 1
	} else {
		// need to assign a and b to a new group (increment I)
		g.I += 1
		g.Groups[a] = g.I
		g.Groups[b] = g.I

		// tracking that we created a new region, and added two never before seen nodes
		g.GroupCount += 1
		g.NodeCount += 2
	}

	// part 2 answer

	if g.GroupCount != 1 || g.NodeCount != g.DesiredNodeCount {
		return 0
	}

	return a.x * b.x
}

// solve part 1
func (g Graph) Solve() int {
	// iterate over kvps, and count for each group ID in map
	freq := make(map[int]int)
	for _, v := range g.Groups {
		freq[v] += 1
	}

	counts := slices.Collect(maps.Values(freq))
	slices.Sort(counts)
	slices.Reverse(counts)

	product := 1
	for _, v := range counts[:3] {
		product *= v
	}

	return product
}

func (a Vec3) Dist(b Vec3) float64 {
	abx := b.x - a.x
	abx *= abx

	aby := b.y - a.y
	aby *= aby

	abz := b.z - a.z
	abz *= abz

	return math.Sqrt(float64(abx + aby + abz))
}

func ParseVec3(line string) Vec3 {
	d := strings.Split(strings.TrimSpace(line), ",")
	x, _ := strconv.Atoi(d[0])
	y, _ := strconv.Atoi(d[1])
	z, _ := strconv.Atoi(d[2])
	return Vec3{x, y, z}
}

type Coupling struct {
	a, b Vec3
	dist float64
}

// solves part 1 and 2 as part of the same iterative process
func solve(lines []string) {
	all_nodes := make([]Vec3, len(lines))
	for i, line := range lines {
		all_nodes[i] = ParseVec3(line)
	}

	graph := Graph{
		I:                0,
		Groups:           make(map[Vec3]int),
		DesiredNodeCount: len(lines),
	}
	pairs := make([]Coupling, get_pair_count(len(lines)))

	idx := 0
	for i, a := range all_nodes {
		for j, b := range all_nodes {
			if j <= i {
				continue
			}
			dist := a.Dist(b)
			pairs[idx] = Coupling{a, b, dist}
			idx += 1
		}
	}

	slices.SortFunc(pairs, func(a, b Coupling) int {
		return cmp.Compare(a.dist, b.dist)
	})

	for _, c := range pairs[:1000] {
		graph.Connect(c.a, c.b)
	}

	solution := graph.Solve()
	fmt.Printf("part 1: %d\n", solution)

	// after finding the first answer, we cont. connecting nodes until we get part 2 answer back from Connect

	for _, c := range pairs[1000:] {
		answer := graph.Connect(c.a, c.b)
		if answer > 0 {
			fmt.Printf("part 2: %d\n", answer)
			break
		}
	}
}

func get_pair_count(n int) int {
	return n * (n - 1) / 2
}

func readFileToStr(fname string) string {
	data, _ := os.ReadFile(fname)
	return string(data)
}
