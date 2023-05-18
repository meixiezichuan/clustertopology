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

// ClusterTopologyListLister helps list ClusterTopologyLists.
// All objects returned here must be treated as read-only.
type ClusterTopologyListLister interface {
	// List lists all ClusterTopologyLists in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterTopologyList, err error)
	// ClusterTopologyLists returns an object that can list and get ClusterTopologyLists.
	ClusterTopologyLists(namespace string) ClusterTopologyListNamespaceLister
	ClusterTopologyListListerExpansion
}

// clusterTopologyListLister implements the ClusterTopologyListLister interface.
type clusterTopologyListLister struct {
	indexer cache.Indexer
}

// NewClusterTopologyListLister returns a new ClusterTopologyListLister.
func NewClusterTopologyListLister(indexer cache.Indexer) ClusterTopologyListLister {
	return &clusterTopologyListLister{indexer: indexer}
}

// List lists all ClusterTopologyLists in the indexer.
func (s *clusterTopologyListLister) List(selector labels.Selector) (ret []*v1.ClusterTopologyList, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterTopologyList))
	})
	return ret, err
}

// ClusterTopologyLists returns an object that can list and get ClusterTopologyLists.
func (s *clusterTopologyListLister) ClusterTopologyLists(namespace string) ClusterTopologyListNamespaceLister {
	return clusterTopologyListNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterTopologyListNamespaceLister helps list and get ClusterTopologyLists.
// All objects returned here must be treated as read-only.
type ClusterTopologyListNamespaceLister interface {
	// List lists all ClusterTopologyLists in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterTopologyList, err error)
	// Get retrieves the ClusterTopologyList from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterTopologyList, error)
	ClusterTopologyListNamespaceListerExpansion
}

// clusterTopologyListNamespaceLister implements the ClusterTopologyListNamespaceLister
// interface.
type clusterTopologyListNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterTopologyLists in the indexer for a given namespace.
func (s clusterTopologyListNamespaceLister) List(selector labels.Selector) (ret []*v1.ClusterTopologyList, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterTopologyList))
	})
	return ret, err
}

// Get retrieves the ClusterTopologyList from the indexer for a given namespace and name.
func (s clusterTopologyListNamespaceLister) Get(name string) (*v1.ClusterTopologyList, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clustertopologylist"), name)
	}
	return obj.(*v1.ClusterTopologyList), nil
}
