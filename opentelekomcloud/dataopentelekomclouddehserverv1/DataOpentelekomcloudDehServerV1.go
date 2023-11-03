// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomclouddehserverv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v10/dataopentelekomclouddehserverv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/deh_server_v1 opentelekomcloud_deh_server_v1}.
type DataOpentelekomcloudDehServerV1 interface {
	cdktf.TerraformDataSource
	Addresses() DataOpentelekomcloudDehServerV1AddressesList
	AddressesInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DedicatedHostId() *string
	SetDedicatedHostId(val *string)
	DedicatedHostIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Flavor() cdktf.StringMap
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
	Metadata() cdktf.StringMap
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ServerId() *string
	SetServerId(val *string)
	ServerIdInput() *string
	Status() *string
	TenantId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserId() *string
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
	PutAddresses(value interface{})
	ResetAddresses()
	ResetDedicatedHostId()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetServerId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataOpentelekomcloudDehServerV1
type jsiiProxy_DataOpentelekomcloudDehServerV1 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) Addresses() DataOpentelekomcloudDehServerV1AddressesList {
	var returns DataOpentelekomcloudDehServerV1AddressesList
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) DedicatedHostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedHostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) DedicatedHostIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedHostIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) Flavor() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"flavor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) Metadata() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) ServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) ServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/deh_server_v1 opentelekomcloud_deh_server_v1} Data Source.
func NewDataOpentelekomcloudDehServerV1(scope constructs.Construct, id *string, config *DataOpentelekomcloudDehServerV1Config) DataOpentelekomcloudDehServerV1 {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudDehServerV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudDehServerV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudDehServerV1.DataOpentelekomcloudDehServerV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/deh_server_v1 opentelekomcloud_deh_server_v1} Data Source.
func NewDataOpentelekomcloudDehServerV1_Override(d DataOpentelekomcloudDehServerV1, scope constructs.Construct, id *string, config *DataOpentelekomcloudDehServerV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudDehServerV1.DataOpentelekomcloudDehServerV1",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1)SetDedicatedHostId(val *string) {
	if err := j.validateSetDedicatedHostIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedHostId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudDehServerV1)SetServerId(val *string) {
	if err := j.validateSetServerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverId",
		val,
	)
}

// Generates CDKTF code for importing a DataOpentelekomcloudDehServerV1 resource upon running "cdktf plan <stack-name>".
func DataOpentelekomcloudDehServerV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudDehServerV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudDehServerV1.DataOpentelekomcloudDehServerV1",
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
func DataOpentelekomcloudDehServerV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudDehServerV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudDehServerV1.DataOpentelekomcloudDehServerV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudDehServerV1_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudDehServerV1_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudDehServerV1.DataOpentelekomcloudDehServerV1",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudDehServerV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudDehServerV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudDehServerV1.DataOpentelekomcloudDehServerV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudDehServerV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudDehServerV1.DataOpentelekomcloudDehServerV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) PutAddresses(value interface{}) {
	if err := d.validatePutAddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAddresses",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) ResetAddresses() {
	_jsii_.InvokeVoid(
		d,
		"resetAddresses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) ResetDedicatedHostId() {
	_jsii_.InvokeVoid(
		d,
		"resetDedicatedHostId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) ResetServerId() {
	_jsii_.InvokeVoid(
		d,
		"resetServerId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudDehServerV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

