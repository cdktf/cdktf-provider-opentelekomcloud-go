// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicateddomainv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/wafdedicateddomainv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafDedicatedDomainV1TimeoutConfigOutputReference interface {
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
	ConnectTimeout() *float64
	SetConnectTimeout(val *float64)
	ConnectTimeoutInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WafDedicatedDomainV1TimeoutConfig
	SetInternalValue(val *WafDedicatedDomainV1TimeoutConfig)
	ReadTimeout() *float64
	SetReadTimeout(val *float64)
	ReadTimeoutInput() *float64
	SendTimeout() *float64
	SetSendTimeout(val *float64)
	SendTimeoutInput() *float64
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
	ResetConnectTimeout()
	ResetReadTimeout()
	ResetSendTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WafDedicatedDomainV1TimeoutConfigOutputReference
type jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) ConnectTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) ConnectTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) InternalValue() *WafDedicatedDomainV1TimeoutConfig {
	var returns *WafDedicatedDomainV1TimeoutConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) ReadTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) ReadTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) SendTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sendTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) SendTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sendTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafDedicatedDomainV1TimeoutConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WafDedicatedDomainV1TimeoutConfigOutputReference {
	_init_.Initialize()

	if err := validateNewWafDedicatedDomainV1TimeoutConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafDedicatedDomainV1.WafDedicatedDomainV1TimeoutConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafDedicatedDomainV1TimeoutConfigOutputReference_Override(w WafDedicatedDomainV1TimeoutConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafDedicatedDomainV1.WafDedicatedDomainV1TimeoutConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference)SetConnectTimeout(val *float64) {
	if err := j.validateSetConnectTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference)SetInternalValue(val *WafDedicatedDomainV1TimeoutConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference)SetReadTimeout(val *float64) {
	if err := j.validateSetReadTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readTimeout",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference)SetSendTimeout(val *float64) {
	if err := j.validateSetSendTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendTimeout",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		w,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) ResetReadTimeout() {
	_jsii_.InvokeVoid(
		w,
		"resetReadTimeout",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) ResetSendTimeout() {
	_jsii_.InvokeVoid(
		w,
		"resetSendTimeout",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WafDedicatedDomainV1TimeoutConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

