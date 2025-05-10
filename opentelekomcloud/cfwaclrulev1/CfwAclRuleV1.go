// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwaclrulev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/cfwaclrulev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cfw_acl_rule_v1 opentelekomcloud_cfw_acl_rule_v1}.
type CfwAclRuleV1 interface {
	cdktf.TerraformResource
	ActionType() *float64
	SetActionType(val *float64)
	ActionTypeInput() *float64
	AddressType() *float64
	SetAddressType(val *float64)
	AddressTypeInput() *float64
	Applications() *[]*string
	SetApplications(val *[]*string)
	ApplicationsInput() *[]*string
	ApplicationsJsonString() *string
	SetApplicationsJsonString(val *string)
	ApplicationsJsonStringInput() *string
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
	CreatedDate() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Destination() CfwAclRuleV1DestinationOutputReference
	DestinationInput() *CfwAclRuleV1Destination
	Direction() *float64
	SetDirection(val *float64)
	DirectionInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	LastOpenTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LongConnectEnable() *float64
	SetLongConnectEnable(val *float64)
	LongConnectEnableInput() *float64
	LongConnectTime() *float64
	SetLongConnectTime(val *float64)
	LongConnectTimeHour() *float64
	SetLongConnectTimeHour(val *float64)
	LongConnectTimeHourInput() *float64
	LongConnectTimeInput() *float64
	LongConnectTimeMinute() *float64
	SetLongConnectTimeMinute(val *float64)
	LongConnectTimeMinuteInput() *float64
	LongConnectTimeSecond() *float64
	SetLongConnectTimeSecond(val *float64)
	LongConnectTimeSecondInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ObjectId() *string
	SetObjectId(val *string)
	ObjectIdInput() *string
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
	Sequence() CfwAclRuleV1SequenceOutputReference
	SequenceInput() *CfwAclRuleV1Sequence
	Service() CfwAclRuleV1ServiceOutputReference
	ServiceInput() *CfwAclRuleV1Service
	Source() CfwAclRuleV1SourceOutputReference
	SourceInput() *CfwAclRuleV1Source
	Status() *float64
	SetStatus(val *float64)
	StatusInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CfwAclRuleV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *float64
	SetType(val *float64)
	TypeInput() *float64
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
	PutDestination(value *CfwAclRuleV1Destination)
	PutSequence(value *CfwAclRuleV1Sequence)
	PutService(value *CfwAclRuleV1Service)
	PutSource(value *CfwAclRuleV1Source)
	PutTimeouts(value *CfwAclRuleV1Timeouts)
	ResetApplications()
	ResetApplicationsJsonString()
	ResetDescription()
	ResetDirection()
	ResetLongConnectTime()
	ResetLongConnectTimeHour()
	ResetLongConnectTimeMinute()
	ResetLongConnectTimeSecond()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for CfwAclRuleV1
type jsiiProxy_CfwAclRuleV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CfwAclRuleV1) ActionType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) ActionTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) AddressType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) AddressTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Applications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) ApplicationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) ApplicationsJsonString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationsJsonString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) ApplicationsJsonStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationsJsonStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Destination() CfwAclRuleV1DestinationOutputReference {
	var returns CfwAclRuleV1DestinationOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) DestinationInput() *CfwAclRuleV1Destination {
	var returns *CfwAclRuleV1Destination
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Direction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) DirectionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) LastOpenTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastOpenTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) LongConnectEnable() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longConnectEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) LongConnectEnableInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longConnectEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) LongConnectTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longConnectTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) LongConnectTimeHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longConnectTimeHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) LongConnectTimeHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longConnectTimeHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) LongConnectTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longConnectTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) LongConnectTimeMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longConnectTimeMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) LongConnectTimeMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longConnectTimeMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) LongConnectTimeSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longConnectTimeSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) LongConnectTimeSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longConnectTimeSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) ObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Sequence() CfwAclRuleV1SequenceOutputReference {
	var returns CfwAclRuleV1SequenceOutputReference
	_jsii_.Get(
		j,
		"sequence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) SequenceInput() *CfwAclRuleV1Sequence {
	var returns *CfwAclRuleV1Sequence
	_jsii_.Get(
		j,
		"sequenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Service() CfwAclRuleV1ServiceOutputReference {
	var returns CfwAclRuleV1ServiceOutputReference
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) ServiceInput() *CfwAclRuleV1Service {
	var returns *CfwAclRuleV1Service
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Source() CfwAclRuleV1SourceOutputReference {
	var returns CfwAclRuleV1SourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) SourceInput() *CfwAclRuleV1Source {
	var returns *CfwAclRuleV1Source
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Status() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) StatusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Timeouts() CfwAclRuleV1TimeoutsOutputReference {
	var returns CfwAclRuleV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) Type() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfwAclRuleV1) TypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cfw_acl_rule_v1 opentelekomcloud_cfw_acl_rule_v1} Resource.
func NewCfwAclRuleV1(scope constructs.Construct, id *string, config *CfwAclRuleV1Config) CfwAclRuleV1 {
	_init_.Initialize()

	if err := validateNewCfwAclRuleV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfwAclRuleV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cfw_acl_rule_v1 opentelekomcloud_cfw_acl_rule_v1} Resource.
func NewCfwAclRuleV1_Override(c CfwAclRuleV1, scope constructs.Construct, id *string, config *CfwAclRuleV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetActionType(val *float64) {
	if err := j.validateSetActionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionType",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetAddressType(val *float64) {
	if err := j.validateSetAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressType",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetApplications(val *[]*string) {
	if err := j.validateSetApplicationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applications",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetApplicationsJsonString(val *string) {
	if err := j.validateSetApplicationsJsonStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationsJsonString",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetDirection(val *float64) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetLongConnectEnable(val *float64) {
	if err := j.validateSetLongConnectEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longConnectEnable",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetLongConnectTime(val *float64) {
	if err := j.validateSetLongConnectTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longConnectTime",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetLongConnectTimeHour(val *float64) {
	if err := j.validateSetLongConnectTimeHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longConnectTimeHour",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetLongConnectTimeMinute(val *float64) {
	if err := j.validateSetLongConnectTimeMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longConnectTimeMinute",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetLongConnectTimeSecond(val *float64) {
	if err := j.validateSetLongConnectTimeSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longConnectTimeSecond",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetObjectId(val *string) {
	if err := j.validateSetObjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectId",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetStatus(val *float64) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CfwAclRuleV1)SetType(val *float64) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a CfwAclRuleV1 resource upon running "cdktf plan <stack-name>".
func CfwAclRuleV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCfwAclRuleV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1",
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
func CfwAclRuleV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfwAclRuleV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CfwAclRuleV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfwAclRuleV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CfwAclRuleV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfwAclRuleV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfwAclRuleV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.cfwAclRuleV1.CfwAclRuleV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CfwAclRuleV1) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfwAclRuleV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CfwAclRuleV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CfwAclRuleV1) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CfwAclRuleV1) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CfwAclRuleV1) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfwAclRuleV1) PutDestination(value *CfwAclRuleV1Destination) {
	if err := c.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDestination",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CfwAclRuleV1) PutSequence(value *CfwAclRuleV1Sequence) {
	if err := c.validatePutSequenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSequence",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CfwAclRuleV1) PutService(value *CfwAclRuleV1Service) {
	if err := c.validatePutServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putService",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CfwAclRuleV1) PutSource(value *CfwAclRuleV1Source) {
	if err := c.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CfwAclRuleV1) PutTimeouts(value *CfwAclRuleV1Timeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CfwAclRuleV1) ResetApplications() {
	_jsii_.InvokeVoid(
		c,
		"resetApplications",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1) ResetApplicationsJsonString() {
	_jsii_.InvokeVoid(
		c,
		"resetApplicationsJsonString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1) ResetDirection() {
	_jsii_.InvokeVoid(
		c,
		"resetDirection",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1) ResetLongConnectTime() {
	_jsii_.InvokeVoid(
		c,
		"resetLongConnectTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1) ResetLongConnectTimeHour() {
	_jsii_.InvokeVoid(
		c,
		"resetLongConnectTimeHour",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1) ResetLongConnectTimeMinute() {
	_jsii_.InvokeVoid(
		c,
		"resetLongConnectTimeMinute",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1) ResetLongConnectTimeSecond() {
	_jsii_.InvokeVoid(
		c,
		"resetLongConnectTimeSecond",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfwAclRuleV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfwAclRuleV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

