// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwapiv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/apigwapiv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigwApiV2MockPolicyOutputReference interface {
	cdktf.ComplexObject
	AuthorizerId() *string
	SetAuthorizerId(val *string)
	AuthorizerIdInput() *string
	BackendParams() ApigwApiV2MockPolicyBackendParamsList
	BackendParamsInput() interface{}
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
	Conditions() ApigwApiV2MockPolicyConditionsList
	ConditionsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EffectiveMode() *string
	SetEffectiveMode(val *string)
	EffectiveModeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Response() *string
	SetResponse(val *string)
	ResponseInput() *string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutBackendParams(value interface{})
	PutConditions(value interface{})
	ResetAuthorizerId()
	ResetBackendParams()
	ResetEffectiveMode()
	ResetResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigwApiV2MockPolicyOutputReference
type jsiiProxy_ApigwApiV2MockPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) AuthorizerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) BackendParams() ApigwApiV2MockPolicyBackendParamsList {
	var returns ApigwApiV2MockPolicyBackendParamsList
	_jsii_.Get(
		j,
		"backendParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) BackendParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) Conditions() ApigwApiV2MockPolicyConditionsList {
	var returns ApigwApiV2MockPolicyConditionsList
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) ConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) EffectiveMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) EffectiveModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) Response() *string {
	var returns *string
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) ResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApigwApiV2MockPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApigwApiV2MockPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewApigwApiV2MockPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigwApiV2MockPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2MockPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApigwApiV2MockPolicyOutputReference_Override(a ApigwApiV2MockPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2MockPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference)SetAuthorizerId(val *string) {
	if err := j.validateSetAuthorizerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerId",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference)SetEffectiveMode(val *string) {
	if err := j.validateSetEffectiveModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveMode",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference)SetResponse(val *string) {
	if err := j.validateSetResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"response",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2MockPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) PutBackendParams(value interface{}) {
	if err := a.validatePutBackendParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackendParams",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) PutConditions(value interface{}) {
	if err := a.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConditions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) ResetAuthorizerId() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) ResetBackendParams() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) ResetEffectiveMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEffectiveMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) ResetResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2MockPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

