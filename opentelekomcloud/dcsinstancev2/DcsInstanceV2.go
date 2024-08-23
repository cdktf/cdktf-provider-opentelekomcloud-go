// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcsinstancev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/dcsinstancev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.17/docs/resources/dcs_instance_v2 opentelekomcloud_dcs_instance_v2}.
type DcsInstanceV2 interface {
	cdktf.TerraformResource
	AccessUser() *string
	SetAccessUser(val *string)
	AccessUserInput() *string
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	BackupPolicy() DcsInstanceV2BackupPolicyOutputReference
	BackupPolicyInput() *DcsInstanceV2BackupPolicy
	BandwidthInfo() DcsInstanceV2BandwidthInfoList
	CacheMode() *string
	Capacity() *float64
	SetCapacity(val *float64)
	CapacityInput() *float64
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
	CpuType() *string
	CreatedAt() *string
	DeletedNodes() *[]*string
	SetDeletedNodes(val *[]*string)
	DeletedNodesInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DomainName() *string
	EnableWhitelist() interface{}
	SetEnableWhitelist(val interface{})
	EnableWhitelistInput() interface{}
	Engine() *string
	SetEngine(val *string)
	EngineInput() *string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
	Flavor() *string
	SetFlavor(val *string)
	FlavorInput() *string
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
	LaunchedAt() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintainBegin() *string
	SetMaintainBegin(val *string)
	MaintainBeginInput() *string
	MaintainEnd() *string
	SetMaintainEnd(val *string)
	MaintainEndInput() *string
	MaxMemory() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Parameters() DcsInstanceV2ParametersList
	ParametersInput() interface{}
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PrivateIp() *string
	SetPrivateIp(val *string)
	PrivateIpInput() *string
	ProductType() *string
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
	ReadonlyDomainName() *string
	Region() *string
	RenameCommands() *map[string]*string
	SetRenameCommands(val *map[string]*string)
	RenameCommandsInput() *map[string]*string
	ReplicaCount() *float64
	ReservedIps() *[]*string
	SetReservedIps(val *[]*string)
	ReservedIpsInput() *[]*string
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	SecurityGroupName() *string
	ShardingCount() *float64
	SslEnable() interface{}
	SetSslEnable(val interface{})
	SslEnableInput() interface{}
	Status() *string
	SubnetCidr() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	SubnetName() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TemplateId() *string
	SetTemplateId(val *string)
	TemplateIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DcsInstanceV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	TransparentClientIpEnable() cdktf.IResolvable
	UsedMemory() *float64
	UserId() *string
	UserName() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	VpcName() *string
	Whitelist() DcsInstanceV2WhitelistStructList
	WhitelistInput() interface{}
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
	PutBackupPolicy(value *DcsInstanceV2BackupPolicy)
	PutParameters(value interface{})
	PutTimeouts(value *DcsInstanceV2Timeouts)
	PutWhitelist(value interface{})
	ResetAccessUser()
	ResetBackupPolicy()
	ResetDeletedNodes()
	ResetDescription()
	ResetEnableWhitelist()
	ResetEngineVersion()
	ResetId()
	ResetMaintainBegin()
	ResetMaintainEnd()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetPassword()
	ResetPort()
	ResetPrivateIp()
	ResetRenameCommands()
	ResetReservedIps()
	ResetSecurityGroupId()
	ResetSslEnable()
	ResetTags()
	ResetTemplateId()
	ResetTimeouts()
	ResetWhitelist()
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

// The jsii proxy struct for DcsInstanceV2
type jsiiProxy_DcsInstanceV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DcsInstanceV2) AccessUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) AccessUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) BackupPolicy() DcsInstanceV2BackupPolicyOutputReference {
	var returns DcsInstanceV2BackupPolicyOutputReference
	_jsii_.Get(
		j,
		"backupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) BackupPolicyInput() *DcsInstanceV2BackupPolicy {
	var returns *DcsInstanceV2BackupPolicy
	_jsii_.Get(
		j,
		"backupPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) BandwidthInfo() DcsInstanceV2BandwidthInfoList {
	var returns DcsInstanceV2BandwidthInfoList
	_jsii_.Get(
		j,
		"bandwidthInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) CacheMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) CapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) CpuType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) DeletedNodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deletedNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) DeletedNodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deletedNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) EnableWhitelist() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) EnableWhitelistInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Flavor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) FlavorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) LaunchedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) MaintainBegin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainBegin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) MaintainBeginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainBeginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) MaintainEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) MaintainEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) MaxMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Parameters() DcsInstanceV2ParametersList {
	var returns DcsInstanceV2ParametersList
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) ProductType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) ReadonlyDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readonlyDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) RenameCommands() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"renameCommands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) RenameCommandsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"renameCommandsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) ReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) ReservedIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) ReservedIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) SecurityGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) ShardingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shardingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) SslEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) SslEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) SubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) SubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) TemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) TemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Timeouts() DcsInstanceV2TimeoutsOutputReference {
	var returns DcsInstanceV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) TransparentClientIpEnable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"transparentClientIpEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) UsedMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) VpcName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) Whitelist() DcsInstanceV2WhitelistStructList {
	var returns DcsInstanceV2WhitelistStructList
	_jsii_.Get(
		j,
		"whitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2) WhitelistInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"whitelistInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.17/docs/resources/dcs_instance_v2 opentelekomcloud_dcs_instance_v2} Resource.
