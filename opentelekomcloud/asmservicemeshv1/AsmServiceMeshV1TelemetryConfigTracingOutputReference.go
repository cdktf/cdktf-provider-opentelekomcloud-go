// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asmservicemeshv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/asmservicemeshv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AsmServiceMeshV1TelemetryConfigTracingOutputReference interface {
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
	DefaultProviders() *[]*string
	SetDefaultProviders(val *[]*string)
	DefaultProvidersInput() *[]*string
	ExtensionProviders() AsmServiceMeshV1TelemetryConfigTracingExtensionProvidersList
	ExtensionProvidersInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *AsmServiceMeshV1TelemetryConfigTracing
	SetInternalValue(val *AsmServiceMeshV1TelemetryConfigTracing)
	RandomSamplingPercentage() *float64
	SetRandomSamplingPercentage(val *float64)
	RandomSamplingPercentageInput() *float64
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
	PutExtensionProviders(value interface{})
	ResetDefaultProviders()
	ResetExtensionProviders()
	ResetRandomSamplingPercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AsmServiceMeshV1TelemetryConfigTracingOutputReference
type jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) DefaultProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) DefaultProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) ExtensionProviders() AsmServiceMeshV1TelemetryConfigTracingExtensionProvidersList {
	var returns AsmServiceMeshV1TelemetryConfigTracingExtensionProvidersList
	_jsii_.Get(
		j,
		"extensionProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) ExtensionProvidersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) InternalValue() *AsmServiceMeshV1TelemetryConfigTracing {
	var returns *AsmServiceMeshV1TelemetryConfigTracing
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) RandomSamplingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomSamplingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) RandomSamplingPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomSamplingPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAsmServiceMeshV1TelemetryConfigTracingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AsmServiceMeshV1TelemetryConfigTracingOutputReference {
	_init_.Initialize()

	if err := validateNewAsmServiceMeshV1TelemetryConfigTracingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.asmServiceMeshV1.AsmServiceMeshV1TelemetryConfigTracingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAsmServiceMeshV1TelemetryConfigTracingOutputReference_Override(a AsmServiceMeshV1TelemetryConfigTracingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.asmServiceMeshV1.AsmServiceMeshV1TelemetryConfigTracingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference)SetDefaultProviders(val *[]*string) {
	if err := j.validateSetDefaultProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultProviders",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference)SetInternalValue(val *AsmServiceMeshV1TelemetryConfigTracing) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference)SetRandomSamplingPercentage(val *float64) {
	if err := j.validateSetRandomSamplingPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"randomSamplingPercentage",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) PutExtensionProviders(value interface{}) {
	if err := a.validatePutExtensionProvidersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExtensionProviders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) ResetDefaultProviders() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultProviders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) ResetExtensionProviders() {
	_jsii_.InvokeVoid(
		a,
		"resetExtensionProviders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) ResetRandomSamplingPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetRandomSamplingPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AsmServiceMeshV1TelemetryConfigTracingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

