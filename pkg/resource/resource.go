package resource

type ResourceTree struct {
	Resource  Resource       `yaml:"resource"`
	Resources []ResourceTree `yaml:"resources"`
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
