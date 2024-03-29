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

// GetTerraformResourceType returns Terraform resource type for this LogsCluster
func (mg *LogsCluster) GetTerraformResourceType() string {
	return "ovh_dbaas_logs_cluster"
}

// GetConnectionDetailsMapping for this LogsCluster
func (tr *LogsCluster) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"dedicated_input_pem": "status.atProvider.dedicatedInputPem", "direct_input_pem": "status.atProvider.directInputPem", "initial_archive_allowed_networks[*]": "status.atProvider.initialArchiveAllowedNetworks[*]", "initial_direct_input_allowed_networks[*]": "status.atProvider.initialDirectInputAllowedNetworks[*]", "initial_query_allowed_networks[*]": "status.atProvider.initialQueryAllowedNetworks[*]"}
}

// GetObservation of this LogsCluster
func (tr *LogsCluster) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LogsCluster
func (tr *LogsCluster) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LogsCluster
func (tr *LogsCluster) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LogsCluster
func (tr *LogsCluster) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LogsCluster
func (tr *LogsCluster) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this LogsCluster
func (tr *LogsCluster) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this LogsCluster using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LogsCluster) LateInitialize(attrs []byte) (bool, error) {
	params := &LogsClusterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LogsCluster) GetTerraformSchemaVersion() int {
	return 0
}
