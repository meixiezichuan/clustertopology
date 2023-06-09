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
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterTopologySpecs implements ClusterTopologySpecInterface
type FakeClusterTopologySpecs struct {
	Fake *FakeEdgeV1
	ns   string
}

var clustertopologyspecsResource = schema.GroupVersionResource{Group: "edge.fdse.lab", Version: "v1", Resource: "clustertopologyspecs"}

var clustertopologyspecsKind = schema.GroupVersionKind{Group: "edge.fdse.lab", Version: "v1", Kind: "ClusterTopologySpec"}

// Get takes name of the clusterTopologySpec, and returns the corresponding clusterTopologySpec object, and an error if there is any.
func (c *FakeClusterTopologySpecs) Get(ctx context.Context, name string, options v1.GetOptions) (result *edgev1.ClusterTopologySpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustertopologyspecsResource, c.ns, name), &edgev1.ClusterTopologySpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopologySpec), err
}

// List takes label and field selectors, and returns the list of ClusterTopologySpecs that match those selectors.
func (c *FakeClusterTopologySpecs) List(ctx context.Context, opts v1.ListOptions) (result *edgev1.ClusterTopologySpecList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustertopologyspecsResource, clustertopologyspecsKind, c.ns, opts), &edgev1.ClusterTopologySpecList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopologySpecList), err
}

// Watch returns a watch.Interface that watches the requested clusterTopologySpecs.
func (c *FakeClusterTopologySpecs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustertopologyspecsResource, c.ns, opts))

}

// Create takes the representation of a clusterTopologySpec and creates it.  Returns the server's representation of the clusterTopologySpec, and an error, if there is any.
func (c *FakeClusterTopologySpecs) Create(ctx context.Context, clusterTopologySpec *edgev1.ClusterTopologySpec, opts v1.CreateOptions) (result *edgev1.ClusterTopologySpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustertopologyspecsResource, c.ns, clusterTopologySpec), &edgev1.ClusterTopologySpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopologySpec), err
}

// Update takes the representation of a clusterTopologySpec and updates it. Returns the server's representation of the clusterTopologySpec, and an error, if there is any.
func (c *FakeClusterTopologySpecs) Update(ctx context.Context, clusterTopologySpec *edgev1.ClusterTopologySpec, opts v1.UpdateOptions) (result *edgev1.ClusterTopologySpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustertopologyspecsResource, c.ns, clusterTopologySpec), &edgev1.ClusterTopologySpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopologySpec), err
}

// Delete takes name of the clusterTopologySpec and deletes it. Returns an error if one occurs.
func (c *FakeClusterTopologySpecs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clustertopologyspecsResource, c.ns, name), &edgev1.ClusterTopologySpec{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTopologySpecs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustertopologyspecsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &edgev1.ClusterTopologySpecList{})
	return err
}

// Patch applies the patch and returns the patched clusterTopologySpec.
func (c *FakeClusterTopologySpecs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *edgev1.ClusterTopologySpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustertopologyspecsResource, c.ns, name, pt, data, subresources...), &edgev1.ClusterTopologySpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopologySpec), err
}
