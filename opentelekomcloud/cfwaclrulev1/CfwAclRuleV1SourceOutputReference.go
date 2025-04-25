// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwaclrulev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/cfwaclrulev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CfwAclRuleV1SourceOutputReference interface {
	cdktf.ComplexObject
	Address() *string
	SetAddress(val *string)
	AddressGroup() *[]*string
	SetAddressGroup(val *[]*string)
	AddressGroupInput() *[]*string
	AddressInput() *string
	AddressSetId() *string
	SetAddressSetId(val *string)
	AddressSetIdInput() *string
	AddressSetName() *string
	SetAddressSetName(val *string)
	AddressSetNameInput() *string
	AddressSetType() *float64
	SetAddressSetType(val *float64)
	AddressSetTypeInput() *float64
	AddressType() *float64
	SetAddressType(val *float64)
	AddressTypeInput() *float64
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
	DomainAddressName() *string
	SetDomainAddressName(val *string)
	DomainAddressNameInput() *string
	DomainSetId() *string
	SetDomainSetId(val *string)
	DomainSetIdInput() *string
	DomainSetName() *string
	SetDomainSetName(val *string)
	DomainSetNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CfwAclRuleV1Source
	SetInternalValue(val *CfwAclRuleV1Source)
	IpAddress() *[]*string
	SetIpAddress(val *[]*string)
	IpAddressInput() *[]*string
	PredefinedGroup() *[]*string
	SetPredefinedGroup(val *[]*string)
	PredefinedGroupInput() *[]*string
	RegionList() CfwAclRuleV1SourceRegionListStructList
	RegionListInput() interface{}
	RegionListJson() *string
	SetRegionListJson(val *string)
	RegionListJsonInput() *string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutRegionList(value interface{})
	ResetAddress()
	ResetAddressGroup()
	ResetAddressSetId()
	ResetAddressSetName()
	ResetAddressSetType()
	ResetAddressType()
	ResetDomainAddressName()
	ResetDomainSetId()
	ResetDomainSetName()
	ResetIpAddress()
	ResetPredefinedGroup()
	ResetRegionList()
	ResetRegionListJson()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CfwAclRuleV1SourceOutputReference
type jsiiProxy_CfwAclRuleV1SourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) AddressGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) AddressGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) AddressSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) AddressSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) AddressSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) AddressSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) AddressSetType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addressSetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) AddressSetTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addressSetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) AddressType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) AddressTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) DomainAddressName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAddressName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) DomainAddressNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAddressNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) DomainSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) DomainSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) DomainSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) DomainSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) InternalValue() *CfwAclRuleV1Source {
	var returns *CfwAclRuleV1Source
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) IpAddress() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) IpAddressInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) PredefinedGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"predefinedGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) PredefinedGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"predefinedGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) RegionList() CfwAclRuleV1SourceRegionListStructList {
	var returns CfwAclRuleV1SourceRegionListStructList
	_jsii_.Get(
		j,
		"regionList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) RegionListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) RegionListJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionListJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) RegionListJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionListJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) Type() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference) TypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCfwAclRuleV1SourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CfwAclRuleV1SourceOutputReference {
	_init_.Initialize()

	if err := validateNewCfwAclRuleV1SourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfwAclRuleV1SourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1SourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCfwAclRuleV1SourceOutputReference_Override(c CfwAclRuleV1SourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1SourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetAddress(val *string) {
	if err := j.validateSetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetAddressGroup(val *[]*string) {
	if err := j.validateSetAddressGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressGroup",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetAddressSetId(val *string) {
	if err := j.validateSetAddressSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressSetId",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetAddressSetName(val *string) {
	if err := j.validateSetAddressSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressSetName",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetAddressSetType(val *float64) {
	if err := j.validateSetAddressSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressSetType",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetAddressType(val *float64) {
	if err := j.validateSetAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressType",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetDomainAddressName(val *string) {
	if err := j.validateSetDomainAddressNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAddressName",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetDomainSetId(val *string) {
	if err := j.validateSetDomainSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSetId",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetDomainSetName(val *string) {
	if err := j.validateSetDomainSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSetName",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetInternalValue(val *CfwAclRuleV1Source) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetIpAddress(val *[]*string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetPredefinedGroup(val *[]*string) {
	if err := j.validateSetPredefinedGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predefinedGroup",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetRegionListJson(val *string) {
	if err := j.validateSetRegionListJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionListJson",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1SourceOutputReference)SetType(val *float64) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) PutRegionList(value interface{}) {
	if err := c.validatePutRegionListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRegionList",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetAddressGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetAddressGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetAddressSetId() {
	_jsii_.InvokeVoid(
		c,
		"resetAddressSetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetAddressSetName() {
	_jsii_.InvokeVoid(
		c,
		"resetAddressSetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetAddressSetType() {
	_jsii_.InvokeVoid(
		c,
		"resetAddressSetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetAddressType() {
	_jsii_.InvokeVoid(
		c,
		"resetAddressType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetDomainAddressName() {
	_jsii_.InvokeVoid(
		c,
		"resetDomainAddressName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetDomainSetId() {
	_jsii_.InvokeVoid(
		c,
		"resetDomainSetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetDomainSetName() {
	_jsii_.InvokeVoid(
		c,
		"resetDomainSetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetPredefinedGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetPredefinedGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetRegionList() {
	_jsii_.InvokeVoid(
		c,
		"resetRegionList",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ResetRegionListJson() {
	_jsii_.InvokeVoid(
		c,
		"resetRegionListJson",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1SourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

