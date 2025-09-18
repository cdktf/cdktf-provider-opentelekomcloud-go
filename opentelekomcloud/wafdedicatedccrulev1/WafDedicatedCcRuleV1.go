// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedccrulev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/wafdedicatedccrulev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/waf_dedicated_cc_rule_v1 opentelekomcloud_waf_dedicated_cc_rule_v1}.
type WafDedicatedCcRuleV1 interface {
	cdktf.TerraformResource
	Action() WafDedicatedCcRuleV1ActionList
	ActionInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Conditions() WafDedicatedCcRuleV1ConditionsList
	ConditionsInput() interface{}
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
	CreatedAt() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	LimitNum() *float64
	SetLimitNum(val *float64)
	LimitNumInput() *float64
	LimitPeriod() *float64
	SetLimitPeriod(val *float64)
	LimitPeriodInput() *float64
	LockTime() *float64
	SetLockTime(val *float64)
	LockTimeInput() *float64
	Mode() *float64
	SetMode(val *float64)
	ModeInput() *float64
	// The tree node.
	Node() constructs.Node
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
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
	Status() *float64
	TagCategory() *string
	SetTagCategory(val *string)
	TagCategoryInput() *string
	TagContents() *[]*string
	SetTagContents(val *[]*string)
	TagContentsInput() *[]*string
	TagIndex() *string
	SetTagIndex(val *string)
	TagIndexInput() *string
	TagType() *string
	SetTagType(val *string)
	TagTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() WafDedicatedCcRuleV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	UnlockNum() *float64
	SetUnlockNum(val *float64)
	UnlockNumInput() *float64
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	PutAction(value interface{})
	PutConditions(value interface{})
	PutTimeouts(value *WafDedicatedCcRuleV1Timeouts)
	ResetConditions()
	ResetDescription()
	ResetId()
	ResetLockTime()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTagCategory()
	ResetTagContents()
	ResetTagIndex()
	ResetTimeouts()
	ResetUnlockNum()
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

// The jsii proxy struct for WafDedicatedCcRuleV1
type jsiiProxy_WafDedicatedCcRuleV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Action() WafDedicatedCcRuleV1ActionList {
	var returns WafDedicatedCcRuleV1ActionList
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) ActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Conditions() WafDedicatedCcRuleV1ConditionsList {
	var returns WafDedicatedCcRuleV1ConditionsList
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) ConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) LimitNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) LimitNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) LimitPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) LimitPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) LockTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) LockTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Mode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) ModeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Status() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) TagCategory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) TagCategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) TagContents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagContents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) TagContentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagContentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) TagIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) TagIndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) TagType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) TagTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Timeouts() WafDedicatedCcRuleV1TimeoutsOutputReference {
	var returns WafDedicatedCcRuleV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) UnlockNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unlockNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) UnlockNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unlockNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedCcRuleV1) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/waf_dedicated_cc_rule_v1 opentelekomcloud_waf_dedicated_cc_rule_v1} Resource.
func NewWafDedicatedCcRuleV1(scope constructs.Construct, id *string, config *WafDedicatedCcRuleV1Config) WafDedicatedCcRuleV1 {
	_init_.Initialize()

	if err := validateNewWafDedicatedCcRuleV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafDedicatedCcRuleV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafDedicatedCcRuleV1.WafDedicatedCcRuleV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/waf_dedicated_cc_rule_v1 opentelekomcloud_waf_dedicated_cc_rule_v1} Resource.
func NewWafDedicatedCcRuleV1_Override(w WafDedicatedCcRuleV1, scope constructs.Construct, id *string, config *WafDedicatedCcRuleV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafDedicatedCcRuleV1.WafDedicatedCcRuleV1",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetLimitNum(val *float64) {
	if err := j.validateSetLimitNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitNum",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetLimitPeriod(val *float64) {
	if err := j.validateSetLimitPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitPeriod",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetLockTime(val *float64) {
	if err := j.validateSetLockTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockTime",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetMode(val *float64) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetTagCategory(val *string) {
	if err := j.validateSetTagCategoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagCategory",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetTagContents(val *[]*string) {
	if err := j.validateSetTagContentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagContents",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetTagIndex(val *string) {
	if err := j.validateSetTagIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagIndex",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetTagType(val *string) {
	if err := j.validateSetTagTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagType",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetUnlockNum(val *float64) {
	if err := j.validateSetUnlockNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unlockNum",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedCcRuleV1)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

// Generates CDKTF code for importing a WafDedicatedCcRuleV1 resource upon running "cdktf plan <stack-name>".
func WafDedicatedCcRuleV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWafDedicatedCcRuleV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.wafDedicatedCcRuleV1.WafDedicatedCcRuleV1",
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
func WafDedicatedCcRuleV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafDedicatedCcRuleV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.wafDedicatedCcRuleV1.WafDedicatedCcRuleV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WafDedicatedCcRuleV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafDedicatedCcRuleV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.wafDedicatedCcRuleV1.WafDedicatedCcRuleV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WafDedicatedCcRuleV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafDedicatedCcRuleV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.wafDedicatedCcRuleV1.WafDedicatedCcRuleV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WafDedicatedCcRuleV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.wafDedicatedCcRuleV1.WafDedicatedCcRuleV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) PutAction(value interface{}) {
	if err := w.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAction",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) PutConditions(value interface{}) {
	if err := w.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putConditions",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) PutTimeouts(value *WafDedicatedCcRuleV1Timeouts) {
	if err := w.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ResetConditions() {
	_jsii_.InvokeVoid(
		w,
		"resetConditions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ResetDescription() {
	_jsii_.InvokeVoid(
		w,
		"resetDescription",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ResetLockTime() {
	_jsii_.InvokeVoid(
		w,
		"resetLockTime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ResetTagCategory() {
	_jsii_.InvokeVoid(
		w,
		"resetTagCategory",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ResetTagContents() {
	_jsii_.InvokeVoid(
		w,
		"resetTagContents",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ResetTagIndex() {
	_jsii_.InvokeVoid(
		w,
		"resetTagIndex",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ResetUnlockNum() {
	_jsii_.InvokeVoid(
		w,
		"resetUnlockNum",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedCcRuleV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

