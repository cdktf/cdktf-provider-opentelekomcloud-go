// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisevpnconnectionv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/enterprisevpnconnectionv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnterpriseVpnConnectionV5IpsecpolicyOutputReference interface {
	cdktf.ComplexObject
	AuthenticationAlgorithm() *string
	SetAuthenticationAlgorithm(val *string)
	AuthenticationAlgorithmInput() *string
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
	EncapsulationMode() *string
	SetEncapsulationMode(val *string)
	EncapsulationModeInput() *string
	EncryptionAlgorithm() *string
	SetEncryptionAlgorithm(val *string)
	EncryptionAlgorithmInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *EnterpriseVpnConnectionV5Ipsecpolicy
	SetInternalValue(val *EnterpriseVpnConnectionV5Ipsecpolicy)
	LifetimeSeconds() *float64
	SetLifetimeSeconds(val *float64)
	LifetimeSecondsInput() *float64
	Pfs() *string
	SetPfs(val *string)
	PfsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransformProtocol() *string
	SetTransformProtocol(val *string)
	TransformProtocolInput() *string
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
	ResetAuthenticationAlgorithm()
	ResetEncapsulationMode()
	ResetEncryptionAlgorithm()
	ResetLifetimeSeconds()
	ResetPfs()
	ResetTransformProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EnterpriseVpnConnectionV5IpsecpolicyOutputReference
type jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) AuthenticationAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) AuthenticationAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) EncapsulationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encapsulationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) EncapsulationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encapsulationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) EncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) EncryptionAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) InternalValue() *EnterpriseVpnConnectionV5Ipsecpolicy {
	var returns *EnterpriseVpnConnectionV5Ipsecpolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) Pfs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) PfsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) TransformProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) TransformProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformProtocolInput",
		&returns,
	)
	return returns
}


func NewEnterpriseVpnConnectionV5IpsecpolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EnterpriseVpnConnectionV5IpsecpolicyOutputReference {
	_init_.Initialize()

	if err := validateNewEnterpriseVpnConnectionV5IpsecpolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnConnectionV5.EnterpriseVpnConnectionV5IpsecpolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEnterpriseVpnConnectionV5IpsecpolicyOutputReference_Override(e EnterpriseVpnConnectionV5IpsecpolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnConnectionV5.EnterpriseVpnConnectionV5IpsecpolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference)SetAuthenticationAlgorithm(val *string) {
	if err := j.validateSetAuthenticationAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationAlgorithm",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference)SetEncapsulationMode(val *string) {
	if err := j.validateSetEncapsulationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encapsulationMode",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference)SetEncryptionAlgorithm(val *string) {
	if err := j.validateSetEncryptionAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAlgorithm",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference)SetInternalValue(val *EnterpriseVpnConnectionV5Ipsecpolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference)SetLifetimeSeconds(val *float64) {
	if err := j.validateSetLifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference)SetPfs(val *string) {
	if err := j.validateSetPfsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pfs",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference)SetTransformProtocol(val *string) {
	if err := j.validateSetTransformProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformProtocol",
		val,
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) ResetAuthenticationAlgorithm() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthenticationAlgorithm",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) ResetEncapsulationMode() {
	_jsii_.InvokeVoid(
		e,
		"resetEncapsulationMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) ResetEncryptionAlgorithm() {
	_jsii_.InvokeVoid(
		e,
		"resetEncryptionAlgorithm",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) ResetLifetimeSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetLifetimeSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) ResetPfs() {
	_jsii_.InvokeVoid(
		e,
		"resetPfs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) ResetTransformProtocol() {
	_jsii_.InvokeVoid(
		e,
		"resetTransformProtocol",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IpsecpolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

