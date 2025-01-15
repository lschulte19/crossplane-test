// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	bucket "github.com/lschulte19/crossplane-test/internal/controller/objectstorage/bucket"
	object "github.com/lschulte19/crossplane-test/internal/controller/objectstorage/object"
	objectlifecyclepolicy "github.com/lschulte19/crossplane-test/internal/controller/objectstorage/objectlifecyclepolicy"
	providerconfig "github.com/lschulte19/crossplane-test/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		object.Setup,
		objectlifecyclepolicy.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
