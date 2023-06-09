/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	edgev1 "github.com/meixiezichuan/clustertopology/api/edge/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterTopologies implements ClusterTopologyInterface
type FakeClusterTopologies struct {
	Fake *FakeEdgeV1
	ns   string
}

var clustertopologiesResource = schema.GroupVersionResource{Group: "edge.fdse.lab", Version: "v1", Resource: "clustertopologies"}

var clustertopologiesKind = schema.GroupVersionKind{Group: "edge.fdse.lab", Version: "v1", Kind: "ClusterTopology"}

// Get takes name of the clusterTopology, and returns the corresponding clusterTopology object, and an error if there is any.
func (c *FakeClusterTopologies) Get(ctx context.Context, name string, options v1.GetOptions) (result *edgev1.ClusterTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustertopologiesResource, c.ns, name), &edgev1.ClusterTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopology), err
}

// List takes label and field selectors, and returns the list of ClusterTopologies that match those selectors.
func (c *FakeClusterTopologies) List(ctx context.Context, opts v1.ListOptions) (result *edgev1.ClusterTopologyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustertopologiesResource, clustertopologiesKind, c.ns, opts), &edgev1.ClusterTopologyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &edgev1.ClusterTopologyList{ListMeta: obj.(*edgev1.ClusterTopologyList).ListMeta}
	for _, item := range obj.(*edgev1.ClusterTopologyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterTopologies.
func (c *FakeClusterTopologies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustertopologiesResource, c.ns, opts))

}

// Create takes the representation of a clusterTopology and creates it.  Returns the server's representation of the clusterTopology, and an error, if there is any.
func (c *FakeClusterTopologies) Create(ctx context.Context, clusterTopology *edgev1.ClusterTopology, opts v1.CreateOptions) (result *edgev1.ClusterTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustertopologiesResource, c.ns, clusterTopology), &edgev1.ClusterTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopology), err
}

// Update takes the representation of a clusterTopology and updates it. Returns the server's representation of the clusterTopology, and an error, if there is any.
func (c *FakeClusterTopologies) Update(ctx context.Context, clusterTopology *edgev1.ClusterTopology, opts v1.UpdateOptions) (result *edgev1.ClusterTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustertopologiesResource, c.ns, clusterTopology), &edgev1.ClusterTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopology), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterTopologies) UpdateStatus(ctx context.Context, clusterTopology *edgev1.ClusterTopology, opts v1.UpdateOptions) (*edgev1.ClusterTopology, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clustertopologiesResource, "status", c.ns, clusterTopology), &edgev1.ClusterTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopology), err
}

// Delete takes name of the clusterTopology and deletes it. Returns an error if one occurs.
func (c *FakeClusterTopologies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clustertopologiesResource, c.ns, name), &edgev1.ClusterTopology{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTopologies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustertopologiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &edgev1.ClusterTopologyList{})
	return err
}

// Patch applies the patch and returns the patched clusterTopology.
func (c *FakeClusterTopologies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *edgev1.ClusterTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustertopologiesResource, c.ns, name, pt, data, subresources...), &edgev1.ClusterTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopology), err
}
