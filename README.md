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
- Generates a `.dot` File representing the graph structure which can be later visualized using [`graphviz`](https://graphviz.gitlab.io/).

## How to Use It

### Prerequisites

- A Kubernetes cluster managed by Flux GitOps.
- [`kubectl`](https://kubernetes.io/docs/tasks/tools/) configured to access your cluster.
- [`flux`](https://fluxcd.io/flux/cmd/)
- [`go`](https://go.dev/doc/install) (only if building from source).
- [`dot`](https://graphviz.gitlab.io/download/) for visualizing the generated `.dot` file.

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
  -f, --file string     Specify input file
  -h, --help            help for flux-graph
  -o, --output string   Specify output file (default "graph.dot")
```

Either pipe the Flux tree output directly to `flux-graph`:
```bash
flux tree ks flux-system -o yaml | flux-graph -o graph.dot
```
Or, specify a YAML file to read from:
```bash
flux tree ks flux-system -o yaml > tree.yaml
flux-graph -f tree.yaml -o graph.dot
```

Then generate an SVG file using the `dot` tool:
```bash
# SVG
dot -Tsvg graph.dot > graph.svg
# PNG
dot -Tpng graph.dot > graph.png
```

## License
This project is licensed under the Apache 2.0 License - see the [LICENSE](./LICENSE) file for details.