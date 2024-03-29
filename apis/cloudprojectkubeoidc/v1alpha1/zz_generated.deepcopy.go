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
func (in *ProjectKubeOidc) DeepCopyInto(out *ProjectKubeOidc) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectKubeOidc.
func (in *ProjectKubeOidc) DeepCopy() *ProjectKubeOidc {
	if in == nil {
		return nil
	}
	out := new(ProjectKubeOidc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectKubeOidc) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectKubeOidcInitParameters) DeepCopyInto(out *ProjectKubeOidcInitParameters) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.IssuerURL != nil {
		in, out := &in.IssuerURL, &out.IssuerURL
		*out = new(string)
		**out = **in
	}
	if in.KubeID != nil {
		in, out := &in.KubeID, &out.KubeID
		*out = new(string)
		**out = **in
	}
	if in.OidcCAContent != nil {
		in, out := &in.OidcCAContent, &out.OidcCAContent
		*out = new(string)
		**out = **in
	}
	if in.OidcGroupsClaim != nil {
		in, out := &in.OidcGroupsClaim, &out.OidcGroupsClaim
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OidcGroupsPrefix != nil {
		in, out := &in.OidcGroupsPrefix, &out.OidcGroupsPrefix
		*out = new(string)
		**out = **in
	}
	if in.OidcRequiredClaim != nil {
		in, out := &in.OidcRequiredClaim, &out.OidcRequiredClaim
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OidcSigningAlgs != nil {
		in, out := &in.OidcSigningAlgs, &out.OidcSigningAlgs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OidcUsernameClaim != nil {
		in, out := &in.OidcUsernameClaim, &out.OidcUsernameClaim
		*out = new(string)
		**out = **in
	}
	if in.OidcUsernamePrefix != nil {
		in, out := &in.OidcUsernamePrefix, &out.OidcUsernamePrefix
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectKubeOidcInitParameters.
func (in *ProjectKubeOidcInitParameters) DeepCopy() *ProjectKubeOidcInitParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectKubeOidcInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectKubeOidcList) DeepCopyInto(out *ProjectKubeOidcList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProjectKubeOidc, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectKubeOidcList.
func (in *ProjectKubeOidcList) DeepCopy() *ProjectKubeOidcList {
	if in == nil {
		return nil
	}
	out := new(ProjectKubeOidcList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectKubeOidcList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectKubeOidcObservation) DeepCopyInto(out *ProjectKubeOidcObservation) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IssuerURL != nil {
		in, out := &in.IssuerURL, &out.IssuerURL
		*out = new(string)
		**out = **in
	}
	if in.KubeID != nil {
		in, out := &in.KubeID, &out.KubeID
		*out = new(string)
		**out = **in
	}
	if in.OidcCAContent != nil {
		in, out := &in.OidcCAContent, &out.OidcCAContent
		*out = new(string)
		**out = **in
	}
	if in.OidcGroupsClaim != nil {
		in, out := &in.OidcGroupsClaim, &out.OidcGroupsClaim
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OidcGroupsPrefix != nil {
		in, out := &in.OidcGroupsPrefix, &out.OidcGroupsPrefix
		*out = new(string)
		**out = **in
	}
	if in.OidcRequiredClaim != nil {
		in, out := &in.OidcRequiredClaim, &out.OidcRequiredClaim
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OidcSigningAlgs != nil {
		in, out := &in.OidcSigningAlgs, &out.OidcSigningAlgs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OidcUsernameClaim != nil {
		in, out := &in.OidcUsernameClaim, &out.OidcUsernameClaim
		*out = new(string)
		**out = **in
	}
	if in.OidcUsernamePrefix != nil {
		in, out := &in.OidcUsernamePrefix, &out.OidcUsernamePrefix
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectKubeOidcObservation.
func (in *ProjectKubeOidcObservation) DeepCopy() *ProjectKubeOidcObservation {
	if in == nil {
		return nil
	}
	out := new(ProjectKubeOidcObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectKubeOidcParameters) DeepCopyInto(out *ProjectKubeOidcParameters) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.IssuerURL != nil {
		in, out := &in.IssuerURL, &out.IssuerURL
		*out = new(string)
		**out = **in
	}
	if in.KubeID != nil {
		in, out := &in.KubeID, &out.KubeID
		*out = new(string)
		**out = **in
	}
	if in.OidcCAContent != nil {
		in, out := &in.OidcCAContent, &out.OidcCAContent
		*out = new(string)
		**out = **in
	}
	if in.OidcGroupsClaim != nil {
		in, out := &in.OidcGroupsClaim, &out.OidcGroupsClaim
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OidcGroupsPrefix != nil {
		in, out := &in.OidcGroupsPrefix, &out.OidcGroupsPrefix
		*out = new(string)
		**out = **in
	}
	if in.OidcRequiredClaim != nil {
		in, out := &in.OidcRequiredClaim, &out.OidcRequiredClaim
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OidcSigningAlgs != nil {
		in, out := &in.OidcSigningAlgs, &out.OidcSigningAlgs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OidcUsernameClaim != nil {
		in, out := &in.OidcUsernameClaim, &out.OidcUsernameClaim
		*out = new(string)
		**out = **in
	}
	if in.OidcUsernamePrefix != nil {
		in, out := &in.OidcUsernamePrefix, &out.OidcUsernamePrefix
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectKubeOidcParameters.
func (in *ProjectKubeOidcParameters) DeepCopy() *ProjectKubeOidcParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectKubeOidcParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectKubeOidcSpec) DeepCopyInto(out *ProjectKubeOidcSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectKubeOidcSpec.
func (in *ProjectKubeOidcSpec) DeepCopy() *ProjectKubeOidcSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectKubeOidcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectKubeOidcStatus) DeepCopyInto(out *ProjectKubeOidcStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectKubeOidcStatus.
func (in *ProjectKubeOidcStatus) DeepCopy() *ProjectKubeOidcStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectKubeOidcStatus)
	in.DeepCopyInto(out)
	return out
}
