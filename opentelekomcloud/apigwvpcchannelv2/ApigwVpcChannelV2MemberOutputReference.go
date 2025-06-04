// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwvpcchannelv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/apigwvpcchannelv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigwVpcChannelV2MemberOutputReference interface {
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
	// Experimental.
	Fqn() *string
	GroupName() *string
	SetGroupName(val *string)
	GroupNameInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsBackup() interface{}
	SetIsBackup(val interface{})
	IsBackupInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Status() *float64
	SetStatus(val *float64)
	StatusInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
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
	ResetGroupName()
	ResetHost()
	ResetId()
	ResetIsBackup()
	ResetName()
	ResetPort()
	ResetStatus()
	ResetWeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigwVpcChannelV2MemberOutputReference
type jsiiProxy_ApigwVpcChannelV2MemberOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) GroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) IsBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) IsBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) Status() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) StatusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewApigwVpcChannelV2MemberOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApigwVpcChannelV2MemberOutputReference {
	_init_.Initialize()

	if err := validateNewApigwVpcChannelV2MemberOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigwVpcChannelV2MemberOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwVpcChannelV2.ApigwVpcChannelV2MemberOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApigwVpcChannelV2MemberOutputReference_Override(a ApigwVpcChannelV2MemberOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwVpcChannelV2.ApigwVpcChannelV2MemberOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetGroupName(val *string) {
	if err := j.validateSetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupName",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetIsBackup(val interface{}) {
	if err := j.validateSetIsBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isBackup",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetStatus(val *float64) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigwVpcChannelV2MemberOutputReference)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) ResetGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		a,
		"resetHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) ResetIsBackup() {
	_jsii_.InvokeVoid(
		a,
		"resetIsBackup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		a,
		"resetWeight",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigwVpcChannelV2MemberOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

