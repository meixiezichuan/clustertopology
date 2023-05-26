/*
Copyright 2023.

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

package v1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const NetTopolgKey = "topology.kubernetes.io/network"

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ClusterTopology is the Schema for the clustertopologies API
type ClusterTopology struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterTopologySpec   `json:"spec,omitempty"`
	Status ClusterTopologyStatus `json:"status,omitempty"`
}

// ClusterTopologySpec defines the desired state of ClusterTopology
type ClusterTopologySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Cluster name
	Cluster string `json:"cluster"`

	// Topologys
	Topologys TopologyList `json:"topologys" protobuf:"bytes,1,opt,name=weights,casttype=TopologyList"`
}

// ClusterTopologyStatus defines the observed state of ClusterTopology
type ClusterTopologyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The total number of nodes in the cluster
	// +kubebuilder:validation:Minimum=0
	NodeCount int64 `json:"nodeCount,omitempty" protobuf:"bytes,1,opt,name=nodeCount"`

	// The calculation time for the weights in the  topology CRD
	CalculationTime metav1.Time `json:"weightCalculationTime,omitempty" protobuf:"bytes,2,opt,name=weightCalculationTime"`
}

//+kubebuilder:object:root=true

// ClusterTopologyList contains a list of ClusterTopology
type ClusterTopologyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterTopology `json:"items"`
}

// TopologyKey is the key of a OriginList in a NetworkTopology.
type TopologyKey string

// TopologyList contains an array of OriginInfo objects.
// +protobuf=true
type TopologyList []TopologyInfo

// TopologyInfo contains information about network costs for a particular Topology Key.
// +protobuf=true
type TopologyInfo struct {
	// Topology key (e.g., "topology.kubernetes.io/network").
	// +required
	TopologyKey TopologyKey `json:"topologyKey" protobuf:"bytes,1,opt,name=topologyKey"` // add as enum instead of string

	// OriginList for a particular origin.
	// +required
	OriginList OriginList `json:"originList" protobuf:"bytes,2,rep,name=originList,casttype=OriginList"`
}

// OriginList contains an array of OriginInfo objects.
// +protobuf=true
type OriginList []OriginInfo

type Properties map[string]string

// OriginInfo contains information about network costs for a particular Origin.
// +protobuf=true
type OriginInfo struct {
	// Name of the origin (e.g., Region Name, Zone Name).
	// +required
	Origin string `json:"origin" protobuf:"bytes,1,opt,name=origin"`

	// properties of the origin`
	Properties Properties `json:"properties,omitempty"`

	// Costs for the particular origin.
	CostList CostList `json:"costList,omitempty" protobuf:"bytes,2,rep,name=costList,casttype=CostList"`
}

// CostList contains an array of CostInfo objects.
// +protobuf=true
type CostList []CostInfo

// CostInfo contains information about networkCosts.
// +protobuf=true
type CostInfo struct {
	// Name of the destination (e.g., Region Name, Zone Name).
	// +required
	Destination string `json:"destination" protobuf:"bytes,1,opt,name=destination"`

	// Bandwidth capacity between origin and destination.
	// +optional
	BandwidthCapacity resource.Quantity `json:"bandwidthCapacity,omitempty" protobuf:"bytes,2,opt,name=bandwidthCapacity"`

	// Bandwidth allocated between origin and destination.
	// +optional
	BandwidthAllocated resource.Quantity `json:"bandwidthAllocated,omitempty" protobuf:"bytes,3,opt,name=bandwidthAllocated"`

	// Network Cost between origin and destination (e.g., Dijkstra shortest path, etc)
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Default=0
	// +required
	NetworkCost int64 `json:"networkCost" protobuf:"bytes,4,opt,name=networkCost"`
}

func (c *ClusterTopology) GetNetOriginList() OriginList {
	for _, t := range c.Spec.Topologys {
		if t.TopologyKey == NetTopolgKey {
			return t.OriginList
		}
	}
	return nil
}

func (c *ClusterTopology) SetNetOriginList(originList OriginList) {
	for _, t := range c.Spec.Topologys {
		if t.TopologyKey == NetTopolgKey {
			t.OriginList = originList
		}
	}
}

func (c *ClusterTopology) GetNetOriginInfoByNode(node string) *OriginInfo {
	var originInfo OriginInfo
	for _, t := range c.Spec.Topologys {
		if t.TopologyKey == NetTopolgKey {
			for _, o := range t.OriginList {
				if o.Origin == node {
					originInfo = *o.DeepCopy()
					return &originInfo
				}
			}
		}
	}
	return nil
}

func (c *ClusterTopology) SetNetOriginInfo(originInfo *OriginInfo) {
	for _, t := range c.Spec.Topologys {
		if t.TopologyKey == NetTopolgKey {
			for i, o := range t.OriginList {
				if o.Origin == originInfo.Origin {
					t.OriginList[i] = *originInfo
				}
			}
		}
	}
}

func init() {
	SchemeBuilder.Register(&ClusterTopology{}, &ClusterTopologyList{})
}
