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

type PrivatedatabaseDatabaseInitParameters struct {

	// The internal name of your private database
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type PrivatedatabaseDatabaseObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The internal name of your private database
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type PrivatedatabaseDatabaseParameters struct {

	// The internal name of your private database
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// PrivatedatabaseDatabaseSpec defines the desired state of PrivatedatabaseDatabase
type PrivatedatabaseDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivatedatabaseDatabaseParameters `json:"forProvider"`
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
	InitProvider PrivatedatabaseDatabaseInitParameters `json:"initProvider,omitempty"`
}

// PrivatedatabaseDatabaseStatus defines the observed state of PrivatedatabaseDatabase.
type PrivatedatabaseDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivatedatabaseDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivatedatabaseDatabase is the Schema for the PrivatedatabaseDatabases API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type PrivatedatabaseDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   PrivatedatabaseDatabaseSpec   `json:"spec"`
	Status PrivatedatabaseDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivatedatabaseDatabaseList contains a list of PrivatedatabaseDatabases
type PrivatedatabaseDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivatedatabaseDatabase `json:"items"`
}

// Repository type metadata.
var (
	PrivatedatabaseDatabase_Kind             = "PrivatedatabaseDatabase"
	PrivatedatabaseDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivatedatabaseDatabase_Kind}.String()
	PrivatedatabaseDatabase_KindAPIVersion   = PrivatedatabaseDatabase_Kind + "." + CRDGroupVersion.String()
	PrivatedatabaseDatabase_GroupVersionKind = CRDGroupVersion.WithKind(PrivatedatabaseDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivatedatabaseDatabase{}, &PrivatedatabaseDatabaseList{})
}
