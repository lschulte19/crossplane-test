// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/lschulte19/crossplane-test/apis/identity/v1alpha1"
	v1alpha11 "github.com/lschulte19/crossplane-test/apis/kms/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Bucket.
func (mg *Bucket) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha11.KeyList{},
			Managed: &v1alpha11.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.CompartmentIDRef,
		Selector:     mg.Spec.InitProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CompartmentID")
	}
	mg.Spec.InitProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.KMSKeyIDRef,
		Selector:     mg.Spec.InitProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha11.KeyList{},
			Managed: &v1alpha11.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyID")
	}
	mg.Spec.InitProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Object.
func (mg *Object) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OpcSseKMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OpcSseKMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.OpcSseKMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha11.KeyList{},
			Managed: &v1alpha11.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OpcSseKMSKeyID")
	}
	mg.Spec.ForProvider.OpcSseKMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OpcSseKMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.OpcSseKMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.OpcSseKMSKeyIDRef,
		Selector:     mg.Spec.InitProvider.OpcSseKMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha11.KeyList{},
			Managed: &v1alpha11.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.OpcSseKMSKeyID")
	}
	mg.Spec.InitProvider.OpcSseKMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.OpcSseKMSKeyIDRef = rsp.ResolvedReference

	return nil
}