func NewDcsInstanceV2(scope constructs.Construct, id *string, config *DcsInstanceV2Config) DcsInstanceV2 {
	_init_.Initialize()

	if err := validateNewDcsInstanceV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DcsInstanceV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV2.DcsInstanceV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.17/docs/resources/dcs_instance_v2 opentelekomcloud_dcs_instance_v2} Resource.
func NewDcsInstanceV2_Override(d DcsInstanceV2, scope constructs.Construct, id *string, config *DcsInstanceV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV2.DcsInstanceV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetAccessUser(val *string) {
	if err := j.validateSetAccessUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessUser",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetCapacity(val *float64) {
	if err := j.validateSetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetDeletedNodes(val *[]*string) {
	if err := j.validateSetDeletedNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletedNodes",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetEnableWhitelist(val interface{}) {
	if err := j.validateSetEnableWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableWhitelist",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetFlavor(val *string) {
	if err := j.validateSetFlavorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flavor",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetMaintainBegin(val *string) {
	if err := j.validateSetMaintainBeginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainBegin",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetMaintainEnd(val *string) {
	if err := j.validateSetMaintainEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainEnd",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetRenameCommands(val *map[string]*string) {
	if err := j.validateSetRenameCommandsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renameCommands",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetReservedIps(val *[]*string) {
	if err := j.validateSetReservedIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedIps",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetSecurityGroupId(val *string) {
	if err := j.validateSetSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetSslEnable(val interface{}) {
	if err := j.validateSetSslEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslEnable",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetTemplateId(val *string) {
	if err := j.validateSetTemplateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateId",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a DcsInstanceV2 resource upon running "cdktf plan <stack-name>".
func DcsInstanceV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDcsInstanceV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV2.DcsInstanceV2",
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
func DcsInstanceV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDcsInstanceV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV2.DcsInstanceV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DcsInstanceV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDcsInstanceV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV2.DcsInstanceV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DcsInstanceV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDcsInstanceV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV2.DcsInstanceV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DcsInstanceV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV2.DcsInstanceV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DcsInstanceV2) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DcsInstanceV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DcsInstanceV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DcsInstanceV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DcsInstanceV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DcsInstanceV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DcsInstanceV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DcsInstanceV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DcsInstanceV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DcsInstanceV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DcsInstanceV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DcsInstanceV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DcsInstanceV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DcsInstanceV2) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DcsInstanceV2) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DcsInstanceV2) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DcsInstanceV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DcsInstanceV2) PutBackupPolicy(value *DcsInstanceV2BackupPolicy) {
	if err := d.validatePutBackupPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBackupPolicy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DcsInstanceV2) PutParameters(value interface{}) {
	if err := d.validatePutParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putParameters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DcsInstanceV2) PutTimeouts(value *DcsInstanceV2Timeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DcsInstanceV2) PutWhitelist(value interface{}) {
	if err := d.validatePutWhitelistParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWhitelist",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetAccessUser() {
	_jsii_.InvokeVoid(
		d,
		"resetAccessUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetBackupPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetDeletedNodes() {
	_jsii_.InvokeVoid(
		d,
		"resetDeletedNodes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetEnableWhitelist() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableWhitelist",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetMaintainBegin() {
	_jsii_.InvokeVoid(
		d,
		"resetMaintainBegin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetMaintainEnd() {
	_jsii_.InvokeVoid(
		d,
		"resetMaintainEnd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetRenameCommands() {
	_jsii_.InvokeVoid(
		d,
		"resetRenameCommands",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetReservedIps() {
	_jsii_.InvokeVoid(
		d,
		"resetReservedIps",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetSecurityGroupId() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityGroupId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetSslEnable() {
	_jsii_.InvokeVoid(
		d,
		"resetSslEnable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetTemplateId() {
	_jsii_.InvokeVoid(
		d,
		"resetTemplateId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) ResetWhitelist() {
	_jsii_.InvokeVoid(
		d,
		"resetWhitelist",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

