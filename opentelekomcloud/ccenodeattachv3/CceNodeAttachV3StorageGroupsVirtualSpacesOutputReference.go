// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodeattachv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/ccenodeattachv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference interface {
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
	LvmLvType() *string
	SetLvmLvType(val *string)
	LvmLvTypeInput() *string
	LvmPath() *string
	SetLvmPath(val *string)
	LvmPathInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	RuntimeLvType() *string
	SetRuntimeLvType(val *string)
	RuntimeLvTypeInput() *string
	Size() *string
	SetSize(val *string)
	SizeInput() *string
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
	ResetLvmLvType()
	ResetLvmPath()
	ResetRuntimeLvType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference
type jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) LvmLvType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lvmLvType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) LvmLvTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lvmLvTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) LvmPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lvmPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) LvmPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lvmPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) RuntimeLvType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeLvType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) RuntimeLvTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeLvTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) Size() *string {
	var returns *string
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) SizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCceNodeAttachV3StorageGroupsVirtualSpacesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference {
	_init_.Initialize()

	if err := validateNewCceNodeAttachV3StorageGroupsVirtualSpacesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cceNodeAttachV3.CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCceNodeAttachV3StorageGroupsVirtualSpacesOutputReference_Override(c CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cceNodeAttachV3.CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference)SetLvmLvType(val *string) {
	if err := j.validateSetLvmLvTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lvmLvType",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference)SetLvmPath(val *string) {
	if err := j.validateSetLvmPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lvmPath",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference)SetRuntimeLvType(val *string) {
	if err := j.validateSetRuntimeLvTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeLvType",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference)SetSize(val *string) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) ResetLvmLvType() {
	_jsii_.InvokeVoid(
		c,
		"resetLvmLvType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) ResetLvmPath() {
	_jsii_.InvokeVoid(
		c,
		"resetLvmPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) ResetRuntimeLvType() {
	_jsii_.InvokeVoid(
		c,
		"resetRuntimeLvType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CceNodeAttachV3StorageGroupsVirtualSpacesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

