package graph

import (
	"testing"

	"github.com/rishnai1/flux-ks-graph/pkg/resource"
	"github.com/stretchr/testify/assert"
)

func Test_setNodeColorAndStyle(t *testing.T) {
	type args struct {
		attrs map[string]string
		kind  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should set color to blue for Kustomization",
			args: args{
				attrs: make(map[string]string),
				kind:  "Kustomization",
			},
			want: "blue",
		},
		{
			name: "should set color to brown for HelmRelease",
			args: args{
				attrs: make(map[string]string),
				kind:  "HelmRelease",
			},
			want: "brown",
		},
		{
			name: "should set color to green for GitRepository",
			args: args{
				attrs: make(map[string]string),
				kind:  "GitRepository",
			},
			want: "green",
		},
		{
			name: "should set color to green for HelmRepository",
			args: args{
				attrs: make(map[string]string),
				kind:  "HelmRepository",
			},
			want: "green",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAttrs := setAttrsColorAndStyle(tt.args.attrs, tt.args.kind)
			got, ok := gotAttrs["color"]
			assert.True(t, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getName(t *testing.T) {
	tests := []struct {
		name     string
		resource resource.Resource
		want     string
	}{
		{
			name: "should return name with kind and namespace",
			resource: resource.Resource{
				Name:      "name",
				Namespace: "namespace",
				GroupKind: resource.GroupKind{
					Kind: "kind",
				},
			},
			want: "\"namespace/kind/name\"",
		},
		{
			name: "should return name with only kind",
			resource: resource.Resource{
				Name: "name",
				GroupKind: resource.GroupKind{
					Kind: "kind",
				},
			},
			want: "\"kind/name\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getName(tt.resource))
		})
	}
}
