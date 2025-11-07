// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwaclrulev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/cfwaclrulev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CfwAclRuleV1ServiceOutputReference interface {
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
	CustomService() CfwAclRuleV1ServiceCustomServiceList
	CustomServiceInput() interface{}
	DestPort() *string
	SetDestPort(val *string)
	DestPortInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CfwAclRuleV1Service
	SetInternalValue(val *CfwAclRuleV1Service)
	PredefinedGroup() *[]*string
	SetPredefinedGroup(val *[]*string)
	PredefinedGroupInput() *[]*string
	Protocol() *float64
	SetProtocol(val *float64)
	ProtocolInput() *float64
	Protocols() *[]*float64
	SetProtocols(val *[]*float64)
	ProtocolsInput() *[]*float64
	ServiceGroup() *[]*string
	SetServiceGroup(val *[]*string)
	ServiceGroupInput() *[]*string
	ServiceGroupNames() CfwAclRuleV1ServiceServiceGroupNamesList
	ServiceGroupNamesInput() interface{}
	ServiceSetId() *string
	SetServiceSetId(val *string)
	ServiceSetIdInput() *string
	ServiceSetName() *string
	SetServiceSetName(val *string)
	ServiceSetNameInput() *string
	ServiceSetType() *float64
	SetServiceSetType(val *float64)
	ServiceSetTypeInput() *float64
	SourcePort() *string
	SetSourcePort(val *string)
	SourcePortInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *float64
	SetType(val *float64)
	TypeInput() *float64
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
	PutCustomService(value interface{})
	PutServiceGroupNames(value interface{})
	ResetCustomService()
	ResetDestPort()
	ResetPredefinedGroup()
	ResetProtocol()
	ResetProtocols()
	ResetServiceGroup()
	ResetServiceGroupNames()
	ResetServiceSetId()
	ResetServiceSetName()
	ResetServiceSetType()
	ResetSourcePort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CfwAclRuleV1ServiceOutputReference
type jsiiProxy_CfwAclRuleV1ServiceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) CustomService() CfwAclRuleV1ServiceCustomServiceList {
	var returns CfwAclRuleV1ServiceCustomServiceList
	_jsii_.Get(
		j,
		"customService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) CustomServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) DestPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) DestPortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) InternalValue() *CfwAclRuleV1Service {
	var returns *CfwAclRuleV1Service
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) PredefinedGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"predefinedGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) PredefinedGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"predefinedGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) Protocol() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ProtocolInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) Protocols() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ProtocolsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ServiceGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ServiceGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ServiceGroupNames() CfwAclRuleV1ServiceServiceGroupNamesList {
	var returns CfwAclRuleV1ServiceServiceGroupNamesList
	_jsii_.Get(
		j,
		"serviceGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ServiceGroupNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceGroupNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ServiceSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ServiceSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ServiceSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ServiceSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ServiceSetType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serviceSetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ServiceSetTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serviceSetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) SourcePort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) SourcePortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) Type() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference) TypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCfwAclRuleV1ServiceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CfwAclRuleV1ServiceOutputReference {
	_init_.Initialize()

	if err := validateNewCfwAclRuleV1ServiceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfwAclRuleV1ServiceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1ServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCfwAclRuleV1ServiceOutputReference_Override(c CfwAclRuleV1ServiceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1ServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetDestPort(val *string) {
	if err := j.validateSetDestPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destPort",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetInternalValue(val *CfwAclRuleV1Service) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetPredefinedGroup(val *[]*string) {
	if err := j.validateSetPredefinedGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predefinedGroup",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetProtocol(val *float64) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetProtocols(val *[]*float64) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetServiceGroup(val *[]*string) {
	if err := j.validateSetServiceGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceGroup",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetServiceSetId(val *string) {
	if err := j.validateSetServiceSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceSetId",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetServiceSetName(val *string) {
	if err := j.validateSetServiceSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceSetName",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetServiceSetType(val *float64) {
	if err := j.validateSetServiceSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceSetType",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetSourcePort(val *string) {
	if err := j.validateSetSourcePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePort",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1ServiceOutputReference)SetType(val *float64) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) PutCustomService(value interface{}) {
	if err := c.validatePutCustomServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomService",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) PutServiceGroupNames(value interface{}) {
	if err := c.validatePutServiceGroupNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceGroupNames",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ResetCustomService() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomService",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ResetDestPort() {
	_jsii_.InvokeVoid(
		c,
		"resetDestPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ResetPredefinedGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetPredefinedGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ResetProtocols() {
	_jsii_.InvokeVoid(
		c,
		"resetProtocols",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ResetServiceGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ResetServiceGroupNames() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceGroupNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ResetServiceSetId() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceSetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ResetServiceSetName() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceSetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ResetServiceSetType() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceSetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ResetSourcePort() {
	_jsii_.InvokeVoid(
		c,
		"resetSourcePort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1ServiceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

