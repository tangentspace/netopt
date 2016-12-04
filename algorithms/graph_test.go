package algorithms

import (
	"reflect"
	"testing"
)

var graphs = map[string][]Arc{
	"g1": {
		{0, 2, 1},
		{0, 1, 1},
		{1, 3, 1},
		{2, 3, 1},
	},
}

func TestNewWeightedGraphFromArcs(t *testing.T) {
	g := NewWeightedGraphFromArcs(graphs["g1"])

	if g.NumNodes != 4 {
		t.Errorf("Loaded %d nodes, expected %d nodes", g.NumNodes, 4)
	}
	if g.NumArcs != 4 {
		t.Errorf("Loaded %d arcs, expected %d arcs", g.NumArcs, 4)
	}

	result := g.Nodes.Sorted()
	expected := NodeIdSlice{0, 1, 2, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nResult: %v\nExpected: %v\n", result, expected)
	}
}
