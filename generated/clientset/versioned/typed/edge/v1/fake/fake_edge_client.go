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
	v1 "github.com/meixiezichuan/clustertopology/generated/clientset/versioned/typed/edge/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeExampleV1 struct {
	*testing.Fake
}

func (c *FakeExampleV1) ClusterTopologies(namespace string) v1.ClusterTopologyInterface {
	return &FakeClusterTopologies{c, namespace}
}

func (c *FakeExampleV1) ClusterTopologyLists(namespace string) v1.ClusterTopologyListInterface {
	return &FakeClusterTopologyLists{c, namespace}
}

func (c *FakeExampleV1) ClusterTopologySpecs(namespace string) v1.ClusterTopologySpecInterface {
	return &FakeClusterTopologySpecs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeExampleV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
