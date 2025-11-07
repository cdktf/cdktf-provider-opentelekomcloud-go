// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwapiv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/apigwapiv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigwApiV2HttpPolicyOutputReference interface {
	cdktf.ComplexObject
	AuthorizerId() *string
	SetAuthorizerId(val *string)
	AuthorizerIdInput() *string
	BackendParams() ApigwApiV2HttpPolicyBackendParamsList
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
	Conditions() ApigwApiV2HttpPolicyConditionsList
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
	RequestMethod() *string
	SetRequestMethod(val *string)
	RequestMethodInput() *string
	RequestProtocol() *string
	SetRequestProtocol(val *string)
	RequestProtocolInput() *string
	RequestUri() *string
	SetRequestUri(val *string)
	RequestUriInput() *string
	RetryCount() *float64
	SetRetryCount(val *float64)
	RetryCountInput() *float64
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
	UrlDomain() *string
	SetUrlDomain(val *string)
	UrlDomainInput() *string
	VpcChannelId() *string
	SetVpcChannelId(val *string)
	VpcChannelIdInput() *string
	VpcChannelProxyHost() *string
	SetVpcChannelProxyHost(val *string)
	VpcChannelProxyHostInput() *string
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
	ResetRequestProtocol()
	ResetRetryCount()
	ResetTimeout()
	ResetUrlDomain()
	ResetVpcChannelId()
	ResetVpcChannelProxyHost()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigwApiV2HttpPolicyOutputReference
type jsiiProxy_ApigwApiV2HttpPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) AuthorizerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) BackendParams() ApigwApiV2HttpPolicyBackendParamsList {
	var returns ApigwApiV2HttpPolicyBackendParamsList
	_jsii_.Get(
		j,
		"backendParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) BackendParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) Conditions() ApigwApiV2HttpPolicyConditionsList {
	var returns ApigwApiV2HttpPolicyConditionsList
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) EffectiveMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) EffectiveModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) RequestMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) RequestMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) RequestProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) RequestProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) RequestUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) RequestUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) RetryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) RetryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) UrlDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) UrlDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) VpcChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcChannelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) VpcChannelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcChannelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) VpcChannelProxyHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcChannelProxyHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) VpcChannelProxyHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcChannelProxyHostInput",
		&returns,
	)
	return returns
}


func NewApigwApiV2HttpPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApigwApiV2HttpPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewApigwApiV2HttpPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigwApiV2HttpPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2HttpPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApigwApiV2HttpPolicyOutputReference_Override(a ApigwApiV2HttpPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2HttpPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetAuthorizerId(val *string) {
	if err := j.validateSetAuthorizerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerId",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetEffectiveMode(val *string) {
	if err := j.validateSetEffectiveModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveMode",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetRequestMethod(val *string) {
	if err := j.validateSetRequestMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMethod",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetRequestProtocol(val *string) {
	if err := j.validateSetRequestProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestProtocol",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetRequestUri(val *string) {
	if err := j.validateSetRequestUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestUri",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetRetryCount(val *float64) {
	if err := j.validateSetRetryCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryCount",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetUrlDomain(val *string) {
	if err := j.validateSetUrlDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlDomain",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetVpcChannelId(val *string) {
	if err := j.validateSetVpcChannelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcChannelId",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpPolicyOutputReference)SetVpcChannelProxyHost(val *string) {
	if err := j.validateSetVpcChannelProxyHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcChannelProxyHost",
		val,
	)
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) PutBackendParams(value interface{}) {
	if err := a.validatePutBackendParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackendParams",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) PutConditions(value interface{}) {
	if err := a.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConditions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ResetAuthorizerId() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ResetBackendParams() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ResetEffectiveMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEffectiveMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ResetRequestProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ResetRetryCount() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ResetUrlDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetUrlDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ResetVpcChannelId() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcChannelId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ResetVpcChannelProxyHost() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcChannelProxyHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigwApiV2HttpPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

