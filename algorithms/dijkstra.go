package algorithms

import (
	"bytes"
	"container/heap"
	"fmt"
	"math"
)

type ShortestPathTree struct {
	Nodes      NodeIdSlice
	Pred       map[NodeId]NodeId
	DistValues map[NodeId]float64
}

func (s ShortestPathTree) String() string {
	var buffer bytes.Buffer
	for _, n := range s.Nodes.Sorted() {
		buffer.WriteString(fmt.Sprintf("%d %d %f\n", n, s.Pred[n], s.DistValues[n]))
	}
	return buffer.String()
}

func (g *WeightedGraph) DijkstraShortestPathTree(source uint64) ShortestPathTree {

	pred := make(map[NodeId]NodeId)

	distValues := make(map[NodeId]float64, g.NumNodes)

	distLabelsTemp := make(NodeDistanceLabelQueue, g.NumNodes)

	// Initialize data structures
	pred[NodeId(source)] = NodeId(source)
	var d float64
	for i, nodeId := range g.Nodes {
		if i == 0 {
			d = 0
		} else {
			d = math.Inf(1)
		}
		distLabelsTemp[i] = &NodeDistanceLabel{nodeId, d}
	}
	heap.Init(&distLabelsTemp)

	for distLabelsTemp.Len() > 0 {
		currentNode := distLabelsTemp.Pop().(*NodeDistanceLabel)
		currentNodeId := currentNode.nodeId

		distValues[currentNodeId] = currentNode.distance

		for _, arc := range g.Arcs[currentNodeId] {

			if d0 := distValues[currentNodeId] + arc.weight; d0 < distValues[arc.tail] {
				distValues[arc.tail] = d0
				pred[arc.tail] = currentNodeId
				heap.Push(&distLabelsTemp, &NodeDistanceLabel{arc.tail, d0})
			}
		}

	}

	return ShortestPathTree{g.Nodes, pred, distValues}

}
