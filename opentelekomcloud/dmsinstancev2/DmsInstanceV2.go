// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsinstancev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/dmsinstancev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.16/docs/resources/dms_instance_v2 opentelekomcloud_dms_instance_v2}.
type DmsInstanceV2 interface {
	cdktf.TerraformResource
	AccessUser() *string
	SetAccessUser(val *string)
	AccessUserInput() *string
	AvailableZones() *[]*string
	SetAvailableZones(val *[]*string)
	AvailableZonesInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConnectAddress() *string
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
	CreatedAt() *string
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
	EnablePublicip() interface{}
	SetEnablePublicip(val interface{})
	EnablePublicipInput() interface{}
	Engine() *string
	SetEngine(val *string)
	EngineInput() *string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
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
	MaintainBegin() *string
	SetMaintainBegin(val *string)
	MaintainBeginInput() *string
	MaintainEnd() *string
	SetMaintainEnd(val *string)
	MaintainEndInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeNum() *float64
	PartitionNum() *float64
	SetPartitionNum(val *float64)
	PartitionNumInput() *float64
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	ProductId() *string
	SetProductId(val *string)
	ProductIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicAccessEnabled() *string
	PublicBandwidth() *float64
	SetPublicBandwidth(val *float64)
	PublicBandwidthInput() *float64
	PublicConnectAddress() *[]*string
	PublicipId() *[]*string
	SetPublicipId(val *[]*string)
	PublicipIdInput() *[]*string
	// Experimental.
	RawOverrides() interface{}
	ResourceSpecCode() *string
	RetentionPolicy() *string
	SetRetentionPolicy(val *string)
	RetentionPolicyInput() *string
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	SecurityGroupName() *string
	Specification() *string
	SetSpecification(val *string)
	SpecificationInput() *string
	SslEnable() cdktf.IResolvable
	Status() *string
	StorageResourceId() *string
	StorageSpace() *float64
	SetStorageSpace(val *float64)
	StorageSpaceInput() *float64
	StorageSpecCode() *string
	SetStorageSpecCode(val *string)
	StorageSpecCodeInput() *string
	SubnetCidr() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	SubnetName() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DmsInstanceV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	TotalStorageSpace() *float64
	Type() *string
	UsedStorageSpace() *float64
	UserId() *string
	UserName() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	VpcName() *string
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
	PutTimeouts(value *DmsInstanceV2Timeouts)
	ResetAccessUser()
	ResetDescription()
	ResetDiskEncryptedEnable()
	ResetDiskEncryptedKey()
	ResetEnablePublicip()
	ResetId()
	ResetMaintainBegin()
	ResetMaintainEnd()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartitionNum()
	ResetPassword()
	ResetPublicBandwidth()
	ResetPublicipId()
	ResetRetentionPolicy()
	ResetSpecification()
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

// The jsii proxy struct for DmsInstanceV2
type jsiiProxy_DmsInstanceV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsInstanceV2) AccessUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) AccessUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) AvailableZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) AvailableZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) ConnectAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) DiskEncryptedEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskEncryptedEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) DiskEncryptedEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskEncryptedEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) DiskEncryptedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) DiskEncryptedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) EnablePublicip() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePublicip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) EnablePublicipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePublicipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) MaintainBegin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainBegin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) MaintainBeginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainBeginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) MaintainEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) MaintainEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) NodeNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) PartitionNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) PartitionNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) PublicAccessEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) PublicBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"publicBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) PublicBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"publicBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) PublicConnectAddress() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicConnectAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) PublicipId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicipId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) PublicipIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicipIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) ResourceSpecCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceSpecCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) RetentionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) RetentionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) SecurityGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Specification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) SpecificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) SslEnable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"sslEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) StorageResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) StorageSpace() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) StorageSpaceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) StorageSpecCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSpecCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) StorageSpecCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSpecCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) SubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) SubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Timeouts() DmsInstanceV2TimeoutsOutputReference {
	var returns DmsInstanceV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) TotalStorageSpace() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalStorageSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) UsedStorageSpace() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedStorageSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsInstanceV2) VpcName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcName",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.16/docs/resources/dms_instance_v2 opentelekomcloud_dms_instance_v2} Resource.
