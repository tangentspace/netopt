package algorithms

import "sort"

type NodeId int64

type NodeIdSlice []NodeId

func (n NodeIdSlice) Sorted() NodeIdSlice {
	n_sorted := make([]NodeId, len(n))
	copy(n_sorted, n)
	sort.Sort(NodeIdSlice(n_sorted))
	return n_sorted
}

func (n NodeIdSlice) Len() int {
	return len(n)
}
func (n NodeIdSlice) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n NodeIdSlice) Less(i, j int) bool {
	return n[i] < n[j]
}

type Arc struct {
	head   NodeId
	tail   NodeId
	weight float64
}

type WeightedGraph struct {
	NumNodes   uint64
	NumArcs    uint64
	Arcs       map[NodeId][]Arc
	Nodes      NodeIdSlice
	nodeExists map[NodeId]bool
}

func NewWeightedGraph() WeightedGraph {
	var g WeightedGraph
	g.Arcs = make(map[NodeId][]Arc)
	g.NumArcs = 0
	g.NumNodes = 0
	g.nodeExists = make(map[NodeId]bool)
	return g
}

func NewWeightedGraphFromArcs(arcs []Arc) WeightedGraph {
	g := NewWeightedGraph()
	for _, a := range arcs {
		g.InsertArc(a)
	}
	return g
}

func (g *WeightedGraph) InsertArc(arc Arc) {
	for _, n := range []NodeId{arc.head, arc.tail} {
		if !g.nodeExists[n] {
			g.Nodes = append(g.Nodes, n)
			g.nodeExists[n] = true
			g.NumNodes++
		}
	}

	g.Arcs[arc.head] = append(g.Arcs[arc.head], arc)
	g.NumArcs++
}

func (g *WeightedGraph) Insert(head, tail uint64, weight float64) {
	g.InsertArc(Arc{NodeId(head), NodeId(tail), weight})
}
