package main

import (
	"log"

	"github.com/rishinair11/flux-ks-graph/pkg/graph"
	"github.com/rishinair11/flux-ks-graph/pkg/resource"
	"github.com/rishinair11/flux-ks-graph/pkg/serve"
	"github.com/rishinair11/flux-ks-graph/pkg/util"
	"github.com/spf13/cobra"
)

func main() {
	var (
		inputFile      string
		outputFile     string
		serverPort     string
		graphDirection string
		noServe        bool
	)

	rootCmd := &cobra.Command{
		Use:   "flux-graph",
		Short: "Processes a Flux Kustomization tree and generates a graph",
		Run: func(cmd *cobra.Command, _ []string) {
			rt, err := resource.NewResourceTree(inputFile)
			if err != nil {
				log.Fatalf("Failed to initialize ResourceTree: %v", err)
			}

			// Process the graph
			graph, err := graph.ProcessGraph(rt, graphDirection)
			if err != nil {
				log.Fatalf("Failed to construct graph: %v", err)
			}

			if err := util.GenerateGraphSVG(graph, outputFile); err != nil {
				log.Fatalf("Failed to generate graph SVG: %v", err)
			}

			log.Println("Generated graph:", outputFile)

			if !noServe {
				serve.ServeAssets(outputFile, serverPort)
			}
		},
	}

	rootCmd.Flags().StringVarP(&inputFile, "file", "f", "", "Specify input file")
	rootCmd.Flags().StringVarP(&graphDirection, "direction", "d", "TB", "Specify direction of graph (https://graphviz.gitlab.io/docs/attrs/rankdir)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "graph.svg", "Specify output file")
	rootCmd.Flags().StringVarP(&serverPort, "port", "p", "9000", "Specify web server port")
	rootCmd.Flags().BoolVarP(&noServe, "no-serve", "n", false, "Don't serve the graph on a web server")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
