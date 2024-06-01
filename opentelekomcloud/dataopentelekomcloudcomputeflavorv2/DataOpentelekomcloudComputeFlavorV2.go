// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudcomputeflavorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/dataopentelekomcloudcomputeflavorv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/data-sources/compute_flavor_v2 opentelekomcloud_compute_flavor_v2}.
type DataOpentelekomcloudComputeFlavorV2 interface {
	cdktf.TerraformDataSource
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
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
	Disk() *float64
	SetDisk(val *float64)
	DiskInput() *float64
	ExtraSpecs() cdktf.StringMap
	FlavorId() *string
	SetFlavorId(val *string)
	FlavorIdInput() *string
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
	MinDisk() *float64
	SetMinDisk(val *float64)
	MinDiskInput() *float64
	MinRam() *float64
	SetMinRam(val *float64)
	MinRamInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	Ram() *float64
	SetRam(val *float64)
	RamInput() *float64
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	RxTxFactor() *float64
	SetRxTxFactor(val *float64)
	RxTxFactorInput() *float64
	Swap() *float64
	SetSwap(val *float64)
	SwapInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Vcpus() *float64
	SetVcpus(val *float64)
	VcpusInput() *float64
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
	ResetAvailabilityZone()
	ResetDisk()
	ResetFlavorId()
	ResetId()
	ResetMinDisk()
	ResetMinRam()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRam()
	ResetRegion()
	ResetResourceType()
	ResetRxTxFactor()
	ResetSwap()
	ResetVcpus()
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

// The jsii proxy struct for DataOpentelekomcloudComputeFlavorV2
type jsiiProxy_DataOpentelekomcloudComputeFlavorV2 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) Disk() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) DiskInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ExtraSpecs() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"extraSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) FlavorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) FlavorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) MinDisk() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) MinDiskInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) MinRam() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) MinRamInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) Ram() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) RamInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ramInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) RxTxFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxTxFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) RxTxFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxTxFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) Swap() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) SwapInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) Vcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) VcpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpusInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/data-sources/compute_flavor_v2 opentelekomcloud_compute_flavor_v2} Data Source.
func NewDataOpentelekomcloudComputeFlavorV2(scope constructs.Construct, id *string, config *DataOpentelekomcloudComputeFlavorV2Config) DataOpentelekomcloudComputeFlavorV2 {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudComputeFlavorV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudComputeFlavorV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudComputeFlavorV2.DataOpentelekomcloudComputeFlavorV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/data-sources/compute_flavor_v2 opentelekomcloud_compute_flavor_v2} Data Source.
func NewDataOpentelekomcloudComputeFlavorV2_Override(d DataOpentelekomcloudComputeFlavorV2, scope constructs.Construct, id *string, config *DataOpentelekomcloudComputeFlavorV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudComputeFlavorV2.DataOpentelekomcloudComputeFlavorV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetDisk(val *float64) {
	if err := j.validateSetDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disk",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetFlavorId(val *string) {
	if err := j.validateSetFlavorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flavorId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetMinDisk(val *float64) {
	if err := j.validateSetMinDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDisk",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetMinRam(val *float64) {
	if err := j.validateSetMinRamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRam",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetRam(val *float64) {
	if err := j.validateSetRamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ram",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetRxTxFactor(val *float64) {
	if err := j.validateSetRxTxFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rxTxFactor",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetSwap(val *float64) {
	if err := j.validateSetSwapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"swap",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudComputeFlavorV2)SetVcpus(val *float64) {
	if err := j.validateSetVcpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcpus",
		val,
	)
}

// Generates CDKTF code for importing a DataOpentelekomcloudComputeFlavorV2 resource upon running "cdktf plan <stack-name>".
func DataOpentelekomcloudComputeFlavorV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudComputeFlavorV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudComputeFlavorV2.DataOpentelekomcloudComputeFlavorV2",
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
func DataOpentelekomcloudComputeFlavorV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudComputeFlavorV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudComputeFlavorV2.DataOpentelekomcloudComputeFlavorV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudComputeFlavorV2_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudComputeFlavorV2_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudComputeFlavorV2.DataOpentelekomcloudComputeFlavorV2",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudComputeFlavorV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudComputeFlavorV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudComputeFlavorV2.DataOpentelekomcloudComputeFlavorV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudComputeFlavorV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudComputeFlavorV2.DataOpentelekomcloudComputeFlavorV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetFlavorId() {
	_jsii_.InvokeVoid(
		d,
		"resetFlavorId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetMinDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetMinDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetMinRam() {
	_jsii_.InvokeVoid(
		d,
		"resetMinRam",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetRam() {
	_jsii_.InvokeVoid(
		d,
		"resetRam",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetResourceType() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetRxTxFactor() {
	_jsii_.InvokeVoid(
		d,
		"resetRxTxFactor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetSwap() {
	_jsii_.InvokeVoid(
		d,
		"resetSwap",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ResetVcpus() {
	_jsii_.InvokeVoid(
		d,
		"resetVcpus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudComputeFlavorV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

