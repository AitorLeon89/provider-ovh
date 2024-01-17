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
func (in *HTTPFarmServer) DeepCopyInto(out *HTTPFarmServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFarmServer.
func (in *HTTPFarmServer) DeepCopy() *HTTPFarmServer {
	if in == nil {
		return nil
	}
	out := new(HTTPFarmServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPFarmServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPFarmServerInitParameters) DeepCopyInto(out *HTTPFarmServerInitParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(bool)
		**out = **in
	}
	if in.Chain != nil {
		in, out := &in.Chain, &out.Chain
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FarmID != nil {
		in, out := &in.FarmID, &out.FarmID
		*out = new(float64)
		**out = **in
	}
	if in.OnMarkedDown != nil {
		in, out := &in.OnMarkedDown, &out.OnMarkedDown
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.Probe != nil {
		in, out := &in.Probe, &out.Probe
		*out = new(bool)
		**out = **in
	}
	if in.ProxyProtocolVersion != nil {
		in, out := &in.ProxyProtocolVersion, &out.ProxyProtocolVersion
		*out = new(string)
		**out = **in
	}
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(bool)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFarmServerInitParameters.
func (in *HTTPFarmServerInitParameters) DeepCopy() *HTTPFarmServerInitParameters {
	if in == nil {
		return nil
	}
	out := new(HTTPFarmServerInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPFarmServerList) DeepCopyInto(out *HTTPFarmServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HTTPFarmServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFarmServerList.
func (in *HTTPFarmServerList) DeepCopy() *HTTPFarmServerList {
	if in == nil {
		return nil
	}
	out := new(HTTPFarmServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPFarmServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPFarmServerObservation) DeepCopyInto(out *HTTPFarmServerObservation) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(bool)
		**out = **in
	}
	if in.Chain != nil {
		in, out := &in.Chain, &out.Chain
		*out = new(string)
		**out = **in
	}
	if in.Cookie != nil {
		in, out := &in.Cookie, &out.Cookie
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FarmID != nil {
		in, out := &in.FarmID, &out.FarmID
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.OnMarkedDown != nil {
		in, out := &in.OnMarkedDown, &out.OnMarkedDown
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.Probe != nil {
		in, out := &in.Probe, &out.Probe
		*out = new(bool)
		**out = **in
	}
	if in.ProxyProtocolVersion != nil {
		in, out := &in.ProxyProtocolVersion, &out.ProxyProtocolVersion
		*out = new(string)
		**out = **in
	}
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(bool)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFarmServerObservation.
func (in *HTTPFarmServerObservation) DeepCopy() *HTTPFarmServerObservation {
	if in == nil {
		return nil
	}
	out := new(HTTPFarmServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPFarmServerParameters) DeepCopyInto(out *HTTPFarmServerParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(bool)
		**out = **in
	}
	if in.Chain != nil {
		in, out := &in.Chain, &out.Chain
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FarmID != nil {
		in, out := &in.FarmID, &out.FarmID
		*out = new(float64)
		**out = **in
	}
	if in.OnMarkedDown != nil {
		in, out := &in.OnMarkedDown, &out.OnMarkedDown
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.Probe != nil {
		in, out := &in.Probe, &out.Probe
		*out = new(bool)
		**out = **in
	}
	if in.ProxyProtocolVersion != nil {
		in, out := &in.ProxyProtocolVersion, &out.ProxyProtocolVersion
		*out = new(string)
		**out = **in
	}
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(bool)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFarmServerParameters.
func (in *HTTPFarmServerParameters) DeepCopy() *HTTPFarmServerParameters {
	if in == nil {
		return nil
	}
	out := new(HTTPFarmServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPFarmServerSpec) DeepCopyInto(out *HTTPFarmServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFarmServerSpec.
func (in *HTTPFarmServerSpec) DeepCopy() *HTTPFarmServerSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPFarmServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPFarmServerStatus) DeepCopyInto(out *HTTPFarmServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFarmServerStatus.
func (in *HTTPFarmServerStatus) DeepCopy() *HTTPFarmServerStatus {
	if in == nil {
		return nil
	}
	out := new(HTTPFarmServerStatus)
	in.DeepCopyInto(out)
	return out
}
