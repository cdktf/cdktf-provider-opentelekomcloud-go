// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdomainv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/wafdomainv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafDomainV1ServerOutputReference interface {
	cdktf.ComplexObject
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
	BackProtocol() *string
	SetBackProtocol(val *string)
	BackProtocolInput() *string
	ClientProtocol() *string
	SetClientProtocol(val *string)
	ClientProtocolInput() *string
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
	FrontProtocol() *string
	SetFrontProtocol(val *string)
	FrontProtocolInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Port() *string
	SetPort(val *string)
	PortInput() *string
	ServerProtocol() *string
	SetServerProtocol(val *string)
	ServerProtocolInput() *string
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
	ResetBackProtocol()
	ResetClientProtocol()
	ResetFrontProtocol()
	ResetServerProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WafDomainV1ServerOutputReference
type jsiiProxy_WafDomainV1ServerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) BackProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) BackProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) ClientProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) ClientProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) FrontProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) FrontProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) PortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) ServerProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) ServerProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafDomainV1ServerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WafDomainV1ServerOutputReference {
	_init_.Initialize()

	if err := validateNewWafDomainV1ServerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafDomainV1ServerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafDomainV1.WafDomainV1ServerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafDomainV1ServerOutputReference_Override(w WafDomainV1ServerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafDomainV1.WafDomainV1ServerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference)SetAddress(val *string) {
	if err := j.validateSetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference)SetBackProtocol(val *string) {
	if err := j.validateSetBackProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backProtocol",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference)SetClientProtocol(val *string) {
	if err := j.validateSetClientProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientProtocol",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference)SetFrontProtocol(val *string) {
	if err := j.validateSetFrontProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontProtocol",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference)SetPort(val *string) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference)SetServerProtocol(val *string) {
	if err := j.validateSetServerProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverProtocol",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1ServerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) ResetBackProtocol() {
	_jsii_.InvokeVoid(
		w,
		"resetBackProtocol",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) ResetClientProtocol() {
	_jsii_.InvokeVoid(
		w,
		"resetClientProtocol",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) ResetFrontProtocol() {
	_jsii_.InvokeVoid(
		w,
		"resetFrontProtocol",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) ResetServerProtocol() {
	_jsii_.InvokeVoid(
		w,
		"resetServerProtocol",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1ServerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

