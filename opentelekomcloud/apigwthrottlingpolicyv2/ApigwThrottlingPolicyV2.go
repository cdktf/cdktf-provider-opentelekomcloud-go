// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwthrottlingpolicyv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/apigwthrottlingpolicyv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/apigw_throttling_policy_v2 opentelekomcloud_apigw_throttling_policy_v2}.
type ApigwThrottlingPolicyV2 interface {
	cdktf.TerraformResource
	AppThrottles() ApigwThrottlingPolicyV2AppThrottlesList
	AppThrottlesInput() interface{}
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
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxApiRequests() *float64
	SetMaxApiRequests(val *float64)
	MaxApiRequestsInput() *float64
	MaxAppRequests() *float64
	SetMaxAppRequests(val *float64)
	MaxAppRequestsInput() *float64
	MaxIpRequests() *float64
	SetMaxIpRequests(val *float64)
	MaxIpRequestsInput() *float64
	MaxUserRequests() *float64
	SetMaxUserRequests(val *float64)
	MaxUserRequestsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Period() *float64
	SetPeriod(val *float64)
	PeriodInput() *float64
	PeriodUnit() *string
	SetPeriodUnit(val *string)
	PeriodUnitInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UserThrottles() ApigwThrottlingPolicyV2UserThrottlesList
	UserThrottlesInput() interface{}
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
	PutAppThrottles(value interface{})
	PutUserThrottles(value interface{})
	ResetAppThrottles()
	ResetDescription()
	ResetId()
	ResetMaxAppRequests()
	ResetMaxIpRequests()
	ResetMaxUserRequests()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeriodUnit()
	ResetType()
	ResetUserThrottles()
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

// The jsii proxy struct for ApigwThrottlingPolicyV2
type jsiiProxy_ApigwThrottlingPolicyV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) AppThrottles() ApigwThrottlingPolicyV2AppThrottlesList {
	var returns ApigwThrottlingPolicyV2AppThrottlesList
	_jsii_.Get(
		j,
		"appThrottles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) AppThrottlesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appThrottlesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) MaxApiRequests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxApiRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) MaxApiRequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxApiRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) MaxAppRequests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAppRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) MaxAppRequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAppRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) MaxIpRequests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIpRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) MaxIpRequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIpRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) MaxUserRequests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUserRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) MaxUserRequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUserRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) PeriodUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) PeriodUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) UserThrottles() ApigwThrottlingPolicyV2UserThrottlesList {
	var returns ApigwThrottlingPolicyV2UserThrottlesList
	_jsii_.Get(
		j,
		"userThrottles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2) UserThrottlesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userThrottlesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/apigw_throttling_policy_v2 opentelekomcloud_apigw_throttling_policy_v2} Resource.
func NewApigwThrottlingPolicyV2(scope constructs.Construct, id *string, config *ApigwThrottlingPolicyV2Config) ApigwThrottlingPolicyV2 {
	_init_.Initialize()

	if err := validateNewApigwThrottlingPolicyV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigwThrottlingPolicyV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwThrottlingPolicyV2.ApigwThrottlingPolicyV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/apigw_throttling_policy_v2 opentelekomcloud_apigw_throttling_policy_v2} Resource.
func NewApigwThrottlingPolicyV2_Override(a ApigwThrottlingPolicyV2, scope constructs.Construct, id *string, config *ApigwThrottlingPolicyV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwThrottlingPolicyV2.ApigwThrottlingPolicyV2",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetMaxApiRequests(val *float64) {
	if err := j.validateSetMaxApiRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxApiRequests",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetMaxAppRequests(val *float64) {
	if err := j.validateSetMaxAppRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAppRequests",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetMaxIpRequests(val *float64) {
	if err := j.validateSetMaxIpRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIpRequests",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetMaxUserRequests(val *float64) {
	if err := j.validateSetMaxUserRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUserRequests",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetPeriod(val *float64) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetPeriodUnit(val *string) {
	if err := j.validateSetPeriodUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodUnit",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApigwThrottlingPolicyV2)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a ApigwThrottlingPolicyV2 resource upon running "cdktf plan <stack-name>".
func ApigwThrottlingPolicyV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApigwThrottlingPolicyV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.apigwThrottlingPolicyV2.ApigwThrottlingPolicyV2",
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
func ApigwThrottlingPolicyV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigwThrottlingPolicyV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.apigwThrottlingPolicyV2.ApigwThrottlingPolicyV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigwThrottlingPolicyV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigwThrottlingPolicyV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.apigwThrottlingPolicyV2.ApigwThrottlingPolicyV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigwThrottlingPolicyV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigwThrottlingPolicyV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.apigwThrottlingPolicyV2.ApigwThrottlingPolicyV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApigwThrottlingPolicyV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.apigwThrottlingPolicyV2.ApigwThrottlingPolicyV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigwThrottlingPolicyV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwThrottlingPolicyV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigwThrottlingPolicyV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigwThrottlingPolicyV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigwThrottlingPolicyV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigwThrottlingPolicyV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigwThrottlingPolicyV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigwThrottlingPolicyV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigwThrottlingPolicyV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigwThrottlingPolicyV2) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) PutAppThrottles(value interface{}) {
	if err := a.validatePutAppThrottlesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAppThrottles",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) PutUserThrottles(value interface{}) {
	if err := a.validatePutUserThrottlesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserThrottles",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ResetAppThrottles() {
	_jsii_.InvokeVoid(
		a,
		"resetAppThrottles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ResetMaxAppRequests() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxAppRequests",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ResetMaxIpRequests() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxIpRequests",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ResetMaxUserRequests() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxUserRequests",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ResetPeriodUnit() {
	_jsii_.InvokeVoid(
		a,
		"resetPeriodUnit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ResetUserThrottles() {
	_jsii_.InvokeVoid(
		a,
		"resetUserThrottles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwThrottlingPolicyV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

