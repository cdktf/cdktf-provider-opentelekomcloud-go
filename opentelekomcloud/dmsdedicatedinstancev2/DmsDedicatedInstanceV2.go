// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsdedicatedinstancev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/dmsdedicatedinstancev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/dms_dedicated_instance_v2 opentelekomcloud_dms_dedicated_instance_v2}.
type DmsDedicatedInstanceV2 interface {
	cdktf.TerraformResource
	AccessUser() *string
	SetAccessUser(val *string)
	AccessUserInput() *string
	ArchType() *string
	SetArchType(val *string)
	ArchTypeInput() *string
	AvailableZones() *[]*string
	SetAvailableZones(val *[]*string)
	AvailableZonesInput() *[]*string
	Bandwidth() *string
	BrokerNum() *float64
	SetBrokerNum(val *float64)
	BrokerNumInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertReplaced() cdktf.IResolvable
	ConnectAddress() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectorNodeNum() *float64
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	CrossVpcAccesses() DmsDedicatedInstanceV2CrossVpcAccessesList
	CrossVpcAccessesInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DiskEncryptedEnable() interface{}
	SetDiskEncryptedEnable(val interface{})
	DiskEncryptedEnableInput() interface{}
	DiskEncryptedKey() *string
	SetDiskEncryptedKey(val *string)
	DiskEncryptedKeyInput() *string
	Dumping() cdktf.IResolvable
	EnabledMechanisms() *[]*string
	SetEnabledMechanisms(val *[]*string)
	EnabledMechanismsInput() *[]*string
	EnablePublicip() interface{}
	SetEnablePublicip(val interface{})
	EnablePublicipInput() interface{}
	Engine() *string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
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
	Ipv6Enable() interface{}
	SetIpv6Enable(val interface{})
	Ipv6EnableInput() interface{}
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
	NewTenantIps() *[]*string
	SetNewTenantIps(val *[]*string)
	NewTenantIpsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	NodeNum() *float64
	PartitionNum() *float64
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PodConnectAddress() *string
	Port() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicBandwidth() *float64
	PublicIpAddress() *string
	PublicipId() *[]*string
	SetPublicipId(val *[]*string)
	PublicipIdInput() *[]*string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	ResourceSpecCode() *string
	RetentionPolicy() *string
	SetRetentionPolicy(val *string)
	RetentionPolicyInput() *string
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	SecurityProtocol() *string
	SetSecurityProtocol(val *string)
	SecurityProtocolInput() *string
	SslEnable() interface{}
	SetSslEnable(val interface{})
	SslEnableInput() interface{}
	SslTwoWayEnable() cdktf.IResolvable
	Status() *string
	StorageResourceId() *string
	StorageSpace() *float64
	SetStorageSpace(val *float64)
	StorageSpaceInput() *float64
	StorageSpecCode() *string
	SetStorageSpecCode(val *string)
	StorageSpecCodeInput() *string
	StorageType() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DmsDedicatedInstanceV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	UsedStorageSpace() *float64
	UserId() *string
	UserName() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	PutCrossVpcAccesses(value interface{})
	PutTimeouts(value *DmsDedicatedInstanceV2Timeouts)
	ResetAccessUser()
	ResetArchType()
	ResetAvailableZones()
	ResetCrossVpcAccesses()
	ResetDescription()
	ResetDiskEncryptedEnable()
	ResetDiskEncryptedKey()
	ResetEnabledMechanisms()
	ResetEnablePublicip()
	ResetId()
	ResetIpv6Enable()
	ResetMaintainBegin()
	ResetMaintainEnd()
	ResetNewTenantIps()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPublicipId()
	ResetRetentionPolicy()
	ResetSecurityProtocol()
	ResetSslEnable()
	ResetTags()
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

// The jsii proxy struct for DmsDedicatedInstanceV2
type jsiiProxy_DmsDedicatedInstanceV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) AccessUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) AccessUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) ArchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) ArchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) AvailableZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) AvailableZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Bandwidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) BrokerNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"brokerNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) BrokerNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"brokerNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) CertReplaced() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"certReplaced",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) ConnectAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) ConnectorNodeNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectorNodeNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) CrossVpcAccesses() DmsDedicatedInstanceV2CrossVpcAccessesList {
	var returns DmsDedicatedInstanceV2CrossVpcAccessesList
	_jsii_.Get(
		j,
		"crossVpcAccesses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) CrossVpcAccessesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossVpcAccessesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) DiskEncryptedEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskEncryptedEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) DiskEncryptedEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskEncryptedEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) DiskEncryptedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) DiskEncryptedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Dumping() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"dumping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) EnabledMechanisms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledMechanisms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) EnabledMechanismsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledMechanismsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) EnablePublicip() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePublicip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) EnablePublicipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePublicipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) FlavorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) FlavorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Ipv6Enable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Ipv6EnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6EnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) MaintainBegin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainBegin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) MaintainBeginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainBeginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) MaintainEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) MaintainEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) NewTenantIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newTenantIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) NewTenantIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newTenantIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) NodeNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) PartitionNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) PodConnectAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podConnectAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) PublicBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"publicBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) PublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) PublicipId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicipId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) PublicipIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicipIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) ResourceSpecCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceSpecCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) RetentionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) RetentionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) SslEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) SslEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) SslTwoWayEnable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"sslTwoWayEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) StorageResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) StorageSpace() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) StorageSpaceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) StorageSpecCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSpecCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) StorageSpecCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSpecCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Timeouts() DmsDedicatedInstanceV2TimeoutsOutputReference {
	var returns DmsDedicatedInstanceV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) UsedStorageSpace() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedStorageSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDedicatedInstanceV2) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/dms_dedicated_instance_v2 opentelekomcloud_dms_dedicated_instance_v2} Resource.
