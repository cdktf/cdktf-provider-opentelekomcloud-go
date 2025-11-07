// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsfunctionv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/fgsfunctionv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FgsFunctionV2FuncMountsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalMountPath() *string
	SetLocalMountPath(val *string)
	LocalMountPathInput() *string
	MountResource() *string
	SetMountResource(val *string)
	MountResourceInput() *string
	MountSharePath() *string
	SetMountSharePath(val *string)
	MountSharePathInput() *string
	MountType() *string
	SetMountType(val *string)
	MountTypeInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FgsFunctionV2FuncMountsOutputReference
type jsiiProxy_FgsFunctionV2FuncMountsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) LocalMountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localMountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) LocalMountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localMountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) MountResource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) MountResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) MountSharePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountSharePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) MountSharePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountSharePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) MountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) MountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFgsFunctionV2FuncMountsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FgsFunctionV2FuncMountsOutputReference {
	_init_.Initialize()

	if err := validateNewFgsFunctionV2FuncMountsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FgsFunctionV2FuncMountsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2FuncMountsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFgsFunctionV2FuncMountsOutputReference_Override(f FgsFunctionV2FuncMountsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2FuncMountsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference)SetLocalMountPath(val *string) {
	if err := j.validateSetLocalMountPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localMountPath",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference)SetMountResource(val *string) {
	if err := j.validateSetMountResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountResource",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference)SetMountSharePath(val *string) {
	if err := j.validateSetMountSharePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountSharePath",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference)SetMountType(val *string) {
	if err := j.validateSetMountTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountType",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2FuncMountsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (f *jsiiProxy_FgsFunctionV2FuncMountsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

