// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisevpngatewayv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/enterprisevpngatewayv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnterpriseVpnGatewayV5Eip1OutputReference interface {
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
	InternalValue() *EnterpriseVpnGatewayV5Eip1
	SetInternalValue(val *EnterpriseVpnGatewayV5Eip1)
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetBandwidthName()
	ResetBandwidthSize()
	ResetChargeMode()
	ResetId()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EnterpriseVpnGatewayV5Eip1OutputReference
type jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) BandwidthId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) BandwidthName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) BandwidthNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) BandwidthSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) BandwidthSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) ChargeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) ChargeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) InternalValue() *EnterpriseVpnGatewayV5Eip1 {
	var returns *EnterpriseVpnGatewayV5Eip1
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) IpVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewEnterpriseVpnGatewayV5Eip1OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EnterpriseVpnGatewayV5Eip1OutputReference {
	_init_.Initialize()

	if err := validateNewEnterpriseVpnGatewayV5Eip1OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnGatewayV5.EnterpriseVpnGatewayV5Eip1OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEnterpriseVpnGatewayV5Eip1OutputReference_Override(e EnterpriseVpnGatewayV5Eip1OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnGatewayV5.EnterpriseVpnGatewayV5Eip1OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference)SetBandwidthName(val *string) {
	if err := j.validateSetBandwidthNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthName",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference)SetBandwidthSize(val *float64) {
	if err := j.validateSetBandwidthSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthSize",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference)SetChargeMode(val *string) {
	if err := j.validateSetChargeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chargeMode",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference)SetInternalValue(val *EnterpriseVpnGatewayV5Eip1) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) ResetBandwidthName() {
	_jsii_.InvokeVoid(
		e,
		"resetBandwidthName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) ResetBandwidthSize() {
	_jsii_.InvokeVoid(
		e,
		"resetBandwidthSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) ResetChargeMode() {
	_jsii_.InvokeVoid(
		e,
		"resetChargeMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) ResetType() {
	_jsii_.InvokeVoid(
		e,
		"resetType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5Eip1OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

