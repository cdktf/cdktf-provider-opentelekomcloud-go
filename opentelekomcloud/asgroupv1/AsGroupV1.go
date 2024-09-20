// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asgroupv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/asgroupv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/as_group_v1 opentelekomcloud_as_group_v1}.
type AsGroupV1 interface {
	cdktf.TerraformResource
	AvailableZones() *[]*string
	SetAvailableZones(val *[]*string)
	AvailableZonesInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CoolDownTime() *float64
	SetCoolDownTime(val *float64)
	CoolDownTimeInput() *float64
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CurrentInstanceNumber() *float64
	DeleteInstances() *string
	SetDeleteInstances(val *string)
	DeleteInstancesInput() *string
	DeletePublicip() interface{}
	SetDeletePublicip(val interface{})
	DeletePublicipInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesireInstanceNumber() *float64
	SetDesireInstanceNumber(val *float64)
	DesireInstanceNumberInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthPeriodicAuditGracePeriod() *float64
	SetHealthPeriodicAuditGracePeriod(val *float64)
	HealthPeriodicAuditGracePeriodInput() *float64
	HealthPeriodicAuditMethod() *string
	SetHealthPeriodicAuditMethod(val *string)
	HealthPeriodicAuditMethodInput() *string
	HealthPeriodicAuditTime() *float64
	SetHealthPeriodicAuditTime(val *float64)
	HealthPeriodicAuditTimeInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	Instances() *[]*string
	InstanceTerminatePolicy() *string
	SetInstanceTerminatePolicy(val *string)
	InstanceTerminatePolicyInput() *string
	LbaasListeners() AsGroupV1LbaasListenersList
	LbaasListenersInput() interface{}
	LbListenerId() *string
	SetLbListenerId(val *string)
	LbListenerIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxInstanceNumber() *float64
	SetMaxInstanceNumber(val *float64)
	MaxInstanceNumberInput() *float64
	MinInstanceNumber() *float64
	SetMinInstanceNumber(val *float64)
	MinInstanceNumberInput() *float64
	Networks() AsGroupV1NetworksList
	NetworksInput() interface{}
	// The tree node.
	Node() constructs.Node
	Notifications() *[]*string
	SetNotifications(val *[]*string)
	NotificationsInput() *[]*string
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
	ScalingConfigurationId() *string
	SetScalingConfigurationId(val *string)
	ScalingConfigurationIdInput() *string
	ScalingGroupName() *string
	SetScalingGroupName(val *string)
	ScalingGroupNameInput() *string
	SecurityGroups() AsGroupV1SecurityGroupsOutputReference
	SecurityGroupsInput() *AsGroupV1SecurityGroups
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
	Timeouts() AsGroupV1TimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutLbaasListeners(value interface{})
	PutNetworks(value interface{})
	PutSecurityGroups(value *AsGroupV1SecurityGroups)
	PutTimeouts(value *AsGroupV1Timeouts)
	ResetAvailableZones()
	ResetCoolDownTime()
	ResetDesireInstanceNumber()
	ResetHealthPeriodicAuditGracePeriod()
	ResetHealthPeriodicAuditMethod()
	ResetHealthPeriodicAuditTime()
	ResetId()
	ResetInstanceTerminatePolicy()
	ResetLbaasListeners()
	ResetLbListenerId()
	ResetMaxInstanceNumber()
	ResetMinInstanceNumber()
	ResetNotifications()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetScalingConfigurationId()
	ResetSecurityGroups()
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

// The jsii proxy struct for AsGroupV1
type jsiiProxy_AsGroupV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AsGroupV1) AvailableZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) AvailableZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) CoolDownTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coolDownTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) CoolDownTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coolDownTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) CurrentInstanceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentInstanceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) DeleteInstances() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) DeleteInstancesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) DeletePublicip() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletePublicip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) DeletePublicipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletePublicipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) DesireInstanceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desireInstanceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) DesireInstanceNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desireInstanceNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) HealthPeriodicAuditGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthPeriodicAuditGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) HealthPeriodicAuditGracePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthPeriodicAuditGracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) HealthPeriodicAuditMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthPeriodicAuditMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) HealthPeriodicAuditMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthPeriodicAuditMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) HealthPeriodicAuditTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthPeriodicAuditTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) HealthPeriodicAuditTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthPeriodicAuditTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Instances() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) InstanceTerminatePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTerminatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) InstanceTerminatePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTerminatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) LbaasListeners() AsGroupV1LbaasListenersList {
	var returns AsGroupV1LbaasListenersList
	_jsii_.Get(
		j,
		"lbaasListeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) LbaasListenersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbaasListenersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) LbListenerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lbListenerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) LbListenerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lbListenerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) MaxInstanceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) MaxInstanceNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) MinInstanceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstanceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) MinInstanceNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstanceNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Networks() AsGroupV1NetworksList {
	var returns AsGroupV1NetworksList
	_jsii_.Get(
		j,
		"networks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) NetworksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Notifications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) NotificationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) ScalingConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) ScalingConfigurationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingConfigurationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) ScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) ScalingGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) SecurityGroups() AsGroupV1SecurityGroupsOutputReference {
	var returns AsGroupV1SecurityGroupsOutputReference
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) SecurityGroupsInput() *AsGroupV1SecurityGroups {
	var returns *AsGroupV1SecurityGroups
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) Timeouts() AsGroupV1TimeoutsOutputReference {
	var returns AsGroupV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsGroupV1) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/as_group_v1 opentelekomcloud_as_group_v1} Resource.
func NewAsGroupV1(scope constructs.Construct, id *string, config *AsGroupV1Config) AsGroupV1 {
	_init_.Initialize()

	if err := validateNewAsGroupV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AsGroupV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.asGroupV1.AsGroupV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/as_group_v1 opentelekomcloud_as_group_v1} Resource.
func NewAsGroupV1_Override(a AsGroupV1, scope constructs.Construct, id *string, config *AsGroupV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.asGroupV1.AsGroupV1",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AsGroupV1)SetAvailableZones(val *[]*string) {
	if err := j.validateSetAvailableZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableZones",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetCoolDownTime(val *float64) {
	if err := j.validateSetCoolDownTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coolDownTime",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetDeleteInstances(val *string) {
	if err := j.validateSetDeleteInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteInstances",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetDeletePublicip(val interface{}) {
	if err := j.validateSetDeletePublicipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletePublicip",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetDesireInstanceNumber(val *float64) {
	if err := j.validateSetDesireInstanceNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desireInstanceNumber",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetHealthPeriodicAuditGracePeriod(val *float64) {
	if err := j.validateSetHealthPeriodicAuditGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthPeriodicAuditGracePeriod",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetHealthPeriodicAuditMethod(val *string) {
	if err := j.validateSetHealthPeriodicAuditMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthPeriodicAuditMethod",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetHealthPeriodicAuditTime(val *float64) {
	if err := j.validateSetHealthPeriodicAuditTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthPeriodicAuditTime",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetInstanceTerminatePolicy(val *string) {
	if err := j.validateSetInstanceTerminatePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTerminatePolicy",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetLbListenerId(val *string) {
	if err := j.validateSetLbListenerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lbListenerId",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetMaxInstanceNumber(val *float64) {
	if err := j.validateSetMaxInstanceNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstanceNumber",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetMinInstanceNumber(val *float64) {
	if err := j.validateSetMinInstanceNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInstanceNumber",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetNotifications(val *[]*string) {
	if err := j.validateSetNotificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifications",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetScalingConfigurationId(val *string) {
	if err := j.validateSetScalingConfigurationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingConfigurationId",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetScalingGroupName(val *string) {
	if err := j.validateSetScalingGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingGroupName",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AsGroupV1)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a AsGroupV1 resource upon running "cdktf plan <stack-name>".
func AsGroupV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAsGroupV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.asGroupV1.AsGroupV1",
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
func AsGroupV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAsGroupV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.asGroupV1.AsGroupV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AsGroupV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAsGroupV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.asGroupV1.AsGroupV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AsGroupV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAsGroupV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.asGroupV1.AsGroupV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AsGroupV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.asGroupV1.AsGroupV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AsGroupV1) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AsGroupV1) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AsGroupV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AsGroupV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AsGroupV1) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AsGroupV1) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AsGroupV1) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AsGroupV1) PutLbaasListeners(value interface{}) {
	if err := a.validatePutLbaasListenersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLbaasListeners",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AsGroupV1) PutNetworks(value interface{}) {
	if err := a.validatePutNetworksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworks",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AsGroupV1) PutSecurityGroups(value *AsGroupV1SecurityGroups) {
	if err := a.validatePutSecurityGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecurityGroups",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AsGroupV1) PutTimeouts(value *AsGroupV1Timeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AsGroupV1) ResetAvailableZones() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailableZones",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetCoolDownTime() {
	_jsii_.InvokeVoid(
		a,
		"resetCoolDownTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetDesireInstanceNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetDesireInstanceNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetHealthPeriodicAuditGracePeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthPeriodicAuditGracePeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetHealthPeriodicAuditMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthPeriodicAuditMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetHealthPeriodicAuditTime() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthPeriodicAuditTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetInstanceTerminatePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceTerminatePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetLbaasListeners() {
	_jsii_.InvokeVoid(
		a,
		"resetLbaasListeners",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetLbListenerId() {
	_jsii_.InvokeVoid(
		a,
		"resetLbListenerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetMaxInstanceNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxInstanceNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetMinInstanceNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetMinInstanceNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetNotifications() {
	_jsii_.InvokeVoid(
		a,
		"resetNotifications",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetScalingConfigurationId() {
	_jsii_.InvokeVoid(
		a,
		"resetScalingConfigurationId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsGroupV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsGroupV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

