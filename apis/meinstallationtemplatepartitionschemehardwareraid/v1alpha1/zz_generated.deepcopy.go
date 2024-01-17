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
func (in *InstallationTemplatePartitionSchemeHardwareRaid) DeepCopyInto(out *InstallationTemplatePartitionSchemeHardwareRaid) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemeHardwareRaid.
func (in *InstallationTemplatePartitionSchemeHardwareRaid) DeepCopy() *InstallationTemplatePartitionSchemeHardwareRaid {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemeHardwareRaid)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallationTemplatePartitionSchemeHardwareRaid) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplatePartitionSchemeHardwareRaidInitParameters) DeepCopyInto(out *InstallationTemplatePartitionSchemeHardwareRaidInitParameters) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SchemeName != nil {
		in, out := &in.SchemeName, &out.SchemeName
		*out = new(string)
		**out = **in
	}
	if in.Step != nil {
		in, out := &in.Step, &out.Step
		*out = new(float64)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemeHardwareRaidInitParameters.
func (in *InstallationTemplatePartitionSchemeHardwareRaidInitParameters) DeepCopy() *InstallationTemplatePartitionSchemeHardwareRaidInitParameters {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemeHardwareRaidInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplatePartitionSchemeHardwareRaidList) DeepCopyInto(out *InstallationTemplatePartitionSchemeHardwareRaidList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstallationTemplatePartitionSchemeHardwareRaid, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemeHardwareRaidList.
func (in *InstallationTemplatePartitionSchemeHardwareRaidList) DeepCopy() *InstallationTemplatePartitionSchemeHardwareRaidList {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemeHardwareRaidList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallationTemplatePartitionSchemeHardwareRaidList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplatePartitionSchemeHardwareRaidObservation) DeepCopyInto(out *InstallationTemplatePartitionSchemeHardwareRaidObservation) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SchemeName != nil {
		in, out := &in.SchemeName, &out.SchemeName
		*out = new(string)
		**out = **in
	}
	if in.Step != nil {
		in, out := &in.Step, &out.Step
		*out = new(float64)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemeHardwareRaidObservation.
func (in *InstallationTemplatePartitionSchemeHardwareRaidObservation) DeepCopy() *InstallationTemplatePartitionSchemeHardwareRaidObservation {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemeHardwareRaidObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplatePartitionSchemeHardwareRaidParameters) DeepCopyInto(out *InstallationTemplatePartitionSchemeHardwareRaidParameters) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SchemeName != nil {
		in, out := &in.SchemeName, &out.SchemeName
		*out = new(string)
		**out = **in
	}
	if in.Step != nil {
		in, out := &in.Step, &out.Step
		*out = new(float64)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemeHardwareRaidParameters.
func (in *InstallationTemplatePartitionSchemeHardwareRaidParameters) DeepCopy() *InstallationTemplatePartitionSchemeHardwareRaidParameters {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemeHardwareRaidParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplatePartitionSchemeHardwareRaidSpec) DeepCopyInto(out *InstallationTemplatePartitionSchemeHardwareRaidSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemeHardwareRaidSpec.
func (in *InstallationTemplatePartitionSchemeHardwareRaidSpec) DeepCopy() *InstallationTemplatePartitionSchemeHardwareRaidSpec {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemeHardwareRaidSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplatePartitionSchemeHardwareRaidStatus) DeepCopyInto(out *InstallationTemplatePartitionSchemeHardwareRaidStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemeHardwareRaidStatus.
func (in *InstallationTemplatePartitionSchemeHardwareRaidStatus) DeepCopy() *InstallationTemplatePartitionSchemeHardwareRaidStatus {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemeHardwareRaidStatus)
	in.DeepCopyInto(out)
	return out
}
