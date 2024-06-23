package main

import (
	"io"
	"log"
	"os"

	"github.com/rishinair11/flux-ks-graph/pkg/graph"
	"github.com/rishinair11/flux-ks-graph/pkg/resource"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func main() {
	var inputFile string
	var outputFile string

	rootCmd := &cobra.Command{
		Use:   "flux-graph",
		Short: "Processes a Flux Kustomization tree and generates a graph",
		Run: func(cmd *cobra.Command, _ []string) {
			var yamlBytes []byte
			var err error

			// Read YAML input
			yamlBytes, err = readInput(inputFile)
			if err != nil {
				log.Fatalf("Failed to read YAML: %v", err)
			}

			// Unmarshal YAML into ResourceTree
			t := &resource.ResourceTree{}
			if err := yaml.Unmarshal(yamlBytes, t); err != nil {
				log.Fatalf("Failed to unmarshal YAML: %v", err)
			}

			// Process the graph
			graph, err := graph.ProcessGraph(t)
			if err != nil {
				log.Fatal(err)
			}

			// Write the graph to the output file
			if err := os.WriteFile(outputFile, []byte(graph.String()), 0o755); err != nil {
				log.Fatalf("Failed to write output file: %v", err)
			}

			log.Println("Generated graph:", outputFile)
		},
	}

	rootCmd.Flags().StringVarP(&inputFile, "file", "f", "", "Specify input file")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "graph.dot", "Specify output file")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// readInput reads the YAML input from a file or stdin
func readInput(inputFile string) ([]byte, error) {
	if inputFile != "" {
		log.Println("Reading from file:", inputFile)
		return os.ReadFile(inputFile)
	}
	log.Println("Reading from STDIN...")
	return io.ReadAll(os.Stdin)
}
