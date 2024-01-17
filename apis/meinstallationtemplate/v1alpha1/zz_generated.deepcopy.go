//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizationInitParameters) DeepCopyInto(out *CustomizationInitParameters) {
	*out = *in
	if in.ChangeLog != nil {
		in, out := &in.ChangeLog, &out.ChangeLog
		*out = new(string)
		**out = **in
	}
	if in.CustomHostname != nil {
		in, out := &in.CustomHostname, &out.CustomHostname
		*out = new(string)
		**out = **in
	}
	if in.PostInstallationScriptLink != nil {
		in, out := &in.PostInstallationScriptLink, &out.PostInstallationScriptLink
		*out = new(string)
		**out = **in
	}
	if in.PostInstallationScriptReturn != nil {
		in, out := &in.PostInstallationScriptReturn, &out.PostInstallationScriptReturn
		*out = new(string)
		**out = **in
	}
	if in.Rating != nil {
		in, out := &in.Rating, &out.Rating
		*out = new(float64)
		**out = **in
	}
	if in.SSHKeyName != nil {
		in, out := &in.SSHKeyName, &out.SSHKeyName
		*out = new(string)
		**out = **in
	}
	if in.UseDistributionKernel != nil {
		in, out := &in.UseDistributionKernel, &out.UseDistributionKernel
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizationInitParameters.
func (in *CustomizationInitParameters) DeepCopy() *CustomizationInitParameters {
	if in == nil {
		return nil
	}
	out := new(CustomizationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizationObservation) DeepCopyInto(out *CustomizationObservation) {
	*out = *in
	if in.ChangeLog != nil {
		in, out := &in.ChangeLog, &out.ChangeLog
		*out = new(string)
		**out = **in
	}
	if in.CustomHostname != nil {
		in, out := &in.CustomHostname, &out.CustomHostname
		*out = new(string)
		**out = **in
	}
	if in.PostInstallationScriptLink != nil {
		in, out := &in.PostInstallationScriptLink, &out.PostInstallationScriptLink
		*out = new(string)
		**out = **in
	}
	if in.PostInstallationScriptReturn != nil {
		in, out := &in.PostInstallationScriptReturn, &out.PostInstallationScriptReturn
		*out = new(string)
		**out = **in
	}
	if in.Rating != nil {
		in, out := &in.Rating, &out.Rating
		*out = new(float64)
		**out = **in
	}
	if in.SSHKeyName != nil {
		in, out := &in.SSHKeyName, &out.SSHKeyName
		*out = new(string)
		**out = **in
	}
	if in.UseDistributionKernel != nil {
		in, out := &in.UseDistributionKernel, &out.UseDistributionKernel
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizationObservation.
func (in *CustomizationObservation) DeepCopy() *CustomizationObservation {
	if in == nil {
		return nil
	}
	out := new(CustomizationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizationParameters) DeepCopyInto(out *CustomizationParameters) {
	*out = *in
	if in.ChangeLog != nil {
		in, out := &in.ChangeLog, &out.ChangeLog
		*out = new(string)
		**out = **in
	}
	if in.CustomHostname != nil {
		in, out := &in.CustomHostname, &out.CustomHostname
		*out = new(string)
		**out = **in
	}
	if in.PostInstallationScriptLink != nil {
		in, out := &in.PostInstallationScriptLink, &out.PostInstallationScriptLink
		*out = new(string)
		**out = **in
	}
	if in.PostInstallationScriptReturn != nil {
		in, out := &in.PostInstallationScriptReturn, &out.PostInstallationScriptReturn
		*out = new(string)
		**out = **in
	}
	if in.Rating != nil {
		in, out := &in.Rating, &out.Rating
		*out = new(float64)
		**out = **in
	}
	if in.SSHKeyName != nil {
		in, out := &in.SSHKeyName, &out.SSHKeyName
		*out = new(string)
		**out = **in
	}
	if in.UseDistributionKernel != nil {
		in, out := &in.UseDistributionKernel, &out.UseDistributionKernel
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizationParameters.
func (in *CustomizationParameters) DeepCopy() *CustomizationParameters {
	if in == nil {
		return nil
	}
	out := new(CustomizationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplate) DeepCopyInto(out *InstallationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplate.
func (in *InstallationTemplate) DeepCopy() *InstallationTemplate {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplateInitParameters) DeepCopyInto(out *InstallationTemplateInitParameters) {
	*out = *in
	if in.BaseTemplateName != nil {
		in, out := &in.BaseTemplateName, &out.BaseTemplateName
		*out = new(string)
		**out = **in
	}
	if in.Customization != nil {
		in, out := &in.Customization, &out.Customization
		*out = make([]CustomizationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultLanguage != nil {
		in, out := &in.DefaultLanguage, &out.DefaultLanguage
		*out = new(string)
		**out = **in
	}
	if in.RemoveDefaultPartitionSchemes != nil {
		in, out := &in.RemoveDefaultPartitionSchemes, &out.RemoveDefaultPartitionSchemes
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplateInitParameters.
func (in *InstallationTemplateInitParameters) DeepCopy() *InstallationTemplateInitParameters {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplateInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplateList) DeepCopyInto(out *InstallationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstallationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplateList.
func (in *InstallationTemplateList) DeepCopy() *InstallationTemplateList {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplateObservation) DeepCopyInto(out *InstallationTemplateObservation) {
	*out = *in
	if in.AvailableLanguages != nil {
		in, out := &in.AvailableLanguages, &out.AvailableLanguages
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BaseTemplateName != nil {
		in, out := &in.BaseTemplateName, &out.BaseTemplateName
		*out = new(string)
		**out = **in
	}
	if in.Beta != nil {
		in, out := &in.Beta, &out.Beta
		*out = new(bool)
		**out = **in
	}
	if in.BitFormat != nil {
		in, out := &in.BitFormat, &out.BitFormat
		*out = new(float64)
		**out = **in
	}
	if in.Category != nil {
		in, out := &in.Category, &out.Category
		*out = new(string)
		**out = **in
	}
	if in.Customization != nil {
		in, out := &in.Customization, &out.Customization
		*out = make([]CustomizationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultLanguage != nil {
		in, out := &in.DefaultLanguage, &out.DefaultLanguage
		*out = new(string)
		**out = **in
	}
	if in.Deprecated != nil {
		in, out := &in.Deprecated, &out.Deprecated
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Distribution != nil {
		in, out := &in.Distribution, &out.Distribution
		*out = new(string)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
	}
	if in.Filesystems != nil {
		in, out := &in.Filesystems, &out.Filesystems
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.HardRaidConfiguration != nil {
		in, out := &in.HardRaidConfiguration, &out.HardRaidConfiguration
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LastModification != nil {
		in, out := &in.LastModification, &out.LastModification
		*out = new(string)
		**out = **in
	}
	if in.LvmReady != nil {
		in, out := &in.LvmReady, &out.LvmReady
		*out = new(bool)
		**out = **in
	}
	if in.RemoveDefaultPartitionSchemes != nil {
		in, out := &in.RemoveDefaultPartitionSchemes, &out.RemoveDefaultPartitionSchemes
		*out = new(bool)
		**out = **in
	}
	if in.SupportsDistributionKernel != nil {
		in, out := &in.SupportsDistributionKernel, &out.SupportsDistributionKernel
		*out = new(bool)
		**out = **in
	}
	if in.SupportsRtm != nil {
		in, out := &in.SupportsRtm, &out.SupportsRtm
		*out = new(bool)
		**out = **in
	}
	if in.SupportsSQLServer != nil {
		in, out := &in.SupportsSQLServer, &out.SupportsSQLServer
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplateObservation.
func (in *InstallationTemplateObservation) DeepCopy() *InstallationTemplateObservation {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplateParameters) DeepCopyInto(out *InstallationTemplateParameters) {
	*out = *in
	if in.BaseTemplateName != nil {
		in, out := &in.BaseTemplateName, &out.BaseTemplateName
		*out = new(string)
		**out = **in
	}
	if in.Customization != nil {
		in, out := &in.Customization, &out.Customization
		*out = make([]CustomizationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultLanguage != nil {
		in, out := &in.DefaultLanguage, &out.DefaultLanguage
		*out = new(string)
		**out = **in
	}
	if in.RemoveDefaultPartitionSchemes != nil {
		in, out := &in.RemoveDefaultPartitionSchemes, &out.RemoveDefaultPartitionSchemes
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplateParameters.
func (in *InstallationTemplateParameters) DeepCopy() *InstallationTemplateParameters {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplateSpec) DeepCopyInto(out *InstallationTemplateSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplateSpec.
func (in *InstallationTemplateSpec) DeepCopy() *InstallationTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplateStatus) DeepCopyInto(out *InstallationTemplateStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplateStatus.
func (in *InstallationTemplateStatus) DeepCopy() *InstallationTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplateStatus)
	in.DeepCopyInto(out)
	return out
}
