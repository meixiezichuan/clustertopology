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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/meixiezichuan/clustertopology/api/edge/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterTopologySpecLister helps list ClusterTopologySpecs.
// All objects returned here must be treated as read-only.
type ClusterTopologySpecLister interface {
	// List lists all ClusterTopologySpecs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterTopologySpec, err error)
	// ClusterTopologySpecs returns an object that can list and get ClusterTopologySpecs.
	ClusterTopologySpecs(namespace string) ClusterTopologySpecNamespaceLister
	ClusterTopologySpecListerExpansion
}

// clusterTopologySpecLister implements the ClusterTopologySpecLister interface.
type clusterTopologySpecLister struct {
	indexer cache.Indexer
}

// NewClusterTopologySpecLister returns a new ClusterTopologySpecLister.
func NewClusterTopologySpecLister(indexer cache.Indexer) ClusterTopologySpecLister {
	return &clusterTopologySpecLister{indexer: indexer}
}

// List lists all ClusterTopologySpecs in the indexer.
func (s *clusterTopologySpecLister) List(selector labels.Selector) (ret []*v1.ClusterTopologySpec, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterTopologySpec))
	})
	return ret, err
}

// ClusterTopologySpecs returns an object that can list and get ClusterTopologySpecs.
func (s *clusterTopologySpecLister) ClusterTopologySpecs(namespace string) ClusterTopologySpecNamespaceLister {
	return clusterTopologySpecNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterTopologySpecNamespaceLister helps list and get ClusterTopologySpecs.
// All objects returned here must be treated as read-only.
type ClusterTopologySpecNamespaceLister interface {
	// List lists all ClusterTopologySpecs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterTopologySpec, err error)
	// Get retrieves the ClusterTopologySpec from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterTopologySpec, error)
	ClusterTopologySpecNamespaceListerExpansion
}

// clusterTopologySpecNamespaceLister implements the ClusterTopologySpecNamespaceLister
// interface.
type clusterTopologySpecNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterTopologySpecs in the indexer for a given namespace.
func (s clusterTopologySpecNamespaceLister) List(selector labels.Selector) (ret []*v1.ClusterTopologySpec, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterTopologySpec))
	})
	return ret, err
}

// Get retrieves the ClusterTopologySpec from the indexer for a given namespace and name.
func (s clusterTopologySpecNamespaceLister) Get(name string) (*v1.ClusterTopologySpec, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clustertopologyspec"), name)
	}
	return obj.(*v1.ClusterTopologySpec), nil
}
