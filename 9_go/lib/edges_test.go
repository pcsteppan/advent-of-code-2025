package lib

import (
	"slices"
	"testing"
)

func TestGetEdges(t *testing.T) {
	tests := []struct {
		name     string
		edges    []Edge
		expected []Edge
	}{
		{
			"| |",
			Edges(Both, Both),
			Edges(Both, Both),
		},
		{
			"up up",
			Edges(Up, Up),
			Edges(Up, Up),
		},
		{
			"up down",
			Edges(Up, Down),
			nil,
		},
		{
			"down up",
			Edges(Down, Up),
			nil,
		},
		{
			"down down",
			Edges(Down, Down),
			Edges(Down, Down),
		},
		// more complex
		{
			"down down down down",
			Edges(Down, Down, Down, Down),
			Edges(Down, Down, Down, Down),
		},
		{
			"up up up up",
			Edges(Up, Up, Up, Up),
			Edges(Up, Up, Up, Up),
		},
		{
			"| up down | |",
			Edges(Both, Up, Down, Both, Both),
			Edges(Both, Down, Both, Both),
		},
		// double curve.. tricky
		{
			"down up down up",
			Edges(Down, Up, Down, Up),
			Edges(Down, Up),
		},
		{
			"down up up down",
			Edges(Down, Up, Up, Down),
			Edges(Down, Down),
		},
		{
			"| down up",
			Edges(Both, Down, Up),
			Edges(Both, Up),
		},
		{
			"up down down down |",
			Edges(Up, Down, Down, Down, Both),
			Edges(Up, Both),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := GetEdges(tc.edges)
			if !slices.EqualFunc(result, tc.expected, func(a, b Edge) bool {
				return a.Type == b.Type
			}) {
				t.Errorf("Did not match")
			}
		})
	}
}

func Edges(edges ...EdgeType) []Edge {
	arr := make([]Edge, len(edges))
	for i, e := range edges {
		arr[i] = Edge{e, i}
	}
	return arr
}
