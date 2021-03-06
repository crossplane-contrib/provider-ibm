/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	"context"

	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

// ResolveReferences of this GroupMembership
func (mg *GroupMembership) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessGroupID),
		Reference:    mg.Spec.ForProvider.AccessGroupIDRef,
		Selector:     mg.Spec.ForProvider.AccessGroupIDSelector,
		To:           reference.To{Managed: &AccessGroup{}, List: &AccessGroupList{}},
		Extract:      AccessGroupID(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.accessGroupId")
	}
	mg.Spec.ForProvider.AccessGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccessGroupIDRef = rsp.ResolvedReference
	return nil
}

// ResolveReferences of this AccessGroupRule
func (mg *AccessGroupRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessGroupID),
		Reference:    mg.Spec.ForProvider.AccessGroupIDRef,
		Selector:     mg.Spec.ForProvider.AccessGroupIDSelector,
		To:           reference.To{Managed: &AccessGroup{}, List: &AccessGroupList{}},
		Extract:      AccessGroupID(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.accessGroupId")
	}
	mg.Spec.ForProvider.AccessGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccessGroupIDRef = rsp.ResolvedReference
	return nil
}

// AccessGroupID extracts the resolved AccessGroupID
func AccessGroupID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		cr, ok := mg.(*AccessGroup)
		if !ok {
			return ""
		}
		return cr.Status.AtProvider.ID
	}
}
