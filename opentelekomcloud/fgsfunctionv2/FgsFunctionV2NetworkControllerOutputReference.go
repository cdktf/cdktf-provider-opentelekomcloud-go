// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsfunctionv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/fgsfunctionv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FgsFunctionV2NetworkControllerOutputReference interface {
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
	DisablePublicNetwork() interface{}
	SetDisablePublicNetwork(val interface{})
	DisablePublicNetworkInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *FgsFunctionV2NetworkController
	SetInternalValue(val *FgsFunctionV2NetworkController)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TriggerAccessVpcs() FgsFunctionV2NetworkControllerTriggerAccessVpcsList
	TriggerAccessVpcsInput() interface{}
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
	PutTriggerAccessVpcs(value interface{})
	ResetDisablePublicNetwork()
	ResetTriggerAccessVpcs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FgsFunctionV2NetworkControllerOutputReference
type jsiiProxy_FgsFunctionV2NetworkControllerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) DisablePublicNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePublicNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) DisablePublicNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePublicNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) InternalValue() *FgsFunctionV2NetworkController {
	var returns *FgsFunctionV2NetworkController
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) TriggerAccessVpcs() FgsFunctionV2NetworkControllerTriggerAccessVpcsList {
	var returns FgsFunctionV2NetworkControllerTriggerAccessVpcsList
	_jsii_.Get(
		j,
		"triggerAccessVpcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) TriggerAccessVpcsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerAccessVpcsInput",
		&returns,
	)
	return returns
}


func NewFgsFunctionV2NetworkControllerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FgsFunctionV2NetworkControllerOutputReference {
	_init_.Initialize()

	if err := validateNewFgsFunctionV2NetworkControllerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FgsFunctionV2NetworkControllerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2NetworkControllerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFgsFunctionV2NetworkControllerOutputReference_Override(f FgsFunctionV2NetworkControllerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2NetworkControllerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference)SetDisablePublicNetwork(val interface{}) {
	if err := j.validateSetDisablePublicNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePublicNetwork",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference)SetInternalValue(val *FgsFunctionV2NetworkController) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) PutTriggerAccessVpcs(value interface{}) {
	if err := f.validatePutTriggerAccessVpcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTriggerAccessVpcs",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) ResetDisablePublicNetwork() {
	_jsii_.InvokeVoid(
		f,
		"resetDisablePublicNetwork",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) ResetTriggerAccessVpcs() {
	_jsii_.InvokeVoid(
		f,
		"resetTriggerAccessVpcs",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (f *jsiiProxy_FgsFunctionV2NetworkControllerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

