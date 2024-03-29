/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
