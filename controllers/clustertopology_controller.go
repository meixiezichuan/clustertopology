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

package controllers

import (
	"context"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	edgev1 "github.com/meixiezichuan/clustertopology/api/edge/v1"
)

// ClusterTopologyReconciler reconciles a ClusterTopology object
type ClusterTopologyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//+kubebuilder:rbac:groups=edge.fdse.lab,resources=clustertopologies,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=edge.fdse.lab,resources=clustertopologies/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=edge.fdse.lab,resources=clustertopologies/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=nodes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=nodes/status,verbs=get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ClusterTopology object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *ClusterTopologyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// get clustertopology
	var c edgev1.ClusterTopology
	err := r.Get(ctx, req.NamespacedName, &c)
	if err != nil {
		if errors.IsNotFound(err) {
			r.Log.Info("CR ClusterTopology not exist")
			return ctrl.Result{}, nil
		}
		r.Log.Error(err, "error on getting cr clustertopology")
		return ctrl.Result{}, err
	}

	// check node list
	nodes := v1.NodeList{}
	listErr := r.List(ctx, &nodes)
	if listErr != nil {
		return ctrl.Result{}, listErr
	}
	var availableNodes []string
	for _, node := range nodes.Items {
		for _, condition := range node.Status.Conditions {
			if condition.Type == "Ready" && condition.Reason == "True" {
				availableNodes = append(availableNodes, node.Name)
			}
		}
	}

	originList := c.GetNetOriginList()
	var newOriginList edgev1.OriginList
	for _, o := range originList {
		if contains(availableNodes, o.Origin) {
			newOrigin := edgev1.OriginInfo{Origin: o.Origin, CostList: edgev1.CostList{}}
			for _, cost := range o.CostList {
				if contains(availableNodes, cost.Destination) {
					newOrigin.CostList = append(newOrigin.CostList, cost)
				}
			}
			newOriginList = append(newOriginList, newOrigin)
		}
	}
	c.SetNetOriginList(newOriginList)
	err = r.Update(ctx, &c)
	if err != nil {
		r.Log.Error(err, "error on update cr clustertopology")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ClusterTopologyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&edgev1.ClusterTopology{}).
		Owns(&v1.Node{}).
		Complete(r)
}
