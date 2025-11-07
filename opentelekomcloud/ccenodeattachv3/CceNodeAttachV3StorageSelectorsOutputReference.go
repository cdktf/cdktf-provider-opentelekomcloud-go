// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodeattachv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/ccenodeattachv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CceNodeAttachV3StorageSelectorsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MatchLabelCount() *string
	SetMatchLabelCount(val *string)
	MatchLabelCountInput() *string
	MatchLabelMetadataCmkid() *string
	SetMatchLabelMetadataCmkid(val *string)
	MatchLabelMetadataCmkidInput() *string
	MatchLabelMetadataEncrypted() *string
	SetMatchLabelMetadataEncrypted(val *string)
	MatchLabelMetadataEncryptedInput() *string
	MatchLabelSize() *string
	SetMatchLabelSize(val *string)
	MatchLabelSizeInput() *string
	MatchLabelVolumeType() *string
	SetMatchLabelVolumeType(val *string)
	MatchLabelVolumeTypeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetMatchLabelCount()
	ResetMatchLabelMetadataCmkid()
	ResetMatchLabelMetadataEncrypted()
	ResetMatchLabelSize()
	ResetMatchLabelVolumeType()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CceNodeAttachV3StorageSelectorsOutputReference
type jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) MatchLabelCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchLabelCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) MatchLabelCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchLabelCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) MatchLabelMetadataCmkid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchLabelMetadataCmkid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) MatchLabelMetadataCmkidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchLabelMetadataCmkidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) MatchLabelMetadataEncrypted() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchLabelMetadataEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) MatchLabelMetadataEncryptedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchLabelMetadataEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) MatchLabelSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchLabelSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) MatchLabelSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchLabelSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) MatchLabelVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchLabelVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) MatchLabelVolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchLabelVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCceNodeAttachV3StorageSelectorsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CceNodeAttachV3StorageSelectorsOutputReference {
	_init_.Initialize()

	if err := validateNewCceNodeAttachV3StorageSelectorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cceNodeAttachV3.CceNodeAttachV3StorageSelectorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCceNodeAttachV3StorageSelectorsOutputReference_Override(c CceNodeAttachV3StorageSelectorsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cceNodeAttachV3.CceNodeAttachV3StorageSelectorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference)SetMatchLabelCount(val *string) {
	if err := j.validateSetMatchLabelCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchLabelCount",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference)SetMatchLabelMetadataCmkid(val *string) {
	if err := j.validateSetMatchLabelMetadataCmkidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchLabelMetadataCmkid",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference)SetMatchLabelMetadataEncrypted(val *string) {
	if err := j.validateSetMatchLabelMetadataEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchLabelMetadataEncrypted",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference)SetMatchLabelSize(val *string) {
	if err := j.validateSetMatchLabelSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchLabelSize",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference)SetMatchLabelVolumeType(val *string) {
	if err := j.validateSetMatchLabelVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchLabelVolumeType",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) ResetMatchLabelCount() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchLabelCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) ResetMatchLabelMetadataCmkid() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchLabelMetadataCmkid",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) ResetMatchLabelMetadataEncrypted() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchLabelMetadataEncrypted",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) ResetMatchLabelSize() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchLabelSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) ResetMatchLabelVolumeType() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchLabelVolumeType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CceNodeAttachV3StorageSelectorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

