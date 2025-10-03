// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudceseventdetailsv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudceseventdetailsv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/data-sources/ces_event_details_v1 opentelekomcloud_ces_event_details_v1}.
type DataOpentelekomcloudCesEventDetailsV1 interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EventInfo() DataOpentelekomcloudCesEventDetailsV1EventInfoList
	EventLevel() *string
	SetEventLevel(val *string)
	EventLevelInput() *string
	EventName() *string
	SetEventName(val *string)
	EventNameInput() *string
	EventSource() *string
	SetEventSource(val *string)
	EventSourceInput() *string
	EventSources() *[]*string
	EventState() *string
	SetEventState(val *string)
	EventStateInput() *string
	EventType() *string
	SetEventType(val *string)
	EventTypeInput() *string
	EventUser() *string
	SetEventUser(val *string)
	EventUserInput() *string
	EventUsers() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	From() *float64
	SetFrom(val *float64)
	FromInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Limit() *float64
	SetLimit(val *float64)
	LimitInput() *float64
	MetaData() DataOpentelekomcloudCesEventDetailsV1MetaDataList
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	To() *float64
	SetTo(val *float64)
	ToInput() *float64
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
	ResetEventLevel()
	ResetEventSource()
	ResetEventState()
	ResetEventUser()
	ResetFrom()
	ResetId()
	ResetLimit()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTo()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataOpentelekomcloudCesEventDetailsV1
type jsiiProxy_DataOpentelekomcloudCesEventDetailsV1 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventInfo() DataOpentelekomcloudCesEventDetailsV1EventInfoList {
	var returns DataOpentelekomcloudCesEventDetailsV1EventInfoList
	_jsii_.Get(
		j,
		"eventInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) EventUsers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) From() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"from",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) FromInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) MetaData() DataOpentelekomcloudCesEventDetailsV1MetaDataList {
	var returns DataOpentelekomcloudCesEventDetailsV1MetaDataList
	_jsii_.Get(
		j,
		"metaData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) To() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"to",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ToInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/data-sources/ces_event_details_v1 opentelekomcloud_ces_event_details_v1} Data Source.
func NewDataOpentelekomcloudCesEventDetailsV1(scope constructs.Construct, id *string, config *DataOpentelekomcloudCesEventDetailsV1Config) DataOpentelekomcloudCesEventDetailsV1 {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudCesEventDetailsV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudCesEventDetailsV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCesEventDetailsV1.DataOpentelekomcloudCesEventDetailsV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/data-sources/ces_event_details_v1 opentelekomcloud_ces_event_details_v1} Data Source.
func NewDataOpentelekomcloudCesEventDetailsV1_Override(d DataOpentelekomcloudCesEventDetailsV1, scope constructs.Construct, id *string, config *DataOpentelekomcloudCesEventDetailsV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCesEventDetailsV1.DataOpentelekomcloudCesEventDetailsV1",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetEventLevel(val *string) {
	if err := j.validateSetEventLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventLevel",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetEventName(val *string) {
	if err := j.validateSetEventNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventName",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetEventSource(val *string) {
	if err := j.validateSetEventSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSource",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetEventState(val *string) {
	if err := j.validateSetEventStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventState",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetEventType(val *string) {
	if err := j.validateSetEventTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventType",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetEventUser(val *string) {
	if err := j.validateSetEventUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventUser",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetFrom(val *float64) {
	if err := j.validateSetFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"from",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetLimit(val *float64) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1)SetTo(val *float64) {
	if err := j.validateSetToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"to",
		val,
	)
}

// Generates CDKTF code for importing a DataOpentelekomcloudCesEventDetailsV1 resource upon running "cdktf plan <stack-name>".
func DataOpentelekomcloudCesEventDetailsV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudCesEventDetailsV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCesEventDetailsV1.DataOpentelekomcloudCesEventDetailsV1",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func DataOpentelekomcloudCesEventDetailsV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudCesEventDetailsV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCesEventDetailsV1.DataOpentelekomcloudCesEventDetailsV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudCesEventDetailsV1_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudCesEventDetailsV1_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCesEventDetailsV1.DataOpentelekomcloudCesEventDetailsV1",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudCesEventDetailsV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudCesEventDetailsV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCesEventDetailsV1.DataOpentelekomcloudCesEventDetailsV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudCesEventDetailsV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCesEventDetailsV1.DataOpentelekomcloudCesEventDetailsV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ResetEventLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetEventLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ResetEventSource() {
	_jsii_.InvokeVoid(
		d,
		"resetEventSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ResetEventState() {
	_jsii_.InvokeVoid(
		d,
		"resetEventState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ResetEventUser() {
	_jsii_.InvokeVoid(
		d,
		"resetEventUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ResetFrom() {
	_jsii_.InvokeVoid(
		d,
		"resetFrom",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ResetLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ResetTo() {
	_jsii_.InvokeVoid(
		d,
		"resetTo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCesEventDetailsV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

