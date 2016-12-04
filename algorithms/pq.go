package algorithms

type NodeDistanceLabel struct {
	nodeId   NodeId
	distance float64
}

type NodeDistanceLabelQueue []*NodeDistanceLabel

func (pq NodeDistanceLabelQueue) Len() int { return len(pq) }

func (pq NodeDistanceLabelQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq NodeDistanceLabelQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *NodeDistanceLabelQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*NodeDistanceLabel))
}

func (pq *NodeDistanceLabelQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	label := old[n-1]
	*pq = old[0 : n-1]
	return label
}