func NewDmsDedicatedInstanceV2(scope constructs.Construct, id *string, config *DmsDedicatedInstanceV2Config) DmsDedicatedInstanceV2 {
	_init_.Initialize()

	if err := validateNewDmsDedicatedInstanceV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsDedicatedInstanceV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dmsDedicatedInstanceV2.DmsDedicatedInstanceV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/dms_dedicated_instance_v2 opentelekomcloud_dms_dedicated_instance_v2} Resource.
func NewDmsDedicatedInstanceV2_Override(d DmsDedicatedInstanceV2, scope constructs.Construct, id *string, config *DmsDedicatedInstanceV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dmsDedicatedInstanceV2.DmsDedicatedInstanceV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetAccessUser(val *string) {
	if err := j.validateSetAccessUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessUser",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetArchType(val *string) {
	if err := j.validateSetArchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archType",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetAvailableZones(val *[]*string) {
	if err := j.validateSetAvailableZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableZones",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetBrokerNum(val *float64) {
	if err := j.validateSetBrokerNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"brokerNum",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetDiskEncryptedEnable(val interface{}) {
	if err := j.validateSetDiskEncryptedEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptedEnable",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetDiskEncryptedKey(val *string) {
	if err := j.validateSetDiskEncryptedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptedKey",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetEnabledMechanisms(val *[]*string) {
	if err := j.validateSetEnabledMechanismsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledMechanisms",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetEnablePublicip(val interface{}) {
	if err := j.validateSetEnablePublicipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePublicip",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetFlavorId(val *string) {
	if err := j.validateSetFlavorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flavorId",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetIpv6Enable(val interface{}) {
	if err := j.validateSetIpv6EnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Enable",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetMaintainBegin(val *string) {
	if err := j.validateSetMaintainBeginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainBegin",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetMaintainEnd(val *string) {
	if err := j.validateSetMaintainEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainEnd",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetNetworkId(val *string) {
	if err := j.validateSetNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetNewTenantIps(val *[]*string) {
	if err := j.validateSetNewTenantIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newTenantIps",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetPublicipId(val *[]*string) {
	if err := j.validateSetPublicipIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicipId",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetRetentionPolicy(val *string) {
	if err := j.validateSetRetentionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPolicy",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetSecurityGroupId(val *string) {
	if err := j.validateSetSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetSecurityProtocol(val *string) {
	if err := j.validateSetSecurityProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetSslEnable(val interface{}) {
	if err := j.validateSetSslEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslEnable",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetStorageSpace(val *float64) {
	if err := j.validateSetStorageSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSpace",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetStorageSpecCode(val *string) {
	if err := j.validateSetStorageSpecCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSpecCode",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DmsDedicatedInstanceV2)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a DmsDedicatedInstanceV2 resource upon running "cdktf plan <stack-name>".
func DmsDedicatedInstanceV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDmsDedicatedInstanceV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dmsDedicatedInstanceV2.DmsDedicatedInstanceV2",
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
func DmsDedicatedInstanceV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsDedicatedInstanceV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dmsDedicatedInstanceV2.DmsDedicatedInstanceV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsDedicatedInstanceV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsDedicatedInstanceV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dmsDedicatedInstanceV2.DmsDedicatedInstanceV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsDedicatedInstanceV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsDedicatedInstanceV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dmsDedicatedInstanceV2.DmsDedicatedInstanceV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsDedicatedInstanceV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dmsDedicatedInstanceV2.DmsDedicatedInstanceV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsDedicatedInstanceV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsDedicatedInstanceV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsDedicatedInstanceV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsDedicatedInstanceV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsDedicatedInstanceV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsDedicatedInstanceV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsDedicatedInstanceV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsDedicatedInstanceV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsDedicatedInstanceV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsDedicatedInstanceV2) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) PutCrossVpcAccesses(value interface{}) {
	if err := d.validatePutCrossVpcAccessesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCrossVpcAccesses",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) PutTimeouts(value *DmsDedicatedInstanceV2Timeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetAccessUser() {
	_jsii_.InvokeVoid(
		d,
		"resetAccessUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetArchType() {
	_jsii_.InvokeVoid(
		d,
		"resetArchType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetAvailableZones() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailableZones",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetCrossVpcAccesses() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossVpcAccesses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetDiskEncryptedEnable() {
	_jsii_.InvokeVoid(
		d,
		"resetDiskEncryptedEnable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetDiskEncryptedKey() {
	_jsii_.InvokeVoid(
		d,
		"resetDiskEncryptedKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetEnabledMechanisms() {
	_jsii_.InvokeVoid(
		d,
		"resetEnabledMechanisms",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetEnablePublicip() {
	_jsii_.InvokeVoid(
		d,
		"resetEnablePublicip",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetIpv6Enable() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6Enable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetMaintainBegin() {
	_jsii_.InvokeVoid(
		d,
		"resetMaintainBegin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetMaintainEnd() {
	_jsii_.InvokeVoid(
		d,
		"resetMaintainEnd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetNewTenantIps() {
	_jsii_.InvokeVoid(
		d,
		"resetNewTenantIps",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetPublicipId() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicipId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetRetentionPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetSslEnable() {
	_jsii_.InvokeVoid(
		d,
		"resetSslEnable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDedicatedInstanceV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

