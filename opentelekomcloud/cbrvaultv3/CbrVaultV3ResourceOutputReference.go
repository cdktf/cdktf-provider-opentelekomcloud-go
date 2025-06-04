// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cbrvaultv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/cbrvaultv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CbrVaultV3ResourceOutputReference interface {
	cdktf.ComplexObject
	BackupCount() *float64
	SetBackupCount(val *float64)
	BackupCountInput() *float64
	BackupSize() *float64
	SetBackupSize(val *float64)
	BackupSizeInput() *float64
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
	ExcludeVolumes() *[]*string
	SetExcludeVolumes(val *[]*string)
	ExcludeVolumesInput() *[]*string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IncludeVolumes() *[]*string
	SetIncludeVolumes(val *[]*string)
	IncludeVolumesInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	ProtectStatus() *string
	SetProtectStatus(val *string)
	ProtectStatusInput() *string
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetBackupCount()
	ResetBackupSize()
	ResetExcludeVolumes()
	ResetId()
	ResetIncludeVolumes()
	ResetName()
	ResetProtectStatus()
	ResetSize()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CbrVaultV3ResourceOutputReference
type jsiiProxy_CbrVaultV3ResourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) BackupCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) BackupCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) BackupSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) BackupSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) ExcludeVolumes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) ExcludeVolumesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) IncludeVolumes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) IncludeVolumesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) ProtectStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) ProtectStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCbrVaultV3ResourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CbrVaultV3ResourceOutputReference {
	_init_.Initialize()

	if err := validateNewCbrVaultV3ResourceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CbrVaultV3ResourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cbrVaultV3.CbrVaultV3ResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCbrVaultV3ResourceOutputReference_Override(c CbrVaultV3ResourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cbrVaultV3.CbrVaultV3ResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetBackupCount(val *float64) {
	if err := j.validateSetBackupCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupCount",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetBackupSize(val *float64) {
	if err := j.validateSetBackupSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupSize",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetExcludeVolumes(val *[]*string) {
	if err := j.validateSetExcludeVolumesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeVolumes",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetIncludeVolumes(val *[]*string) {
	if err := j.validateSetIncludeVolumesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeVolumes",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetProtectStatus(val *string) {
	if err := j.validateSetProtectStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectStatus",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3ResourceOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) ResetBackupCount() {
	_jsii_.InvokeVoid(
		c,
		"resetBackupCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) ResetBackupSize() {
	_jsii_.InvokeVoid(
		c,
		"resetBackupSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) ResetExcludeVolumes() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludeVolumes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) ResetIncludeVolumes() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeVolumes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) ResetProtectStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetProtectStatus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) ResetSize() {
	_jsii_.InvokeVoid(
		c,
		"resetSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CbrVaultV3ResourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

