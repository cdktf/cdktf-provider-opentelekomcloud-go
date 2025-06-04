// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/identityprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityProviderAccessConfigOutputReference interface {
	cdktf.ComplexObject
	AccessType() *string
	SetAccessType(val *string)
	AccessTypeInput() *string
	AuthorizationEndpoint() *string
	SetAuthorizationEndpoint(val *string)
	AuthorizationEndpointInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
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
	InternalValue() *IdentityProviderAccessConfig
	SetInternalValue(val *IdentityProviderAccessConfig)
	ProviderUrl() *string
	SetProviderUrl(val *string)
	ProviderUrlInput() *string
	ResponseMode() *string
	SetResponseMode(val *string)
	ResponseModeInput() *string
	ResponseType() *string
	SetResponseType(val *string)
	ResponseTypeInput() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	SigningKey() *string
	SetSigningKey(val *string)
	SigningKeyInput() *string
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
	ResetAuthorizationEndpoint()
	ResetResponseMode()
	ResetResponseType()
	ResetScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IdentityProviderAccessConfigOutputReference
type jsiiProxy_IdentityProviderAccessConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) AccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) AccessTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) AuthorizationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) AuthorizationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) InternalValue() *IdentityProviderAccessConfig {
	var returns *IdentityProviderAccessConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) ProviderUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) ProviderUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) ResponseMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) ResponseModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) ResponseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) ResponseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) SigningKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) SigningKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIdentityProviderAccessConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IdentityProviderAccessConfigOutputReference {
	_init_.Initialize()

	if err := validateNewIdentityProviderAccessConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityProviderAccessConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.identityProvider.IdentityProviderAccessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIdentityProviderAccessConfigOutputReference_Override(i IdentityProviderAccessConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.identityProvider.IdentityProviderAccessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetAccessType(val *string) {
	if err := j.validateSetAccessTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessType",
		val,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetAuthorizationEndpoint(val *string) {
	if err := j.validateSetAuthorizationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationEndpoint",
		val,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetInternalValue(val *IdentityProviderAccessConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetProviderUrl(val *string) {
	if err := j.validateSetProviderUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerUrl",
		val,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetResponseMode(val *string) {
	if err := j.validateSetResponseModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseMode",
		val,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetResponseType(val *string) {
	if err := j.validateSetResponseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseType",
		val,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetSigningKey(val *string) {
	if err := j.validateSetSigningKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signingKey",
		val,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IdentityProviderAccessConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) ResetAuthorizationEndpoint() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthorizationEndpoint",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) ResetResponseMode() {
	_jsii_.InvokeVoid(
		i,
		"resetResponseMode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) ResetResponseType() {
	_jsii_.InvokeVoid(
		i,
		"resetResponseType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) ResetScopes() {
	_jsii_.InvokeVoid(
		i,
		"resetScopes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProviderAccessConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

