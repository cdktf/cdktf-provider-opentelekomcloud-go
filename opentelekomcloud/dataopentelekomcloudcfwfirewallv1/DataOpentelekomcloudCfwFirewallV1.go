// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudcfwfirewallv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudcfwfirewallv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/cfw_firewall_v1 opentelekomcloud_cfw_firewall_v1}.
type DataOpentelekomcloudCfwFirewallV1 interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChargeMode() *float64
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
	EngineType() *float64
	EnterpriseProjectId() *string
	FeatureToggle() cdktf.BooleanMap
	Flavor() DataOpentelekomcloudCfwFirewallV1FlavorList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HaType() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsAvailableObs() cdktf.IResolvable
	IsOldFirewallInstance() cdktf.IResolvable
	IsSupportThreatTags() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	ProtectObjects() DataOpentelekomcloudCfwFirewallV1ProtectObjectsList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ResourceId() *string
	Resources() DataOpentelekomcloudCfwFirewallV1ResourcesList
	ServiceType() *string
	SetServiceType(val *string)
	ServiceTypeInput() *string
	Status() *float64
	SupportIpv6() cdktf.IResolvable
	SupportUrlFiltering() cdktf.IResolvable
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceType()
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

// The jsii proxy struct for DataOpentelekomcloudCfwFirewallV1
type jsiiProxy_DataOpentelekomcloudCfwFirewallV1 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ChargeMode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"chargeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) EngineType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) EnterpriseProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) FeatureToggle() cdktf.BooleanMap {
	var returns cdktf.BooleanMap
	_jsii_.Get(
		j,
		"featureToggle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) Flavor() DataOpentelekomcloudCfwFirewallV1FlavorList {
	var returns DataOpentelekomcloudCfwFirewallV1FlavorList
	_jsii_.Get(
		j,
		"flavor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) HaType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) IsAvailableObs() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isAvailableObs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) IsOldFirewallInstance() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isOldFirewallInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) IsSupportThreatTags() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isSupportThreatTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ProtectObjects() DataOpentelekomcloudCfwFirewallV1ProtectObjectsList {
	var returns DataOpentelekomcloudCfwFirewallV1ProtectObjectsList
	_jsii_.Get(
		j,
		"protectObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) Resources() DataOpentelekomcloudCfwFirewallV1ResourcesList {
	var returns DataOpentelekomcloudCfwFirewallV1ResourcesList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ServiceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ServiceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) Status() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) SupportIpv6() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"supportIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) SupportUrlFiltering() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"supportUrlFiltering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/cfw_firewall_v1 opentelekomcloud_cfw_firewall_v1} Data Source.
func NewDataOpentelekomcloudCfwFirewallV1(scope constructs.Construct, id *string, config *DataOpentelekomcloudCfwFirewallV1Config) DataOpentelekomcloudCfwFirewallV1 {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudCfwFirewallV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudCfwFirewallV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCfwFirewallV1.DataOpentelekomcloudCfwFirewallV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/cfw_firewall_v1 opentelekomcloud_cfw_firewall_v1} Data Source.
func NewDataOpentelekomcloudCfwFirewallV1_Override(d DataOpentelekomcloudCfwFirewallV1, scope constructs.Construct, id *string, config *DataOpentelekomcloudCfwFirewallV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCfwFirewallV1.DataOpentelekomcloudCfwFirewallV1",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCfwFirewallV1)SetServiceType(val *string) {
	if err := j.validateSetServiceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceType",
		val,
	)
}

// Generates CDKTF code for importing a DataOpentelekomcloudCfwFirewallV1 resource upon running "cdktf plan <stack-name>".
func DataOpentelekomcloudCfwFirewallV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudCfwFirewallV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCfwFirewallV1.DataOpentelekomcloudCfwFirewallV1",
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
func DataOpentelekomcloudCfwFirewallV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudCfwFirewallV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCfwFirewallV1.DataOpentelekomcloudCfwFirewallV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudCfwFirewallV1_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudCfwFirewallV1_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCfwFirewallV1.DataOpentelekomcloudCfwFirewallV1",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudCfwFirewallV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudCfwFirewallV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCfwFirewallV1.DataOpentelekomcloudCfwFirewallV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudCfwFirewallV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCfwFirewallV1.DataOpentelekomcloudCfwFirewallV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ResetServiceType() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCfwFirewallV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

