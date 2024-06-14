// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwvpcchannelv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/apigwvpcchannelv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigwVpcChannelV2HealthCheckOutputReference interface {
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
	EnableClientSsl() interface{}
	SetEnableClientSsl(val interface{})
	EnableClientSslInput() interface{}
	// Experimental.
	Fqn() *string
	HttpCodes() *string
	SetHttpCodes(val *string)
	HttpCodesInput() *string
	InternalValue() *ApigwVpcChannelV2HealthCheck
	SetInternalValue(val *ApigwVpcChannelV2HealthCheck)
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	Method() *string
	SetMethod(val *string)
	MethodInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	Status() *float64
	SetStatus(val *float64)
	StatusInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThresholdAbnormal() *float64
	SetThresholdAbnormal(val *float64)
	ThresholdAbnormalInput() *float64
	ThresholdNormal() *float64
	SetThresholdNormal(val *float64)
	ThresholdNormalInput() *float64
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
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
	ResetEnableClientSsl()
	ResetHttpCodes()
	ResetMethod()
	ResetPath()
	ResetPort()
	ResetStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigwVpcChannelV2HealthCheckOutputReference
type jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) EnableClientSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableClientSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) EnableClientSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableClientSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) HttpCodes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) HttpCodesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) InternalValue() *ApigwVpcChannelV2HealthCheck {
	var returns *ApigwVpcChannelV2HealthCheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) Status() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) StatusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ThresholdAbnormal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdAbnormal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ThresholdAbnormalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdAbnormalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ThresholdNormal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdNormal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ThresholdNormalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdNormalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


func NewApigwVpcChannelV2HealthCheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApigwVpcChannelV2HealthCheckOutputReference {
	_init_.Initialize()

	if err := validateNewApigwVpcChannelV2HealthCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwVpcChannelV2.ApigwVpcChannelV2HealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigwVpcChannelV2HealthCheckOutputReference_Override(a ApigwVpcChannelV2HealthCheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwVpcChannelV2.ApigwVpcChannelV2HealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetEnableClientSsl(val interface{}) {
	if err := j.validateSetEnableClientSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableClientSsl",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetHttpCodes(val *string) {
	if err := j.validateSetHttpCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCodes",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetInternalValue(val *ApigwVpcChannelV2HealthCheck) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetStatus(val *float64) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetThresholdAbnormal(val *float64) {
	if err := j.validateSetThresholdAbnormalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdAbnormal",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetThresholdNormal(val *float64) {
	if err := j.validateSetThresholdNormalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdNormal",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ResetEnableClientSsl() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableClientSsl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ResetHttpCodes() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpCodes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigwVpcChannelV2HealthCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

