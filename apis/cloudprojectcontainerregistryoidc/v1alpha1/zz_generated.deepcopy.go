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
func (in *ProjectContainerregistryOidc) DeepCopyInto(out *ProjectContainerregistryOidc) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectContainerregistryOidc.
func (in *ProjectContainerregistryOidc) DeepCopy() *ProjectContainerregistryOidc {
	if in == nil {
		return nil
	}
	out := new(ProjectContainerregistryOidc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectContainerregistryOidc) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectContainerregistryOidcInitParameters) DeepCopyInto(out *ProjectContainerregistryOidcInitParameters) {
	*out = *in
	if in.DeleteUsers != nil {
		in, out := &in.DeleteUsers, &out.DeleteUsers
		*out = new(bool)
		**out = **in
	}
	if in.OidcAdminGroup != nil {
		in, out := &in.OidcAdminGroup, &out.OidcAdminGroup
		*out = new(string)
		**out = **in
	}
	if in.OidcAutoOnboard != nil {
		in, out := &in.OidcAutoOnboard, &out.OidcAutoOnboard
		*out = new(bool)
		**out = **in
	}
	if in.OidcClientID != nil {
		in, out := &in.OidcClientID, &out.OidcClientID
		*out = new(string)
		**out = **in
	}
	if in.OidcEndpoint != nil {
		in, out := &in.OidcEndpoint, &out.OidcEndpoint
		*out = new(string)
		**out = **in
	}
	if in.OidcGroupsClaim != nil {
		in, out := &in.OidcGroupsClaim, &out.OidcGroupsClaim
		*out = new(string)
		**out = **in
	}
	if in.OidcName != nil {
		in, out := &in.OidcName, &out.OidcName
		*out = new(string)
		**out = **in
	}
	if in.OidcScope != nil {
		in, out := &in.OidcScope, &out.OidcScope
		*out = new(string)
		**out = **in
	}
	if in.OidcUserClaim != nil {
		in, out := &in.OidcUserClaim, &out.OidcUserClaim
		*out = new(string)
		**out = **in
	}
	if in.OidcVerifyCert != nil {
		in, out := &in.OidcVerifyCert, &out.OidcVerifyCert
		*out = new(bool)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectContainerregistryOidcInitParameters.
func (in *ProjectContainerregistryOidcInitParameters) DeepCopy() *ProjectContainerregistryOidcInitParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectContainerregistryOidcInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectContainerregistryOidcList) DeepCopyInto(out *ProjectContainerregistryOidcList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProjectContainerregistryOidc, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectContainerregistryOidcList.
func (in *ProjectContainerregistryOidcList) DeepCopy() *ProjectContainerregistryOidcList {
	if in == nil {
		return nil
	}
	out := new(ProjectContainerregistryOidcList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectContainerregistryOidcList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectContainerregistryOidcObservation) DeepCopyInto(out *ProjectContainerregistryOidcObservation) {
	*out = *in
	if in.DeleteUsers != nil {
		in, out := &in.DeleteUsers, &out.DeleteUsers
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.OidcAdminGroup != nil {
		in, out := &in.OidcAdminGroup, &out.OidcAdminGroup
		*out = new(string)
		**out = **in
	}
	if in.OidcAutoOnboard != nil {
		in, out := &in.OidcAutoOnboard, &out.OidcAutoOnboard
		*out = new(bool)
		**out = **in
	}
	if in.OidcClientID != nil {
		in, out := &in.OidcClientID, &out.OidcClientID
		*out = new(string)
		**out = **in
	}
	if in.OidcEndpoint != nil {
		in, out := &in.OidcEndpoint, &out.OidcEndpoint
		*out = new(string)
		**out = **in
	}
	if in.OidcGroupsClaim != nil {
		in, out := &in.OidcGroupsClaim, &out.OidcGroupsClaim
		*out = new(string)
		**out = **in
	}
	if in.OidcName != nil {
		in, out := &in.OidcName, &out.OidcName
		*out = new(string)
		**out = **in
	}
	if in.OidcScope != nil {
		in, out := &in.OidcScope, &out.OidcScope
		*out = new(string)
		**out = **in
	}
	if in.OidcUserClaim != nil {
		in, out := &in.OidcUserClaim, &out.OidcUserClaim
		*out = new(string)
		**out = **in
	}
	if in.OidcVerifyCert != nil {
		in, out := &in.OidcVerifyCert, &out.OidcVerifyCert
		*out = new(bool)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectContainerregistryOidcObservation.
func (in *ProjectContainerregistryOidcObservation) DeepCopy() *ProjectContainerregistryOidcObservation {
	if in == nil {
		return nil
	}
	out := new(ProjectContainerregistryOidcObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectContainerregistryOidcParameters) DeepCopyInto(out *ProjectContainerregistryOidcParameters) {
	*out = *in
	if in.DeleteUsers != nil {
		in, out := &in.DeleteUsers, &out.DeleteUsers
		*out = new(bool)
		**out = **in
	}
	if in.OidcAdminGroup != nil {
		in, out := &in.OidcAdminGroup, &out.OidcAdminGroup
		*out = new(string)
		**out = **in
	}
	if in.OidcAutoOnboard != nil {
		in, out := &in.OidcAutoOnboard, &out.OidcAutoOnboard
		*out = new(bool)
		**out = **in
	}
	if in.OidcClientID != nil {
		in, out := &in.OidcClientID, &out.OidcClientID
		*out = new(string)
		**out = **in
	}
	out.OidcClientSecretSecretRef = in.OidcClientSecretSecretRef
	if in.OidcEndpoint != nil {
		in, out := &in.OidcEndpoint, &out.OidcEndpoint
		*out = new(string)
		**out = **in
	}
	if in.OidcGroupsClaim != nil {
		in, out := &in.OidcGroupsClaim, &out.OidcGroupsClaim
		*out = new(string)
		**out = **in
	}
	if in.OidcName != nil {
		in, out := &in.OidcName, &out.OidcName
		*out = new(string)
		**out = **in
	}
	if in.OidcScope != nil {
		in, out := &in.OidcScope, &out.OidcScope
		*out = new(string)
		**out = **in
	}
	if in.OidcUserClaim != nil {
		in, out := &in.OidcUserClaim, &out.OidcUserClaim
		*out = new(string)
		**out = **in
	}
	if in.OidcVerifyCert != nil {
		in, out := &in.OidcVerifyCert, &out.OidcVerifyCert
		*out = new(bool)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectContainerregistryOidcParameters.
func (in *ProjectContainerregistryOidcParameters) DeepCopy() *ProjectContainerregistryOidcParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectContainerregistryOidcParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectContainerregistryOidcSpec) DeepCopyInto(out *ProjectContainerregistryOidcSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectContainerregistryOidcSpec.
func (in *ProjectContainerregistryOidcSpec) DeepCopy() *ProjectContainerregistryOidcSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectContainerregistryOidcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectContainerregistryOidcStatus) DeepCopyInto(out *ProjectContainerregistryOidcStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectContainerregistryOidcStatus.
func (in *ProjectContainerregistryOidcStatus) DeepCopy() *ProjectContainerregistryOidcStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectContainerregistryOidcStatus)
	in.DeepCopyInto(out)
	return out
}
