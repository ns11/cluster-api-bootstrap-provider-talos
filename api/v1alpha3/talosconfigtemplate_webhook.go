// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package v1alpha3

import (
	"github.com/google/go-cmp/cmp"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	runtime "k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *TalosConfigTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=update,path=/validate-bootstrap-cluster-x-k8s-io-v1alpha3-talosconfigtemplate,mutating=false,failurePolicy=fail,groups=bootstrap.cluster.x-k8s.io,resources=talosconfigtemplates,versions=v1alpha3,name=vtalosconfigtemplate.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &TalosConfigTemplate{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *TalosConfigTemplate) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *TalosConfigTemplate) ValidateUpdate(oldRaw runtime.Object) error {
	old := oldRaw.(*TalosConfigTemplate)

	if !cmp.Equal(r.Spec, old.Spec) {
		return apierrors.NewBadRequest("TalosConfigTemplate.Spec is immutable")
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *TalosConfigTemplate) ValidateDelete() error {
	return nil
}
