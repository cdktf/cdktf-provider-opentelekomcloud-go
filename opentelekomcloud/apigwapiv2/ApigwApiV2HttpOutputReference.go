// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwapiv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/apigwapiv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigwApiV2HttpOutputReference interface {
	cdktf.ComplexObject
	AuthorizerId() *string
	SetAuthorizerId(val *string)
	AuthorizerIdInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ApigwApiV2Http
	SetInternalValue(val *ApigwApiV2Http)
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
	SslEnable() interface{}
	SetSslEnable(val interface{})
	SslEnableInput() interface{}
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
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetAuthorizerId()
	ResetDescription()
	ResetRequestProtocol()
	ResetRetryCount()
	ResetSslEnable()
	ResetTimeout()
	ResetUrlDomain()
	ResetVersion()
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

// The jsii proxy struct for ApigwApiV2HttpOutputReference
type jsiiProxy_ApigwApiV2HttpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) AuthorizerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) InternalValue() *ApigwApiV2Http {
	var returns *ApigwApiV2Http
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) RequestMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) RequestMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) RequestProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) RequestProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) RequestUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) RequestUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) RetryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) RetryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) SslEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) SslEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) UrlDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) UrlDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) VpcChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcChannelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) VpcChannelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcChannelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) VpcChannelProxyHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcChannelProxyHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference) VpcChannelProxyHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcChannelProxyHostInput",
		&returns,
	)
	return returns
}


func NewApigwApiV2HttpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApigwApiV2HttpOutputReference {
	_init_.Initialize()

	if err := validateNewApigwApiV2HttpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigwApiV2HttpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2HttpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigwApiV2HttpOutputReference_Override(a ApigwApiV2HttpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2HttpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetAuthorizerId(val *string) {
	if err := j.validateSetAuthorizerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerId",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetInternalValue(val *ApigwApiV2Http) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetRequestMethod(val *string) {
	if err := j.validateSetRequestMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMethod",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetRequestProtocol(val *string) {
	if err := j.validateSetRequestProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestProtocol",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetRequestUri(val *string) {
	if err := j.validateSetRequestUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestUri",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetRetryCount(val *float64) {
	if err := j.validateSetRetryCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryCount",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetSslEnable(val interface{}) {
	if err := j.validateSetSslEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslEnable",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetUrlDomain(val *string) {
	if err := j.validateSetUrlDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlDomain",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetVpcChannelId(val *string) {
	if err := j.validateSetVpcChannelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcChannelId",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2HttpOutputReference)SetVpcChannelProxyHost(val *string) {
	if err := j.validateSetVpcChannelProxyHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcChannelProxyHost",
		val,
	)
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) ResetAuthorizerId() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) ResetRequestProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) ResetRetryCount() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) ResetSslEnable() {
	_jsii_.InvokeVoid(
		a,
		"resetSslEnable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) ResetUrlDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetUrlDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) ResetVpcChannelId() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcChannelId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) ResetVpcChannelProxyHost() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcChannelProxyHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigwApiV2HttpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

