// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwvpcchannelv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/apigwvpcchannelv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigwVpcChannelV2MicroserviceOutputReference interface {
	cdktf.ComplexObject
	CceConfig() ApigwVpcChannelV2MicroserviceCceConfigOutputReference
	CceConfigInput() *ApigwVpcChannelV2MicroserviceCceConfig
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
	CseConfig() ApigwVpcChannelV2MicroserviceCseConfigOutputReference
	CseConfigInput() *ApigwVpcChannelV2MicroserviceCseConfig
	// Experimental.
	Fqn() *string
	InternalValue() *ApigwVpcChannelV2Microservice
	SetInternalValue(val *ApigwVpcChannelV2Microservice)
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
	PutCceConfig(value *ApigwVpcChannelV2MicroserviceCceConfig)
	PutCseConfig(value *ApigwVpcChannelV2MicroserviceCseConfig)
	ResetCceConfig()
	ResetCseConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigwVpcChannelV2MicroserviceOutputReference
type jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) CceConfig() ApigwVpcChannelV2MicroserviceCceConfigOutputReference {
	var returns ApigwVpcChannelV2MicroserviceCceConfigOutputReference
	_jsii_.Get(
		j,
		"cceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) CceConfigInput() *ApigwVpcChannelV2MicroserviceCceConfig {
	var returns *ApigwVpcChannelV2MicroserviceCceConfig
	_jsii_.Get(
		j,
		"cceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) CseConfig() ApigwVpcChannelV2MicroserviceCseConfigOutputReference {
	var returns ApigwVpcChannelV2MicroserviceCseConfigOutputReference
	_jsii_.Get(
		j,
		"cseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) CseConfigInput() *ApigwVpcChannelV2MicroserviceCseConfig {
	var returns *ApigwVpcChannelV2MicroserviceCseConfig
	_jsii_.Get(
		j,
		"cseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) InternalValue() *ApigwVpcChannelV2Microservice {
	var returns *ApigwVpcChannelV2Microservice
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApigwVpcChannelV2MicroserviceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApigwVpcChannelV2MicroserviceOutputReference {
	_init_.Initialize()

	if err := validateNewApigwVpcChannelV2MicroserviceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwVpcChannelV2.ApigwVpcChannelV2MicroserviceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigwVpcChannelV2MicroserviceOutputReference_Override(a ApigwVpcChannelV2MicroserviceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwVpcChannelV2.ApigwVpcChannelV2MicroserviceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference)SetInternalValue(val *ApigwVpcChannelV2Microservice) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) PutCceConfig(value *ApigwVpcChannelV2MicroserviceCceConfig) {
	if err := a.validatePutCceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCceConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) PutCseConfig(value *ApigwVpcChannelV2MicroserviceCseConfig) {
	if err := a.validatePutCseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) ResetCceConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCceConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) ResetCseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