func NewDmsInstanceV2(scope constructs.Construct, id *string, config *DmsInstanceV2Config) DmsInstanceV2 {
	_init_.Initialize()

	if err := validateNewDmsInstanceV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsInstanceV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dmsInstanceV2.DmsInstanceV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.16/docs/resources/dms_instance_v2 opentelekomcloud_dms_instance_v2} Resource.
func NewDmsInstanceV2_Override(d DmsInstanceV2, scope constructs.Construct, id *string, config *DmsInstanceV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dmsInstanceV2.DmsInstanceV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetAccessUser(val *string) {
	if err := j.validateSetAccessUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessUser",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetAvailableZones(val *[]*string) {
	if err := j.validateSetAvailableZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableZones",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetDiskEncryptedEnable(val interface{}) {
	if err := j.validateSetDiskEncryptedEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptedEnable",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetDiskEncryptedKey(val *string) {
	if err := j.validateSetDiskEncryptedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptedKey",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetEnablePublicip(val interface{}) {
	if err := j.validateSetEnablePublicipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePublicip",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetMaintainBegin(val *string) {
	if err := j.validateSetMaintainBeginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainBegin",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetMaintainEnd(val *string) {
	if err := j.validateSetMaintainEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainEnd",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetPartitionNum(val *float64) {
	if err := j.validateSetPartitionNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionNum",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetProductId(val *string) {
	if err := j.validateSetProductIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetPublicBandwidth(val *float64) {
	if err := j.validateSetPublicBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicBandwidth",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetPublicipId(val *[]*string) {
	if err := j.validateSetPublicipIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicipId",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetRetentionPolicy(val *string) {
	if err := j.validateSetRetentionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPolicy",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetSecurityGroupId(val *string) {
	if err := j.validateSetSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetSpecification(val *string) {
	if err := j.validateSetSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"specification",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetStorageSpace(val *float64) {
	if err := j.validateSetStorageSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSpace",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetStorageSpecCode(val *string) {
	if err := j.validateSetStorageSpecCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSpecCode",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DmsInstanceV2)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a DmsInstanceV2 resource upon running "cdktf plan <stack-name>".
func DmsInstanceV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDmsInstanceV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dmsInstanceV2.DmsInstanceV2",
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
func DmsInstanceV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsInstanceV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dmsInstanceV2.DmsInstanceV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsInstanceV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsInstanceV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dmsInstanceV2.DmsInstanceV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsInstanceV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsInstanceV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dmsInstanceV2.DmsInstanceV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsInstanceV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dmsInstanceV2.DmsInstanceV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DmsInstanceV2) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DmsInstanceV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DmsInstanceV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsInstanceV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsInstanceV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsInstanceV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsInstanceV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsInstanceV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsInstanceV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsInstanceV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsInstanceV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsInstanceV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsInstanceV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DmsInstanceV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsInstanceV2) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsInstanceV2) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DmsInstanceV2) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsInstanceV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsInstanceV2) PutTimeouts(value *DmsInstanceV2Timeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetAccessUser() {
	_jsii_.InvokeVoid(
		d,
		"resetAccessUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetDiskEncryptedEnable() {
	_jsii_.InvokeVoid(
		d,
		"resetDiskEncryptedEnable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetDiskEncryptedKey() {
	_jsii_.InvokeVoid(
		d,
		"resetDiskEncryptedKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetEnablePublicip() {
	_jsii_.InvokeVoid(
		d,
		"resetEnablePublicip",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetMaintainBegin() {
	_jsii_.InvokeVoid(
		d,
		"resetMaintainBegin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetMaintainEnd() {
	_jsii_.InvokeVoid(
		d,
		"resetMaintainEnd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetPartitionNum() {
	_jsii_.InvokeVoid(
		d,
		"resetPartitionNum",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetPublicBandwidth() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicBandwidth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetPublicipId() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicipId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetRetentionPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsInstanceV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsInstanceV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsInstanceV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsInstanceV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsInstanceV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsInstanceV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

