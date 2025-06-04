// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltshostaccessv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/ltshostaccessv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsHostAccessV3AccessConfigOutputReference interface {
	cdktf.ComplexObject
	BlackPaths() *[]*string
	SetBlackPaths(val *[]*string)
	BlackPathsInput() *[]*string
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
	// Experimental.
	Fqn() *string
	InternalValue() *LtsHostAccessV3AccessConfig
	SetInternalValue(val *LtsHostAccessV3AccessConfig)
	MultiLogFormat() LtsHostAccessV3AccessConfigMultiLogFormatOutputReference
	MultiLogFormatInput() *LtsHostAccessV3AccessConfigMultiLogFormat
	Paths() *[]*string
	SetPaths(val *[]*string)
	PathsInput() *[]*string
	SingleLogFormat() LtsHostAccessV3AccessConfigSingleLogFormatOutputReference
	SingleLogFormatInput() *LtsHostAccessV3AccessConfigSingleLogFormat
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsLogInfo() LtsHostAccessV3AccessConfigWindowsLogInfoOutputReference
	WindowsLogInfoInput() *LtsHostAccessV3AccessConfigWindowsLogInfo
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutMultiLogFormat(value *LtsHostAccessV3AccessConfigMultiLogFormat)
	PutSingleLogFormat(value *LtsHostAccessV3AccessConfigSingleLogFormat)
	PutWindowsLogInfo(value *LtsHostAccessV3AccessConfigWindowsLogInfo)
	ResetBlackPaths()
	ResetMultiLogFormat()
	ResetSingleLogFormat()
	ResetWindowsLogInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LtsHostAccessV3AccessConfigOutputReference
type jsiiProxy_LtsHostAccessV3AccessConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) BlackPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blackPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) BlackPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blackPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) InternalValue() *LtsHostAccessV3AccessConfig {
	var returns *LtsHostAccessV3AccessConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) MultiLogFormat() LtsHostAccessV3AccessConfigMultiLogFormatOutputReference {
	var returns LtsHostAccessV3AccessConfigMultiLogFormatOutputReference
	_jsii_.Get(
		j,
		"multiLogFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) MultiLogFormatInput() *LtsHostAccessV3AccessConfigMultiLogFormat {
	var returns *LtsHostAccessV3AccessConfigMultiLogFormat
	_jsii_.Get(
		j,
		"multiLogFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) Paths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"paths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) PathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) SingleLogFormat() LtsHostAccessV3AccessConfigSingleLogFormatOutputReference {
	var returns LtsHostAccessV3AccessConfigSingleLogFormatOutputReference
	_jsii_.Get(
		j,
		"singleLogFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) SingleLogFormatInput() *LtsHostAccessV3AccessConfigSingleLogFormat {
	var returns *LtsHostAccessV3AccessConfigSingleLogFormat
	_jsii_.Get(
		j,
		"singleLogFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) WindowsLogInfo() LtsHostAccessV3AccessConfigWindowsLogInfoOutputReference {
	var returns LtsHostAccessV3AccessConfigWindowsLogInfoOutputReference
	_jsii_.Get(
		j,
		"windowsLogInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) WindowsLogInfoInput() *LtsHostAccessV3AccessConfigWindowsLogInfo {
	var returns *LtsHostAccessV3AccessConfigWindowsLogInfo
	_jsii_.Get(
		j,
		"windowsLogInfoInput",
		&returns,
	)
	return returns
}


func NewLtsHostAccessV3AccessConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LtsHostAccessV3AccessConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLtsHostAccessV3AccessConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LtsHostAccessV3AccessConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsHostAccessV3.LtsHostAccessV3AccessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLtsHostAccessV3AccessConfigOutputReference_Override(l LtsHostAccessV3AccessConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsHostAccessV3.LtsHostAccessV3AccessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference)SetBlackPaths(val *[]*string) {
	if err := j.validateSetBlackPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blackPaths",
		val,
	)
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference)SetInternalValue(val *LtsHostAccessV3AccessConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference)SetPaths(val *[]*string) {
	if err := j.validateSetPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paths",
		val,
	)
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) PutMultiLogFormat(value *LtsHostAccessV3AccessConfigMultiLogFormat) {
	if err := l.validatePutMultiLogFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMultiLogFormat",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) PutSingleLogFormat(value *LtsHostAccessV3AccessConfigSingleLogFormat) {
	if err := l.validatePutSingleLogFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSingleLogFormat",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) PutWindowsLogInfo(value *LtsHostAccessV3AccessConfigWindowsLogInfo) {
	if err := l.validatePutWindowsLogInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putWindowsLogInfo",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) ResetBlackPaths() {
	_jsii_.InvokeVoid(
		l,
		"resetBlackPaths",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) ResetMultiLogFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetMultiLogFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) ResetSingleLogFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetSingleLogFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) ResetWindowsLogInfo() {
	_jsii_.InvokeVoid(
		l,
		"resetWindowsLogInfo",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsHostAccessV3AccessConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

