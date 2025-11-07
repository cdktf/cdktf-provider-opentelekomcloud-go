// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisevpngatewayv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/enterprisevpngatewayv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnterpriseVpnGatewayV5Eip2OutputReference interface {
	cdktf.ComplexObject
	BandwidthId() *string
	BandwidthName() *string
	SetBandwidthName(val *string)
	BandwidthNameInput() *string
	BandwidthSize() *float64
	SetBandwidthSize(val *float64)
	BandwidthSizeInput() *float64
	ChargeMode() *string
	SetChargeMode(val *string)
	ChargeModeInput() *string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() *EnterpriseVpnGatewayV5Eip2
	SetInternalValue(val *EnterpriseVpnGatewayV5Eip2)
	IpAddress() *string
	IpVersion() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetBandwidthName()
	ResetBandwidthSize()
	ResetChargeMode()
	ResetId()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EnterpriseVpnGatewayV5Eip2OutputReference
type jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) BandwidthId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) BandwidthName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) BandwidthNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) BandwidthSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) BandwidthSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) ChargeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) ChargeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) InternalValue() *EnterpriseVpnGatewayV5Eip2 {
	var returns *EnterpriseVpnGatewayV5Eip2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) IpVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewEnterpriseVpnGatewayV5Eip2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EnterpriseVpnGatewayV5Eip2OutputReference {
	_init_.Initialize()

	if err := validateNewEnterpriseVpnGatewayV5Eip2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnGatewayV5.EnterpriseVpnGatewayV5Eip2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEnterpriseVpnGatewayV5Eip2OutputReference_Override(e EnterpriseVpnGatewayV5Eip2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnGatewayV5.EnterpriseVpnGatewayV5Eip2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference)SetBandwidthName(val *string) {
	if err := j.validateSetBandwidthNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthName",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference)SetBandwidthSize(val *float64) {
	if err := j.validateSetBandwidthSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthSize",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference)SetChargeMode(val *string) {
	if err := j.validateSetChargeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chargeMode",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference)SetInternalValue(val *EnterpriseVpnGatewayV5Eip2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) ResetBandwidthName() {
	_jsii_.InvokeVoid(
		e,
		"resetBandwidthName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) ResetBandwidthSize() {
	_jsii_.InvokeVoid(
		e,
		"resetBandwidthSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) ResetChargeMode() {
	_jsii_.InvokeVoid(
		e,
		"resetChargeMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) ResetType() {
	_jsii_.InvokeVoid(
		e,
		"resetType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

