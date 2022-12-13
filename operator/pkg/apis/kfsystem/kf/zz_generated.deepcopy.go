//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package kf

import (
	v1 "k8s.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in BuildNodeSelectors) DeepCopyInto(out *BuildNodeSelectors) {
	{
		in := &in
		*out = make(BuildNodeSelectors, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildNodeSelectors.
func (in BuildNodeSelectors) DeepCopy() BuildNodeSelectors {
	if in == nil {
		return nil
	}
	out := new(BuildNodeSelectors)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildpackV2Definition) DeepCopyInto(out *BuildpackV2Definition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildpackV2Definition.
func (in *BuildpackV2Definition) DeepCopy() *BuildpackV2Definition {
	if in == nil {
		return nil
	}
	out := new(BuildpackV2Definition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in BuildpackV2List) DeepCopyInto(out *BuildpackV2List) {
	{
		in := &in
		*out = make(BuildpackV2List, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildpackV2List.
func (in BuildpackV2List) DeepCopy() BuildpackV2List {
	if in == nil {
		return nil
	}
	out := new(BuildpackV2List)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultsConfig) DeepCopyInto(out *DefaultsConfig) {
	*out = *in
	if in.SpaceClusterDomains != nil {
		in, out := &in.SpaceClusterDomains, &out.SpaceClusterDomains
		*out = make([]DomainTemplate, len(*in))
		copy(*out, *in)
	}
	if in.SpaceBuildpacksV2 != nil {
		in, out := &in.SpaceBuildpacksV2, &out.SpaceBuildpacksV2
		*out = make(BuildpackV2List, len(*in))
		copy(*out, *in)
	}
	if in.SpaceStacksV2 != nil {
		in, out := &in.SpaceStacksV2, &out.SpaceStacksV2
		*out = make(StackV2List, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SpaceStacksV3 != nil {
		in, out := &in.SpaceStacksV3, &out.SpaceStacksV3
		*out = make(StackV3List, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BuildPodResources != nil {
		in, out := &in.BuildPodResources, &out.BuildPodResources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.BuildRetentionCount != nil {
		in, out := &in.BuildRetentionCount, &out.BuildRetentionCount
		*out = new(uint)
		**out = **in
	}
	if in.BuildNodeSelectors != nil {
		in, out := &in.BuildNodeSelectors, &out.BuildNodeSelectors
		*out = make(BuildNodeSelectors, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FeatureFlags != nil {
		in, out := &in.FeatureFlags, &out.FeatureFlags
		*out = make(FeatureFlagToggles, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultsConfig.
func (in *DefaultsConfig) DeepCopy() *DefaultsConfig {
	if in == nil {
		return nil
	}
	out := new(DefaultsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainTemplate) DeepCopyInto(out *DomainTemplate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainTemplate.
func (in *DomainTemplate) DeepCopy() *DomainTemplate {
	if in == nil {
		return nil
	}
	out := new(DomainTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FeatureFlagToggles) DeepCopyInto(out *FeatureFlagToggles) {
	{
		in := &in
		*out = make(FeatureFlagToggles, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagToggles.
func (in FeatureFlagToggles) DeepCopy() FeatureFlagToggles {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagToggles)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackV2Definition) DeepCopyInto(out *StackV2Definition) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackV2Definition.
func (in *StackV2Definition) DeepCopy() *StackV2Definition {
	if in == nil {
		return nil
	}
	out := new(StackV2Definition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in StackV2List) DeepCopyInto(out *StackV2List) {
	{
		in := &in
		*out = make(StackV2List, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackV2List.
func (in StackV2List) DeepCopy() StackV2List {
	if in == nil {
		return nil
	}
	out := new(StackV2List)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackV3Definition) DeepCopyInto(out *StackV3Definition) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackV3Definition.
func (in *StackV3Definition) DeepCopy() *StackV3Definition {
	if in == nil {
		return nil
	}
	out := new(StackV3Definition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in StackV3List) DeepCopyInto(out *StackV3List) {
	{
		in := &in
		*out = make(StackV3List, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackV3List.
func (in StackV3List) DeepCopy() StackV3List {
	if in == nil {
		return nil
	}
	out := new(StackV3List)
	in.DeepCopyInto(out)
	return *out
}
