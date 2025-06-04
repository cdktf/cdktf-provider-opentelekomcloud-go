// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwaclrulev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/cfwaclrulev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CfwAclRuleV1DestinationOutputReference interface {
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
	InternalValue() *CfwAclRuleV1Destination
	SetInternalValue(val *CfwAclRuleV1Destination)
	IpAddress() *[]*string
	SetIpAddress(val *[]*string)
	IpAddressInput() *[]*string
	PredefinedGroup() *[]*string
	SetPredefinedGroup(val *[]*string)
	PredefinedGroupInput() *[]*string
	RegionList() CfwAclRuleV1DestinationRegionListStructList
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

// The jsii proxy struct for CfwAclRuleV1DestinationOutputReference
type jsiiProxy_CfwAclRuleV1DestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) AddressGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) AddressGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) AddressSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) AddressSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) AddressSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) AddressSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) AddressSetType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addressSetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) AddressSetTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addressSetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) AddressType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) AddressTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) DomainAddressName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAddressName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) DomainAddressNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAddressNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) DomainSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) DomainSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) DomainSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) DomainSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) InternalValue() *CfwAclRuleV1Destination {
	var returns *CfwAclRuleV1Destination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) IpAddress() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) IpAddressInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) PredefinedGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"predefinedGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) PredefinedGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"predefinedGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) RegionList() CfwAclRuleV1DestinationRegionListStructList {
	var returns CfwAclRuleV1DestinationRegionListStructList
	_jsii_.Get(
		j,
		"regionList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) RegionListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) RegionListJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionListJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) RegionListJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionListJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) Type() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference) TypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCfwAclRuleV1DestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CfwAclRuleV1DestinationOutputReference {
	_init_.Initialize()

	if err := validateNewCfwAclRuleV1DestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfwAclRuleV1DestinationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1DestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCfwAclRuleV1DestinationOutputReference_Override(c CfwAclRuleV1DestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1DestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetAddress(val *string) {
	if err := j.validateSetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetAddressGroup(val *[]*string) {
	if err := j.validateSetAddressGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressGroup",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetAddressSetId(val *string) {
	if err := j.validateSetAddressSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressSetId",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetAddressSetName(val *string) {
	if err := j.validateSetAddressSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressSetName",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetAddressSetType(val *float64) {
	if err := j.validateSetAddressSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressSetType",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetAddressType(val *float64) {
	if err := j.validateSetAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressType",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetDomainAddressName(val *string) {
	if err := j.validateSetDomainAddressNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAddressName",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetDomainSetId(val *string) {
	if err := j.validateSetDomainSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSetId",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetDomainSetName(val *string) {
	if err := j.validateSetDomainSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSetName",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetInternalValue(val *CfwAclRuleV1Destination) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetIpAddress(val *[]*string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetPredefinedGroup(val *[]*string) {
	if err := j.validateSetPredefinedGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predefinedGroup",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetRegionListJson(val *string) {
	if err := j.validateSetRegionListJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionListJson",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1DestinationOutputReference)SetType(val *float64) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) PutRegionList(value interface{}) {
	if err := c.validatePutRegionListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRegionList",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetAddressGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetAddressGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetAddressSetId() {
	_jsii_.InvokeVoid(
		c,
		"resetAddressSetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetAddressSetName() {
	_jsii_.InvokeVoid(
		c,
		"resetAddressSetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetAddressSetType() {
	_jsii_.InvokeVoid(
		c,
		"resetAddressSetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetAddressType() {
	_jsii_.InvokeVoid(
		c,
		"resetAddressType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetDomainAddressName() {
	_jsii_.InvokeVoid(
		c,
		"resetDomainAddressName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetDomainSetId() {
	_jsii_.InvokeVoid(
		c,
		"resetDomainSetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetDomainSetName() {
	_jsii_.InvokeVoid(
		c,
		"resetDomainSetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetPredefinedGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetPredefinedGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetRegionList() {
	_jsii_.InvokeVoid(
		c,
		"resetRegionList",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ResetRegionListJson() {
	_jsii_.InvokeVoid(
		c,
		"resetRegionListJson",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CfwAclRuleV1DestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

