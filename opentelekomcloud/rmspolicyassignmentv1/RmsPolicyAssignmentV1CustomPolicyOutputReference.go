// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rmspolicyassignmentv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/rmspolicyassignmentv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RmsPolicyAssignmentV1CustomPolicyOutputReference interface {
	cdktf.ComplexObject
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
	AuthValue() *map[string]*string
	SetAuthValue(val *map[string]*string)
	AuthValueInput() *map[string]*string
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
	FunctionUrn() *string
	SetFunctionUrn(val *string)
	FunctionUrnInput() *string
	InternalValue() *RmsPolicyAssignmentV1CustomPolicy
	SetInternalValue(val *RmsPolicyAssignmentV1CustomPolicy)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetAuthValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RmsPolicyAssignmentV1CustomPolicyOutputReference
type jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) AuthValue() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"authValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) AuthValueInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"authValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) FunctionUrn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionUrn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) FunctionUrnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionUrnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) InternalValue() *RmsPolicyAssignmentV1CustomPolicy {
	var returns *RmsPolicyAssignmentV1CustomPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRmsPolicyAssignmentV1CustomPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RmsPolicyAssignmentV1CustomPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewRmsPolicyAssignmentV1CustomPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.rmsPolicyAssignmentV1.RmsPolicyAssignmentV1CustomPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRmsPolicyAssignmentV1CustomPolicyOutputReference_Override(r RmsPolicyAssignmentV1CustomPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.rmsPolicyAssignmentV1.RmsPolicyAssignmentV1CustomPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference)SetAuthValue(val *map[string]*string) {
	if err := j.validateSetAuthValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authValue",
		val,
	)
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference)SetFunctionUrn(val *string) {
	if err := j.validateSetFunctionUrnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionUrn",
		val,
	)
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference)SetInternalValue(val *RmsPolicyAssignmentV1CustomPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) ResetAuthValue() {
	_jsii_.InvokeVoid(
		r,
		"resetAuthValue",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RmsPolicyAssignmentV1CustomPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

