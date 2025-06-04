// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/computeinstancev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceV2BlockDeviceOutputReference interface {
	cdktf.ComplexObject
	BootIndex() *float64
	SetBootIndex(val *float64)
	BootIndexInput() *float64
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
	DeleteOnTermination() interface{}
	SetDeleteOnTermination(val interface{})
	DeleteOnTerminationInput() interface{}
	DestinationType() *string
	SetDestinationType(val *string)
	DestinationTypeInput() *string
	DeviceName() *string
	SetDeviceName(val *string)
	DeviceNameInput() *string
	// Experimental.
	Fqn() *string
	GuestFormat() *string
	SetGuestFormat(val *string)
	GuestFormatInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SourceType() *string
	SetSourceType(val *string)
	SourceTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uuid() *string
	SetUuid(val *string)
	UuidInput() *string
	VolumeSize() *float64
	SetVolumeSize(val *float64)
	VolumeSizeInput() *float64
	VolumeType() *string
	SetVolumeType(val *string)
	VolumeTypeInput() *string
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
	ResetBootIndex()
	ResetDeleteOnTermination()
	ResetDestinationType()
	ResetDeviceName()
	ResetGuestFormat()
	ResetUuid()
	ResetVolumeSize()
	ResetVolumeType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceV2BlockDeviceOutputReference
type jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) BootIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) BootIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) DeleteOnTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) DeleteOnTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) DestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) DestinationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) DeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) GuestFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) GuestFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) SourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) SourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) UuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) VolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) VolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}


func NewComputeInstanceV2BlockDeviceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeInstanceV2BlockDeviceOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInstanceV2BlockDeviceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.computeInstanceV2.ComputeInstanceV2BlockDeviceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeInstanceV2BlockDeviceOutputReference_Override(c ComputeInstanceV2BlockDeviceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.computeInstanceV2.ComputeInstanceV2BlockDeviceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetBootIndex(val *float64) {
	if err := j.validateSetBootIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetDeleteOnTermination(val interface{}) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetDestinationType(val *string) {
	if err := j.validateSetDestinationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationType",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetDeviceName(val *string) {
	if err := j.validateSetDeviceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceName",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetGuestFormat(val *string) {
	if err := j.validateSetGuestFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestFormat",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetSourceType(val *string) {
	if err := j.validateSetSourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceType",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetUuid(val *string) {
	if err := j.validateSetUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uuid",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetVolumeSize(val *float64) {
	if err := j.validateSetVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSize",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference)SetVolumeType(val *string) {
	if err := j.validateSetVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) ResetBootIndex() {
	_jsii_.InvokeVoid(
		c,
		"resetBootIndex",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		c,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) ResetDestinationType() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) ResetDeviceName() {
	_jsii_.InvokeVoid(
		c,
		"resetDeviceName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) ResetGuestFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetGuestFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) ResetUuid() {
	_jsii_.InvokeVoid(
		c,
		"resetUuid",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) ResetVolumeSize() {
	_jsii_.InvokeVoid(
		c,
		"resetVolumeSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) ResetVolumeType() {
	_jsii_.InvokeVoid(
		c,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeInstanceV2BlockDeviceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

