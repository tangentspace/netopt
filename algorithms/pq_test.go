package algorithms

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestNodeDistanceQueue_Pop(t *testing.T) {
	labels := map[NodeId]float64{
		0: 3, 1: 2, 2: 4,
	}

	pq := make(NodeDistanceLabelQueue, len(labels))
	i := 0
	for nodeId, distance := range labels {
		pq[i] = &NodeDistanceLabel{nodeId, distance}
		i++
	}
	heap.Init(&pq)

	// Insert a new label and then modify its priority.
	label := &NodeDistanceLabel{
		nodeId:   3,
		distance: 1,
	}
	heap.Push(&pq, label)

	// Take the labels out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		label := heap.Pop(&pq).(*NodeDistanceLabel)
		fmt.Printf("%.2f:%d ", label.distance, label.nodeId)
	}
}
