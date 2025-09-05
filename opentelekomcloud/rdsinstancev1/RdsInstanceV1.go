// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsinstancev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/rdsinstancev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/rds_instance_v1 opentelekomcloud_rds_instance_v1}.
type RdsInstanceV1 interface {
	cdktf.TerraformResource
	Availabilityzone() *string
	SetAvailabilityzone(val *string)
	AvailabilityzoneInput() *string
	Backupstrategy() RdsInstanceV1BackupstrategyOutputReference
	BackupstrategyInput() *RdsInstanceV1Backupstrategy
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
	Datastore() RdsInstanceV1DatastoreOutputReference
	DatastoreInput() *RdsInstanceV1Datastore
	Dbport() *string
	SetDbport(val *string)
	DbportInput() *string
	Dbrtpd() *string
	SetDbrtpd(val *string)
	DbrtpdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Flavorref() *string
	SetFlavorref(val *string)
	FlavorrefInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Ha() RdsInstanceV1HaOutputReference
	HaInput() *RdsInstanceV1Ha
	Hostname() *string
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
	Nics() RdsInstanceV1NicsOutputReference
	NicsInput() *RdsInstanceV1Nics
	// The tree node.
	Node() constructs.Node
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
	Securitygroup() RdsInstanceV1SecuritygroupOutputReference
	SecuritygroupInput() *RdsInstanceV1Securitygroup
	Status() *string
	Tag() *map[string]*string
	SetTag(val *map[string]*string)
	TagInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() RdsInstanceV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	Updated() *string
	Volume() RdsInstanceV1VolumeOutputReference
	VolumeInput() *RdsInstanceV1Volume
	Vpc() *string
	SetVpc(val *string)
	VpcInput() *string
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
	PutBackupstrategy(value *RdsInstanceV1Backupstrategy)
	PutDatastore(value *RdsInstanceV1Datastore)
	PutHa(value *RdsInstanceV1Ha)
	PutNics(value *RdsInstanceV1Nics)
	PutSecuritygroup(value *RdsInstanceV1Securitygroup)
	PutTimeouts(value *RdsInstanceV1Timeouts)
	PutVolume(value *RdsInstanceV1Volume)
	ResetBackupstrategy()
	ResetDbport()
	ResetHa()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetTag()
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

// The jsii proxy struct for RdsInstanceV1
type jsiiProxy_RdsInstanceV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RdsInstanceV1) Availabilityzone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityzone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) AvailabilityzoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityzoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Backupstrategy() RdsInstanceV1BackupstrategyOutputReference {
	var returns RdsInstanceV1BackupstrategyOutputReference
	_jsii_.Get(
		j,
		"backupstrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) BackupstrategyInput() *RdsInstanceV1Backupstrategy {
	var returns *RdsInstanceV1Backupstrategy
	_jsii_.Get(
		j,
		"backupstrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Datastore() RdsInstanceV1DatastoreOutputReference {
	var returns RdsInstanceV1DatastoreOutputReference
	_jsii_.Get(
		j,
		"datastore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) DatastoreInput() *RdsInstanceV1Datastore {
	var returns *RdsInstanceV1Datastore
	_jsii_.Get(
		j,
		"datastoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Dbport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) DbportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Dbrtpd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbrtpd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) DbrtpdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbrtpdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Flavorref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) FlavorrefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorrefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Ha() RdsInstanceV1HaOutputReference {
	var returns RdsInstanceV1HaOutputReference
	_jsii_.Get(
		j,
		"ha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) HaInput() *RdsInstanceV1Ha {
	var returns *RdsInstanceV1Ha
	_jsii_.Get(
		j,
		"haInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Nics() RdsInstanceV1NicsOutputReference {
	var returns RdsInstanceV1NicsOutputReference
	_jsii_.Get(
		j,
		"nics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) NicsInput() *RdsInstanceV1Nics {
	var returns *RdsInstanceV1Nics
	_jsii_.Get(
		j,
		"nicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Securitygroup() RdsInstanceV1SecuritygroupOutputReference {
	var returns RdsInstanceV1SecuritygroupOutputReference
	_jsii_.Get(
		j,
		"securitygroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) SecuritygroupInput() *RdsInstanceV1Securitygroup {
	var returns *RdsInstanceV1Securitygroup
	_jsii_.Get(
		j,
		"securitygroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Tag() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) TagInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Timeouts() RdsInstanceV1TimeoutsOutputReference {
	var returns RdsInstanceV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Updated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Volume() RdsInstanceV1VolumeOutputReference {
	var returns RdsInstanceV1VolumeOutputReference
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) VolumeInput() *RdsInstanceV1Volume {
	var returns *RdsInstanceV1Volume
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) Vpc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsInstanceV1) VpcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/rds_instance_v1 opentelekomcloud_rds_instance_v1} Resource.
func NewRdsInstanceV1(scope constructs.Construct, id *string, config *RdsInstanceV1Config) RdsInstanceV1 {
	_init_.Initialize()

	if err := validateNewRdsInstanceV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsInstanceV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV1.RdsInstanceV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/rds_instance_v1 opentelekomcloud_rds_instance_v1} Resource.
func NewRdsInstanceV1_Override(r RdsInstanceV1, scope constructs.Construct, id *string, config *RdsInstanceV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV1.RdsInstanceV1",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetAvailabilityzone(val *string) {
	if err := j.validateSetAvailabilityzoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityzone",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetDbport(val *string) {
	if err := j.validateSetDbportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbport",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetDbrtpd(val *string) {
	if err := j.validateSetDbrtpdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbrtpd",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetFlavorref(val *string) {
	if err := j.validateSetFlavorrefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flavorref",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetTag(val *map[string]*string) {
	if err := j.validateSetTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_RdsInstanceV1)SetVpc(val *string) {
	if err := j.validateSetVpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpc",
		val,
	)
}

// Generates CDKTF code for importing a RdsInstanceV1 resource upon running "cdktf plan <stack-name>".
func RdsInstanceV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRdsInstanceV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV1.RdsInstanceV1",
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
func RdsInstanceV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsInstanceV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV1.RdsInstanceV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsInstanceV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsInstanceV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV1.RdsInstanceV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsInstanceV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsInstanceV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV1.RdsInstanceV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RdsInstanceV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.rdsInstanceV1.RdsInstanceV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RdsInstanceV1) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RdsInstanceV1) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RdsInstanceV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RdsInstanceV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RdsInstanceV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RdsInstanceV1) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RdsInstanceV1) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RdsInstanceV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RdsInstanceV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RdsInstanceV1) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RdsInstanceV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RdsInstanceV1) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV1) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RdsInstanceV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RdsInstanceV1) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsInstanceV1) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RdsInstanceV1) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsInstanceV1) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RdsInstanceV1) PutBackupstrategy(value *RdsInstanceV1Backupstrategy) {
	if err := r.validatePutBackupstrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putBackupstrategy",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV1) PutDatastore(value *RdsInstanceV1Datastore) {
	if err := r.validatePutDatastoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putDatastore",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV1) PutHa(value *RdsInstanceV1Ha) {
	if err := r.validatePutHaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putHa",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV1) PutNics(value *RdsInstanceV1Nics) {
	if err := r.validatePutNicsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putNics",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV1) PutSecuritygroup(value *RdsInstanceV1Securitygroup) {
	if err := r.validatePutSecuritygroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSecuritygroup",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV1) PutTimeouts(value *RdsInstanceV1Timeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV1) PutVolume(value *RdsInstanceV1Volume) {
	if err := r.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putVolume",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsInstanceV1) ResetBackupstrategy() {
	_jsii_.InvokeVoid(
		r,
		"resetBackupstrategy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV1) ResetDbport() {
	_jsii_.InvokeVoid(
		r,
		"resetDbport",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV1) ResetHa() {
	_jsii_.InvokeVoid(
		r,
		"resetHa",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV1) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV1) ResetName() {
	_jsii_.InvokeVoid(
		r,
		"resetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV1) ResetRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV1) ResetTag() {
	_jsii_.InvokeVoid(
		r,
		"resetTag",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsInstanceV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV1) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV1) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsInstanceV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

