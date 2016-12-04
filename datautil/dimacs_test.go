package datautil_test

import (
	"github.com/tangentspace/netopt/datautil"
	"testing"
)

func TestLoadDIMACSGraphFile(t *testing.T) {
	g, err := datautil.LoadDIMACSGraphFile("testdata/graph_g1.gr")
	if err != nil {
		t.Errorf(err.Error())
	}

	if g.NumNodes != 6 {
		t.Errorf("Loaded %d nodes, expected %d nodes", g.NumNodes, 5)
	}
	if g.NumArcs != 9 {
		t.Errorf("Loaded %d arcs, expected %d arcs", g.NumArcs, 9)
	}

	_, err = datautil.LoadDIMACSGraphFile("testdata/graph_truncated.gr")
	if err == nil {
		t.Error("Loading truncated graph did not produce error")
	}

}
