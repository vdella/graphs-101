package graph

import (
	"parzival/graphs-101/src/graph"
	"testing"
)

const expected = "Expected %v, got %v"

func Test_element(t *testing.T) {
	testNode := graph.NewNode(1)
	if testNode.GetElement() != 1 {
		t.Errorf(expected, 1, testNode.GetElement())
	}
}

func Test_zero_neighbors(t *testing.T) {
	testNode := graph.NewNode(1)
	if len(testNode.GetNeighbors()) != 0 {
		t.Errorf(expected, true, len(testNode.GetNeighbors()))
	}
}