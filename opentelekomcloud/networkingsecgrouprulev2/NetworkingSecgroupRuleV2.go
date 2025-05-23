// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkingsecgrouprulev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/networkingsecgrouprulev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/networking_secgroup_rule_v2 opentelekomcloud_networking_secgroup_rule_v2}.
type NetworkingSecgroupRuleV2 interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	Ethertype() *string
	SetEthertype(val *string)
	EthertypeInput() *string
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
	// The tree node.
	Node() constructs.Node
	PortRangeMax() *float64
	SetPortRangeMax(val *float64)
	PortRangeMaxInput() *float64
	PortRangeMin() *float64
	SetPortRangeMin(val *float64)
	PortRangeMinInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RemoteGroupId() *string
	SetRemoteGroupId(val *string)
	RemoteGroupIdInput() *string
	RemoteIpPrefix() *string
	SetRemoteIpPrefix(val *string)
	RemoteIpPrefixInput() *string
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetworkingSecgroupRuleV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *NetworkingSecgroupRuleV2Timeouts)
	ResetDescription()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortRangeMax()
	ResetPortRangeMin()
	ResetProtocol()
	ResetRegion()
	ResetRemoteGroupId()
	ResetRemoteIpPrefix()
	ResetTenantId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for NetworkingSecgroupRuleV2
type jsiiProxy_NetworkingSecgroupRuleV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Ethertype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ethertype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) EthertypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ethertypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) PortRangeMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) PortRangeMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) PortRangeMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) PortRangeMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRangeMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) RemoteGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) RemoteGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) RemoteIpPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) RemoteIpPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) Timeouts() NetworkingSecgroupRuleV2TimeoutsOutputReference {
	var returns NetworkingSecgroupRuleV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/networking_secgroup_rule_v2 opentelekomcloud_networking_secgroup_rule_v2} Resource.
func NewNetworkingSecgroupRuleV2(scope constructs.Construct, id *string, config *NetworkingSecgroupRuleV2Config) NetworkingSecgroupRuleV2 {
	_init_.Initialize()

	if err := validateNewNetworkingSecgroupRuleV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkingSecgroupRuleV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.networkingSecgroupRuleV2.NetworkingSecgroupRuleV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/networking_secgroup_rule_v2 opentelekomcloud_networking_secgroup_rule_v2} Resource.
func NewNetworkingSecgroupRuleV2_Override(n NetworkingSecgroupRuleV2, scope constructs.Construct, id *string, config *NetworkingSecgroupRuleV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.networkingSecgroupRuleV2.NetworkingSecgroupRuleV2",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetEthertype(val *string) {
	if err := j.validateSetEthertypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ethertype",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetPortRangeMax(val *float64) {
	if err := j.validateSetPortRangeMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRangeMax",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetPortRangeMin(val *float64) {
	if err := j.validateSetPortRangeMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRangeMin",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetRemoteGroupId(val *string) {
	if err := j.validateSetRemoteGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteGroupId",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetRemoteIpPrefix(val *string) {
	if err := j.validateSetRemoteIpPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteIpPrefix",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetSecurityGroupId(val *string) {
	if err := j.validateSetSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_NetworkingSecgroupRuleV2)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

// Generates CDKTF code for importing a NetworkingSecgroupRuleV2 resource upon running "cdktf plan <stack-name>".
func NetworkingSecgroupRuleV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkingSecgroupRuleV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.networkingSecgroupRuleV2.NetworkingSecgroupRuleV2",
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
func NetworkingSecgroupRuleV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkingSecgroupRuleV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.networkingSecgroupRuleV2.NetworkingSecgroupRuleV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkingSecgroupRuleV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkingSecgroupRuleV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.networkingSecgroupRuleV2.NetworkingSecgroupRuleV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkingSecgroupRuleV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkingSecgroupRuleV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.networkingSecgroupRuleV2.NetworkingSecgroupRuleV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkingSecgroupRuleV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.networkingSecgroupRuleV2.NetworkingSecgroupRuleV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) PutTimeouts(value *NetworkingSecgroupRuleV2Timeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ResetPortRangeMax() {
	_jsii_.InvokeVoid(
		n,
		"resetPortRangeMax",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ResetPortRangeMin() {
	_jsii_.InvokeVoid(
		n,
		"resetPortRangeMin",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ResetProtocol() {
	_jsii_.InvokeVoid(
		n,
		"resetProtocol",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ResetRegion() {
	_jsii_.InvokeVoid(
		n,
		"resetRegion",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ResetRemoteGroupId() {
	_jsii_.InvokeVoid(
		n,
		"resetRemoteGroupId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ResetRemoteIpPrefix() {
	_jsii_.InvokeVoid(
		n,
		"resetRemoteIpPrefix",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ResetTenantId() {
	_jsii_.InvokeVoid(
		n,
		"resetTenantId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingSecgroupRuleV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

