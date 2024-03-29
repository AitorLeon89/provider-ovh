// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ProjectKubeIprestrictions
func (mg *ProjectKubeIprestrictions) GetTerraformResourceType() string {
	return "ovh_cloud_project_kube_iprestrictions"
}

// GetConnectionDetailsMapping for this ProjectKubeIprestrictions
func (tr *ProjectKubeIprestrictions) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ProjectKubeIprestrictions
func (tr *ProjectKubeIprestrictions) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ProjectKubeIprestrictions
func (tr *ProjectKubeIprestrictions) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ProjectKubeIprestrictions
func (tr *ProjectKubeIprestrictions) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ProjectKubeIprestrictions
func (tr *ProjectKubeIprestrictions) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ProjectKubeIprestrictions
func (tr *ProjectKubeIprestrictions) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ProjectKubeIprestrictions
func (tr *ProjectKubeIprestrictions) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this ProjectKubeIprestrictions using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ProjectKubeIprestrictions) LateInitialize(attrs []byte) (bool, error) {
	params := &ProjectKubeIprestrictionsParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ProjectKubeIprestrictions) GetTerraformSchemaVersion() int {
	return 0
}
