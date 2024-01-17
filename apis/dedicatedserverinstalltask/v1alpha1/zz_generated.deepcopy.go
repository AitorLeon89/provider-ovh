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
func (in *DetailsInitParameters) DeepCopyInto(out *DetailsInitParameters) {
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
	if in.DiskGroupID != nil {
		in, out := &in.DiskGroupID, &out.DiskGroupID
		*out = new(float64)
		**out = **in
	}
	if in.InstallRtm != nil {
		in, out := &in.InstallRtm, &out.InstallRtm
		*out = new(bool)
		**out = **in
	}
	if in.InstallSQLServer != nil {
		in, out := &in.InstallSQLServer, &out.InstallSQLServer
		*out = new(bool)
		**out = **in
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(string)
		**out = **in
	}
	if in.NoRaid != nil {
		in, out := &in.NoRaid, &out.NoRaid
		*out = new(bool)
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
	if in.ResetHwRaid != nil {
		in, out := &in.ResetHwRaid, &out.ResetHwRaid
		*out = new(bool)
		**out = **in
	}
	if in.SSHKeyName != nil {
		in, out := &in.SSHKeyName, &out.SSHKeyName
		*out = new(string)
		**out = **in
	}
	if in.SoftRaidDevices != nil {
		in, out := &in.SoftRaidDevices, &out.SoftRaidDevices
		*out = new(float64)
		**out = **in
	}
	if in.UseDistribKernel != nil {
		in, out := &in.UseDistribKernel, &out.UseDistribKernel
		*out = new(bool)
		**out = **in
	}
	if in.UseSpla != nil {
		in, out := &in.UseSpla, &out.UseSpla
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetailsInitParameters.
func (in *DetailsInitParameters) DeepCopy() *DetailsInitParameters {
	if in == nil {
		return nil
	}
	out := new(DetailsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetailsObservation) DeepCopyInto(out *DetailsObservation) {
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
	if in.DiskGroupID != nil {
		in, out := &in.DiskGroupID, &out.DiskGroupID
		*out = new(float64)
		**out = **in
	}
	if in.InstallRtm != nil {
		in, out := &in.InstallRtm, &out.InstallRtm
		*out = new(bool)
		**out = **in
	}
	if in.InstallSQLServer != nil {
		in, out := &in.InstallSQLServer, &out.InstallSQLServer
		*out = new(bool)
		**out = **in
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(string)
		**out = **in
	}
	if in.NoRaid != nil {
		in, out := &in.NoRaid, &out.NoRaid
		*out = new(bool)
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
	if in.ResetHwRaid != nil {
		in, out := &in.ResetHwRaid, &out.ResetHwRaid
		*out = new(bool)
		**out = **in
	}
	if in.SSHKeyName != nil {
		in, out := &in.SSHKeyName, &out.SSHKeyName
		*out = new(string)
		**out = **in
	}
	if in.SoftRaidDevices != nil {
		in, out := &in.SoftRaidDevices, &out.SoftRaidDevices
		*out = new(float64)
		**out = **in
	}
	if in.UseDistribKernel != nil {
		in, out := &in.UseDistribKernel, &out.UseDistribKernel
		*out = new(bool)
		**out = **in
	}
	if in.UseSpla != nil {
		in, out := &in.UseSpla, &out.UseSpla
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetailsObservation.
func (in *DetailsObservation) DeepCopy() *DetailsObservation {
	if in == nil {
		return nil
	}
	out := new(DetailsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetailsParameters) DeepCopyInto(out *DetailsParameters) {
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
	if in.DiskGroupID != nil {
		in, out := &in.DiskGroupID, &out.DiskGroupID
		*out = new(float64)
		**out = **in
	}
	if in.InstallRtm != nil {
		in, out := &in.InstallRtm, &out.InstallRtm
		*out = new(bool)
		**out = **in
	}
	if in.InstallSQLServer != nil {
		in, out := &in.InstallSQLServer, &out.InstallSQLServer
		*out = new(bool)
		**out = **in
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(string)
		**out = **in
	}
	if in.NoRaid != nil {
		in, out := &in.NoRaid, &out.NoRaid
		*out = new(bool)
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
	if in.ResetHwRaid != nil {
		in, out := &in.ResetHwRaid, &out.ResetHwRaid
		*out = new(bool)
		**out = **in
	}
	if in.SSHKeyName != nil {
		in, out := &in.SSHKeyName, &out.SSHKeyName
		*out = new(string)
		**out = **in
	}
	if in.SoftRaidDevices != nil {
		in, out := &in.SoftRaidDevices, &out.SoftRaidDevices
		*out = new(float64)
		**out = **in
	}
	if in.UseDistribKernel != nil {
		in, out := &in.UseDistribKernel, &out.UseDistribKernel
		*out = new(bool)
		**out = **in
	}
	if in.UseSpla != nil {
		in, out := &in.UseSpla, &out.UseSpla
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetailsParameters.
func (in *DetailsParameters) DeepCopy() *DetailsParameters {
	if in == nil {
		return nil
	}
	out := new(DetailsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerInstallTask) DeepCopyInto(out *ServerInstallTask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerInstallTask.
func (in *ServerInstallTask) DeepCopy() *ServerInstallTask {
	if in == nil {
		return nil
	}
	out := new(ServerInstallTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerInstallTask) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerInstallTaskInitParameters) DeepCopyInto(out *ServerInstallTaskInitParameters) {
	*out = *in
	if in.BootidOnDestroy != nil {
		in, out := &in.BootidOnDestroy, &out.BootidOnDestroy
		*out = new(float64)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]DetailsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PartitionSchemeName != nil {
		in, out := &in.PartitionSchemeName, &out.PartitionSchemeName
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerInstallTaskInitParameters.
func (in *ServerInstallTaskInitParameters) DeepCopy() *ServerInstallTaskInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServerInstallTaskInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerInstallTaskList) DeepCopyInto(out *ServerInstallTaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServerInstallTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerInstallTaskList.
func (in *ServerInstallTaskList) DeepCopy() *ServerInstallTaskList {
	if in == nil {
		return nil
	}
	out := new(ServerInstallTaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerInstallTaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerInstallTaskObservation) DeepCopyInto(out *ServerInstallTaskObservation) {
	*out = *in
	if in.BootidOnDestroy != nil {
		in, out := &in.BootidOnDestroy, &out.BootidOnDestroy
		*out = new(float64)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]DetailsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DoneDate != nil {
		in, out := &in.DoneDate, &out.DoneDate
		*out = new(string)
		**out = **in
	}
	if in.Function != nil {
		in, out := &in.Function, &out.Function
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = new(string)
		**out = **in
	}
	if in.PartitionSchemeName != nil {
		in, out := &in.PartitionSchemeName, &out.PartitionSchemeName
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.StartDate != nil {
		in, out := &in.StartDate, &out.StartDate
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerInstallTaskObservation.
func (in *ServerInstallTaskObservation) DeepCopy() *ServerInstallTaskObservation {
	if in == nil {
		return nil
	}
	out := new(ServerInstallTaskObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerInstallTaskParameters) DeepCopyInto(out *ServerInstallTaskParameters) {
	*out = *in
	if in.BootidOnDestroy != nil {
		in, out := &in.BootidOnDestroy, &out.BootidOnDestroy
		*out = new(float64)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]DetailsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PartitionSchemeName != nil {
		in, out := &in.PartitionSchemeName, &out.PartitionSchemeName
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerInstallTaskParameters.
func (in *ServerInstallTaskParameters) DeepCopy() *ServerInstallTaskParameters {
	if in == nil {
		return nil
	}
	out := new(ServerInstallTaskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerInstallTaskSpec) DeepCopyInto(out *ServerInstallTaskSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerInstallTaskSpec.
func (in *ServerInstallTaskSpec) DeepCopy() *ServerInstallTaskSpec {
	if in == nil {
		return nil
	}
	out := new(ServerInstallTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerInstallTaskStatus) DeepCopyInto(out *ServerInstallTaskStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerInstallTaskStatus.
func (in *ServerInstallTaskStatus) DeepCopy() *ServerInstallTaskStatus {
	if in == nil {
		return nil
	}
	out := new(ServerInstallTaskStatus)
	in.DeepCopyInto(out)
	return out
}