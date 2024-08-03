package resource

import (
	"github.com/rishinair11/flux-ks-graph/pkg/util"
	"gopkg.in/yaml.v3"
)

type ResourceTree struct {
	Resource  Resource       `yaml:"resource"`
	Resources []ResourceTree `yaml:"resources"`
}

// NewResourceTree parses a Flux 'tree' YAML file and returns a ResourceTree
func NewResourceTree(fileName string) (*ResourceTree, error) {
	yamlBytes, err := util.ReadInput(fileName)
	if err != nil {
		return nil, err
	}

	rt := &ResourceTree{}
	if err := yaml.Unmarshal(yamlBytes, rt); err != nil {
		return nil, err
	}

	return rt, nil
}

type Resource struct {
	GroupKind GroupKind `yaml:"GroupKind"`
	Name      string    `yaml:"Name"`
	Namespace string    `yaml:"Namespace"`
}

type GroupKind struct {
	Group string `yaml:"Group"`
	Kind  string `yaml:"Kind"`
}

func (r Resource) GetKind() string {
	return r.GroupKind.Kind
}
