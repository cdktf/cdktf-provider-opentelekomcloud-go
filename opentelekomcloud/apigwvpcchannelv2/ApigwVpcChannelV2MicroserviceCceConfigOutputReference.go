// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwvpcchannelv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/apigwvpcchannelv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigwVpcChannelV2MicroserviceCceConfigOutputReference interface {
	cdktf.ComplexObject
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
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
	InternalValue() *ApigwVpcChannelV2MicroserviceCceConfig
	SetInternalValue(val *ApigwVpcChannelV2MicroserviceCceConfig)
	LabelKey() *string
	SetLabelKey(val *string)
	LabelKeyInput() *string
	LabelValue() *string
	SetLabelValue(val *string)
	LabelValueInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkloadName() *string
	SetWorkloadName(val *string)
	WorkloadNameInput() *string
	WorkloadType() *string
	SetWorkloadType(val *string)
	WorkloadTypeInput() *string
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
	ResetLabelKey()
	ResetLabelValue()
	ResetWorkloadName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigwVpcChannelV2MicroserviceCceConfigOutputReference
type jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) InternalValue() *ApigwVpcChannelV2MicroserviceCceConfig {
	var returns *ApigwVpcChannelV2MicroserviceCceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) LabelKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) LabelKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) LabelValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) LabelValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) WorkloadName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) WorkloadNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) WorkloadType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) WorkloadTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadTypeInput",
		&returns,
	)
	return returns
}


func NewApigwVpcChannelV2MicroserviceCceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApigwVpcChannelV2MicroserviceCceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewApigwVpcChannelV2MicroserviceCceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwVpcChannelV2.ApigwVpcChannelV2MicroserviceCceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigwVpcChannelV2MicroserviceCceConfigOutputReference_Override(a ApigwVpcChannelV2MicroserviceCceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwVpcChannelV2.ApigwVpcChannelV2MicroserviceCceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference)SetInternalValue(val *ApigwVpcChannelV2MicroserviceCceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference)SetLabelKey(val *string) {
	if err := j.validateSetLabelKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelKey",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference)SetLabelValue(val *string) {
	if err := j.validateSetLabelValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelValue",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference)SetWorkloadName(val *string) {
	if err := j.validateSetWorkloadNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadName",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference)SetWorkloadType(val *string) {
	if err := j.validateSetWorkloadTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadType",
		val,
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) ResetLabelKey() {
	_jsii_.InvokeVoid(
		a,
		"resetLabelKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) ResetLabelValue() {
	_jsii_.InvokeVoid(
		a,
		"resetLabelValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) ResetWorkloadName() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkloadName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwVpcChannelV2MicroserviceCceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

