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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	edgev1 "github.com/meixiezichuan/clustertopology/api/edge/v1"
	versioned "github.com/meixiezichuan/clustertopology/generated/clientset/versioned"
	internalinterfaces "github.com/meixiezichuan/clustertopology/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/meixiezichuan/clustertopology/generated/listers/edge/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterTopologySpecInformer provides access to a shared informer and lister for
// ClusterTopologySpecs.
type ClusterTopologySpecInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterTopologySpecLister
}

type clusterTopologySpecInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterTopologySpecInformer constructs a new informer for ClusterTopologySpec type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterTopologySpecInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterTopologySpecInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterTopologySpecInformer constructs a new informer for ClusterTopologySpec type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterTopologySpecInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdgeV1().ClusterTopologySpecs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdgeV1().ClusterTopologySpecs(namespace).Watch(context.TODO(), options)
			},
		},
		&edgev1.ClusterTopologySpec{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterTopologySpecInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterTopologySpecInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterTopologySpecInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&edgev1.ClusterTopologySpec{}, f.defaultInformer)
}

func (f *clusterTopologySpecInformer) Lister() v1.ClusterTopologySpecLister {
	return v1.NewClusterTopologySpecLister(f.Informer().GetIndexer())
}
