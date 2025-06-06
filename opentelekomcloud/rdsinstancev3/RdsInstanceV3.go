// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsinstancev3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/rdsinstancev3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/rds_instance_v3 opentelekomcloud_rds_instance_v3}.
type RdsInstanceV3 interface {
	cdktf.TerraformResource
	AutoscalingEnabled() cdktf.IResolvable
	AvailabilityZone() *[]*string
	SetAvailabilityZone(val *[]*string)
	AvailabilityZoneInput() *[]*string
	AvailabilityZones() *[]*string
	BackupStrategy() RdsInstanceV3BackupStrategyOutputReference
	BackupStrategyInput() *RdsInstanceV3BackupStrategy
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
	Created() *string
	Db() RdsInstanceV3DbOutputReference
	DbInput() *RdsInstanceV3Db
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	HaReplicationMode() *string
	SetHaReplicationMode(val *string)
	HaReplicationModeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LowerCaseTableNames() *string
	SetLowerCaseTableNames(val *string)
	LowerCaseTableNamesInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Nodes() RdsInstanceV3NodesList
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	ParamGroupId() *string
	SetParamGroupId(val *string)
	ParamGroupIdInput() *string
	PrivateIps() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicIps() *[]*string
	SetPublicIps(val *[]*string)
	PublicIpsInput() *[]*string
	// Experimental.
	RawOverrides() interface{}
	RestoredBackupId() *string
	RestoreFromBackup() RdsInstanceV3RestoreFromBackupOutputReference
	RestoreFromBackupInput() *RdsInstanceV3RestoreFromBackup
	RestorePoint() RdsInstanceV3RestorePointOutputReference
	RestorePointInput() *RdsInstanceV3RestorePoint
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	SslEnable() interface{}
	SetSslEnable(val interface{})
	SslEnableInput() interface{}
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tag() *map[string]*string
	SetTag(val *map[string]*string)
	TagInput() *map[string]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() RdsInstanceV3TimeoutsOutputReference
	TimeoutsInput() interface{}
	Volume() RdsInstanceV3VolumeOutputReference
	VolumeInput() *RdsInstanceV3Volume
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
	PutBackupStrategy(value *RdsInstanceV3BackupStrategy)
	PutDb(value *RdsInstanceV3Db)
	PutRestoreFromBackup(value *RdsInstanceV3RestoreFromBackup)
	PutRestorePoint(value *RdsInstanceV3RestorePoint)
	PutTimeouts(value *RdsInstanceV3Timeouts)
	PutVolume(value *RdsInstanceV3Volume)
	ResetBackupStrategy()
	ResetHaReplicationMode()
	ResetId()
	ResetLowerCaseTableNames()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetParamGroupId()
	ResetPublicIps()
	ResetRestoreFromBackup()
	ResetRestorePoint()
	ResetSslEnable()
	ResetTag()
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

// The jsii proxy struct for RdsInstanceV3
type jsiiProxy_RdsInstanceV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RdsInstanceV3) AutoscalingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoscalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) AvailabilityZone() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) AvailabilityZoneInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) BackupStrategy() RdsInstanceV3BackupStrategyOutputReference {
	var returns RdsInstanceV3BackupStrategyOutputReference
	_jsii_.Get(
		j,
		"backupStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) BackupStrategyInput() *RdsInstanceV3BackupStrategy {
	var returns *RdsInstanceV3BackupStrategy
	_jsii_.Get(
		j,
		"backupStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Db() RdsInstanceV3DbOutputReference {
	var returns RdsInstanceV3DbOutputReference
	_jsii_.Get(
		j,
		"db",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) DbInput() *RdsInstanceV3Db {
	var returns *RdsInstanceV3Db
	_jsii_.Get(
		j,
		"dbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Flavor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) FlavorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) HaReplicationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haReplicationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) HaReplicationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haReplicationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) LowerCaseTableNames() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lowerCaseTableNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) LowerCaseTableNamesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lowerCaseTableNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Nodes() RdsInstanceV3NodesList {
	var returns RdsInstanceV3NodesList
	_jsii_.Get(
		j,
		"nodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) ParamGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paramGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) ParamGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paramGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) PrivateIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) PublicIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) PublicIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) RestoredBackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoredBackupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) RestoreFromBackup() RdsInstanceV3RestoreFromBackupOutputReference {
	var returns RdsInstanceV3RestoreFromBackupOutputReference
	_jsii_.Get(
		j,
		"restoreFromBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) RestoreFromBackupInput() *RdsInstanceV3RestoreFromBackup {
	var returns *RdsInstanceV3RestoreFromBackup
	_jsii_.Get(
		j,
		"restoreFromBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) RestorePoint() RdsInstanceV3RestorePointOutputReference {
	var returns RdsInstanceV3RestorePointOutputReference
	_jsii_.Get(
		j,
		"restorePoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) RestorePointInput() *RdsInstanceV3RestorePoint {
	var returns *RdsInstanceV3RestorePoint
	_jsii_.Get(
		j,
		"restorePointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) SslEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) SslEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Tag() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) TagInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Timeouts() RdsInstanceV3TimeoutsOutputReference {
	var returns RdsInstanceV3TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) Volume() RdsInstanceV3VolumeOutputReference {
	var returns RdsInstanceV3VolumeOutputReference
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) VolumeInput() *RdsInstanceV3Volume {
	var returns *RdsInstanceV3Volume
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV3) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/rds_instance_v3 opentelekomcloud_rds_instance_v3} Resource.
func NewRdsInstanceV3(scope constructs.Construct, id *string, config *RdsInstanceV3Config) RdsInstanceV3 {
	_init_.Initialize()

	if err := validateNewRdsInstanceV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsInstanceV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV3.RdsInstanceV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/rds_instance_v3 opentelekomcloud_rds_instance_v3} Resource.
func NewRdsInstanceV3_Override(r RdsInstanceV3, scope constructs.Construct, id *string, config *RdsInstanceV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV3.RdsInstanceV3",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetAvailabilityZone(val *[]*string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetFlavor(val *string) {
	if err := j.validateSetFlavorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flavor",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetHaReplicationMode(val *string) {
	if err := j.validateSetHaReplicationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haReplicationMode",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetLowerCaseTableNames(val *string) {
	if err := j.validateSetLowerCaseTableNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lowerCaseTableNames",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetParamGroupId(val *string) {
	if err := j.validateSetParamGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paramGroupId",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetPublicIps(val *[]*string) {
	if err := j.validateSetPublicIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIps",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetSecurityGroupId(val *string) {
	if err := j.validateSetSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetSslEnable(val interface{}) {
	if err := j.validateSetSslEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslEnable",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetTag(val *map[string]*string) {
	if err := j.validateSetTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV3)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a RdsInstanceV3 resource upon running "cdktf plan <stack-name>".
func RdsInstanceV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRdsInstanceV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV3.RdsInstanceV3",
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
func RdsInstanceV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsInstanceV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV3.RdsInstanceV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsInstanceV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsInstanceV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV3.RdsInstanceV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsInstanceV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsInstanceV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV3.RdsInstanceV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RdsInstanceV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV3.RdsInstanceV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RdsInstanceV3) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RdsInstanceV3) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RdsInstanceV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RdsInstanceV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsInstanceV3) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RdsInstanceV3) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsInstanceV3) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RdsInstanceV3) PutBackupStrategy(value *RdsInstanceV3BackupStrategy) {
	if err := r.validatePutBackupStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putBackupStrategy",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV3) PutDb(value *RdsInstanceV3Db) {
	if err := r.validatePutDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putDb",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV3) PutRestoreFromBackup(value *RdsInstanceV3RestoreFromBackup) {
	if err := r.validatePutRestoreFromBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRestoreFromBackup",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV3) PutRestorePoint(value *RdsInstanceV3RestorePoint) {
	if err := r.validatePutRestorePointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRestorePoint",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV3) PutTimeouts(value *RdsInstanceV3Timeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV3) PutVolume(value *RdsInstanceV3Volume) {
	if err := r.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putVolume",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetBackupStrategy() {
	_jsii_.InvokeVoid(
		r,
		"resetBackupStrategy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetHaReplicationMode() {
	_jsii_.InvokeVoid(
		r,
		"resetHaReplicationMode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetLowerCaseTableNames() {
	_jsii_.InvokeVoid(
		r,
		"resetLowerCaseTableNames",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetParameters() {
	_jsii_.InvokeVoid(
		r,
		"resetParameters",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetParamGroupId() {
	_jsii_.InvokeVoid(
		r,
		"resetParamGroupId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetPublicIps() {
	_jsii_.InvokeVoid(
		r,
		"resetPublicIps",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetRestoreFromBackup() {
	_jsii_.InvokeVoid(
		r,
		"resetRestoreFromBackup",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetRestorePoint() {
	_jsii_.InvokeVoid(
		r,
		"resetRestorePoint",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetSslEnable() {
	_jsii_.InvokeVoid(
		r,
		"resetSslEnable",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetTag() {
	_jsii_.InvokeVoid(
		r,
		"resetTag",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

