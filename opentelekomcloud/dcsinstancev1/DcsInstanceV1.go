// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcsinstancev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dcsinstancev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/dcs_instance_v1 opentelekomcloud_dcs_instance_v1}.
type DcsInstanceV1 interface {
	cdktf.TerraformResource
	AvailableZones() *[]*string
	SetAvailableZones(val *[]*string)
	AvailableZonesInput() *[]*string
	BackupAt() *[]*float64
	SetBackupAt(val *[]*float64)
	BackupAtInput() *[]*float64
	BackupPolicy() DcsInstanceV1BackupPolicyOutputReference
	BackupPolicyInput() *DcsInstanceV1BackupPolicy
	BackupType() *string
	SetBackupType(val *string)
	BackupTypeInput() *string
	BeginAt() *string
	SetBeginAt(val *string)
	BeginAtInput() *string
	Capacity() *float64
	SetCapacity(val *float64)
	CapacityInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Configuration() DcsInstanceV1ConfigurationList
	ConfigurationInput() interface{}
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
	EnableWhitelist() interface{}
	SetEnableWhitelist(val interface{})
	EnableWhitelistInput() interface{}
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
	InternalVersion() *string
	Ip() *string
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
	NoPasswordAccess() *string
	OrderId() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PeriodType() *string
	SetPeriodType(val *string)
	PeriodTypeInput() *string
	Port() *float64
	PrivateIp() *string
	SetPrivateIp(val *string)
	PrivateIpInput() *string
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
	// Experimental.
	RawOverrides() interface{}
	ResourceSpecCode() *string
	SaveDays() *float64
	SetSaveDays(val *float64)
	SaveDaysInput() *float64
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	SecurityGroupName() *string
	Status() *string
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
	Timeouts() DcsInstanceV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	UsedMemory() *float64
	UserId() *string
	UserName() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	VpcName() *string
	Whitelist() DcsInstanceV1WhitelistStructList
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
	PutBackupPolicy(value *DcsInstanceV1BackupPolicy)
	PutConfiguration(value interface{})
	PutTimeouts(value *DcsInstanceV1Timeouts)
	PutWhitelist(value interface{})
	ResetBackupAt()
	ResetBackupPolicy()
	ResetBackupType()
	ResetBeginAt()
	ResetConfiguration()
	ResetDescription()
	ResetEnableWhitelist()
	ResetId()
	ResetMaintainBegin()
	ResetMaintainEnd()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPeriodType()
	ResetPrivateIp()
	ResetSaveDays()
	ResetSecurityGroupId()
	ResetTags()
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

// The jsii proxy struct for DcsInstanceV1
type jsiiProxy_DcsInstanceV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DcsInstanceV1) AvailableZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) AvailableZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) BackupAt() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"backupAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) BackupAtInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"backupAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) BackupPolicy() DcsInstanceV1BackupPolicyOutputReference {
	var returns DcsInstanceV1BackupPolicyOutputReference
	_jsii_.Get(
		j,
		"backupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) BackupPolicyInput() *DcsInstanceV1BackupPolicy {
	var returns *DcsInstanceV1BackupPolicy
	_jsii_.Get(
		j,
		"backupPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) BackupType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) BackupTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) BeginAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) BeginAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) CapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Configuration() DcsInstanceV1ConfigurationList {
	var returns DcsInstanceV1ConfigurationList
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) EnableWhitelist() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) EnableWhitelistInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) InternalVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Ip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) MaintainBegin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainBegin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) MaintainBeginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainBeginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) MaintainEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) MaintainEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) MaxMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) NoPasswordAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noPasswordAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) OrderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) PeriodType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) PeriodTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) ResourceSpecCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceSpecCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) SaveDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saveDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) SaveDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saveDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) SecurityGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) SubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Timeouts() DcsInstanceV1TimeoutsOutputReference {
	var returns DcsInstanceV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) UsedMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) VpcName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) Whitelist() DcsInstanceV1WhitelistStructList {
	var returns DcsInstanceV1WhitelistStructList
	_jsii_.Get(
		j,
		"whitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1) WhitelistInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"whitelistInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/dcs_instance_v1 opentelekomcloud_dcs_instance_v1} Resource.
func NewDcsInstanceV1(scope constructs.Construct, id *string, config *DcsInstanceV1Config) DcsInstanceV1 {
	_init_.Initialize()

	if err := validateNewDcsInstanceV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DcsInstanceV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV1.DcsInstanceV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/dcs_instance_v1 opentelekomcloud_dcs_instance_v1} Resource.
func NewDcsInstanceV1_Override(d DcsInstanceV1, scope constructs.Construct, id *string, config *DcsInstanceV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV1.DcsInstanceV1",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetAvailableZones(val *[]*string) {
	if err := j.validateSetAvailableZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableZones",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetBackupAt(val *[]*float64) {
	if err := j.validateSetBackupAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupAt",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetBackupType(val *string) {
	if err := j.validateSetBackupTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupType",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetBeginAt(val *string) {
	if err := j.validateSetBeginAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"beginAt",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetCapacity(val *float64) {
	if err := j.validateSetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetEnableWhitelist(val interface{}) {
	if err := j.validateSetEnableWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableWhitelist",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetMaintainBegin(val *string) {
	if err := j.validateSetMaintainBeginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainBegin",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetMaintainEnd(val *string) {
	if err := j.validateSetMaintainEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainEnd",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetPeriodType(val *string) {
	if err := j.validateSetPeriodTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodType",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetProductId(val *string) {
	if err := j.validateSetProductIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetSaveDays(val *float64) {
	if err := j.validateSetSaveDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saveDays",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetSecurityGroupId(val *string) {
	if err := j.validateSetSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a DcsInstanceV1 resource upon running "cdktf plan <stack-name>".
func DcsInstanceV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDcsInstanceV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV1.DcsInstanceV1",
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
func DcsInstanceV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDcsInstanceV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV1.DcsInstanceV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DcsInstanceV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDcsInstanceV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV1.DcsInstanceV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DcsInstanceV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDcsInstanceV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV1.DcsInstanceV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DcsInstanceV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV1.DcsInstanceV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DcsInstanceV1) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DcsInstanceV1) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DcsInstanceV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DcsInstanceV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DcsInstanceV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DcsInstanceV1) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DcsInstanceV1) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DcsInstanceV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DcsInstanceV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DcsInstanceV1) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DcsInstanceV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DcsInstanceV1) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV1) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DcsInstanceV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DcsInstanceV1) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DcsInstanceV1) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DcsInstanceV1) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DcsInstanceV1) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DcsInstanceV1) PutBackupPolicy(value *DcsInstanceV1BackupPolicy) {
	if err := d.validatePutBackupPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBackupPolicy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DcsInstanceV1) PutConfiguration(value interface{}) {
	if err := d.validatePutConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DcsInstanceV1) PutTimeouts(value *DcsInstanceV1Timeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DcsInstanceV1) PutWhitelist(value interface{}) {
	if err := d.validatePutWhitelistParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWhitelist",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetBackupAt() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetBackupPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetBackupType() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetBeginAt() {
	_jsii_.InvokeVoid(
		d,
		"resetBeginAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetEnableWhitelist() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableWhitelist",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetMaintainBegin() {
	_jsii_.InvokeVoid(
		d,
		"resetMaintainBegin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetMaintainEnd() {
	_jsii_.InvokeVoid(
		d,
		"resetMaintainEnd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetPeriodType() {
	_jsii_.InvokeVoid(
		d,
		"resetPeriodType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetSaveDays() {
	_jsii_.InvokeVoid(
		d,
		"resetSaveDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetSecurityGroupId() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityGroupId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) ResetWhitelist() {
	_jsii_.InvokeVoid(
		d,
		"resetWhitelist",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV1) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV1) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

