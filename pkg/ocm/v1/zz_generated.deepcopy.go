//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnParameter) DeepCopyInto(out *AddOnParameter) {
	*out = *in
	if in.Validation != nil {
		in, out := &in.Validation, &out.Validation
		*out = new(string)
		**out = **in
	}
	if in.ValidationErrMsg != nil {
		in, out := &in.ValidationErrMsg, &out.ValidationErrMsg
		*out = new(string)
		**out = **in
	}
	if in.DefaultValue != nil {
		in, out := &in.DefaultValue, &out.DefaultValue
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(int)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new([]AddOnParameterOption)
		if **in != nil {
			in, out := *in, *out
			*out = make([]AddOnParameterOption, len(*in))
			copy(*out, *in)
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = new([]AddOnResourceRequirement)
		if **in != nil {
			in, out := *in, *out
			*out = make([]AddOnResourceRequirement, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnParameter.
func (in *AddOnParameter) DeepCopy() *AddOnParameter {
	if in == nil {
		return nil
	}
	out := new(AddOnParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnParameterOption) DeepCopyInto(out *AddOnParameterOption) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnParameterOption.
func (in *AddOnParameterOption) DeepCopy() *AddOnParameterOption {
	if in == nil {
		return nil
	}
	out := new(AddOnParameterOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnRequirement) DeepCopyInto(out *AddOnRequirement) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(AddOnRequirementData, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(AddOnResourceRequirementStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnRequirement.
func (in *AddOnRequirement) DeepCopy() *AddOnRequirement {
	if in == nil {
		return nil
	}
	out := new(AddOnRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnResourceRequirement) DeepCopyInto(out *AddOnResourceRequirement) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(AddOnRequirementData, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(AddOnResourceRequirementStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnResourceRequirement.
func (in *AddOnResourceRequirement) DeepCopy() *AddOnResourceRequirement {
	if in == nil {
		return nil
	}
	out := new(AddOnResourceRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnResourceRequirementStatus) DeepCopyInto(out *AddOnResourceRequirementStatus) {
	*out = *in
	if in.Fulfilled != nil {
		in, out := &in.Fulfilled, &out.Fulfilled
		*out = new(bool)
		**out = **in
	}
	if in.ErrorMsgs != nil {
		in, out := &in.ErrorMsgs, &out.ErrorMsgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnResourceRequirementStatus.
func (in *AddOnResourceRequirementStatus) DeepCopy() *AddOnResourceRequirementStatus {
	if in == nil {
		return nil
	}
	out := new(AddOnResourceRequirementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnSubOperator) DeepCopyInto(out *AddOnSubOperator) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnSubOperator.
func (in *AddOnSubOperator) DeepCopy() *AddOnSubOperator {
	if in == nil {
		return nil
	}
	out := new(AddOnSubOperator)
	in.DeepCopyInto(out)
	return out
}
