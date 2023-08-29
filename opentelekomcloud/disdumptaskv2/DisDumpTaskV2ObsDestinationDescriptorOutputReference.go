// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package disdumptaskv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v9/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v9/disdumptaskv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DisDumpTaskV2ObsDestinationDescriptorOutputReference interface {
	cdktf.ComplexObject
	AgencyName() *string
	SetAgencyName(val *string)
	AgencyNameInput() *string
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
	ConsumerStrategy() *string
	SetConsumerStrategy(val *string)
	ConsumerStrategyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeliverTimeInterval() *float64
	SetDeliverTimeInterval(val *float64)
	DeliverTimeIntervalInput() *float64
	DestinationFileType() *string
	SetDestinationFileType(val *string)
	DestinationFileTypeInput() *string
	FilePrefix() *string
	SetFilePrefix(val *string)
	FilePrefixInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ObsBucketPath() *string
	SetObsBucketPath(val *string)
	ObsBucketPathInput() *string
	PartitionFormat() *string
	SetPartitionFormat(val *string)
	PartitionFormatInput() *string
	RecordDelimiter() *string
	SetRecordDelimiter(val *string)
	RecordDelimiterInput() *string
	TaskName() *string
	SetTaskName(val *string)
	TaskNameInput() *string
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
	ResetConsumerStrategy()
	ResetDestinationFileType()
	ResetFilePrefix()
	ResetPartitionFormat()
	ResetRecordDelimiter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DisDumpTaskV2ObsDestinationDescriptorOutputReference
type jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) AgencyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) AgencyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ConsumerStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ConsumerStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) DeliverTimeInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deliverTimeInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) DeliverTimeIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deliverTimeIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) DestinationFileType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFileType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) DestinationFileTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFileTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) FilePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) FilePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ObsBucketPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsBucketPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ObsBucketPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsBucketPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) PartitionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) PartitionFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) RecordDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) RecordDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) TaskName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) TaskNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDisDumpTaskV2ObsDestinationDescriptorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DisDumpTaskV2ObsDestinationDescriptorOutputReference {
	_init_.Initialize()

	if err := validateNewDisDumpTaskV2ObsDestinationDescriptorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.disDumpTaskV2.DisDumpTaskV2ObsDestinationDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDisDumpTaskV2ObsDestinationDescriptorOutputReference_Override(d DisDumpTaskV2ObsDestinationDescriptorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.disDumpTaskV2.DisDumpTaskV2ObsDestinationDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetAgencyName(val *string) {
	if err := j.validateSetAgencyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agencyName",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetConsumerStrategy(val *string) {
	if err := j.validateSetConsumerStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerStrategy",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetDeliverTimeInterval(val *float64) {
	if err := j.validateSetDeliverTimeIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliverTimeInterval",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetDestinationFileType(val *string) {
	if err := j.validateSetDestinationFileTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFileType",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetFilePrefix(val *string) {
	if err := j.validateSetFilePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filePrefix",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetObsBucketPath(val *string) {
	if err := j.validateSetObsBucketPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsBucketPath",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetPartitionFormat(val *string) {
	if err := j.validateSetPartitionFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionFormat",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetRecordDelimiter(val *string) {
	if err := j.validateSetRecordDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordDelimiter",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetTaskName(val *string) {
	if err := j.validateSetTaskNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskName",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ResetConsumerStrategy() {
	_jsii_.InvokeVoid(
		d,
		"resetConsumerStrategy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ResetDestinationFileType() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationFileType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ResetFilePrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetFilePrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ResetPartitionFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetPartitionFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ResetRecordDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetRecordDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

