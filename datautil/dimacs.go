package datautil

import (
	"bufio"
	"fmt"
	"github.com/tangentspace/netopt/algorithms"
	"os"
	"strconv"
	"strings"
)

func LoadDIMACSGraphFile(fileName string) (algorithms.WeightedGraph, error) {
	var numNodes, numArcs uint64

	g := algorithms.NewWeightedGraph()

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		switch line[0] {

		// Problem line
		case 'p':
			fields := strings.Fields(line)
			numNodes, err = strconv.ParseUint(fields[2], 10, 64)
			if err != nil {
				panic(err)
			}
			numArcs, err = strconv.ParseUint(fields[3], 10, 64)
			if err != nil {
				panic(err)
			}

		// Arc descriptor line
		case 'a':
			fields := strings.Fields(line)
			head, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				panic(err)
			}
			tail, err := strconv.ParseUint(fields[2], 10, 64)
			if err != nil {
				panic(err)
			}
			weight, err := strconv.ParseFloat(fields[3], 64)
			if err != nil {
				panic(err)
			}
			g.Insert(head, tail, weight)
		}
	}

	if g.NumNodes != numNodes || g.NumArcs != numArcs {
		message := "Error loading file: expected %d nodes and %d arcs, but file contained %d nodes and %d arcs."
		return g, fmt.Errorf(message, numNodes, numArcs, g.NumNodes, g.NumArcs)
	}

	return g, nil
}
