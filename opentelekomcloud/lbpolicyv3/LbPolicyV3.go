// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbpolicyv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/lbpolicyv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/lb_policy_v3 opentelekomcloud_lb_policy_v3}.
type LbPolicyV3 interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	FixedResponseConfig() LbPolicyV3FixedResponseConfigOutputReference
	FixedResponseConfigInput() *LbPolicyV3FixedResponseConfig
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
	ListenerId() *string
	SetListenerId(val *string)
	ListenerIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Position() *float64
	SetPosition(val *float64)
	PositionInput() *float64
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
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
	RedirectListenerId() *string
	SetRedirectListenerId(val *string)
	RedirectListenerIdInput() *string
	RedirectPoolId() *string
	SetRedirectPoolId(val *string)
	RedirectPoolIdInput() *string
	RedirectPoolsConfig() LbPolicyV3RedirectPoolsConfigList
	RedirectPoolsConfigInput() interface{}
	RedirectUrl() *string
	SetRedirectUrl(val *string)
	RedirectUrlConfig() LbPolicyV3RedirectUrlConfigOutputReference
	RedirectUrlConfigInput() *LbPolicyV3RedirectUrlConfig
	RedirectUrlInput() *string
	Rules() LbPolicyV3RulesList
	RulesInput() interface{}
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutFixedResponseConfig(value *LbPolicyV3FixedResponseConfig)
	PutRedirectPoolsConfig(value interface{})
	PutRedirectUrlConfig(value *LbPolicyV3RedirectUrlConfig)
	PutRules(value interface{})
	ResetDescription()
	ResetFixedResponseConfig()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPosition()
	ResetPriority()
	ResetProjectId()
	ResetRedirectListenerId()
	ResetRedirectPoolId()
	ResetRedirectPoolsConfig()
	ResetRedirectUrl()
	ResetRedirectUrlConfig()
	ResetRules()
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

// The jsii proxy struct for LbPolicyV3
type jsiiProxy_LbPolicyV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LbPolicyV3) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) FixedResponseConfig() LbPolicyV3FixedResponseConfigOutputReference {
	var returns LbPolicyV3FixedResponseConfigOutputReference
	_jsii_.Get(
		j,
		"fixedResponseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) FixedResponseConfigInput() *LbPolicyV3FixedResponseConfig {
	var returns *LbPolicyV3FixedResponseConfig
	_jsii_.Get(
		j,
		"fixedResponseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) ListenerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) ListenerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Position() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"position",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) PositionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"positionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) RedirectListenerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectListenerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) RedirectListenerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectListenerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) RedirectPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) RedirectPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) RedirectPoolsConfig() LbPolicyV3RedirectPoolsConfigList {
	var returns LbPolicyV3RedirectPoolsConfigList
	_jsii_.Get(
		j,
		"redirectPoolsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) RedirectPoolsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectPoolsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) RedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) RedirectUrlConfig() LbPolicyV3RedirectUrlConfigOutputReference {
	var returns LbPolicyV3RedirectUrlConfigOutputReference
	_jsii_.Get(
		j,
		"redirectUrlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) RedirectUrlConfigInput() *LbPolicyV3RedirectUrlConfig {
	var returns *LbPolicyV3RedirectUrlConfig
	_jsii_.Get(
		j,
		"redirectUrlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) RedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Rules() LbPolicyV3RulesList {
	var returns LbPolicyV3RulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) RulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbPolicyV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/lb_policy_v3 opentelekomcloud_lb_policy_v3} Resource.
func NewLbPolicyV3(scope constructs.Construct, id *string, config *LbPolicyV3Config) LbPolicyV3 {
	_init_.Initialize()

	if err := validateNewLbPolicyV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbPolicyV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.lbPolicyV3.LbPolicyV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/lb_policy_v3 opentelekomcloud_lb_policy_v3} Resource.
func NewLbPolicyV3_Override(l LbPolicyV3, scope constructs.Construct, id *string, config *LbPolicyV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.lbPolicyV3.LbPolicyV3",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetListenerId(val *string) {
	if err := j.validateSetListenerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listenerId",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetPosition(val *float64) {
	if err := j.validateSetPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"position",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetRedirectListenerId(val *string) {
	if err := j.validateSetRedirectListenerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectListenerId",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetRedirectPoolId(val *string) {
	if err := j.validateSetRedirectPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectPoolId",
		val,
	)
}

func (j *jsiiProxy_LbPolicyV3)SetRedirectUrl(val *string) {
	if err := j.validateSetRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectUrl",
		val,
	)
}

// Generates CDKTF code for importing a LbPolicyV3 resource upon running "cdktf plan <stack-name>".
func LbPolicyV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLbPolicyV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.lbPolicyV3.LbPolicyV3",
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
func LbPolicyV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbPolicyV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.lbPolicyV3.LbPolicyV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbPolicyV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbPolicyV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.lbPolicyV3.LbPolicyV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbPolicyV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbPolicyV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.lbPolicyV3.LbPolicyV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LbPolicyV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.lbPolicyV3.LbPolicyV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LbPolicyV3) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LbPolicyV3) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LbPolicyV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LbPolicyV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbPolicyV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LbPolicyV3) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LbPolicyV3) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LbPolicyV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LbPolicyV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LbPolicyV3) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LbPolicyV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LbPolicyV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbPolicyV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LbPolicyV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbPolicyV3) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LbPolicyV3) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LbPolicyV3) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LbPolicyV3) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LbPolicyV3) PutFixedResponseConfig(value *LbPolicyV3FixedResponseConfig) {
	if err := l.validatePutFixedResponseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFixedResponseConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbPolicyV3) PutRedirectPoolsConfig(value interface{}) {
	if err := l.validatePutRedirectPoolsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRedirectPoolsConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbPolicyV3) PutRedirectUrlConfig(value *LbPolicyV3RedirectUrlConfig) {
	if err := l.validatePutRedirectUrlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRedirectUrlConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbPolicyV3) PutRules(value interface{}) {
	if err := l.validatePutRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRules",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetFixedResponseConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetFixedResponseConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetName() {
	_jsii_.InvokeVoid(
		l,
		"resetName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetPosition() {
	_jsii_.InvokeVoid(
		l,
		"resetPosition",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetPriority() {
	_jsii_.InvokeVoid(
		l,
		"resetPriority",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetProjectId() {
	_jsii_.InvokeVoid(
		l,
		"resetProjectId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetRedirectListenerId() {
	_jsii_.InvokeVoid(
		l,
		"resetRedirectListenerId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetRedirectPoolId() {
	_jsii_.InvokeVoid(
		l,
		"resetRedirectPoolId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetRedirectPoolsConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetRedirectPoolsConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetRedirectUrl() {
	_jsii_.InvokeVoid(
		l,
		"resetRedirectUrl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetRedirectUrlConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetRedirectUrlConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) ResetRules() {
	_jsii_.InvokeVoid(
		l,
		"resetRules",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbPolicyV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbPolicyV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbPolicyV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbPolicyV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbPolicyV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbPolicyV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

