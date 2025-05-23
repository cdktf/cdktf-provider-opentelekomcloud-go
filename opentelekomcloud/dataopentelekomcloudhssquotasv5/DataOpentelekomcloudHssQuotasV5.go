// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudhssquotasv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/dataopentelekomcloudhssquotasv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/data-sources/hss_quotas_v5 opentelekomcloud_hss_quotas_v5}.
type DataOpentelekomcloudHssQuotasV5 interface {
	cdktf.TerraformDataSource
	Category() *string
	SetCategory(val *string)
	CategoryInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChargingMode() *string
	SetChargingMode(val *string)
	ChargingModeInput() *string
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
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
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	Quotas() DataOpentelekomcloudHssQuotasV5QuotasList
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UsedStatus() *string
	SetUsedStatus(val *string)
	UsedStatusInput() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetCategory()
	ResetChargingMode()
	ResetHostName()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceId()
	ResetStatus()
	ResetUsedStatus()
	ResetVersion()
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

// The jsii proxy struct for DataOpentelekomcloudHssQuotasV5
type jsiiProxy_DataOpentelekomcloudHssQuotasV5 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) CategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"categoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ChargingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ChargingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) Quotas() DataOpentelekomcloudHssQuotasV5QuotasList {
	var returns DataOpentelekomcloudHssQuotasV5QuotasList
	_jsii_.Get(
		j,
		"quotas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) UsedStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usedStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) UsedStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usedStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/data-sources/hss_quotas_v5 opentelekomcloud_hss_quotas_v5} Data Source.
func NewDataOpentelekomcloudHssQuotasV5(scope constructs.Construct, id *string, config *DataOpentelekomcloudHssQuotasV5Config) DataOpentelekomcloudHssQuotasV5 {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudHssQuotasV5Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudHssQuotasV5{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssQuotasV5.DataOpentelekomcloudHssQuotasV5",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/data-sources/hss_quotas_v5 opentelekomcloud_hss_quotas_v5} Data Source.
func NewDataOpentelekomcloudHssQuotasV5_Override(d DataOpentelekomcloudHssQuotasV5, scope constructs.Construct, id *string, config *DataOpentelekomcloudHssQuotasV5Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssQuotasV5.DataOpentelekomcloudHssQuotasV5",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetCategory(val *string) {
	if err := j.validateSetCategoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"category",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetChargingMode(val *string) {
	if err := j.validateSetChargingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chargingMode",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetHostName(val *string) {
	if err := j.validateSetHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetUsedStatus(val *string) {
	if err := j.validateSetUsedStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usedStatus",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssQuotasV5)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a DataOpentelekomcloudHssQuotasV5 resource upon running "cdktf plan <stack-name>".
func DataOpentelekomcloudHssQuotasV5_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudHssQuotasV5_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssQuotasV5.DataOpentelekomcloudHssQuotasV5",
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
func DataOpentelekomcloudHssQuotasV5_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudHssQuotasV5_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssQuotasV5.DataOpentelekomcloudHssQuotasV5",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudHssQuotasV5_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudHssQuotasV5_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssQuotasV5.DataOpentelekomcloudHssQuotasV5",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudHssQuotasV5_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudHssQuotasV5_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssQuotasV5.DataOpentelekomcloudHssQuotasV5",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudHssQuotasV5_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssQuotasV5.DataOpentelekomcloudHssQuotasV5",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ResetCategory() {
	_jsii_.InvokeVoid(
		d,
		"resetCategory",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ResetChargingMode() {
	_jsii_.InvokeVoid(
		d,
		"resetChargingMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ResetHostName() {
	_jsii_.InvokeVoid(
		d,
		"resetHostName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ResetResourceId() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ResetUsedStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetUsedStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssQuotasV5) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

