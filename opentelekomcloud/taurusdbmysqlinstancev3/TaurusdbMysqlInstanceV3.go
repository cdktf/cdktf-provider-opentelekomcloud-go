// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package taurusdbmysqlinstancev3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/taurusdbmysqlinstancev3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/taurusdb_mysql_instance_v3 opentelekomcloud_taurusdb_mysql_instance_v3}.
type TaurusdbMysqlInstanceV3 interface {
	cdktf.TerraformResource
	AvailabilityZoneMode() *string
	SetAvailabilityZoneMode(val *string)
	AvailabilityZoneModeInput() *string
	BackupStrategy() TaurusdbMysqlInstanceV3BackupStrategyOutputReference
	BackupStrategyInput() *TaurusdbMysqlInstanceV3BackupStrategy
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigurationId() *string
	SetConfigurationId(val *string)
	ConfigurationIdInput() *string
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
	Datastore() TaurusdbMysqlInstanceV3DatastoreOutputReference
	DatastoreInput() *TaurusdbMysqlInstanceV3Datastore
	DbUserName() *string
	DedicatedResourceId() *string
	SetDedicatedResourceId(val *string)
	DedicatedResourceIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnterpriseProjectId() *string
	SetEnterpriseProjectId(val *string)
	EnterpriseProjectIdInput() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MasterAvailabilityZone() *string
	SetMasterAvailabilityZone(val *string)
	MasterAvailabilityZoneInput() *string
	Mode() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Nodes() TaurusdbMysqlInstanceV3NodesList
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PrivateDnsName() *string
	PrivateWriteIp() *string
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
	ReadReplicas() *float64
	SetReadReplicas(val *float64)
	ReadReplicasInput() *float64
	Region() *string
	SecondsLevelMonitoringEnabled() interface{}
	SetSecondsLevelMonitoringEnabled(val interface{})
	SecondsLevelMonitoringEnabledInput() interface{}
	SecondsLevelMonitoringPeriod() *float64
	SetSecondsLevelMonitoringPeriod(val *float64)
	SecondsLevelMonitoringPeriodInput() *float64
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	Status() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	TableNameCaseSensitivity() interface{}
	SetTableNameCaseSensitivity(val interface{})
	TableNameCaseSensitivityInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() TaurusdbMysqlInstanceV3TimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	UpdatedAt() *string
	VolumeSize() *float64
	SetVolumeSize(val *float64)
	VolumeSizeInput() *float64
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
	PutBackupStrategy(value *TaurusdbMysqlInstanceV3BackupStrategy)
	PutDatastore(value *TaurusdbMysqlInstanceV3Datastore)
	PutTimeouts(value *TaurusdbMysqlInstanceV3Timeouts)
	ResetAvailabilityZoneMode()
	ResetBackupStrategy()
	ResetConfigurationId()
	ResetDatastore()
	ResetDedicatedResourceId()
	ResetEnterpriseProjectId()
	ResetId()
	ResetMasterAvailabilityZone()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetReadReplicas()
	ResetSecondsLevelMonitoringEnabled()
	ResetSecondsLevelMonitoringPeriod()
	ResetSecurityGroupId()
	ResetTableNameCaseSensitivity()
	ResetTimeouts()
	ResetTimeZone()
	ResetVolumeSize()
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

// The jsii proxy struct for TaurusdbMysqlInstanceV3
type jsiiProxy_TaurusdbMysqlInstanceV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) AvailabilityZoneMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) AvailabilityZoneModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) BackupStrategy() TaurusdbMysqlInstanceV3BackupStrategyOutputReference {
	var returns TaurusdbMysqlInstanceV3BackupStrategyOutputReference
	_jsii_.Get(
		j,
		"backupStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) BackupStrategyInput() *TaurusdbMysqlInstanceV3BackupStrategy {
	var returns *TaurusdbMysqlInstanceV3BackupStrategy
	_jsii_.Get(
		j,
		"backupStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) ConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) ConfigurationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Datastore() TaurusdbMysqlInstanceV3DatastoreOutputReference {
	var returns TaurusdbMysqlInstanceV3DatastoreOutputReference
	_jsii_.Get(
		j,
		"datastore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) DatastoreInput() *TaurusdbMysqlInstanceV3Datastore {
	var returns *TaurusdbMysqlInstanceV3Datastore
	_jsii_.Get(
		j,
		"datastoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) DbUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) DedicatedResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) DedicatedResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) EnterpriseProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) EnterpriseProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Flavor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) FlavorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) MasterAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) MasterAvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Nodes() TaurusdbMysqlInstanceV3NodesList {
	var returns TaurusdbMysqlInstanceV3NodesList
	_jsii_.Get(
		j,
		"nodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) PrivateDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) PrivateWriteIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateWriteIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) ReadReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) ReadReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) SecondsLevelMonitoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondsLevelMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) SecondsLevelMonitoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondsLevelMonitoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) SecondsLevelMonitoringPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsLevelMonitoringPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) SecondsLevelMonitoringPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsLevelMonitoringPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) TableNameCaseSensitivity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableNameCaseSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) TableNameCaseSensitivityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableNameCaseSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) Timeouts() TaurusdbMysqlInstanceV3TimeoutsOutputReference {
	var returns TaurusdbMysqlInstanceV3TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) VolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) VolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/taurusdb_mysql_instance_v3 opentelekomcloud_taurusdb_mysql_instance_v3} Resource.
func NewTaurusdbMysqlInstanceV3(scope constructs.Construct, id *string, config *TaurusdbMysqlInstanceV3Config) TaurusdbMysqlInstanceV3 {
	_init_.Initialize()

	if err := validateNewTaurusdbMysqlInstanceV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TaurusdbMysqlInstanceV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlInstanceV3.TaurusdbMysqlInstanceV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/taurusdb_mysql_instance_v3 opentelekomcloud_taurusdb_mysql_instance_v3} Resource.
func NewTaurusdbMysqlInstanceV3_Override(t TaurusdbMysqlInstanceV3, scope constructs.Construct, id *string, config *TaurusdbMysqlInstanceV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlInstanceV3.TaurusdbMysqlInstanceV3",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetAvailabilityZoneMode(val *string) {
	if err := j.validateSetAvailabilityZoneModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneMode",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetConfigurationId(val *string) {
	if err := j.validateSetConfigurationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationId",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetDedicatedResourceId(val *string) {
	if err := j.validateSetDedicatedResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedResourceId",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetEnterpriseProjectId(val *string) {
	if err := j.validateSetEnterpriseProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enterpriseProjectId",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetFlavor(val *string) {
	if err := j.validateSetFlavorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flavor",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetMasterAvailabilityZone(val *string) {
	if err := j.validateSetMasterAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetReadReplicas(val *float64) {
	if err := j.validateSetReadReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readReplicas",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetSecondsLevelMonitoringEnabled(val interface{}) {
	if err := j.validateSetSecondsLevelMonitoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondsLevelMonitoringEnabled",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetSecondsLevelMonitoringPeriod(val *float64) {
	if err := j.validateSetSecondsLevelMonitoringPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondsLevelMonitoringPeriod",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetSecurityGroupId(val *string) {
	if err := j.validateSetSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetTableNameCaseSensitivity(val interface{}) {
	if err := j.validateSetTableNameCaseSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableNameCaseSensitivity",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetVolumeSize(val *float64) {
	if err := j.validateSetVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSize",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlInstanceV3)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a TaurusdbMysqlInstanceV3 resource upon running "cdktf plan <stack-name>".
func TaurusdbMysqlInstanceV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateTaurusdbMysqlInstanceV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlInstanceV3.TaurusdbMysqlInstanceV3",
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
func TaurusdbMysqlInstanceV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTaurusdbMysqlInstanceV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlInstanceV3.TaurusdbMysqlInstanceV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TaurusdbMysqlInstanceV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTaurusdbMysqlInstanceV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlInstanceV3.TaurusdbMysqlInstanceV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TaurusdbMysqlInstanceV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTaurusdbMysqlInstanceV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlInstanceV3.TaurusdbMysqlInstanceV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TaurusdbMysqlInstanceV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlInstanceV3.TaurusdbMysqlInstanceV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) PutBackupStrategy(value *TaurusdbMysqlInstanceV3BackupStrategy) {
	if err := t.validatePutBackupStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBackupStrategy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) PutDatastore(value *TaurusdbMysqlInstanceV3Datastore) {
	if err := t.validatePutDatastoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDatastore",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) PutTimeouts(value *TaurusdbMysqlInstanceV3Timeouts) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetAvailabilityZoneMode() {
	_jsii_.InvokeVoid(
		t,
		"resetAvailabilityZoneMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetBackupStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetBackupStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetConfigurationId() {
	_jsii_.InvokeVoid(
		t,
		"resetConfigurationId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetDatastore() {
	_jsii_.InvokeVoid(
		t,
		"resetDatastore",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetDedicatedResourceId() {
	_jsii_.InvokeVoid(
		t,
		"resetDedicatedResourceId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetEnterpriseProjectId() {
	_jsii_.InvokeVoid(
		t,
		"resetEnterpriseProjectId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetMasterAvailabilityZone() {
	_jsii_.InvokeVoid(
		t,
		"resetMasterAvailabilityZone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetReadReplicas() {
	_jsii_.InvokeVoid(
		t,
		"resetReadReplicas",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetSecondsLevelMonitoringEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondsLevelMonitoringEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetSecondsLevelMonitoringPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondsLevelMonitoringPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetSecurityGroupId() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroupId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetTableNameCaseSensitivity() {
	_jsii_.InvokeVoid(
		t,
		"resetTableNameCaseSensitivity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetTimeZone() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ResetVolumeSize() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumeSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlInstanceV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

