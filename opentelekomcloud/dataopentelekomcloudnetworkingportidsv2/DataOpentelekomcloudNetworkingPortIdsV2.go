// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudnetworkingportidsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudnetworkingportidsv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/data-sources/networking_port_ids_v2 opentelekomcloud_networking_port_ids_v2}.
type DataOpentelekomcloudNetworkingPortIdsV2 interface {
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
	DeviceId() *string
	SetDeviceId(val *string)
	DeviceIdInput() *string
	DeviceOwner() *string
	SetDeviceOwner(val *string)
	DeviceOwnerInput() *string
	FixedIp() *string
	SetFixedIp(val *string)
	FixedIpInput() *string
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
	Ids() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MacAddress() *string
	SetMacAddress(val *string)
	MacAddressInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
	// The tree node.
	Node() constructs.Node
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SortDirection() *string
	SetSortDirection(val *string)
	SortDirectionInput() *string
	SortKey() *string
	SetSortKey(val *string)
	SortKeyInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	ResetDeviceId()
	ResetDeviceOwner()
	ResetFixedIp()
	ResetId()
	ResetMacAddress()
	ResetName()
	ResetNetworkId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectId()
	ResetRegion()
	ResetSecurityGroupIds()
	ResetSortDirection()
	ResetSortKey()
	ResetStatus()
	ResetTenantId()
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

// The jsii proxy struct for DataOpentelekomcloudNetworkingPortIdsV2
type jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) DeviceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) DeviceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) DeviceOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) DeviceOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) FixedIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) FixedIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) Ids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) MacAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) MacAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) SortDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) SortDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) SortKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) SortKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/data-sources/networking_port_ids_v2 opentelekomcloud_networking_port_ids_v2} Data Source.
func NewDataOpentelekomcloudNetworkingPortIdsV2(scope constructs.Construct, id *string, config *DataOpentelekomcloudNetworkingPortIdsV2Config) DataOpentelekomcloudNetworkingPortIdsV2 {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudNetworkingPortIdsV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNetworkingPortIdsV2.DataOpentelekomcloudNetworkingPortIdsV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/data-sources/networking_port_ids_v2 opentelekomcloud_networking_port_ids_v2} Data Source.
func NewDataOpentelekomcloudNetworkingPortIdsV2_Override(d DataOpentelekomcloudNetworkingPortIdsV2, scope constructs.Construct, id *string, config *DataOpentelekomcloudNetworkingPortIdsV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNetworkingPortIdsV2.DataOpentelekomcloudNetworkingPortIdsV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetDeviceId(val *string) {
	if err := j.validateSetDeviceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetDeviceOwner(val *string) {
	if err := j.validateSetDeviceOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceOwner",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetFixedIp(val *string) {
	if err := j.validateSetFixedIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedIp",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetMacAddress(val *string) {
	if err := j.validateSetMacAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macAddress",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetNetworkId(val *string) {
	if err := j.validateSetNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetSortDirection(val *string) {
	if err := j.validateSetSortDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortDirection",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetSortKey(val *string) {
	if err := j.validateSetSortKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortKey",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

// Generates CDKTF code for importing a DataOpentelekomcloudNetworkingPortIdsV2 resource upon running "cdktf plan <stack-name>".
func DataOpentelekomcloudNetworkingPortIdsV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudNetworkingPortIdsV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNetworkingPortIdsV2.DataOpentelekomcloudNetworkingPortIdsV2",
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
func DataOpentelekomcloudNetworkingPortIdsV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudNetworkingPortIdsV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNetworkingPortIdsV2.DataOpentelekomcloudNetworkingPortIdsV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudNetworkingPortIdsV2_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudNetworkingPortIdsV2_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNetworkingPortIdsV2.DataOpentelekomcloudNetworkingPortIdsV2",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudNetworkingPortIdsV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudNetworkingPortIdsV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNetworkingPortIdsV2.DataOpentelekomcloudNetworkingPortIdsV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudNetworkingPortIdsV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNetworkingPortIdsV2.DataOpentelekomcloudNetworkingPortIdsV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetDeviceId() {
	_jsii_.InvokeVoid(
		d,
		"resetDeviceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetDeviceOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetDeviceOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetFixedIp() {
	_jsii_.InvokeVoid(
		d,
		"resetFixedIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetMacAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetMacAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetNetworkId() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetProjectId() {
	_jsii_.InvokeVoid(
		d,
		"resetProjectId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetSortDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetSortDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetSortKey() {
	_jsii_.InvokeVoid(
		d,
		"resetSortKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ResetTenantId() {
	_jsii_.InvokeVoid(
		d,
		"resetTenantId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudNetworkingPortIdsV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

