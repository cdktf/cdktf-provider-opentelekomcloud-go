// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitypasswordpolicyv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/identitypasswordpolicyv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/identity_password_policy_v3 opentelekomcloud_identity_password_policy_v3}.
type IdentityPasswordPolicyV3 interface {
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
	MaximumConsecutiveIdenticalChars() *float64
	SetMaximumConsecutiveIdenticalChars(val *float64)
	MaximumConsecutiveIdenticalCharsInput() *float64
	MaximumPasswordLength() *float64
	MinimumPasswordAge() *float64
	SetMinimumPasswordAge(val *float64)
	MinimumPasswordAgeInput() *float64
	MinimumPasswordLength() *float64
	SetMinimumPasswordLength(val *float64)
	MinimumPasswordLengthInput() *float64
	// The tree node.
	Node() constructs.Node
	NumberOfRecentPasswordsDisallowed() *float64
	SetNumberOfRecentPasswordsDisallowed(val *float64)
	NumberOfRecentPasswordsDisallowedInput() *float64
	PasswordCharCombination() *float64
	SetPasswordCharCombination(val *float64)
	PasswordCharCombinationInput() *float64
	PasswordNotUsernameOrInvert() interface{}
	SetPasswordNotUsernameOrInvert(val interface{})
	PasswordNotUsernameOrInvertInput() interface{}
	PasswordRequirements() *string
	PasswordValidityPeriod() *float64
	SetPasswordValidityPeriod(val *float64)
	PasswordValidityPeriodInput() *float64
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
	ResetId()
	ResetMaximumConsecutiveIdenticalChars()
	ResetMinimumPasswordAge()
	ResetMinimumPasswordLength()
	ResetNumberOfRecentPasswordsDisallowed()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordCharCombination()
	ResetPasswordNotUsernameOrInvert()
	ResetPasswordValidityPeriod()
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

// The jsii proxy struct for IdentityPasswordPolicyV3
type jsiiProxy_IdentityPasswordPolicyV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) MaximumConsecutiveIdenticalChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumConsecutiveIdenticalChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) MaximumConsecutiveIdenticalCharsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumConsecutiveIdenticalCharsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) MaximumPasswordLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumPasswordLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) MinimumPasswordAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumPasswordAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) MinimumPasswordAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumPasswordAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) MinimumPasswordLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumPasswordLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) MinimumPasswordLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumPasswordLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) NumberOfRecentPasswordsDisallowed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfRecentPasswordsDisallowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) NumberOfRecentPasswordsDisallowedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfRecentPasswordsDisallowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) PasswordCharCombination() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordCharCombination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) PasswordCharCombinationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordCharCombinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) PasswordNotUsernameOrInvert() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordNotUsernameOrInvert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) PasswordNotUsernameOrInvertInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordNotUsernameOrInvertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) PasswordRequirements() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) PasswordValidityPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordValidityPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) PasswordValidityPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordValidityPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPasswordPolicyV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/identity_password_policy_v3 opentelekomcloud_identity_password_policy_v3} Resource.
func NewIdentityPasswordPolicyV3(scope constructs.Construct, id *string, config *IdentityPasswordPolicyV3Config) IdentityPasswordPolicyV3 {
	_init_.Initialize()

	if err := validateNewIdentityPasswordPolicyV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityPasswordPolicyV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.identityPasswordPolicyV3.IdentityPasswordPolicyV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/identity_password_policy_v3 opentelekomcloud_identity_password_policy_v3} Resource.
func NewIdentityPasswordPolicyV3_Override(i IdentityPasswordPolicyV3, scope constructs.Construct, id *string, config *IdentityPasswordPolicyV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.identityPasswordPolicyV3.IdentityPasswordPolicyV3",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetMaximumConsecutiveIdenticalChars(val *float64) {
	if err := j.validateSetMaximumConsecutiveIdenticalCharsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumConsecutiveIdenticalChars",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetMinimumPasswordAge(val *float64) {
	if err := j.validateSetMinimumPasswordAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumPasswordAge",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetMinimumPasswordLength(val *float64) {
	if err := j.validateSetMinimumPasswordLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumPasswordLength",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetNumberOfRecentPasswordsDisallowed(val *float64) {
	if err := j.validateSetNumberOfRecentPasswordsDisallowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfRecentPasswordsDisallowed",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetPasswordCharCombination(val *float64) {
	if err := j.validateSetPasswordCharCombinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordCharCombination",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetPasswordNotUsernameOrInvert(val interface{}) {
	if err := j.validateSetPasswordNotUsernameOrInvertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordNotUsernameOrInvert",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetPasswordValidityPeriod(val *float64) {
	if err := j.validateSetPasswordValidityPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordValidityPeriod",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IdentityPasswordPolicyV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a IdentityPasswordPolicyV3 resource upon running "cdktf plan <stack-name>".
func IdentityPasswordPolicyV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIdentityPasswordPolicyV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.identityPasswordPolicyV3.IdentityPasswordPolicyV3",
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
func IdentityPasswordPolicyV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityPasswordPolicyV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.identityPasswordPolicyV3.IdentityPasswordPolicyV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityPasswordPolicyV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityPasswordPolicyV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.identityPasswordPolicyV3.IdentityPasswordPolicyV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityPasswordPolicyV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityPasswordPolicyV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.identityPasswordPolicyV3.IdentityPasswordPolicyV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IdentityPasswordPolicyV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.identityPasswordPolicyV3.IdentityPasswordPolicyV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IdentityPasswordPolicyV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IdentityPasswordPolicyV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IdentityPasswordPolicyV3) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IdentityPasswordPolicyV3) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IdentityPasswordPolicyV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IdentityPasswordPolicyV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IdentityPasswordPolicyV3) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IdentityPasswordPolicyV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IdentityPasswordPolicyV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IdentityPasswordPolicyV3) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ResetMaximumConsecutiveIdenticalChars() {
	_jsii_.InvokeVoid(
		i,
		"resetMaximumConsecutiveIdenticalChars",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ResetMinimumPasswordAge() {
	_jsii_.InvokeVoid(
		i,
		"resetMinimumPasswordAge",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ResetMinimumPasswordLength() {
	_jsii_.InvokeVoid(
		i,
		"resetMinimumPasswordLength",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ResetNumberOfRecentPasswordsDisallowed() {
	_jsii_.InvokeVoid(
		i,
		"resetNumberOfRecentPasswordsDisallowed",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ResetPasswordCharCombination() {
	_jsii_.InvokeVoid(
		i,
		"resetPasswordCharCombination",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ResetPasswordNotUsernameOrInvert() {
	_jsii_.InvokeVoid(
		i,
		"resetPasswordNotUsernameOrInvert",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ResetPasswordValidityPeriod() {
	_jsii_.InvokeVoid(
		i,
		"resetPasswordValidityPeriod",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPasswordPolicyV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

