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

type IdentityUserInitParameters struct {

	// User description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// User's email
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// User's group
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// User's login suffix
	Login *string `json:"login,omitempty" tf:"login,omitempty"`
}

type IdentityUserObservation struct {

	// Creation date of this user
	Creation *string `json:"creation,omitempty" tf:"creation,omitempty"`

	// User description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// User's email
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// User's group
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Last update of this user
	LastUpdate *string `json:"lastUpdate,omitempty" tf:"last_update,omitempty"`

	// User's login suffix
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// When the user changed his password for the last time
	PasswordLastUpdate *string `json:"passwordLastUpdate,omitempty" tf:"password_last_update,omitempty"`

	// Current user's status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Urn *string `json:"urn,omitempty" tf:"urn,omitempty"`
}

type IdentityUserParameters struct {

	// User description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// User's email
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// User's group
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// User's login suffix
	// +kubebuilder:validation:Optional
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// User's password
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`
}

// IdentityUserSpec defines the desired state of IdentityUser
type IdentityUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityUserParameters `json:"forProvider"`
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
	InitProvider IdentityUserInitParameters `json:"initProvider,omitempty"`
}

// IdentityUserStatus defines the observed state of IdentityUser.
type IdentityUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityUser is the Schema for the IdentityUsers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type IdentityUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.login) || (has(self.initProvider) && has(self.initProvider.login))",message="spec.forProvider.login is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	Spec   IdentityUserSpec   `json:"spec"`
	Status IdentityUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityUserList contains a list of IdentityUsers
type IdentityUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityUser `json:"items"`
}

// Repository type metadata.
var (
	IdentityUser_Kind             = "IdentityUser"
	IdentityUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityUser_Kind}.String()
	IdentityUser_KindAPIVersion   = IdentityUser_Kind + "." + CRDGroupVersion.String()
	IdentityUser_GroupVersionKind = CRDGroupVersion.WithKind(IdentityUser_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityUser{}, &IdentityUserList{})
}
