/*
Copyright 2021 The Kupenstack Authors.

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

package cluster

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kstypes "github.com/kupenstack/kupenstack/apis/cluster/v1alpha1"
)

// ProfileReconciler reconciles a Profile object
type ProfileReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cluster.kupenstack.io,resources=profiles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cluster.kupenstack.io,resources=profiles/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cluster.kupenstack.io,resources=profiles/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Profile object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *ProfileReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("profile", req.NamespacedName)

	// your logic here
	// var cr kstypes.Profile
	// err := r.Get(ctx, req.NamespacedName, &cr)
	// if err != nil {
	// 	return ctrl.Result{}, client.IgnoreNotFound(err)
	// }

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "cluster.kupenstack.io",
		Kind:    "Profile",
		Version: "v1alpha1",
	})
	err := r.Get(ctx, req.NamespacedName, u)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{RequeueAfter: 1000000000}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProfileReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kstypes.Profile{}).
		Complete(r)
}
