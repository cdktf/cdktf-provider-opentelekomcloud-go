// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsinstancev3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/rdsinstancev3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsInstanceV3VolumeOutputReference interface {
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
	DiskEncryptionId() *string
	SetDiskEncryptionId(val *string)
	DiskEncryptionIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *RdsInstanceV3Volume
	SetInternalValue(val *RdsInstanceV3Volume)
	LimitSize() *float64
	SetLimitSize(val *float64)
	LimitSizeInput() *float64
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
	TriggerThreshold() *float64
	SetTriggerThreshold(val *float64)
	TriggerThresholdInput() *float64
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetDiskEncryptionId()
	ResetLimitSize()
	ResetTriggerThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RdsInstanceV3VolumeOutputReference
type jsiiProxy_RdsInstanceV3VolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) DiskEncryptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) DiskEncryptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) InternalValue() *RdsInstanceV3Volume {
	var returns *RdsInstanceV3Volume
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) LimitSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) LimitSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) TriggerThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"triggerThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) TriggerThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"triggerThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewRdsInstanceV3VolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RdsInstanceV3VolumeOutputReference {
	_init_.Initialize()

	if err := validateNewRdsInstanceV3VolumeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsInstanceV3VolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV3.RdsInstanceV3VolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRdsInstanceV3VolumeOutputReference_Override(r RdsInstanceV3VolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV3.RdsInstanceV3VolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference)SetDiskEncryptionId(val *string) {
	if err := j.validateSetDiskEncryptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionId",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference)SetInternalValue(val *RdsInstanceV3Volume) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference)SetLimitSize(val *float64) {
	if err := j.validateSetLimitSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitSize",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference)SetTriggerThreshold(val *float64) {
	if err := j.validateSetTriggerThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerThreshold",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3VolumeOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) ResetDiskEncryptionId() {
	_jsii_.InvokeVoid(
		r,
		"resetDiskEncryptionId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) ResetLimitSize() {
	_jsii_.InvokeVoid(
		r,
		"resetLimitSize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) ResetTriggerThreshold() {
	_jsii_.InvokeVoid(
		r,
		"resetTriggerThreshold",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3VolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

