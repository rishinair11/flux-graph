package graph

import (
	"fmt"
	"log"

	"github.com/awalterschulze/gographviz"
	"github.com/golang-collections/collections/stack"
	"github.com/rishnai1/flux-ks-graph/pkg/resource"
)

// ProcessGraph creates a Graphviz graph from the given ResourceTree
func ProcessGraph(t *resource.ResourceTree) *gographviz.Graph {
	graphName := "fluxgraph"
	g := gographviz.NewGraph()
	g.SetDir(true)
	g.SetName(graphName)
	g.AddAttr(graphName, string(gographviz.RankDir), "LR")
	g.AddAttr(graphName, string(gographviz.RankSep), "5.0")

	// Initialize stack and process the root nodes
	s := stack.New()
	processRootNodes(g, t, s)

	// Process the graph nodes
	processOtherNodes(g, s)

	return g
}

// processRootNodes processes the root nodes of the resource tree
func processRootNodes(g *gographviz.Graph, root *resource.ResourceTree, s *stack.Stack) {
	// Add root node to graph
	g.AddNode(g.Name, getName(root.Resource), getNodeAttrs(root.Resource, root.Resource))
	for _, child := range root.Resources {
		// Add child nodes of root node to graph
		g.AddNode(g.Name, getName(child.Resource), getNodeAttrs(child.Resource, root.Resource))
		// Add edge from root node to child node
		g.AddEdge(getName(root.Resource), getName(child.Resource), true, setAttrsColorAndStyle(make(map[string]string), root.Resource.GetKind()))

		s.Push(child)
	}
}

// processOtherNodes processes the nodes in the graph using a stack
func processOtherNodes(g *gographviz.Graph, s *stack.Stack) {
	for {
		poppedResource := s.Pop()
		if poppedResource == nil {
			log.Println("Done!")
			break
		}

		parent, ok := poppedResource.(resource.ResourceTree)
		if !ok {
			log.Fatalf("error during popping resource from stack")
		}

		for _, child := range parent.Resources {
			// Add child node of parent node to graph
			g.AddNode(g.Name, getName(child.Resource), getNodeAttrs(child.Resource, parent.Resource))
			// Add edge from parent node to child node
			g.AddEdge(getName(parent.Resource), getName(child.Resource), true, setAttrsColorAndStyle(make(map[string]string), parent.Resource.GetKind()))

			s.Push(child)
		}
	}
}

// getName returns the formatted name for a resource
func getName(resource resource.Resource) string {
	if resource.Namespace == "" {
		return fmt.Sprintf("\"%s/%s\"", resource.GroupKind.Kind, resource.Name)
	}
	return fmt.Sprintf("\"%s/%s/%s\"", resource.Namespace, resource.GroupKind.Kind, resource.Name)
}

// getNodeAttrs returns the attributes for a graph node
func getNodeAttrs(child, parent resource.Resource) map[string]string {
	childNamespaceRow := ""
	if child.Namespace != "" {
		childNamespaceRow = fmt.Sprintf(`<tr><td><b>namespace:</b></td><td>%s</td></tr>`, child.Namespace)
	}

	parentNamespaceRow := ""
	if parent.Namespace != "" {
		parentNamespaceRow = fmt.Sprintf(`<tr><td>ownerNamespace:</td><td>%s</td></tr>`, parent.Namespace)
	}

	attrs := map[string]string{
		"shape": "none",
		"label": fmt.Sprintf(`<<table border="2" cellborder="0">
                    <tr><td><b>kind:</b></td><td>%s</td></tr>
                    <tr><td><b>name:</b></td><td>%s</td></tr>
                    %s
					<hr/>
					<tr><td>ownerKind:</td><td>%s</td></tr>
					<tr><td>ownerName:</td><td>%s</td></tr>
					%s
                  </table>>`, child.GroupKind.Kind, child.Name, childNamespaceRow, parent.GroupKind.Kind, parent.Name, parentNamespaceRow),
	}

	return setAttrsColorAndStyle(attrs, child.GroupKind.Kind)
}

// setAttrsColorAndStyle sets the color and style of the node based on the kind
func setAttrsColorAndStyle(attrs map[string]string, kind string) map[string]string {
	switch kind {
	case "Kustomization":
		attrs["color"] = "blue"
		attrs["style"] = "bold"
	case "HelmRelease":
		attrs["color"] = "brown"
		attrs["style"] = "bold"
	case "GitRepository", "HelmRepository":
		attrs["color"] = "green"
		attrs["style"] = "bold"
	}
	return attrs
}
