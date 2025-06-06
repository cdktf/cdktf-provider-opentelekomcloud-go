// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodepoolv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/ccenodepoolv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CceNodePoolV3RootVolumeOutputReference interface {
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
	ExtendParam() *string
	SetExtendParam(val *string)
	ExtendParamInput() *string
	ExtendParams() *map[string]*string
	SetExtendParams(val *map[string]*string)
	ExtendParamsInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CceNodePoolV3RootVolume
	SetInternalValue(val *CceNodePoolV3RootVolume)
	KmsId() *string
	SetKmsId(val *string)
	KmsIdInput() *string
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volumetype() *string
	SetVolumetype(val *string)
	VolumetypeInput() *string
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
	ResetExtendParam()
	ResetExtendParams()
	ResetKmsId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CceNodePoolV3RootVolumeOutputReference
type jsiiProxy_CceNodePoolV3RootVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) ExtendParam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) ExtendParamInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) ExtendParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extendParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) ExtendParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extendParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) InternalValue() *CceNodePoolV3RootVolume {
	var returns *CceNodePoolV3RootVolume
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) KmsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) KmsIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) Volumetype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumetype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) VolumetypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumetypeInput",
		&returns,
	)
	return returns
}


func NewCceNodePoolV3RootVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CceNodePoolV3RootVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewCceNodePoolV3RootVolumeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CceNodePoolV3RootVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cceNodePoolV3.CceNodePoolV3RootVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCceNodePoolV3RootVolumeOutputReference_Override(c CceNodePoolV3RootVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cceNodePoolV3.CceNodePoolV3RootVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference)SetExtendParam(val *string) {
	if err := j.validateSetExtendParamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendParam",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference)SetExtendParams(val *map[string]*string) {
	if err := j.validateSetExtendParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendParams",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference)SetInternalValue(val *CceNodePoolV3RootVolume) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference)SetKmsId(val *string) {
	if err := j.validateSetKmsIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsId",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3RootVolumeOutputReference)SetVolumetype(val *string) {
	if err := j.validateSetVolumetypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumetype",
		val,
	)
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) ResetExtendParam() {
	_jsii_.InvokeVoid(
		c,
		"resetExtendParam",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) ResetExtendParams() {
	_jsii_.InvokeVoid(
		c,
		"resetExtendParams",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) ResetKmsId() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3RootVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

