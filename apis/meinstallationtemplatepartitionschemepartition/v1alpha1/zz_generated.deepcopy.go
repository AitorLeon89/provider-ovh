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
func (in *InstallationTemplatePartitionSchemePartition) DeepCopyInto(out *InstallationTemplatePartitionSchemePartition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemePartition.
func (in *InstallationTemplatePartitionSchemePartition) DeepCopy() *InstallationTemplatePartitionSchemePartition {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemePartition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallationTemplatePartitionSchemePartition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplatePartitionSchemePartitionInitParameters) DeepCopyInto(out *InstallationTemplatePartitionSchemePartitionInitParameters) {
	*out = *in
	if in.Filesystem != nil {
		in, out := &in.Filesystem, &out.Filesystem
		*out = new(string)
		**out = **in
	}
	if in.Mountpoint != nil {
		in, out := &in.Mountpoint, &out.Mountpoint
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.Raid != nil {
		in, out := &in.Raid, &out.Raid
		*out = new(string)
		**out = **in
	}
	if in.SchemeName != nil {
		in, out := &in.SchemeName, &out.SchemeName
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.VolumeName != nil {
		in, out := &in.VolumeName, &out.VolumeName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemePartitionInitParameters.
func (in *InstallationTemplatePartitionSchemePartitionInitParameters) DeepCopy() *InstallationTemplatePartitionSchemePartitionInitParameters {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemePartitionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplatePartitionSchemePartitionList) DeepCopyInto(out *InstallationTemplatePartitionSchemePartitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstallationTemplatePartitionSchemePartition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemePartitionList.
func (in *InstallationTemplatePartitionSchemePartitionList) DeepCopy() *InstallationTemplatePartitionSchemePartitionList {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemePartitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallationTemplatePartitionSchemePartitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplatePartitionSchemePartitionObservation) DeepCopyInto(out *InstallationTemplatePartitionSchemePartitionObservation) {
	*out = *in
	if in.Filesystem != nil {
		in, out := &in.Filesystem, &out.Filesystem
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Mountpoint != nil {
		in, out := &in.Mountpoint, &out.Mountpoint
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.Raid != nil {
		in, out := &in.Raid, &out.Raid
		*out = new(string)
		**out = **in
	}
	if in.SchemeName != nil {
		in, out := &in.SchemeName, &out.SchemeName
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.VolumeName != nil {
		in, out := &in.VolumeName, &out.VolumeName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemePartitionObservation.
func (in *InstallationTemplatePartitionSchemePartitionObservation) DeepCopy() *InstallationTemplatePartitionSchemePartitionObservation {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemePartitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplatePartitionSchemePartitionParameters) DeepCopyInto(out *InstallationTemplatePartitionSchemePartitionParameters) {
	*out = *in
	if in.Filesystem != nil {
		in, out := &in.Filesystem, &out.Filesystem
		*out = new(string)
		**out = **in
	}
	if in.Mountpoint != nil {
		in, out := &in.Mountpoint, &out.Mountpoint
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.Raid != nil {
		in, out := &in.Raid, &out.Raid
		*out = new(string)
		**out = **in
	}
	if in.SchemeName != nil {
		in, out := &in.SchemeName, &out.SchemeName
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.VolumeName != nil {
		in, out := &in.VolumeName, &out.VolumeName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemePartitionParameters.
func (in *InstallationTemplatePartitionSchemePartitionParameters) DeepCopy() *InstallationTemplatePartitionSchemePartitionParameters {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemePartitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplatePartitionSchemePartitionSpec) DeepCopyInto(out *InstallationTemplatePartitionSchemePartitionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemePartitionSpec.
func (in *InstallationTemplatePartitionSchemePartitionSpec) DeepCopy() *InstallationTemplatePartitionSchemePartitionSpec {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemePartitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationTemplatePartitionSchemePartitionStatus) DeepCopyInto(out *InstallationTemplatePartitionSchemePartitionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationTemplatePartitionSchemePartitionStatus.
func (in *InstallationTemplatePartitionSchemePartitionStatus) DeepCopy() *InstallationTemplatePartitionSchemePartitionStatus {
	if in == nil {
		return nil
	}
	out := new(InstallationTemplatePartitionSchemePartitionStatus)
	in.DeepCopyInto(out)
	return out
}