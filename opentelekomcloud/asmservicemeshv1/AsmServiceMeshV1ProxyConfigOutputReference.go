// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asmservicemeshv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/asmservicemeshv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AsmServiceMeshV1ProxyConfigOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludeInboundPorts() *string
	SetExcludeInboundPorts(val *string)
	ExcludeInboundPortsInput() *string
	ExcludeIpRanges() *string
	SetExcludeIpRanges(val *string)
	ExcludeIpRangesInput() *string
	ExcludeOutboundPorts() *string
	SetExcludeOutboundPorts(val *string)
	ExcludeOutboundPortsInput() *string
	// Experimental.
	Fqn() *string
	IncludeInboundPorts() *string
	SetIncludeInboundPorts(val *string)
	IncludeInboundPortsInput() *string
	IncludeIpRanges() *string
	SetIncludeIpRanges(val *string)
	IncludeIpRangesInput() *string
	IncludeOutboundPorts() *string
	SetIncludeOutboundPorts(val *string)
	IncludeOutboundPortsInput() *string
	InternalValue() *AsmServiceMeshV1ProxyConfig
	SetInternalValue(val *AsmServiceMeshV1ProxyConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetExcludeInboundPorts()
	ResetExcludeIpRanges()
	ResetExcludeOutboundPorts()
	ResetIncludeInboundPorts()
	ResetIncludeIpRanges()
	ResetIncludeOutboundPorts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AsmServiceMeshV1ProxyConfigOutputReference
type jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ExcludeInboundPorts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeInboundPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ExcludeInboundPortsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeInboundPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ExcludeIpRanges() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ExcludeIpRangesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ExcludeOutboundPorts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeOutboundPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ExcludeOutboundPortsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeOutboundPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) IncludeInboundPorts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeInboundPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) IncludeInboundPortsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeInboundPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) IncludeIpRanges() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) IncludeIpRangesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) IncludeOutboundPorts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeOutboundPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) IncludeOutboundPortsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeOutboundPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) InternalValue() *AsmServiceMeshV1ProxyConfig {
	var returns *AsmServiceMeshV1ProxyConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAsmServiceMeshV1ProxyConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AsmServiceMeshV1ProxyConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAsmServiceMeshV1ProxyConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.asmServiceMeshV1.AsmServiceMeshV1ProxyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAsmServiceMeshV1ProxyConfigOutputReference_Override(a AsmServiceMeshV1ProxyConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.asmServiceMeshV1.AsmServiceMeshV1ProxyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference)SetExcludeInboundPorts(val *string) {
	if err := j.validateSetExcludeInboundPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeInboundPorts",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference)SetExcludeIpRanges(val *string) {
	if err := j.validateSetExcludeIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeIpRanges",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference)SetExcludeOutboundPorts(val *string) {
	if err := j.validateSetExcludeOutboundPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeOutboundPorts",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference)SetIncludeInboundPorts(val *string) {
	if err := j.validateSetIncludeInboundPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeInboundPorts",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference)SetIncludeIpRanges(val *string) {
	if err := j.validateSetIncludeIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeIpRanges",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference)SetIncludeOutboundPorts(val *string) {
	if err := j.validateSetIncludeOutboundPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeOutboundPorts",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference)SetInternalValue(val *AsmServiceMeshV1ProxyConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ResetExcludeInboundPorts() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludeInboundPorts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ResetExcludeIpRanges() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludeIpRanges",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ResetExcludeOutboundPorts() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludeOutboundPorts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ResetIncludeInboundPorts() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeInboundPorts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ResetIncludeIpRanges() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeIpRanges",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ResetIncludeOutboundPorts() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeOutboundPorts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1ProxyConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

