/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
