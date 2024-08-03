package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResource_GetKind(t *testing.T) {
	tests := []struct {
		name     string
		resource Resource
		want     string
	}{
		{
			name: "should return resource kind",
			resource: Resource{
				GroupKind: GroupKind{
					Kind: "dummy",
				},
			},
			want: "dummy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Resource{
				GroupKind: tt.resource.GroupKind,
				Name:      tt.resource.Name,
				Namespace: tt.resource.Namespace,
			}
			assert.Equal(t, tt.want, r.GetKind())
		})
	}
}

func TestNewResourceTree(t *testing.T) {
	testPath := "testdata/test.yaml"

	want := &ResourceTree{
		Resource: Resource{
			GroupKind: GroupKind{
				Group: "parent-group",
				Kind:  "parent-kind",
			},
			Name:      "parent-name",
			Namespace: "parent-namespace",
		},
		Resources: []ResourceTree{
			{
				Resource: Resource{
					GroupKind: GroupKind{
						Group: "child-group",
						Kind:  "child-kind",
					},
					Name:      "child-name",
					Namespace: "child-namespace",
				},
				Resources: []ResourceTree{
					{
						Resource: Resource{
							GroupKind: GroupKind{
								Group: "grandchild-group",
								Kind:  "grandchild-kind",
							},
							Name:      "grandchild-name",
							Namespace: "grandchild-namespace",
						},
					},
				},
			},
		},
	}

	got, err := NewResourceTree(testPath)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
