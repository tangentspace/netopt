package main

import (
	"flag"
	"fmt"
	"github.com/tangentspace/netopt/datautil"
)

func main() {
	inputFile := flag.String("input", "", "Input file name")
	algorithm := flag.String("algorithm", "", "Which algorithm to use.")
	sourceNode := flag.Uint64("source", 0, "Source node")

	flag.Parse()

	fmt.Printf("Loading file: %s\n", *inputFile)
	fmt.Printf("Using algorithm: %s\n", *algorithm)

	g, err := datautil.LoadDIMACSGraphFile(*inputFile)
	if err != nil {
		panic("Unable to load graph file.")
	}

	switch *algorithm {
	case "dijkstra":
		sol := g.DijkstraShortestPathTree(*sourceNode)
		fmt.Println("Solution data:")
		fmt.Println(sol)
	}
}
