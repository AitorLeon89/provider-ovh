/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this InstallationTemplatePartitionSchemePartition.
func (mg *InstallationTemplatePartitionSchemePartition) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this InstallationTemplatePartitionSchemePartition.
func (mg *InstallationTemplatePartitionSchemePartition) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this InstallationTemplatePartitionSchemePartition.
func (mg *InstallationTemplatePartitionSchemePartition) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this InstallationTemplatePartitionSchemePartition.
func (mg *InstallationTemplatePartitionSchemePartition) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this InstallationTemplatePartitionSchemePartition.
func (mg *InstallationTemplatePartitionSchemePartition) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this InstallationTemplatePartitionSchemePartition.
func (mg *InstallationTemplatePartitionSchemePartition) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this InstallationTemplatePartitionSchemePartition.
func (mg *InstallationTemplatePartitionSchemePartition) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this InstallationTemplatePartitionSchemePartition.
func (mg *InstallationTemplatePartitionSchemePartition) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this InstallationTemplatePartitionSchemePartition.
func (mg *InstallationTemplatePartitionSchemePartition) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this InstallationTemplatePartitionSchemePartition.
func (mg *InstallationTemplatePartitionSchemePartition) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this InstallationTemplatePartitionSchemePartition.
func (mg *InstallationTemplatePartitionSchemePartition) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this InstallationTemplatePartitionSchemePartition.
func (mg *InstallationTemplatePartitionSchemePartition) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
