// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityloginpolicyv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/identityloginpolicyv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/identity_login_policy_v3 opentelekomcloud_identity_login_policy_v3}.
type IdentityLoginPolicyV3 interface {
	cdktf.TerraformResource
	AccountValidityPeriod() *float64
	SetAccountValidityPeriod(val *float64)
	AccountValidityPeriodInput() *float64
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
	CustomInfoForLogin() *string
	SetCustomInfoForLogin(val *string)
	CustomInfoForLoginInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	LockoutDuration() *float64
	SetLockoutDuration(val *float64)
	LockoutDurationInput() *float64
	LoginFailedTimes() *float64
	SetLoginFailedTimes(val *float64)
	LoginFailedTimesInput() *float64
	// The tree node.
	Node() constructs.Node
	PeriodWithLoginFailures() *float64
	SetPeriodWithLoginFailures(val *float64)
	PeriodWithLoginFailuresInput() *float64
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
	SessionTimeout() *float64
	SetSessionTimeout(val *float64)
	SessionTimeoutInput() *float64
	ShowRecentLoginInfo() interface{}
	SetShowRecentLoginInfo(val interface{})
	ShowRecentLoginInfoInput() interface{}
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
	ResetAccountValidityPeriod()
	ResetCustomInfoForLogin()
	ResetId()
	ResetLockoutDuration()
	ResetLoginFailedTimes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeriodWithLoginFailures()
	ResetSessionTimeout()
	ResetShowRecentLoginInfo()
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

// The jsii proxy struct for IdentityLoginPolicyV3
type jsiiProxy_IdentityLoginPolicyV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IdentityLoginPolicyV3) AccountValidityPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountValidityPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) AccountValidityPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountValidityPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) CustomInfoForLogin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInfoForLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) CustomInfoForLoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInfoForLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) LockoutDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockoutDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) LockoutDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockoutDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) LoginFailedTimes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loginFailedTimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) LoginFailedTimesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loginFailedTimesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) PeriodWithLoginFailures() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodWithLoginFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) PeriodWithLoginFailuresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodWithLoginFailuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) SessionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) SessionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) ShowRecentLoginInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showRecentLoginInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) ShowRecentLoginInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showRecentLoginInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityLoginPolicyV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/identity_login_policy_v3 opentelekomcloud_identity_login_policy_v3} Resource.
func NewIdentityLoginPolicyV3(scope constructs.Construct, id *string, config *IdentityLoginPolicyV3Config) IdentityLoginPolicyV3 {
	_init_.Initialize()

	if err := validateNewIdentityLoginPolicyV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityLoginPolicyV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.identityLoginPolicyV3.IdentityLoginPolicyV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/identity_login_policy_v3 opentelekomcloud_identity_login_policy_v3} Resource.
func NewIdentityLoginPolicyV3_Override(i IdentityLoginPolicyV3, scope constructs.Construct, id *string, config *IdentityLoginPolicyV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.identityLoginPolicyV3.IdentityLoginPolicyV3",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetAccountValidityPeriod(val *float64) {
	if err := j.validateSetAccountValidityPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountValidityPeriod",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetCustomInfoForLogin(val *string) {
	if err := j.validateSetCustomInfoForLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customInfoForLogin",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetLockoutDuration(val *float64) {
	if err := j.validateSetLockoutDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockoutDuration",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetLoginFailedTimes(val *float64) {
	if err := j.validateSetLoginFailedTimesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginFailedTimes",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetPeriodWithLoginFailures(val *float64) {
	if err := j.validateSetPeriodWithLoginFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodWithLoginFailures",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetSessionTimeout(val *float64) {
	if err := j.validateSetSessionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeout",
		val,
	)
}

func (j *jsiiProxy_IdentityLoginPolicyV3)SetShowRecentLoginInfo(val interface{}) {
	if err := j.validateSetShowRecentLoginInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showRecentLoginInfo",
		val,
	)
}

// Generates CDKTF code for importing a IdentityLoginPolicyV3 resource upon running "cdktf plan <stack-name>".
func IdentityLoginPolicyV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIdentityLoginPolicyV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.identityLoginPolicyV3.IdentityLoginPolicyV3",
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
func IdentityLoginPolicyV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityLoginPolicyV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.identityLoginPolicyV3.IdentityLoginPolicyV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityLoginPolicyV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityLoginPolicyV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.identityLoginPolicyV3.IdentityLoginPolicyV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityLoginPolicyV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityLoginPolicyV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.identityLoginPolicyV3.IdentityLoginPolicyV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IdentityLoginPolicyV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.identityLoginPolicyV3.IdentityLoginPolicyV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ResetAccountValidityPeriod() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountValidityPeriod",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ResetCustomInfoForLogin() {
	_jsii_.InvokeVoid(
		i,
		"resetCustomInfoForLogin",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ResetLockoutDuration() {
	_jsii_.InvokeVoid(
		i,
		"resetLockoutDuration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ResetLoginFailedTimes() {
	_jsii_.InvokeVoid(
		i,
		"resetLoginFailedTimes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ResetPeriodWithLoginFailures() {
	_jsii_.InvokeVoid(
		i,
		"resetPeriodWithLoginFailures",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ResetSessionTimeout() {
	_jsii_.InvokeVoid(
		i,
		"resetSessionTimeout",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ResetShowRecentLoginInfo() {
	_jsii_.InvokeVoid(
		i,
		"resetShowRecentLoginInfo",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityLoginPolicyV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityLoginPolicyV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

