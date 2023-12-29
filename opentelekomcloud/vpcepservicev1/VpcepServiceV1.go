// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcepservicev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v10/vpcepservicev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/vpcep_service_v1 opentelekomcloud_vpcep_service_v1}.
type VpcepServiceV1 interface {
	cdktf.TerraformResource
	ApprovalEnabled() interface{}
	SetApprovalEnabled(val interface{})
	ApprovalEnabledInput() interface{}
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PoolId() *string
	SetPoolId(val *string)
	PoolIdInput() *string
	Port() VpcepServiceV1PortList
	PortId() *string
	SetPortId(val *string)
	PortIdInput() *string
	PortInput() interface{}
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
	ServerType() *string
	SetServerType(val *string)
	ServerTypeInput() *string
	ServiceType() *string
	SetServiceType(val *string)
	ServiceTypeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TcpProxy() *string
	SetTcpProxy(val *string)
	TcpProxyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VpcepServiceV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	VipPortId() *string
	SetVipPortId(val *string)
	VipPortIdInput() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	Whitelist() *[]*string
	SetWhitelist(val *[]*string)
	WhitelistInput() *[]*string
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
	PutPort(value interface{})
	PutTimeouts(value *VpcepServiceV1Timeouts)
	ResetApprovalEnabled()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPoolId()
	ResetServiceType()
	ResetTags()
	ResetTcpProxy()
	ResetTimeouts()
	ResetVipPortId()
	ResetWhitelist()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VpcepServiceV1
type jsiiProxy_VpcepServiceV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpcepServiceV1) ApprovalEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) ApprovalEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) PoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) PoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Port() VpcepServiceV1PortList {
	var returns VpcepServiceV1PortList
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) PortId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) PortIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) PortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) ServerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) ServerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) ServiceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) ServiceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) TcpProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tcpProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) TcpProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tcpProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Timeouts() VpcepServiceV1TimeoutsOutputReference {
	var returns VpcepServiceV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) VipPortId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vipPortId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) VipPortIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vipPortIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) Whitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepServiceV1) WhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/vpcep_service_v1 opentelekomcloud_vpcep_service_v1} Resource.
func NewVpcepServiceV1(scope constructs.Construct, id *string, config *VpcepServiceV1Config) VpcepServiceV1 {
	_init_.Initialize()

	if err := validateNewVpcepServiceV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcepServiceV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.vpcepServiceV1.VpcepServiceV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/vpcep_service_v1 opentelekomcloud_vpcep_service_v1} Resource.
func NewVpcepServiceV1_Override(v VpcepServiceV1, scope constructs.Construct, id *string, config *VpcepServiceV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.vpcepServiceV1.VpcepServiceV1",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetApprovalEnabled(val interface{}) {
	if err := j.validateSetApprovalEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalEnabled",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetPoolId(val *string) {
	if err := j.validateSetPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolId",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetPortId(val *string) {
	if err := j.validateSetPortIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portId",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetServerType(val *string) {
	if err := j.validateSetServerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverType",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetServiceType(val *string) {
	if err := j.validateSetServiceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceType",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetTcpProxy(val *string) {
	if err := j.validateSetTcpProxyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpProxy",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetVipPortId(val *string) {
	if err := j.validateSetVipPortIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vipPortId",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_VpcepServiceV1)SetWhitelist(val *[]*string) {
	if err := j.validateSetWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whitelist",
		val,
	)
}

// Generates CDKTF code for importing a VpcepServiceV1 resource upon running "cdktf plan <stack-name>".
func VpcepServiceV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVpcepServiceV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.vpcepServiceV1.VpcepServiceV1",
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
func VpcepServiceV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcepServiceV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.vpcepServiceV1.VpcepServiceV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcepServiceV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcepServiceV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.vpcepServiceV1.VpcepServiceV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcepServiceV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcepServiceV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.vpcepServiceV1.VpcepServiceV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpcepServiceV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.vpcepServiceV1.VpcepServiceV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpcepServiceV1) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VpcepServiceV1) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpcepServiceV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VpcepServiceV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpcepServiceV1) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VpcepServiceV1) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpcepServiceV1) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpcepServiceV1) PutPort(value interface{}) {
	if err := v.validatePutPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPort",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcepServiceV1) PutTimeouts(value *VpcepServiceV1Timeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcepServiceV1) ResetApprovalEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetApprovalEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepServiceV1) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepServiceV1) ResetName() {
	_jsii_.InvokeVoid(
		v,
		"resetName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepServiceV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepServiceV1) ResetPoolId() {
	_jsii_.InvokeVoid(
		v,
		"resetPoolId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepServiceV1) ResetServiceType() {
	_jsii_.InvokeVoid(
		v,
		"resetServiceType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepServiceV1) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepServiceV1) ResetTcpProxy() {
	_jsii_.InvokeVoid(
		v,
		"resetTcpProxy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepServiceV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepServiceV1) ResetVipPortId() {
	_jsii_.InvokeVoid(
		v,
		"resetVipPortId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepServiceV1) ResetWhitelist() {
	_jsii_.InvokeVoid(
		v,
		"resetWhitelist",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepServiceV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepServiceV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

