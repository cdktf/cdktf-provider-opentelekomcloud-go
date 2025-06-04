// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisevpnconnectionv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/enterprisevpnconnectionv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnterpriseVpnConnectionV5IkepolicyOutputReference interface {
	cdktf.ComplexObject
	AuthenticationAlgorithm() *string
	SetAuthenticationAlgorithm(val *string)
	AuthenticationAlgorithmInput() *string
	AuthenticationMethod() *string
	SetAuthenticationMethod(val *string)
	AuthenticationMethodInput() *string
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
	DhGroup() *string
	SetDhGroup(val *string)
	DhGroupInput() *string
	Dpd() EnterpriseVpnConnectionV5IkepolicyDpdOutputReference
	DpdInput() *EnterpriseVpnConnectionV5IkepolicyDpd
	EncryptionAlgorithm() *string
	SetEncryptionAlgorithm(val *string)
	EncryptionAlgorithmInput() *string
	// Experimental.
	Fqn() *string
	IkeVersion() *string
	SetIkeVersion(val *string)
	IkeVersionInput() *string
	InternalValue() *EnterpriseVpnConnectionV5Ikepolicy
	SetInternalValue(val *EnterpriseVpnConnectionV5Ikepolicy)
	LifetimeSeconds() *float64
	SetLifetimeSeconds(val *float64)
	LifetimeSecondsInput() *float64
	LocalId() *string
	SetLocalId(val *string)
	LocalIdInput() *string
	LocalIdType() *string
	SetLocalIdType(val *string)
	LocalIdTypeInput() *string
	PeerId() *string
	SetPeerId(val *string)
	PeerIdInput() *string
	PeerIdType() *string
	SetPeerIdType(val *string)
	PeerIdTypeInput() *string
	PhaseOneNegotiationMode() *string
	SetPhaseOneNegotiationMode(val *string)
	PhaseOneNegotiationModeInput() *string
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
	PutDpd(value *EnterpriseVpnConnectionV5IkepolicyDpd)
	ResetAuthenticationAlgorithm()
	ResetAuthenticationMethod()
	ResetDhGroup()
	ResetDpd()
	ResetEncryptionAlgorithm()
	ResetIkeVersion()
	ResetLifetimeSeconds()
	ResetLocalId()
	ResetLocalIdType()
	ResetPeerId()
	ResetPeerIdType()
	ResetPhaseOneNegotiationMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EnterpriseVpnConnectionV5IkepolicyOutputReference
type jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) AuthenticationAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) AuthenticationAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) AuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) DhGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) DhGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) Dpd() EnterpriseVpnConnectionV5IkepolicyDpdOutputReference {
	var returns EnterpriseVpnConnectionV5IkepolicyDpdOutputReference
	_jsii_.Get(
		j,
		"dpd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) DpdInput() *EnterpriseVpnConnectionV5IkepolicyDpd {
	var returns *EnterpriseVpnConnectionV5IkepolicyDpd
	_jsii_.Get(
		j,
		"dpdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) EncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) EncryptionAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) IkeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) IkeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) InternalValue() *EnterpriseVpnConnectionV5Ikepolicy {
	var returns *EnterpriseVpnConnectionV5Ikepolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) LocalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) LocalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) LocalIdType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIdType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) LocalIdTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIdTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) PeerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) PeerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) PeerIdType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIdType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) PeerIdTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIdTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) PhaseOneNegotiationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phaseOneNegotiationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) PhaseOneNegotiationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phaseOneNegotiationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEnterpriseVpnConnectionV5IkepolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EnterpriseVpnConnectionV5IkepolicyOutputReference {
	_init_.Initialize()

	if err := validateNewEnterpriseVpnConnectionV5IkepolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnConnectionV5.EnterpriseVpnConnectionV5IkepolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEnterpriseVpnConnectionV5IkepolicyOutputReference_Override(e EnterpriseVpnConnectionV5IkepolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnConnectionV5.EnterpriseVpnConnectionV5IkepolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetAuthenticationAlgorithm(val *string) {
	if err := j.validateSetAuthenticationAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationAlgorithm",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetAuthenticationMethod(val *string) {
	if err := j.validateSetAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMethod",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetDhGroup(val *string) {
	if err := j.validateSetDhGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhGroup",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetEncryptionAlgorithm(val *string) {
	if err := j.validateSetEncryptionAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAlgorithm",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetIkeVersion(val *string) {
	if err := j.validateSetIkeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeVersion",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetInternalValue(val *EnterpriseVpnConnectionV5Ikepolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetLifetimeSeconds(val *float64) {
	if err := j.validateSetLifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetLocalId(val *string) {
	if err := j.validateSetLocalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localId",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetLocalIdType(val *string) {
	if err := j.validateSetLocalIdTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localIdType",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetPeerId(val *string) {
	if err := j.validateSetPeerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerId",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetPeerIdType(val *string) {
	if err := j.validateSetPeerIdTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIdType",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetPhaseOneNegotiationMode(val *string) {
	if err := j.validateSetPhaseOneNegotiationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phaseOneNegotiationMode",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) PutDpd(value *EnterpriseVpnConnectionV5IkepolicyDpd) {
	if err := e.validatePutDpdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDpd",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ResetAuthenticationAlgorithm() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthenticationAlgorithm",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ResetAuthenticationMethod() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthenticationMethod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ResetDhGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetDhGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ResetDpd() {
	_jsii_.InvokeVoid(
		e,
		"resetDpd",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ResetEncryptionAlgorithm() {
	_jsii_.InvokeVoid(
		e,
		"resetEncryptionAlgorithm",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ResetIkeVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetIkeVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ResetLifetimeSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetLifetimeSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ResetLocalId() {
	_jsii_.InvokeVoid(
		e,
		"resetLocalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ResetLocalIdType() {
	_jsii_.InvokeVoid(
		e,
		"resetLocalIdType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ResetPeerId() {
	_jsii_.InvokeVoid(
		e,
		"resetPeerId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ResetPeerIdType() {
	_jsii_.InvokeVoid(
		e,
		"resetPeerIdType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ResetPhaseOneNegotiationMode() {
	_jsii_.InvokeVoid(
		e,
		"resetPhaseOneNegotiationMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EnterpriseVpnConnectionV5IkepolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

