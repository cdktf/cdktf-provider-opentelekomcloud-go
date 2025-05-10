// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudhssintrusioneventsv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/dataopentelekomcloudhssintrusioneventsv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/data-sources/hss_intrusion_events_v5 opentelekomcloud_hss_intrusion_events_v5}.
type DataOpentelekomcloudHssIntrusionEventsV5 interface {
	cdktf.TerraformDataSource
	BeginTime() *string
	SetBeginTime(val *string)
	BeginTimeInput() *string
	Category() *string
	SetCategory(val *string)
	CategoryInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerName() *string
	SetContainerName(val *string)
	ContainerNameInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Days() *float64
	SetDays(val *float64)
	DaysInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	EnterpriseProjectId() *string
	SetEnterpriseProjectId(val *string)
	EnterpriseProjectIdInput() *string
	Events() DataOpentelekomcloudHssIntrusionEventsV5EventsList
	EventTypes() *[]*string
	SetEventTypes(val *[]*string)
	EventTypesInput() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HandleStatus() *string
	SetHandleStatus(val *string)
	HandleStatusInput() *string
	HostId() *string
	SetHostId(val *string)
	HostIdInput() *string
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PrivateIp() *string
	SetPrivateIp(val *string)
	PrivateIpInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	Severity() *string
	SetSeverity(val *string)
	SeverityInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetBeginTime()
	ResetContainerName()
	ResetDays()
	ResetEndTime()
	ResetEnterpriseProjectId()
	ResetEventTypes()
	ResetHandleStatus()
	ResetHostId()
	ResetHostName()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateIp()
	ResetSeverity()
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

// The jsii proxy struct for DataOpentelekomcloudHssIntrusionEventsV5
type jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) BeginTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) BeginTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) CategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"categoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) Days() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"days",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) DaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) EnterpriseProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) EnterpriseProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) Events() DataOpentelekomcloudHssIntrusionEventsV5EventsList {
	var returns DataOpentelekomcloudHssIntrusionEventsV5EventsList
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) EventTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) EventTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) HandleStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handleStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) HandleStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handleStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) HostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) HostIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/data-sources/hss_intrusion_events_v5 opentelekomcloud_hss_intrusion_events_v5} Data Source.
func NewDataOpentelekomcloudHssIntrusionEventsV5(scope constructs.Construct, id *string, config *DataOpentelekomcloudHssIntrusionEventsV5Config) DataOpentelekomcloudHssIntrusionEventsV5 {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudHssIntrusionEventsV5Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/data-sources/hss_intrusion_events_v5 opentelekomcloud_hss_intrusion_events_v5} Data Source.
func NewDataOpentelekomcloudHssIntrusionEventsV5_Override(d DataOpentelekomcloudHssIntrusionEventsV5, scope constructs.Construct, id *string, config *DataOpentelekomcloudHssIntrusionEventsV5Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetBeginTime(val *string) {
	if err := j.validateSetBeginTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"beginTime",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetCategory(val *string) {
	if err := j.validateSetCategoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"category",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetContainerName(val *string) {
	if err := j.validateSetContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerName",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetDays(val *float64) {
	if err := j.validateSetDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"days",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetEndTime(val *string) {
	if err := j.validateSetEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetEnterpriseProjectId(val *string) {
	if err := j.validateSetEnterpriseProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enterpriseProjectId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetEventTypes(val *[]*string) {
	if err := j.validateSetEventTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventTypes",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetHandleStatus(val *string) {
	if err := j.validateSetHandleStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handleStatus",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetHostId(val *string) {
	if err := j.validateSetHostIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetHostName(val *string) {
	if err := j.validateSetHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

// Generates CDKTF code for importing a DataOpentelekomcloudHssIntrusionEventsV5 resource upon running "cdktf plan <stack-name>".
func DataOpentelekomcloudHssIntrusionEventsV5_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudHssIntrusionEventsV5_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5",
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
func DataOpentelekomcloudHssIntrusionEventsV5_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudHssIntrusionEventsV5_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudHssIntrusionEventsV5_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudHssIntrusionEventsV5_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudHssIntrusionEventsV5_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudHssIntrusionEventsV5_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudHssIntrusionEventsV5_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetBeginTime() {
	_jsii_.InvokeVoid(
		d,
		"resetBeginTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetContainerName() {
	_jsii_.InvokeVoid(
		d,
		"resetContainerName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetDays() {
	_jsii_.InvokeVoid(
		d,
		"resetDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetEndTime() {
	_jsii_.InvokeVoid(
		d,
		"resetEndTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetEnterpriseProjectId() {
	_jsii_.InvokeVoid(
		d,
		"resetEnterpriseProjectId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetEventTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetEventTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetHandleStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetHandleStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetHostId() {
	_jsii_.InvokeVoid(
		d,
		"resetHostId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetHostName() {
	_jsii_.InvokeVoid(
		d,
		"resetHostName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ResetSeverity() {
	_jsii_.InvokeVoid(
		d,
		"resetSeverity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

