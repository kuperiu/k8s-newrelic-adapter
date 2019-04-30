// +build !ignore_autogenerated

// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// A copy of the License is located at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dimension) DeepCopyInto(out *Dimension) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dimension.
func (in *Dimension) DeepCopy() *Dimension {
	if in == nil {
		return nil
	}
	out := new(Dimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalMetric) DeepCopyInto(out *ExternalMetric) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalMetric.
func (in *ExternalMetric) DeepCopy() *ExternalMetric {
	if in == nil {
		return nil
	}
	out := new(ExternalMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalMetric) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalMetricList) DeepCopyInto(out *ExternalMetricList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalMetricList.
func (in *ExternalMetricList) DeepCopy() *ExternalMetricList {
	if in == nil {
		return nil
	}
	out := new(ExternalMetricList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalMetricList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metric) DeepCopyInto(out *Metric) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]Dimension, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metric.
func (in *Metric) DeepCopy() *Metric {
	if in == nil {
		return nil
	}
	out := new(Metric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricDataQuery) DeepCopyInto(out *MetricDataQuery) {
	*out = *in
	in.MetricStat.DeepCopyInto(&out.MetricStat)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricDataQuery.
func (in *MetricDataQuery) DeepCopy() *MetricDataQuery {
	if in == nil {
		return nil
	}
	out := new(MetricDataQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSeriesSpec) DeepCopyInto(out *MetricSeriesSpec) {
	*out = *in
	if in.Queries != nil {
		in, out := &in.Queries, &out.Queries
		*out = make([]MetricDataQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSeriesSpec.
func (in *MetricSeriesSpec) DeepCopy() *MetricSeriesSpec {
	if in == nil {
		return nil
	}
	out := new(MetricSeriesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricStat) DeepCopyInto(out *MetricStat) {
	*out = *in
	in.Metric.DeepCopyInto(&out.Metric)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricStat.
func (in *MetricStat) DeepCopy() *MetricStat {
	if in == nil {
		return nil
	}
	out := new(MetricStat)
	in.DeepCopyInto(out)
	return out
}
