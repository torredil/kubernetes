/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// CustomResourceSubresourceScaleApplyConfiguration represents a declarative configuration of the CustomResourceSubresourceScale type for use
// with apply.
type CustomResourceSubresourceScaleApplyConfiguration struct {
	SpecReplicasPath   *string `json:"specReplicasPath,omitempty"`
	StatusReplicasPath *string `json:"statusReplicasPath,omitempty"`
	LabelSelectorPath  *string `json:"labelSelectorPath,omitempty"`
}

// CustomResourceSubresourceScaleApplyConfiguration constructs a declarative configuration of the CustomResourceSubresourceScale type for use with
// apply.
func CustomResourceSubresourceScale() *CustomResourceSubresourceScaleApplyConfiguration {
	return &CustomResourceSubresourceScaleApplyConfiguration{}
}
func (b CustomResourceSubresourceScaleApplyConfiguration) IsApplyConfiguration() {}

// WithSpecReplicasPath sets the SpecReplicasPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SpecReplicasPath field is set to the value of the last call.
func (b *CustomResourceSubresourceScaleApplyConfiguration) WithSpecReplicasPath(value string) *CustomResourceSubresourceScaleApplyConfiguration {
	b.SpecReplicasPath = &value
	return b
}

// WithStatusReplicasPath sets the StatusReplicasPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StatusReplicasPath field is set to the value of the last call.
func (b *CustomResourceSubresourceScaleApplyConfiguration) WithStatusReplicasPath(value string) *CustomResourceSubresourceScaleApplyConfiguration {
	b.StatusReplicasPath = &value
	return b
}

// WithLabelSelectorPath sets the LabelSelectorPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelSelectorPath field is set to the value of the last call.
func (b *CustomResourceSubresourceScaleApplyConfiguration) WithLabelSelectorPath(value string) *CustomResourceSubresourceScaleApplyConfiguration {
	b.LabelSelectorPath = &value
	return b
}
