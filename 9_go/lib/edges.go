package lib

import (
	"errors"
	"fmt"
)

type EdgeType int

const (
	None EdgeType = iota
	Up
	Down
	Both
)

type ShapePosition int

const (
	Out ShapePosition = iota
	In
	OnEdge_WasInside
	OnEdge_WasOutside
)

type Edge struct {
	Type EdgeType
	Id   int
}

func GetEdges(edges []Edge) ([]Edge, error) {
	if len(edges) < 2 {
		return nil, errors.New("Invalid edges, must have at least two edges")
	}

	result := make([]Edge, 0)
	state := Out
	prev := None
	for _, edge := range edges {
		switch state {
		case Out:
			switch edge.Type {
			case Up:
				state = OnEdge_WasOutside
			case Down:
				state = OnEdge_WasOutside
			case Both:
				state = In
			}
			// transitionType = edge.Type
			prev = edge.Type
			result = append(result, edge)
		case In:
			// if we're inside
			// then we can transition to outside
			// or on edge (was inside)
			if edge.Type == Both {
				result = append(result, edge)
				state = Out
			} else {
				state = OnEdge_WasInside
				prev = edge.Type
			}
		// if we're on an edge, then closing the edge
		// results in either going inside or outside
		case OnEdge_WasInside:
			if edge.Type == Both {
				return nil, errors.New("Unexpected Both edge found")
			}

			if edge.Type == prev {
				state = In
			} else {
				// found complementing edge
				// e.g, up for down, or down for up
				// which transitions state to inside
				prev = None
				state = Out
				result = append(result, edge)
			}

		case OnEdge_WasOutside:
			if edge.Type == Both {
				return nil, fmt.Errorf("Unexpected Both edge found with ID: %d", edge.Id)
			}

			if edge.Type == prev {
				prev = None
				state = Out
				result = append(result, edge)
			} else {
				// found complementing edge
				// e.g, up for down, or down for up
				// which transitions state to inside
				prev = edge.Type
				state = In
			}
		}
	}

	if len(result)%2 == 1 || state != Out {
		return nil, errors.New("Edges don't define a bounding shape")
	}

	return result, nil
}
