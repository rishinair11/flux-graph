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
