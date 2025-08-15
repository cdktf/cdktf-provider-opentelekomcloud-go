// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudnatdnatrulesv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudnatdnatrulesv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/data-sources/nat_dnat_rules_v2 opentelekomcloud_nat_dnat_rules_v2}.
type DataOpentelekomcloudNatDnatRulesV2 interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExternalServicePort() *float64
	SetExternalServicePort(val *float64)
	ExternalServicePortInput() *float64
	FloatingIpAddress() *string
	SetFloatingIpAddress(val *string)
	FloatingIpAddressInput() *string
	FloatingIpId() *string
	SetFloatingIpId(val *string)
	FloatingIpIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayId() *string
	SetGatewayId(val *string)
	GatewayIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalServicePort() *float64
	SetInternalServicePort(val *float64)
	InternalServicePortInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PortId() *string
	SetPortId(val *string)
	PortIdInput() *string
	PrivateIp() *string
	SetPrivateIp(val *string)
	PrivateIpInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	RuleId() *string
	SetRuleId(val *string)
	RuleIdInput() *string
	Rules() DataOpentelekomcloudNatDnatRulesV2RulesList
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
	ResetDescription()
	ResetExternalServicePort()
	ResetFloatingIpAddress()
	ResetFloatingIpId()
	ResetGatewayId()
	ResetId()
	ResetInternalServicePort()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortId()
	ResetPrivateIp()
	ResetProtocol()
	ResetRuleId()
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

// The jsii proxy struct for DataOpentelekomcloudNatDnatRulesV2
type jsiiProxy_DataOpentelekomcloudNatDnatRulesV2 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ExternalServicePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalServicePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ExternalServicePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalServicePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) FloatingIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"floatingIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) FloatingIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"floatingIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) FloatingIpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"floatingIpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) FloatingIpIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"floatingIpIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) GatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) InternalServicePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalServicePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) InternalServicePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalServicePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) PortId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) PortIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) RuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) RuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) Rules() DataOpentelekomcloudNatDnatRulesV2RulesList {
	var returns DataOpentelekomcloudNatDnatRulesV2RulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/data-sources/nat_dnat_rules_v2 opentelekomcloud_nat_dnat_rules_v2} Data Source.
func NewDataOpentelekomcloudNatDnatRulesV2(scope constructs.Construct, id *string, config *DataOpentelekomcloudNatDnatRulesV2Config) DataOpentelekomcloudNatDnatRulesV2 {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudNatDnatRulesV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudNatDnatRulesV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNatDnatRulesV2.DataOpentelekomcloudNatDnatRulesV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/data-sources/nat_dnat_rules_v2 opentelekomcloud_nat_dnat_rules_v2} Data Source.
func NewDataOpentelekomcloudNatDnatRulesV2_Override(d DataOpentelekomcloudNatDnatRulesV2, scope constructs.Construct, id *string, config *DataOpentelekomcloudNatDnatRulesV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNatDnatRulesV2.DataOpentelekomcloudNatDnatRulesV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetExternalServicePort(val *float64) {
	if err := j.validateSetExternalServicePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalServicePort",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetFloatingIpAddress(val *string) {
	if err := j.validateSetFloatingIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"floatingIpAddress",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetFloatingIpId(val *string) {
	if err := j.validateSetFloatingIpIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"floatingIpId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetGatewayId(val *string) {
	if err := j.validateSetGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetInternalServicePort(val *float64) {
	if err := j.validateSetInternalServicePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalServicePort",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetPortId(val *string) {
	if err := j.validateSetPortIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetRuleId(val *string) {
	if err := j.validateSetRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

// Generates CDKTF code for importing a DataOpentelekomcloudNatDnatRulesV2 resource upon running "cdktf plan <stack-name>".
func DataOpentelekomcloudNatDnatRulesV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudNatDnatRulesV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNatDnatRulesV2.DataOpentelekomcloudNatDnatRulesV2",
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
func DataOpentelekomcloudNatDnatRulesV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudNatDnatRulesV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNatDnatRulesV2.DataOpentelekomcloudNatDnatRulesV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudNatDnatRulesV2_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudNatDnatRulesV2_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNatDnatRulesV2.DataOpentelekomcloudNatDnatRulesV2",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudNatDnatRulesV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudNatDnatRulesV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNatDnatRulesV2.DataOpentelekomcloudNatDnatRulesV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudNatDnatRulesV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudNatDnatRulesV2.DataOpentelekomcloudNatDnatRulesV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetExternalServicePort() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalServicePort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetFloatingIpAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetFloatingIpAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetFloatingIpId() {
	_jsii_.InvokeVoid(
		d,
		"resetFloatingIpId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetInternalServicePort() {
	_jsii_.InvokeVoid(
		d,
		"resetInternalServicePort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetPortId() {
	_jsii_.InvokeVoid(
		d,
		"resetPortId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetRuleId() {
	_jsii_.InvokeVoid(
		d,
		"resetRuleId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudNatDnatRulesV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

