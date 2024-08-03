package util

import (
	"io"
	"log"
	"os"

	"github.com/awalterschulze/gographviz"
	"github.com/goccy/go-graphviz"
)

// readInput reads the YAML input from a file or stdin
func ReadInput(inputFile string) ([]byte, error) {
	if inputFile != "" {
		log.Println("Reading from file:", inputFile)
		return os.ReadFile(inputFile)
	}
	log.Println("Reading from STDIN...")
	return io.ReadAll(os.Stdin)
}

// GenerateGraphSVG takes a gographviz Graph object, gets the DOT bytes,
// converts it into SVG bytes using the goccy/graphviz library and
// writes it to an .svg file
func GenerateGraphSVG(graph *gographviz.Graph, outputFile string) error {
	// TODO: maybe fix this double parsing by generating the graph using goccy/graphviz?
	// convert gographviz graph to DOT bytes
	gvGraph, err := graphviz.ParseBytes([]byte(graph.String()))
	if err != nil {
		log.Fatalf("Failed to parse graph dot string: %v", err)
	}
	defer gvGraph.Close()

	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer f.Close()

	// convert DOT bytes to SVG file
	return graphviz.New().RenderFilename(gvGraph, graphviz.SVG, f.Name())
}
