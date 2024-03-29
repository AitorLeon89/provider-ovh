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
func (in *Cloudproject) DeepCopyInto(out *Cloudproject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cloudproject.
func (in *Cloudproject) DeepCopy() *Cloudproject {
	if in == nil {
		return nil
	}
	out := new(Cloudproject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cloudproject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudprojectInitParameters) DeepCopyInto(out *CloudprojectInitParameters) {
	*out = *in
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudprojectInitParameters.
func (in *CloudprojectInitParameters) DeepCopy() *CloudprojectInitParameters {
	if in == nil {
		return nil
	}
	out := new(CloudprojectInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudprojectList) DeepCopyInto(out *CloudprojectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cloudproject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudprojectList.
func (in *CloudprojectList) DeepCopy() *CloudprojectList {
	if in == nil {
		return nil
	}
	out := new(CloudprojectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudprojectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudprojectObservation) DeepCopyInto(out *CloudprojectObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudprojectObservation.
func (in *CloudprojectObservation) DeepCopy() *CloudprojectObservation {
	if in == nil {
		return nil
	}
	out := new(CloudprojectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudprojectParameters) DeepCopyInto(out *CloudprojectParameters) {
	*out = *in
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudprojectParameters.
func (in *CloudprojectParameters) DeepCopy() *CloudprojectParameters {
	if in == nil {
		return nil
	}
	out := new(CloudprojectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudprojectSpec) DeepCopyInto(out *CloudprojectSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudprojectSpec.
func (in *CloudprojectSpec) DeepCopy() *CloudprojectSpec {
	if in == nil {
		return nil
	}
	out := new(CloudprojectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudprojectStatus) DeepCopyInto(out *CloudprojectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudprojectStatus.
func (in *CloudprojectStatus) DeepCopy() *CloudprojectStatus {
	if in == nil {
		return nil
	}
	out := new(CloudprojectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationInitParameters) DeepCopyInto(out *ConfigurationInitParameters) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationInitParameters.
func (in *ConfigurationInitParameters) DeepCopy() *ConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationObservation) DeepCopyInto(out *ConfigurationObservation) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationObservation.
func (in *ConfigurationObservation) DeepCopy() *ConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationParameters) DeepCopyInto(out *ConfigurationParameters) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationParameters.
func (in *ConfigurationParameters) DeepCopy() *ConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetailsInitParameters) DeepCopyInto(out *DetailsInitParameters) {
	*out = *in
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
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.OrderDetailID != nil {
		in, out := &in.OrderDetailID, &out.OrderDetailID
		*out = new(float64)
		**out = **in
	}
	if in.Quantity != nil {
		in, out := &in.Quantity, &out.Quantity
		*out = new(string)
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
func (in *OrderInitParameters) DeepCopyInto(out *OrderInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderInitParameters.
func (in *OrderInitParameters) DeepCopy() *OrderInitParameters {
	if in == nil {
		return nil
	}
	out := new(OrderInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderObservation) DeepCopyInto(out *OrderObservation) {
	*out = *in
	if in.Date != nil {
		in, out := &in.Date, &out.Date
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
	if in.ExpirationDate != nil {
		in, out := &in.ExpirationDate, &out.ExpirationDate
		*out = new(string)
		**out = **in
	}
	if in.OrderID != nil {
		in, out := &in.OrderID, &out.OrderID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderObservation.
func (in *OrderObservation) DeepCopy() *OrderObservation {
	if in == nil {
		return nil
	}
	out := new(OrderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderParameters) DeepCopyInto(out *OrderParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderParameters.
func (in *OrderParameters) DeepCopy() *OrderParameters {
	if in == nil {
		return nil
	}
	out := new(OrderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanInitParameters) DeepCopyInto(out *PlanInitParameters) {
	*out = *in
	if in.CatalogName != nil {
		in, out := &in.CatalogName, &out.CatalogName
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]ConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.PlanCode != nil {
		in, out := &in.PlanCode, &out.PlanCode
		*out = new(string)
		**out = **in
	}
	if in.PricingMode != nil {
		in, out := &in.PricingMode, &out.PricingMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanInitParameters.
func (in *PlanInitParameters) DeepCopy() *PlanInitParameters {
	if in == nil {
		return nil
	}
	out := new(PlanInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanObservation) DeepCopyInto(out *PlanObservation) {
	*out = *in
	if in.CatalogName != nil {
		in, out := &in.CatalogName, &out.CatalogName
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]ConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.PlanCode != nil {
		in, out := &in.PlanCode, &out.PlanCode
		*out = new(string)
		**out = **in
	}
	if in.PricingMode != nil {
		in, out := &in.PricingMode, &out.PricingMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanObservation.
func (in *PlanObservation) DeepCopy() *PlanObservation {
	if in == nil {
		return nil
	}
	out := new(PlanObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanOptionConfigurationInitParameters) DeepCopyInto(out *PlanOptionConfigurationInitParameters) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanOptionConfigurationInitParameters.
func (in *PlanOptionConfigurationInitParameters) DeepCopy() *PlanOptionConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(PlanOptionConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanOptionConfigurationObservation) DeepCopyInto(out *PlanOptionConfigurationObservation) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanOptionConfigurationObservation.
func (in *PlanOptionConfigurationObservation) DeepCopy() *PlanOptionConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(PlanOptionConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanOptionConfigurationParameters) DeepCopyInto(out *PlanOptionConfigurationParameters) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanOptionConfigurationParameters.
func (in *PlanOptionConfigurationParameters) DeepCopy() *PlanOptionConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(PlanOptionConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanOptionInitParameters) DeepCopyInto(out *PlanOptionInitParameters) {
	*out = *in
	if in.CatalogName != nil {
		in, out := &in.CatalogName, &out.CatalogName
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]PlanOptionConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.PlanCode != nil {
		in, out := &in.PlanCode, &out.PlanCode
		*out = new(string)
		**out = **in
	}
	if in.PricingMode != nil {
		in, out := &in.PricingMode, &out.PricingMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanOptionInitParameters.
func (in *PlanOptionInitParameters) DeepCopy() *PlanOptionInitParameters {
	if in == nil {
		return nil
	}
	out := new(PlanOptionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanOptionObservation) DeepCopyInto(out *PlanOptionObservation) {
	*out = *in
	if in.CatalogName != nil {
		in, out := &in.CatalogName, &out.CatalogName
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]PlanOptionConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.PlanCode != nil {
		in, out := &in.PlanCode, &out.PlanCode
		*out = new(string)
		**out = **in
	}
	if in.PricingMode != nil {
		in, out := &in.PricingMode, &out.PricingMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanOptionObservation.
func (in *PlanOptionObservation) DeepCopy() *PlanOptionObservation {
	if in == nil {
		return nil
	}
	out := new(PlanOptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanOptionParameters) DeepCopyInto(out *PlanOptionParameters) {
	*out = *in
	if in.CatalogName != nil {
		in, out := &in.CatalogName, &out.CatalogName
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]PlanOptionConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.PlanCode != nil {
		in, out := &in.PlanCode, &out.PlanCode
		*out = new(string)
		**out = **in
	}
	if in.PricingMode != nil {
		in, out := &in.PricingMode, &out.PricingMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanOptionParameters.
func (in *PlanOptionParameters) DeepCopy() *PlanOptionParameters {
	if in == nil {
		return nil
	}
	out := new(PlanOptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanParameters) DeepCopyInto(out *PlanParameters) {
	*out = *in
	if in.CatalogName != nil {
		in, out := &in.CatalogName, &out.CatalogName
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]ConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.PlanCode != nil {
		in, out := &in.PlanCode, &out.PlanCode
		*out = new(string)
		**out = **in
	}
	if in.PricingMode != nil {
		in, out := &in.PricingMode, &out.PricingMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanParameters.
func (in *PlanParameters) DeepCopy() *PlanParameters {
	if in == nil {
		return nil
	}
	out := new(PlanParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vrack) DeepCopyInto(out *Vrack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vrack.
func (in *Vrack) DeepCopy() *Vrack {
	if in == nil {
		return nil
	}
	out := new(Vrack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vrack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VrackInitParameters) DeepCopyInto(out *VrackInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OvhSubsidiary != nil {
		in, out := &in.OvhSubsidiary, &out.OvhSubsidiary
		*out = new(string)
		**out = **in
	}
	if in.PaymentMean != nil {
		in, out := &in.PaymentMean, &out.PaymentMean
		*out = new(string)
		**out = **in
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = make([]PlanInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlanOption != nil {
		in, out := &in.PlanOption, &out.PlanOption
		*out = make([]PlanOptionInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VrackInitParameters.
func (in *VrackInitParameters) DeepCopy() *VrackInitParameters {
	if in == nil {
		return nil
	}
	out := new(VrackInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VrackList) DeepCopyInto(out *VrackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vrack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VrackList.
func (in *VrackList) DeepCopy() *VrackList {
	if in == nil {
		return nil
	}
	out := new(VrackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VrackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VrackObservation) DeepCopyInto(out *VrackObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = make([]OrderObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OvhSubsidiary != nil {
		in, out := &in.OvhSubsidiary, &out.OvhSubsidiary
		*out = new(string)
		**out = **in
	}
	if in.PaymentMean != nil {
		in, out := &in.PaymentMean, &out.PaymentMean
		*out = new(string)
		**out = **in
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = make([]PlanObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlanOption != nil {
		in, out := &in.PlanOption, &out.PlanOption
		*out = make([]PlanOptionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.Urn != nil {
		in, out := &in.Urn, &out.Urn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VrackObservation.
func (in *VrackObservation) DeepCopy() *VrackObservation {
	if in == nil {
		return nil
	}
	out := new(VrackObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VrackParameters) DeepCopyInto(out *VrackParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OvhSubsidiary != nil {
		in, out := &in.OvhSubsidiary, &out.OvhSubsidiary
		*out = new(string)
		**out = **in
	}
	if in.PaymentMean != nil {
		in, out := &in.PaymentMean, &out.PaymentMean
		*out = new(string)
		**out = **in
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = make([]PlanParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlanOption != nil {
		in, out := &in.PlanOption, &out.PlanOption
		*out = make([]PlanOptionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VrackParameters.
func (in *VrackParameters) DeepCopy() *VrackParameters {
	if in == nil {
		return nil
	}
	out := new(VrackParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VrackSpec) DeepCopyInto(out *VrackSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VrackSpec.
func (in *VrackSpec) DeepCopy() *VrackSpec {
	if in == nil {
		return nil
	}
	out := new(VrackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VrackStatus) DeepCopyInto(out *VrackStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VrackStatus.
func (in *VrackStatus) DeepCopy() *VrackStatus {
	if in == nil {
		return nil
	}
	out := new(VrackStatus)
	in.DeepCopyInto(out)
	return out
}
