// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package drstaskv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/drstaskv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/drs_task_v3 opentelekomcloud_drs_task_v3}.
type DrsTaskV3 interface {
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
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DestinationDb() DrsTaskV3DestinationDbOutputReference
	DestinationDbInput() *DrsTaskV3DestinationDb
	DestinationDbReadonly() interface{}
	SetDestinationDbReadonly(val interface{})
	DestinationDbReadonlyInput() interface{}
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	EngineType() *string
	SetEngineType(val *string)
	EngineTypeInput() *string
	ExpiredDays() *float64
	SetExpiredDays(val *float64)
	ExpiredDaysInput() *float64
	ForceDestroy() interface{}
	SetForceDestroy(val interface{})
	ForceDestroyInput() interface{}
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
	LimitSpeed() DrsTaskV3LimitSpeedList
	LimitSpeedInput() interface{}
	MigrateDefiner() interface{}
	SetMigrateDefiner(val interface{})
	MigrateDefinerInput() interface{}
	MigrationType() *string
	SetMigrationType(val *string)
	MigrationTypeInput() *string
	MultiWrite() interface{}
	SetMultiWrite(val interface{})
	MultiWriteInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetType() *string
	SetNetType(val *string)
	NetTypeInput() *string
	// The tree node.
	Node() constructs.Node
	NodeNum() *float64
	SetNodeNum(val *float64)
	NodeNumInput() *float64
	PrivateIp() *string
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
	Region() *string
	SourceDb() DrsTaskV3SourceDbOutputReference
	SourceDbInput() *DrsTaskV3SourceDb
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	Status() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DrsTaskV3TimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutDestinationDb(value *DrsTaskV3DestinationDb)
	PutLimitSpeed(value interface{})
	PutSourceDb(value *DrsTaskV3SourceDb)
	PutTimeouts(value *DrsTaskV3Timeouts)
	ResetDescription()
	ResetDestinationDbReadonly()
	ResetExpiredDays()
	ResetForceDestroy()
	ResetId()
	ResetLimitSpeed()
	ResetMigrateDefiner()
	ResetMigrationType()
	ResetMultiWrite()
	ResetNetType()
	ResetNodeNum()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStartTime()
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

// The jsii proxy struct for DrsTaskV3
type jsiiProxy_DrsTaskV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DrsTaskV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) DestinationDb() DrsTaskV3DestinationDbOutputReference {
	var returns DrsTaskV3DestinationDbOutputReference
	_jsii_.Get(
		j,
		"destinationDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) DestinationDbInput() *DrsTaskV3DestinationDb {
	var returns *DrsTaskV3DestinationDb
	_jsii_.Get(
		j,
		"destinationDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) DestinationDbReadonly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationDbReadonly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) DestinationDbReadonlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationDbReadonlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) EngineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) ExpiredDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expiredDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) ExpiredDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expiredDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) LimitSpeed() DrsTaskV3LimitSpeedList {
	var returns DrsTaskV3LimitSpeedList
	_jsii_.Get(
		j,
		"limitSpeed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) LimitSpeedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"limitSpeedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) MigrateDefiner() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"migrateDefiner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) MigrateDefinerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"migrateDefinerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) MigrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) MigrationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) MultiWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) MultiWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) NetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) NetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) NodeNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) NodeNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) SourceDb() DrsTaskV3SourceDbOutputReference {
	var returns DrsTaskV3SourceDbOutputReference
	_jsii_.Get(
		j,
		"sourceDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) SourceDbInput() *DrsTaskV3SourceDb {
	var returns *DrsTaskV3SourceDb
	_jsii_.Get(
		j,
		"sourceDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Timeouts() DrsTaskV3TimeoutsOutputReference {
	var returns DrsTaskV3TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/drs_task_v3 opentelekomcloud_drs_task_v3} Resource.
func NewDrsTaskV3(scope constructs.Construct, id *string, config *DrsTaskV3Config) DrsTaskV3 {
	_init_.Initialize()

	if err := validateNewDrsTaskV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DrsTaskV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.drsTaskV3.DrsTaskV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/drs_task_v3 opentelekomcloud_drs_task_v3} Resource.
func NewDrsTaskV3_Override(d DrsTaskV3, scope constructs.Construct, id *string, config *DrsTaskV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.drsTaskV3.DrsTaskV3",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetDestinationDbReadonly(val interface{}) {
	if err := j.validateSetDestinationDbReadonlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationDbReadonly",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetEngineType(val *string) {
	if err := j.validateSetEngineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineType",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetExpiredDays(val *float64) {
	if err := j.validateSetExpiredDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiredDays",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetMigrateDefiner(val interface{}) {
	if err := j.validateSetMigrateDefinerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migrateDefiner",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetMigrationType(val *string) {
	if err := j.validateSetMigrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migrationType",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetMultiWrite(val interface{}) {
	if err := j.validateSetMultiWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiWrite",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetNetType(val *string) {
	if err := j.validateSetNetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netType",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetNodeNum(val *float64) {
	if err := j.validateSetNodeNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeNum",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a DrsTaskV3 resource upon running "cdktf plan <stack-name>".
func DrsTaskV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDrsTaskV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.drsTaskV3.DrsTaskV3",
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
func DrsTaskV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDrsTaskV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.drsTaskV3.DrsTaskV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DrsTaskV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDrsTaskV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.drsTaskV3.DrsTaskV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DrsTaskV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDrsTaskV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.drsTaskV3.DrsTaskV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DrsTaskV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.drsTaskV3.DrsTaskV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DrsTaskV3) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DrsTaskV3) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DrsTaskV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DrsTaskV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DrsTaskV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DrsTaskV3) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DrsTaskV3) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DrsTaskV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DrsTaskV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DrsTaskV3) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DrsTaskV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DrsTaskV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsTaskV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DrsTaskV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DrsTaskV3) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DrsTaskV3) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DrsTaskV3) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DrsTaskV3) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DrsTaskV3) PutDestinationDb(value *DrsTaskV3DestinationDb) {
	if err := d.validatePutDestinationDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDestinationDb",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DrsTaskV3) PutLimitSpeed(value interface{}) {
	if err := d.validatePutLimitSpeedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLimitSpeed",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DrsTaskV3) PutSourceDb(value *DrsTaskV3SourceDb) {
	if err := d.validatePutSourceDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSourceDb",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DrsTaskV3) PutTimeouts(value *DrsTaskV3Timeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetDestinationDbReadonly() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationDbReadonly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetExpiredDays() {
	_jsii_.InvokeVoid(
		d,
		"resetExpiredDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		d,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetLimitSpeed() {
	_jsii_.InvokeVoid(
		d,
		"resetLimitSpeed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetMigrateDefiner() {
	_jsii_.InvokeVoid(
		d,
		"resetMigrateDefiner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetMigrationType() {
	_jsii_.InvokeVoid(
		d,
		"resetMigrationType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetMultiWrite() {
	_jsii_.InvokeVoid(
		d,
		"resetMultiWrite",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetNetType() {
	_jsii_.InvokeVoid(
		d,
		"resetNetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetNodeNum() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeNum",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetStartTime() {
	_jsii_.InvokeVoid(
		d,
		"resetStartTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsTaskV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsTaskV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsTaskV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsTaskV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsTaskV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

