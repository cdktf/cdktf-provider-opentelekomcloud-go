// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gaussdbmysqlinstancev3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/gaussdbmysqlinstancev3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/gaussdb_mysql_instance_v3 opentelekomcloud_gaussdb_mysql_instance_v3}.
type GaussdbMysqlInstanceV3 interface {
	cdktf.TerraformResource
	Alias() *string
	AvailabilityZoneMode() *string
	SetAvailabilityZoneMode(val *string)
	AvailabilityZoneModeInput() *string
	BackupStrategy() GaussdbMysqlInstanceV3BackupStrategyOutputReference
	BackupStrategyInput() *GaussdbMysqlInstanceV3BackupStrategy
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChargingMode() *string
	ConfigurationId() *string
	SetConfigurationId(val *string)
	ConfigurationIdInput() *string
	ConfigurationName() *string
	SetConfigurationName(val *string)
	ConfigurationNameInput() *string
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
	Datastore() GaussdbMysqlInstanceV3DatastoreOutputReference
	DatastoreInput() *GaussdbMysqlInstanceV3Datastore
	DbUserName() *string
	DedicatedResourceId() *string
	SetDedicatedResourceId(val *string)
	DedicatedResourceIdInput() *string
	DedicatedResourceName() *string
	SetDedicatedResourceName(val *string)
	DedicatedResourceNameInput() *string
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
	NodeCount() *float64
	Nodes() GaussdbMysqlInstanceV3NodesList
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	PrivateWriteIp() *[]*string
	ProjectId() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicIp() *string
	// Experimental.
	RawOverrides() interface{}
	ReadReplicas() *float64
	SetReadReplicas(val *float64)
	ReadReplicasInput() *float64
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	Status() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GaussdbMysqlInstanceV3TimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	Updated() *string
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
	PutBackupStrategy(value *GaussdbMysqlInstanceV3BackupStrategy)
	PutDatastore(value *GaussdbMysqlInstanceV3Datastore)
	PutTimeouts(value *GaussdbMysqlInstanceV3Timeouts)
	ResetAvailabilityZoneMode()
	ResetBackupStrategy()
	ResetConfigurationId()
	ResetConfigurationName()
	ResetDatastore()
	ResetDedicatedResourceId()
	ResetDedicatedResourceName()
	ResetId()
	ResetMasterAvailabilityZone()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReadReplicas()
	ResetRegion()
	ResetSecurityGroupId()
	ResetTimeouts()
	ResetTimeZone()
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

// The jsii proxy struct for GaussdbMysqlInstanceV3
type jsiiProxy_GaussdbMysqlInstanceV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) AvailabilityZoneMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) AvailabilityZoneModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) BackupStrategy() GaussdbMysqlInstanceV3BackupStrategyOutputReference {
	var returns GaussdbMysqlInstanceV3BackupStrategyOutputReference
	_jsii_.Get(
		j,
		"backupStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) BackupStrategyInput() *GaussdbMysqlInstanceV3BackupStrategy {
	var returns *GaussdbMysqlInstanceV3BackupStrategy
	_jsii_.Get(
		j,
		"backupStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) ChargingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) ConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) ConfigurationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) ConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) ConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Datastore() GaussdbMysqlInstanceV3DatastoreOutputReference {
	var returns GaussdbMysqlInstanceV3DatastoreOutputReference
	_jsii_.Get(
		j,
		"datastore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) DatastoreInput() *GaussdbMysqlInstanceV3Datastore {
	var returns *GaussdbMysqlInstanceV3Datastore
	_jsii_.Get(
		j,
		"datastoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) DbUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) DedicatedResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) DedicatedResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) DedicatedResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) DedicatedResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Flavor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) FlavorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) MasterAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) MasterAvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Nodes() GaussdbMysqlInstanceV3NodesList {
	var returns GaussdbMysqlInstanceV3NodesList
	_jsii_.Get(
		j,
		"nodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) PrivateWriteIp() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateWriteIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) ReadReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) ReadReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Timeouts() GaussdbMysqlInstanceV3TimeoutsOutputReference {
	var returns GaussdbMysqlInstanceV3TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) Updated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/gaussdb_mysql_instance_v3 opentelekomcloud_gaussdb_mysql_instance_v3} Resource.
func NewGaussdbMysqlInstanceV3(scope constructs.Construct, id *string, config *GaussdbMysqlInstanceV3Config) GaussdbMysqlInstanceV3 {
	_init_.Initialize()

	if err := validateNewGaussdbMysqlInstanceV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GaussdbMysqlInstanceV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.gaussdbMysqlInstanceV3.GaussdbMysqlInstanceV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/gaussdb_mysql_instance_v3 opentelekomcloud_gaussdb_mysql_instance_v3} Resource.
func NewGaussdbMysqlInstanceV3_Override(g GaussdbMysqlInstanceV3, scope constructs.Construct, id *string, config *GaussdbMysqlInstanceV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.gaussdbMysqlInstanceV3.GaussdbMysqlInstanceV3",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetAvailabilityZoneMode(val *string) {
	if err := j.validateSetAvailabilityZoneModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneMode",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetConfigurationId(val *string) {
	if err := j.validateSetConfigurationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationId",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetConfigurationName(val *string) {
	if err := j.validateSetConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationName",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetDedicatedResourceId(val *string) {
	if err := j.validateSetDedicatedResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedResourceId",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetDedicatedResourceName(val *string) {
	if err := j.validateSetDedicatedResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedResourceName",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetFlavor(val *string) {
	if err := j.validateSetFlavorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flavor",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetMasterAvailabilityZone(val *string) {
	if err := j.validateSetMasterAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetReadReplicas(val *float64) {
	if err := j.validateSetReadReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readReplicas",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetSecurityGroupId(val *string) {
	if err := j.validateSetSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (j *jsiiProxy_GaussdbMysqlInstanceV3)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a GaussdbMysqlInstanceV3 resource upon running "cdktf plan <stack-name>".
func GaussdbMysqlInstanceV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGaussdbMysqlInstanceV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.gaussdbMysqlInstanceV3.GaussdbMysqlInstanceV3",
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
func GaussdbMysqlInstanceV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGaussdbMysqlInstanceV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.gaussdbMysqlInstanceV3.GaussdbMysqlInstanceV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GaussdbMysqlInstanceV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGaussdbMysqlInstanceV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.gaussdbMysqlInstanceV3.GaussdbMysqlInstanceV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GaussdbMysqlInstanceV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGaussdbMysqlInstanceV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.gaussdbMysqlInstanceV3.GaussdbMysqlInstanceV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GaussdbMysqlInstanceV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.gaussdbMysqlInstanceV3.GaussdbMysqlInstanceV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) PutBackupStrategy(value *GaussdbMysqlInstanceV3BackupStrategy) {
	if err := g.validatePutBackupStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBackupStrategy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) PutDatastore(value *GaussdbMysqlInstanceV3Datastore) {
	if err := g.validatePutDatastoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDatastore",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) PutTimeouts(value *GaussdbMysqlInstanceV3Timeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetAvailabilityZoneMode() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailabilityZoneMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetBackupStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetBackupStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetConfigurationId() {
	_jsii_.InvokeVoid(
		g,
		"resetConfigurationId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetConfigurationName() {
	_jsii_.InvokeVoid(
		g,
		"resetConfigurationName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetDatastore() {
	_jsii_.InvokeVoid(
		g,
		"resetDatastore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetDedicatedResourceId() {
	_jsii_.InvokeVoid(
		g,
		"resetDedicatedResourceId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetDedicatedResourceName() {
	_jsii_.InvokeVoid(
		g,
		"resetDedicatedResourceName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetMasterAvailabilityZone() {
	_jsii_.InvokeVoid(
		g,
		"resetMasterAvailabilityZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetReadReplicas() {
	_jsii_.InvokeVoid(
		g,
		"resetReadReplicas",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetSecurityGroupId() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityGroupId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ResetTimeZone() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GaussdbMysqlInstanceV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

