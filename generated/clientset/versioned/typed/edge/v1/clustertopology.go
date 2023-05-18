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

package v1

import (
	"context"
	"time"

	v1 "github.com/meixiezichuan/clustertopology/api/edge/v1"
	scheme "github.com/meixiezichuan/clustertopology/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterTopologiesGetter has a method to return a ClusterTopologyInterface.
// A group's client should implement this interface.
type ClusterTopologiesGetter interface {
	ClusterTopologies(namespace string) ClusterTopologyInterface
}

// ClusterTopologyInterface has methods to work with ClusterTopology resources.
type ClusterTopologyInterface interface {
	Create(ctx context.Context, clusterTopology *v1.ClusterTopology, opts metav1.CreateOptions) (*v1.ClusterTopology, error)
	Update(ctx context.Context, clusterTopology *v1.ClusterTopology, opts metav1.UpdateOptions) (*v1.ClusterTopology, error)
	UpdateStatus(ctx context.Context, clusterTopology *v1.ClusterTopology, opts metav1.UpdateOptions) (*v1.ClusterTopology, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ClusterTopology, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClusterTopologyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterTopology, err error)
	ClusterTopologyExpansion
}

// clusterTopologies implements ClusterTopologyInterface
type clusterTopologies struct {
	client rest.Interface
	ns     string
}

// newClusterTopologies returns a ClusterTopologies
func newClusterTopologies(c *ExampleV1Client, namespace string) *clusterTopologies {
	return &clusterTopologies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterTopology, and returns the corresponding clusterTopology object, and an error if there is any.
func (c *clusterTopologies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterTopology, err error) {
	result = &v1.ClusterTopology{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustertopologies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterTopologies that match those selectors.
func (c *clusterTopologies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterTopologyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterTopologyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustertopologies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterTopologies.
func (c *clusterTopologies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clustertopologies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterTopology and creates it.  Returns the server's representation of the clusterTopology, and an error, if there is any.
func (c *clusterTopologies) Create(ctx context.Context, clusterTopology *v1.ClusterTopology, opts metav1.CreateOptions) (result *v1.ClusterTopology, err error) {
	result = &v1.ClusterTopology{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clustertopologies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTopology).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterTopology and updates it. Returns the server's representation of the clusterTopology, and an error, if there is any.
func (c *clusterTopologies) Update(ctx context.Context, clusterTopology *v1.ClusterTopology, opts metav1.UpdateOptions) (result *v1.ClusterTopology, err error) {
	result = &v1.ClusterTopology{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustertopologies").
		Name(clusterTopology.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTopology).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterTopologies) UpdateStatus(ctx context.Context, clusterTopology *v1.ClusterTopology, opts metav1.UpdateOptions) (result *v1.ClusterTopology, err error) {
	result = &v1.ClusterTopology{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustertopologies").
		Name(clusterTopology.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTopology).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterTopology and deletes it. Returns an error if one occurs.
func (c *clusterTopologies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustertopologies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterTopologies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustertopologies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterTopology.
func (c *clusterTopologies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterTopology, err error) {
	result = &v1.ClusterTopology{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clustertopologies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}