// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsfunctionv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/fgsfunctionv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FgsFunctionV2ReservedInstancesOutputReference interface {
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IdleMode() interface{}
	SetIdleMode(val interface{})
	IdleModeInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	QualifierName() *string
	SetQualifierName(val *string)
	QualifierNameInput() *string
	QualifierType() *string
	SetQualifierType(val *string)
	QualifierTypeInput() *string
	TacticsConfig() FgsFunctionV2ReservedInstancesTacticsConfigOutputReference
	TacticsConfigInput() *FgsFunctionV2ReservedInstancesTacticsConfig
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
	PutTacticsConfig(value *FgsFunctionV2ReservedInstancesTacticsConfig)
	ResetIdleMode()
	ResetTacticsConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FgsFunctionV2ReservedInstancesOutputReference
type jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) IdleMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idleMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) IdleModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idleModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) QualifierName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) QualifierNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) QualifierType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) QualifierTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) TacticsConfig() FgsFunctionV2ReservedInstancesTacticsConfigOutputReference {
	var returns FgsFunctionV2ReservedInstancesTacticsConfigOutputReference
	_jsii_.Get(
		j,
		"tacticsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) TacticsConfigInput() *FgsFunctionV2ReservedInstancesTacticsConfig {
	var returns *FgsFunctionV2ReservedInstancesTacticsConfig
	_jsii_.Get(
		j,
		"tacticsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFgsFunctionV2ReservedInstancesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FgsFunctionV2ReservedInstancesOutputReference {
	_init_.Initialize()

	if err := validateNewFgsFunctionV2ReservedInstancesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2ReservedInstancesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFgsFunctionV2ReservedInstancesOutputReference_Override(f FgsFunctionV2ReservedInstancesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2ReservedInstancesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference)SetCount(val *float64) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference)SetIdleMode(val interface{}) {
	if err := j.validateSetIdleModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleMode",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference)SetQualifierName(val *string) {
	if err := j.validateSetQualifierNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qualifierName",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference)SetQualifierType(val *string) {
	if err := j.validateSetQualifierTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qualifierType",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) PutTacticsConfig(value *FgsFunctionV2ReservedInstancesTacticsConfig) {
	if err := f.validatePutTacticsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTacticsConfig",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) ResetIdleMode() {
	_jsii_.InvokeVoid(
		f,
		"resetIdleMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) ResetTacticsConfig() {
	_jsii_.InvokeVoid(
		f,
		"resetTacticsConfig",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2ReservedInstancesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

