// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwapiv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/apigwapiv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigwApiV2FuncGraphPolicyOutputReference interface {
	cdktf.ComplexObject
	AuthorizerId() *string
	SetAuthorizerId(val *string)
	AuthorizerIdInput() *string
	BackendParams() ApigwApiV2FuncGraphPolicyBackendParamsList
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
	Conditions() ApigwApiV2FuncGraphPolicyConditionsList
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
	FunctionUrn() *string
	SetFunctionUrn(val *string)
	FunctionUrnInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	InvocationType() *string
	SetInvocationType(val *string)
	InvocationTypeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutBackendParams(value interface{})
	PutConditions(value interface{})
	ResetAuthorizerId()
	ResetBackendParams()
	ResetEffectiveMode()
	ResetInvocationType()
	ResetNetworkType()
	ResetTimeout()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigwApiV2FuncGraphPolicyOutputReference
type jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) AuthorizerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) BackendParams() ApigwApiV2FuncGraphPolicyBackendParamsList {
	var returns ApigwApiV2FuncGraphPolicyBackendParamsList
	_jsii_.Get(
		j,
		"backendParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) BackendParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) Conditions() ApigwApiV2FuncGraphPolicyConditionsList {
	var returns ApigwApiV2FuncGraphPolicyConditionsList
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) ConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) EffectiveMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) EffectiveModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) FunctionUrn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionUrn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) FunctionUrnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionUrnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) InvocationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) InvocationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewApigwApiV2FuncGraphPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApigwApiV2FuncGraphPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewApigwApiV2FuncGraphPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2FuncGraphPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApigwApiV2FuncGraphPolicyOutputReference_Override(a ApigwApiV2FuncGraphPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2FuncGraphPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetAuthorizerId(val *string) {
	if err := j.validateSetAuthorizerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerId",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetEffectiveMode(val *string) {
	if err := j.validateSetEffectiveModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveMode",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetFunctionUrn(val *string) {
	if err := j.validateSetFunctionUrnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionUrn",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetInvocationType(val *string) {
	if err := j.validateSetInvocationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invocationType",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) PutBackendParams(value interface{}) {
	if err := a.validatePutBackendParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackendParams",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) PutConditions(value interface{}) {
	if err := a.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConditions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) ResetAuthorizerId() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) ResetBackendParams() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) ResetEffectiveMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEffectiveMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) ResetInvocationType() {
	_jsii_.InvokeVoid(
		a,
		"resetInvocationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) ResetNetworkType() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigwApiV2FuncGraphPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

