// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudhsshostsv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudhsshostsv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/data-sources/hss_hosts_v5 opentelekomcloud_hss_hosts_v5}.
type DataOpentelekomcloudHssHostsV5 interface {
	cdktf.TerraformDataSource
	AgentStatus() *string
	SetAgentStatus(val *string)
	AgentStatusInput() *string
	AssetValue() *string
	SetAssetValue(val *string)
	AssetValueInput() *string
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
	DetectResult() *string
	SetDetectResult(val *string)
	DetectResultInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupId() *string
	SetGroupId(val *string)
	GroupIdInput() *string
	HostId() *string
	SetHostId(val *string)
	HostIdInput() *string
	Hosts() DataOpentelekomcloudHssHostsV5HostsList
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
	OsType() *string
	SetOsType(val *string)
	OsTypeInput() *string
	PolicyGroupId() *string
	SetPolicyGroupId(val *string)
	PolicyGroupIdInput() *string
	ProtectChargingMode() *string
	SetProtectChargingMode(val *string)
	ProtectChargingModeInput() *string
	ProtectStatus() *string
	SetProtectStatus(val *string)
	ProtectStatusInput() *string
	ProtectVersion() *string
	SetProtectVersion(val *string)
	ProtectVersionInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
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
	ResetAgentStatus()
	ResetAssetValue()
	ResetDetectResult()
	ResetGroupId()
	ResetHostId()
	ResetId()
	ResetName()
	ResetOsType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyGroupId()
	ResetProtectChargingMode()
	ResetProtectStatus()
	ResetProtectVersion()
	ResetStatus()
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

// The jsii proxy struct for DataOpentelekomcloudHssHostsV5
type jsiiProxy_DataOpentelekomcloudHssHostsV5 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) AgentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) AgentStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) AssetValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) AssetValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) DetectResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) DetectResultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectResultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) GroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) HostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) HostIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) Hosts() DataOpentelekomcloudHssHostsV5HostsList {
	var returns DataOpentelekomcloudHssHostsV5HostsList
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) OsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) PolicyGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) PolicyGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) ProtectChargingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectChargingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) ProtectChargingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectChargingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) ProtectStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) ProtectStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) ProtectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) ProtectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/data-sources/hss_hosts_v5 opentelekomcloud_hss_hosts_v5} Data Source.
func NewDataOpentelekomcloudHssHostsV5(scope constructs.Construct, id *string, config *DataOpentelekomcloudHssHostsV5Config) DataOpentelekomcloudHssHostsV5 {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudHssHostsV5Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudHssHostsV5{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssHostsV5.DataOpentelekomcloudHssHostsV5",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/data-sources/hss_hosts_v5 opentelekomcloud_hss_hosts_v5} Data Source.
func NewDataOpentelekomcloudHssHostsV5_Override(d DataOpentelekomcloudHssHostsV5, scope constructs.Construct, id *string, config *DataOpentelekomcloudHssHostsV5Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssHostsV5.DataOpentelekomcloudHssHostsV5",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetAgentStatus(val *string) {
	if err := j.validateSetAgentStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentStatus",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetAssetValue(val *string) {
	if err := j.validateSetAssetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetValue",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetDetectResult(val *string) {
	if err := j.validateSetDetectResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectResult",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetGroupId(val *string) {
	if err := j.validateSetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetHostId(val *string) {
	if err := j.validateSetHostIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetOsType(val *string) {
	if err := j.validateSetOsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetPolicyGroupId(val *string) {
	if err := j.validateSetPolicyGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyGroupId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetProtectChargingMode(val *string) {
	if err := j.validateSetProtectChargingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectChargingMode",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetProtectStatus(val *string) {
	if err := j.validateSetProtectStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectStatus",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetProtectVersion(val *string) {
	if err := j.validateSetProtectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectVersion",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

// Generates CDKTF code for importing a DataOpentelekomcloudHssHostsV5 resource upon running "cdktf plan <stack-name>".
func DataOpentelekomcloudHssHostsV5_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudHssHostsV5_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssHostsV5.DataOpentelekomcloudHssHostsV5",
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
func DataOpentelekomcloudHssHostsV5_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudHssHostsV5_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssHostsV5.DataOpentelekomcloudHssHostsV5",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudHssHostsV5_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudHssHostsV5_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssHostsV5.DataOpentelekomcloudHssHostsV5",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudHssHostsV5_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudHssHostsV5_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssHostsV5.DataOpentelekomcloudHssHostsV5",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudHssHostsV5_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssHostsV5.DataOpentelekomcloudHssHostsV5",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetAgentStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetAgentStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetAssetValue() {
	_jsii_.InvokeVoid(
		d,
		"resetAssetValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetDetectResult() {
	_jsii_.InvokeVoid(
		d,
		"resetDetectResult",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetGroupId() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetHostId() {
	_jsii_.InvokeVoid(
		d,
		"resetHostId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetOsType() {
	_jsii_.InvokeVoid(
		d,
		"resetOsType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetPolicyGroupId() {
	_jsii_.InvokeVoid(
		d,
		"resetPolicyGroupId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetProtectChargingMode() {
	_jsii_.InvokeVoid(
		d,
		"resetProtectChargingMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetProtectStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetProtectStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetProtectVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetProtectVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

