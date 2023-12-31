//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoinWorkflow) DeepCopyInto(out *CoinWorkflow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoinWorkflow.
func (in *CoinWorkflow) DeepCopy() *CoinWorkflow {
	if in == nil {
		return nil
	}
	out := new(CoinWorkflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoinWorkflow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoinWorkflowList) DeepCopyInto(out *CoinWorkflowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CoinWorkflow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoinWorkflowList.
func (in *CoinWorkflowList) DeepCopy() *CoinWorkflowList {
	if in == nil {
		return nil
	}
	out := new(CoinWorkflowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoinWorkflowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoinWorkflowSpec) DeepCopyInto(out *CoinWorkflowSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoinWorkflowSpec.
func (in *CoinWorkflowSpec) DeepCopy() *CoinWorkflowSpec {
	if in == nil {
		return nil
	}
	out := new(CoinWorkflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoinWorkflowStatus) DeepCopyInto(out *CoinWorkflowStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoinWorkflowStatus.
func (in *CoinWorkflowStatus) DeepCopy() *CoinWorkflowStatus {
	if in == nil {
		return nil
	}
	out := new(CoinWorkflowStatus)
	in.DeepCopyInto(out)
	return out
}
