// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmssmartconnecttaskv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dmssmartconnecttaskv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsSmartConnectTaskV2DestinationTaskOutputReference interface {
	cdktf.ComplexObject
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *DmsSmartConnectTaskV2DestinationTask
	SetInternalValue(val *DmsSmartConnectTaskV2DestinationTask)
	ObsBucketName() *string
	SetObsBucketName(val *string)
	ObsBucketNameInput() *string
	ObsPath() *string
	SetObsPath(val *string)
	ObsPathInput() *string
	PartitionFormat() *string
	SetPartitionFormat(val *string)
	PartitionFormatInput() *string
	RecordDelimiter() *string
	SetRecordDelimiter(val *string)
	RecordDelimiterInput() *string
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
	StoreKeys() interface{}
	SetStoreKeys(val interface{})
	StoreKeysInput() interface{}
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
	ResetAccessKey()
	ResetConsumerStrategy()
	ResetDeliverTimeInterval()
	ResetDestinationFileType()
	ResetObsBucketName()
	ResetObsPath()
	ResetPartitionFormat()
	ResetRecordDelimiter()
	ResetSecretKey()
	ResetStoreKeys()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsSmartConnectTaskV2DestinationTaskOutputReference
type jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ConsumerStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ConsumerStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) DeliverTimeInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deliverTimeInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) DeliverTimeIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deliverTimeIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) DestinationFileType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFileType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) DestinationFileTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFileTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) InternalValue() *DmsSmartConnectTaskV2DestinationTask {
	var returns *DmsSmartConnectTaskV2DestinationTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ObsBucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsBucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ObsBucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsBucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ObsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ObsPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) PartitionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) PartitionFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) RecordDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) RecordDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) StoreKeys() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storeKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) StoreKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storeKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsSmartConnectTaskV2DestinationTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DmsSmartConnectTaskV2DestinationTaskOutputReference {
	_init_.Initialize()

	if err := validateNewDmsSmartConnectTaskV2DestinationTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dmsSmartConnectTaskV2.DmsSmartConnectTaskV2DestinationTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsSmartConnectTaskV2DestinationTaskOutputReference_Override(d DmsSmartConnectTaskV2DestinationTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dmsSmartConnectTaskV2.DmsSmartConnectTaskV2DestinationTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetAccessKey(val *string) {
	if err := j.validateSetAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetConsumerStrategy(val *string) {
	if err := j.validateSetConsumerStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerStrategy",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetDeliverTimeInterval(val *float64) {
	if err := j.validateSetDeliverTimeIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliverTimeInterval",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetDestinationFileType(val *string) {
	if err := j.validateSetDestinationFileTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFileType",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetInternalValue(val *DmsSmartConnectTaskV2DestinationTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetObsBucketName(val *string) {
	if err := j.validateSetObsBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsBucketName",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetObsPath(val *string) {
	if err := j.validateSetObsPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsPath",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetPartitionFormat(val *string) {
	if err := j.validateSetPartitionFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionFormat",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetRecordDelimiter(val *string) {
	if err := j.validateSetRecordDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordDelimiter",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetSecretKey(val *string) {
	if err := j.validateSetSecretKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetStoreKeys(val interface{}) {
	if err := j.validateSetStoreKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storeKeys",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ResetAccessKey() {
	_jsii_.InvokeVoid(
		d,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ResetConsumerStrategy() {
	_jsii_.InvokeVoid(
		d,
		"resetConsumerStrategy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ResetDeliverTimeInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetDeliverTimeInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ResetDestinationFileType() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationFileType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ResetObsBucketName() {
	_jsii_.InvokeVoid(
		d,
		"resetObsBucketName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ResetObsPath() {
	_jsii_.InvokeVoid(
		d,
		"resetObsPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ResetPartitionFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetPartitionFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ResetRecordDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetRecordDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ResetSecretKey() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ResetStoreKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetStoreKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2DestinationTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

