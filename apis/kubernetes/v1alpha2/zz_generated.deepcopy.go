//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionDetail) DeepCopyInto(out *ConnectionDetail) {
	*out = *in
	out.ObjectReference = in.ObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionDetail.
func (in *ConnectionDetail) DeepCopy() *ConnectionDetail {
	if in == nil {
		return nil
	}
	out := new(ConnectionDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependsOn) DeepCopyInto(out *DependsOn) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependsOn.
func (in *DependsOn) DeepCopy() *DependsOn {
	if in == nil {
		return nil
	}
	out := new(DependsOn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Object) DeepCopyInto(out *Object) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Object.
func (in *Object) DeepCopy() *Object {
	if in == nil {
		return nil
	}
	out := new(Object)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Object) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectList) DeepCopyInto(out *ObjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Object, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectList.
func (in *ObjectList) DeepCopy() *ObjectList {
	if in == nil {
		return nil
	}
	out := new(ObjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectObservation) DeepCopyInto(out *ObjectObservation) {
	*out = *in
	in.Manifest.DeepCopyInto(&out.Manifest)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectObservation.
func (in *ObjectObservation) DeepCopy() *ObjectObservation {
	if in == nil {
		return nil
	}
	out := new(ObjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectParameters) DeepCopyInto(out *ObjectParameters) {
	*out = *in
	in.Manifest.DeepCopyInto(&out.Manifest)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectParameters.
func (in *ObjectParameters) DeepCopy() *ObjectParameters {
	if in == nil {
		return nil
	}
	out := new(ObjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSpec) DeepCopyInto(out *ObjectSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.ConnectionDetails != nil {
		in, out := &in.ConnectionDetails, &out.ConnectionDetails
		*out = make([]ConnectionDetail, len(*in))
		copy(*out, *in)
	}
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	if in.References != nil {
		in, out := &in.References, &out.References
		*out = make([]Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Readiness = in.Readiness
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSpec.
func (in *ObjectSpec) DeepCopy() *ObjectSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStatus) DeepCopyInto(out *ObjectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStatus.
func (in *ObjectStatus) DeepCopy() *ObjectStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchesFrom) DeepCopyInto(out *PatchesFrom) {
	*out = *in
	out.DependsOn = in.DependsOn
	if in.FieldPath != nil {
		in, out := &in.FieldPath, &out.FieldPath
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchesFrom.
func (in *PatchesFrom) DeepCopy() *PatchesFrom {
	if in == nil {
		return nil
	}
	out := new(PatchesFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Readiness) DeepCopyInto(out *Readiness) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Readiness.
func (in *Readiness) DeepCopy() *Readiness {
	if in == nil {
		return nil
	}
	out := new(Readiness)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reference) DeepCopyInto(out *Reference) {
	*out = *in
	if in.DependsOn != nil {
		in, out := &in.DependsOn, &out.DependsOn
		*out = new(DependsOn)
		**out = **in
	}
	if in.PatchesFrom != nil {
		in, out := &in.PatchesFrom, &out.PatchesFrom
		*out = new(PatchesFrom)
		(*in).DeepCopyInto(*out)
	}
	if in.ToFieldPath != nil {
		in, out := &in.ToFieldPath, &out.ToFieldPath
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reference.
func (in *Reference) DeepCopy() *Reference {
	if in == nil {
		return nil
	}
	out := new(Reference)
	in.DeepCopyInto(out)
	return out
}
