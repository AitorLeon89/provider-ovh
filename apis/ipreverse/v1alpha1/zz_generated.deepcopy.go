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
func (in *Reverse) DeepCopyInto(out *Reverse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reverse.
func (in *Reverse) DeepCopy() *Reverse {
	if in == nil {
		return nil
	}
	out := new(Reverse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Reverse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReverseInitParameters) DeepCopyInto(out *ReverseInitParameters) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.IPReverse != nil {
		in, out := &in.IPReverse, &out.IPReverse
		*out = new(string)
		**out = **in
	}
	if in.Reverse != nil {
		in, out := &in.Reverse, &out.Reverse
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReverseInitParameters.
func (in *ReverseInitParameters) DeepCopy() *ReverseInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReverseInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReverseList) DeepCopyInto(out *ReverseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Reverse, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReverseList.
func (in *ReverseList) DeepCopy() *ReverseList {
	if in == nil {
		return nil
	}
	out := new(ReverseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReverseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReverseObservation) DeepCopyInto(out *ReverseObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.IPReverse != nil {
		in, out := &in.IPReverse, &out.IPReverse
		*out = new(string)
		**out = **in
	}
	if in.Reverse != nil {
		in, out := &in.Reverse, &out.Reverse
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReverseObservation.
func (in *ReverseObservation) DeepCopy() *ReverseObservation {
	if in == nil {
		return nil
	}
	out := new(ReverseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReverseParameters) DeepCopyInto(out *ReverseParameters) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.IPReverse != nil {
		in, out := &in.IPReverse, &out.IPReverse
		*out = new(string)
		**out = **in
	}
	if in.Reverse != nil {
		in, out := &in.Reverse, &out.Reverse
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReverseParameters.
func (in *ReverseParameters) DeepCopy() *ReverseParameters {
	if in == nil {
		return nil
	}
	out := new(ReverseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReverseSpec) DeepCopyInto(out *ReverseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReverseSpec.
func (in *ReverseSpec) DeepCopy() *ReverseSpec {
	if in == nil {
		return nil
	}
	out := new(ReverseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReverseStatus) DeepCopyInto(out *ReverseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReverseStatus.
func (in *ReverseStatus) DeepCopy() *ReverseStatus {
	if in == nil {
		return nil
	}
	out := new(ReverseStatus)
	in.DeepCopyInto(out)
	return out
}
