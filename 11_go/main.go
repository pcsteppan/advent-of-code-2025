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
	graph := make(map[string][]string, 0)
	for _, line := range lines {
		split := strings.Split(line, ":")
		from := strings.TrimSpace(split[0])
		to := strings.Split(split[1], " ")
		cnxs := make([]string, len(to))
		for _, t := range to {
			trimmed := strings.TrimSpace(t)
			if trimmed == "" {
				continue
			}
			cnxs = append(cnxs, trimmed)
		}
		graph[from] = cnxs
	}

	part1(graph)
}

func part1(data map[string][]string) {
	q := make([]string, 0)
	visited := make(map[string]bool, 0)
	q = append(q, "you")
	reachedOuts := 0
	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		visited[head] = true
		cnxs := data[head]
		for _, cnx := range cnxs {
			if cnx == "out" {
				reachedOuts++
				continue
			}
			if !visited[cnx] {
				q = append(q, cnx)
			}
		}
	}

	fmt.Println("part 1:", reachedOuts)
}

func readFileToStr(fname string) string {
	data, _ := os.ReadFile(fname)
	return string(data)
}
