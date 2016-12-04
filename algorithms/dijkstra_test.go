package algorithms_test

import (
	"github.com/tangentspace/netopt/algorithms"
	"reflect"
	"testing"
)

func generateGraph2() algorithms.WeightedGraph {
	g := algorithms.NewWeightedGraph()
	g.Insert(1, 3, 4)
	g.Insert(1, 2, 6)
	g.Insert(2, 3, 2)
	g.Insert(2, 4, 2)
	g.Insert(3, 4, 1)
	g.Insert(3, 5, 2)
	g.Insert(4, 6, 7)
	g.Insert(5, 4, 1)
	g.Insert(5, 6, 3)
	return g
}

func TestWeightedGraph_DijkstraShortestPathTree(t *testing.T) {
	g := generateGraph2()
	result := g.DijkstraShortestPathTree(1)

	expectedDistValues := map[algorithms.NodeId]float64{
		1: 0,
		2: 6,
		3: 4,
		4: 5,
		5: 6,
		6: 9,
	}

	expectedPred := map[algorithms.NodeId]algorithms.NodeId{
		1: 1,
		2: 1,
		3: 1,
		4: 3,
		5: 3,
		6: 5,
	}

	if !reflect.DeepEqual(result.DistValues, expectedDistValues) {
		t.Errorf("\nResult: %v\nExpected: %v\n", result.DistValues, expectedDistValues)
	}

	if !reflect.DeepEqual(result.Pred, expectedPred) {
		t.Errorf("\nResult: %v\nExpected: %v\n", result.Pred, expectedPred)
	}
}
