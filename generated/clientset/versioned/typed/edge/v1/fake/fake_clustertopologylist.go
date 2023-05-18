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

// FakeClusterTopologyLists implements ClusterTopologyListInterface
type FakeClusterTopologyLists struct {
	Fake *FakeExampleV1
	ns   string
}

var clustertopologylistsResource = schema.GroupVersionResource{Group: "example.my.domain", Version: "v1", Resource: "clustertopologylists"}

var clustertopologylistsKind = schema.GroupVersionKind{Group: "example.my.domain", Version: "v1", Kind: "ClusterTopologyList"}

// Get takes name of the clusterTopologyList, and returns the corresponding clusterTopologyList object, and an error if there is any.
func (c *FakeClusterTopologyLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *edgev1.ClusterTopologyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustertopologylistsResource, c.ns, name), &edgev1.ClusterTopologyList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopologyList), err
}

// List takes label and field selectors, and returns the list of ClusterTopologyLists that match those selectors.
func (c *FakeClusterTopologyLists) List(ctx context.Context, opts v1.ListOptions) (result *edgev1.ClusterTopologyListList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustertopologylistsResource, clustertopologylistsKind, c.ns, opts), &edgev1.ClusterTopologyListList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopologyListList), err
}

// Watch returns a watch.Interface that watches the requested clusterTopologyLists.
func (c *FakeClusterTopologyLists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustertopologylistsResource, c.ns, opts))

}

// Create takes the representation of a clusterTopologyList and creates it.  Returns the server's representation of the clusterTopologyList, and an error, if there is any.
func (c *FakeClusterTopologyLists) Create(ctx context.Context, clusterTopologyList *edgev1.ClusterTopologyList, opts v1.CreateOptions) (result *edgev1.ClusterTopologyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustertopologylistsResource, c.ns, clusterTopologyList), &edgev1.ClusterTopologyList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopologyList), err
}

// Update takes the representation of a clusterTopologyList and updates it. Returns the server's representation of the clusterTopologyList, and an error, if there is any.
func (c *FakeClusterTopologyLists) Update(ctx context.Context, clusterTopologyList *edgev1.ClusterTopologyList, opts v1.UpdateOptions) (result *edgev1.ClusterTopologyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustertopologylistsResource, c.ns, clusterTopologyList), &edgev1.ClusterTopologyList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopologyList), err
}

// Delete takes name of the clusterTopologyList and deletes it. Returns an error if one occurs.
func (c *FakeClusterTopologyLists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clustertopologylistsResource, c.ns, name), &edgev1.ClusterTopologyList{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTopologyLists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustertopologylistsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &edgev1.ClusterTopologyListList{})
	return err
}

// Patch applies the patch and returns the patched clusterTopologyList.
func (c *FakeClusterTopologyLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *edgev1.ClusterTopologyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustertopologylistsResource, c.ns, name, pt, data, subresources...), &edgev1.ClusterTopologyList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edgev1.ClusterTopologyList), err
}
