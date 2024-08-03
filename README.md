# Flux Graph

[![PR Build](https://github.com/rishinair11/flux-graph/actions/workflows/pr.yml/badge.svg?branch=main)](https://github.com/rishinair11/flux-graph/actions/workflows/pr.yml)

`flux-graph` is a simple tool that generates a static SVG graph representing the resources managed by Flux Kustomizations and HelmReleases in your Kubernetes cluster managed by Flux GitOps.

## Why
Understanding the relationships between Kubernetes resources managed by Flux can be challenging.\
`flux-graph` provides a visual representation of the dependencies between these resources, helping you to:

- Easily see the structure and dependencies.
- Quickly understand resource relationships.
- Improve overall visibility into your cluster's deployment heirarchy.

## How
`flux-graph` works by:

- Accepts the YAML output of `flux tree ks flux-system -o yaml` either from STDIN or from a YAML file.
- Processes YAML data obtained to create a graph of the relationships between the Flux Kustomizations and child resources.
- Generates a `.svg` file representing the graph structure of the various Flux and Kubernetes resources.
- **NEW!** Serves the SVG file in a local web server using an SVG Image Viewer with zooming and panning capabilities for easier viewing.

## How to Use It

### Prerequisites

- A Kubernetes cluster managed by Flux GitOps.
- [`kubectl`](https://kubernetes.io/docs/tasks/tools/) configured to access your cluster.
- [`flux`](https://fluxcd.io/flux/cmd/)
- [`go`](https://go.dev/doc/install) (only if building from source).

### Installation

#### Clone the Repository:
```bash
git clone https://github.com/rishinair11/flux-graph.git
cd flux-graph
```

#### Build the Application:

```bash
make
```

#### Usage
```console
$ flux-graph -h                                                                                                                  
Processes a Flux Kustomization tree and generates a graph

Usage:
  flux-graph [flags]

Flags:
  -d, --direction string   Specify direction of graph (https://graphviz.gitlab.io/docs/attrs/rankdir) (default "TB")
  -f, --file string        Specify input file
  -h, --help               help for flux-graph
  -n, --no-serve           Don't serve the graph on a web server
  -o, --output string      Specify output file (default "graph.svg")
  -p, --port string        Specify web server port (default "9000")
```

Either pipe the Flux tree output directly to `flux-graph`:
```bash
flux tree ks flux-system -o yaml | flux-graph
```
Or, specify a YAML file to read from:
```bash
flux tree ks flux-system -o yaml > tree.yaml
flux-graph -f tree.yaml
```

Then visit "https://localhost:9000" to view your Flux resources in a SVG graph format.

## License
This project is licensed under the Apache 2.0 License - see the [LICENSE](./LICENSE) file for details.