// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltskeywordsalarmrulev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/ltskeywordsalarmrulev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/lts_keywords_alarm_rule_v2 opentelekomcloud_lts_keywords_alarm_rule_v2}.
type LtsKeywordsAlarmRuleV2 interface {
	cdktf.TerraformResource
	AlarmActionRuleName() *string
	SetAlarmActionRuleName(val *string)
	AlarmActionRuleNameInput() *string
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
	DomainId() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	Frequency() LtsKeywordsAlarmRuleV2FrequencyOutputReference
	FrequencyInput() *LtsKeywordsAlarmRuleV2Frequency
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KeywordsRequests() LtsKeywordsAlarmRuleV2KeywordsRequestsList
	KeywordsRequestsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationFrequency() *float64
	SetNotificationFrequency(val *float64)
	NotificationFrequencyInput() *float64
	NotificationRule() LtsKeywordsAlarmRuleV2NotificationRuleOutputReference
	NotificationRuleInput() *LtsKeywordsAlarmRuleV2NotificationRule
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
	RecoveryPolicy() *float64
	SetRecoveryPolicy(val *float64)
	RecoveryPolicyInput() *float64
	Region() *string
	SendNotifications() interface{}
	SetSendNotifications(val interface{})
	SendNotificationsInput() interface{}
	SendRecoveryNotifications() interface{}
	SetSendRecoveryNotifications(val interface{})
	SendRecoveryNotificationsInput() interface{}
	Severity() *string
	SetSeverity(val *string)
	SeverityInput() *string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TriggerConditionCount() *float64
	SetTriggerConditionCount(val *float64)
	TriggerConditionCountInput() *float64
	TriggerConditionFrequency() *float64
	SetTriggerConditionFrequency(val *float64)
	TriggerConditionFrequencyInput() *float64
	UpdatedAt() *string
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
	PutFrequency(value *LtsKeywordsAlarmRuleV2Frequency)
	PutKeywordsRequests(value interface{})
	PutNotificationRule(value *LtsKeywordsAlarmRuleV2NotificationRule)
	ResetAlarmActionRuleName()
	ResetDescription()
	ResetId()
	ResetNotificationRule()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRecoveryPolicy()
	ResetSendNotifications()
	ResetSendRecoveryNotifications()
	ResetTriggerConditionCount()
	ResetTriggerConditionFrequency()
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

// The jsii proxy struct for LtsKeywordsAlarmRuleV2
type jsiiProxy_LtsKeywordsAlarmRuleV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) AlarmActionRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmActionRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) AlarmActionRuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmActionRuleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Frequency() LtsKeywordsAlarmRuleV2FrequencyOutputReference {
	var returns LtsKeywordsAlarmRuleV2FrequencyOutputReference
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) FrequencyInput() *LtsKeywordsAlarmRuleV2Frequency {
	var returns *LtsKeywordsAlarmRuleV2Frequency
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) KeywordsRequests() LtsKeywordsAlarmRuleV2KeywordsRequestsList {
	var returns LtsKeywordsAlarmRuleV2KeywordsRequestsList
	_jsii_.Get(
		j,
		"keywordsRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) KeywordsRequestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keywordsRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) NotificationFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"notificationFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) NotificationFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"notificationFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) NotificationRule() LtsKeywordsAlarmRuleV2NotificationRuleOutputReference {
	var returns LtsKeywordsAlarmRuleV2NotificationRuleOutputReference
	_jsii_.Get(
		j,
		"notificationRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) NotificationRuleInput() *LtsKeywordsAlarmRuleV2NotificationRule {
	var returns *LtsKeywordsAlarmRuleV2NotificationRule
	_jsii_.Get(
		j,
		"notificationRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) RecoveryPolicy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoveryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) RecoveryPolicyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoveryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) SendNotifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) SendNotificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) SendRecoveryNotifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendRecoveryNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) SendRecoveryNotificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendRecoveryNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) TriggerConditionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"triggerConditionCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) TriggerConditionCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"triggerConditionCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) TriggerConditionFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"triggerConditionFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) TriggerConditionFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"triggerConditionFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/lts_keywords_alarm_rule_v2 opentelekomcloud_lts_keywords_alarm_rule_v2} Resource.
func NewLtsKeywordsAlarmRuleV2(scope constructs.Construct, id *string, config *LtsKeywordsAlarmRuleV2Config) LtsKeywordsAlarmRuleV2 {
	_init_.Initialize()

	if err := validateNewLtsKeywordsAlarmRuleV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LtsKeywordsAlarmRuleV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/lts_keywords_alarm_rule_v2 opentelekomcloud_lts_keywords_alarm_rule_v2} Resource.
func NewLtsKeywordsAlarmRuleV2_Override(l LtsKeywordsAlarmRuleV2, scope constructs.Construct, id *string, config *LtsKeywordsAlarmRuleV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetAlarmActionRuleName(val *string) {
	if err := j.validateSetAlarmActionRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmActionRuleName",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetNotificationFrequency(val *float64) {
	if err := j.validateSetNotificationFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationFrequency",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetRecoveryPolicy(val *float64) {
	if err := j.validateSetRecoveryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryPolicy",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetSendNotifications(val interface{}) {
	if err := j.validateSetSendNotificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendNotifications",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetSendRecoveryNotifications(val interface{}) {
	if err := j.validateSetSendRecoveryNotificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendRecoveryNotifications",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetTriggerConditionCount(val *float64) {
	if err := j.validateSetTriggerConditionCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerConditionCount",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2)SetTriggerConditionFrequency(val *float64) {
	if err := j.validateSetTriggerConditionFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerConditionFrequency",
		val,
	)
}

// Generates CDKTF code for importing a LtsKeywordsAlarmRuleV2 resource upon running "cdktf plan <stack-name>".
func LtsKeywordsAlarmRuleV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLtsKeywordsAlarmRuleV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2",
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
func LtsKeywordsAlarmRuleV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLtsKeywordsAlarmRuleV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LtsKeywordsAlarmRuleV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLtsKeywordsAlarmRuleV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LtsKeywordsAlarmRuleV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLtsKeywordsAlarmRuleV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LtsKeywordsAlarmRuleV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) PutFrequency(value *LtsKeywordsAlarmRuleV2Frequency) {
	if err := l.validatePutFrequencyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFrequency",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) PutKeywordsRequests(value interface{}) {
	if err := l.validatePutKeywordsRequestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putKeywordsRequests",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) PutNotificationRule(value *LtsKeywordsAlarmRuleV2NotificationRule) {
	if err := l.validatePutNotificationRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putNotificationRule",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ResetAlarmActionRuleName() {
	_jsii_.InvokeVoid(
		l,
		"resetAlarmActionRuleName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ResetNotificationRule() {
	_jsii_.InvokeVoid(
		l,
		"resetNotificationRule",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ResetRecoveryPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetRecoveryPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ResetSendNotifications() {
	_jsii_.InvokeVoid(
		l,
		"resetSendNotifications",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ResetSendRecoveryNotifications() {
	_jsii_.InvokeVoid(
		l,
		"resetSendRecoveryNotifications",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ResetTriggerConditionCount() {
	_jsii_.InvokeVoid(
		l,
		"resetTriggerConditionCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ResetTriggerConditionFrequency() {
	_jsii_.InvokeVoid(
		l,
		"resetTriggerConditionFrequency",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

