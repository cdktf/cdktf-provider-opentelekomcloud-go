// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package disstreamv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v9/disstreamv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.7/docs/resources/dis_stream_v2 opentelekomcloud_dis_stream_v2}.
type DisStreamV2 interface {
	cdktf.TerraformResource
	AutoScaleMaxPartitionCount() *float64
	SetAutoScaleMaxPartitionCount(val *float64)
	AutoScaleMaxPartitionCountInput() *float64
	AutoScaleMinPartitionCount() *float64
	SetAutoScaleMinPartitionCount(val *float64)
	AutoScaleMinPartitionCountInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CompressionFormat() *string
	SetCompressionFormat(val *string)
	CompressionFormatInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Created() *float64
	DataType() *string
	SetDataType(val *string)
	DataTypeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PartitionCount() *float64
	SetPartitionCount(val *float64)
	PartitionCountInput() *float64
	Partitions() DisStreamV2PartitionsList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReadablePartitionCount() *float64
	RetentionPeriod() *float64
	SetRetentionPeriod(val *float64)
	RetentionPeriodInput() *float64
	Status() *string
	StreamId() *string
	StreamType() *string
	SetStreamType(val *string)
	StreamTypeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DisStreamV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	WritablePartitionCount() *float64
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *DisStreamV2Timeouts)
	ResetAutoScaleMaxPartitionCount()
	ResetAutoScaleMinPartitionCount()
	ResetCompressionFormat()
	ResetDataType()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetentionPeriod()
	ResetStreamType()
	ResetTags()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DisStreamV2
type jsiiProxy_DisStreamV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DisStreamV2) AutoScaleMaxPartitionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoScaleMaxPartitionCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) AutoScaleMaxPartitionCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoScaleMaxPartitionCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) AutoScaleMinPartitionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoScaleMinPartitionCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) AutoScaleMinPartitionCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoScaleMinPartitionCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) CompressionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) CompressionFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Created() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) DataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) DataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) PartitionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) PartitionCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Partitions() DisStreamV2PartitionsList {
	var returns DisStreamV2PartitionsList
	_jsii_.Get(
		j,
		"partitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) ReadablePartitionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readablePartitionCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) RetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) RetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) StreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) StreamType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) StreamTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) Timeouts() DisStreamV2TimeoutsOutputReference {
	var returns DisStreamV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisStreamV2) WritablePartitionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writablePartitionCount",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.7/docs/resources/dis_stream_v2 opentelekomcloud_dis_stream_v2} Resource.
func NewDisStreamV2(scope constructs.Construct, id *string, config *DisStreamV2Config) DisStreamV2 {
	_init_.Initialize()

	if err := validateNewDisStreamV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DisStreamV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.disStreamV2.DisStreamV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.7/docs/resources/dis_stream_v2 opentelekomcloud_dis_stream_v2} Resource.
func NewDisStreamV2_Override(d DisStreamV2, scope constructs.Construct, id *string, config *DisStreamV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.disStreamV2.DisStreamV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DisStreamV2)SetAutoScaleMaxPartitionCount(val *float64) {
	if err := j.validateSetAutoScaleMaxPartitionCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScaleMaxPartitionCount",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetAutoScaleMinPartitionCount(val *float64) {
	if err := j.validateSetAutoScaleMinPartitionCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScaleMinPartitionCount",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetCompressionFormat(val *string) {
	if err := j.validateSetCompressionFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionFormat",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetDataType(val *string) {
	if err := j.validateSetDataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataType",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetPartitionCount(val *float64) {
	if err := j.validateSetPartitionCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionCount",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetRetentionPeriod(val *float64) {
	if err := j.validateSetRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPeriod",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetStreamType(val *string) {
	if err := j.validateSetStreamTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamType",
		val,
	)
}

func (j *jsiiProxy_DisStreamV2)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DisStreamV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDisStreamV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.disStreamV2.DisStreamV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DisStreamV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDisStreamV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.disStreamV2.DisStreamV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DisStreamV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDisStreamV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.disStreamV2.DisStreamV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DisStreamV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.disStreamV2.DisStreamV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DisStreamV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DisStreamV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DisStreamV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DisStreamV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DisStreamV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DisStreamV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DisStreamV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DisStreamV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DisStreamV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DisStreamV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DisStreamV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisStreamV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DisStreamV2) PutTimeouts(value *DisStreamV2Timeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DisStreamV2) ResetAutoScaleMaxPartitionCount() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoScaleMaxPartitionCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisStreamV2) ResetAutoScaleMinPartitionCount() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoScaleMinPartitionCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisStreamV2) ResetCompressionFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetCompressionFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisStreamV2) ResetDataType() {
	_jsii_.InvokeVoid(
		d,
		"resetDataType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisStreamV2) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisStreamV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisStreamV2) ResetRetentionPeriod() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionPeriod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisStreamV2) ResetStreamType() {
	_jsii_.InvokeVoid(
		d,
		"resetStreamType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisStreamV2) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisStreamV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisStreamV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisStreamV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisStreamV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisStreamV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

