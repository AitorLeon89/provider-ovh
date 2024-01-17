// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProjectNetworkPrivateInitParameters struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	VlanID *float64 `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type ProjectNetworkPrivateObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	RegionsAttributes []RegionsAttributesObservation `json:"regionsAttributes,omitempty" tf:"regions_attributes,omitempty"`

	RegionsStatus []RegionsStatusObservation `json:"regionsStatus,omitempty" tf:"regions_status,omitempty"`

	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	VlanID *float64 `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type ProjectNetworkPrivateParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// Service name of the resource representing the id of the cloud project.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Optional
	VlanID *float64 `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type RegionsAttributesInitParameters struct {
}

type RegionsAttributesObservation struct {
	Openstackid *string `json:"openstackid,omitempty" tf:"openstackid,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RegionsAttributesParameters struct {
}

type RegionsStatusInitParameters struct {
}

type RegionsStatusObservation struct {
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RegionsStatusParameters struct {
}

// ProjectNetworkPrivateSpec defines the desired state of ProjectNetworkPrivate
type ProjectNetworkPrivateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectNetworkPrivateParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ProjectNetworkPrivateInitParameters `json:"initProvider,omitempty"`
}

// ProjectNetworkPrivateStatus defines the observed state of ProjectNetworkPrivate.
type ProjectNetworkPrivateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectNetworkPrivateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectNetworkPrivate is the Schema for the ProjectNetworkPrivates API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type ProjectNetworkPrivate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ProjectNetworkPrivateSpec   `json:"spec"`
	Status ProjectNetworkPrivateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectNetworkPrivateList contains a list of ProjectNetworkPrivates
type ProjectNetworkPrivateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectNetworkPrivate `json:"items"`
}

// Repository type metadata.
var (
	ProjectNetworkPrivate_Kind             = "ProjectNetworkPrivate"
	ProjectNetworkPrivate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectNetworkPrivate_Kind}.String()
	ProjectNetworkPrivate_KindAPIVersion   = ProjectNetworkPrivate_Kind + "." + CRDGroupVersion.String()
	ProjectNetworkPrivate_GroupVersionKind = CRDGroupVersion.WithKind(ProjectNetworkPrivate_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectNetworkPrivate{}, &ProjectNetworkPrivateList{})
}
