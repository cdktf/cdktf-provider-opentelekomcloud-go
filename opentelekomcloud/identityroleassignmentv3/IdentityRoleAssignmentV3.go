// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityroleassignmentv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v10/identityroleassignmentv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.12/docs/resources/identity_role_assignment_v3 opentelekomcloud_identity_role_assignment_v3}.
type IdentityRoleAssignmentV3 interface {
	cdktf.TerraformResource
	AllProjects() interface{}
	SetAllProjects(val interface{})
	AllProjectsInput() interface{}
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
	DomainId() *string
	SetDomainId(val *string)
	DomainIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupId() *string
	SetGroupId(val *string)
	GroupIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
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
	RoleId() *string
	SetRoleId(val *string)
	RoleIdInput() *string
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAllProjects()
	ResetDomainId()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for IdentityRoleAssignmentV3
type jsiiProxy_IdentityRoleAssignmentV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) AllProjects() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) AllProjectsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) DomainIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) GroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) RoleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityRoleAssignmentV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.12/docs/resources/identity_role_assignment_v3 opentelekomcloud_identity_role_assignment_v3} Resource.
func NewIdentityRoleAssignmentV3(scope constructs.Construct, id *string, config *IdentityRoleAssignmentV3Config) IdentityRoleAssignmentV3 {
	_init_.Initialize()

	if err := validateNewIdentityRoleAssignmentV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityRoleAssignmentV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.identityRoleAssignmentV3.IdentityRoleAssignmentV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.12/docs/resources/identity_role_assignment_v3 opentelekomcloud_identity_role_assignment_v3} Resource.
func NewIdentityRoleAssignmentV3_Override(i IdentityRoleAssignmentV3, scope constructs.Construct, id *string, config *IdentityRoleAssignmentV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.identityRoleAssignmentV3.IdentityRoleAssignmentV3",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetAllProjects(val interface{}) {
	if err := j.validateSetAllProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allProjects",
		val,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetDomainId(val *string) {
	if err := j.validateSetDomainIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainId",
		val,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetGroupId(val *string) {
	if err := j.validateSetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IdentityRoleAssignmentV3)SetRoleId(val *string) {
	if err := j.validateSetRoleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleId",
		val,
	)
}

// Generates CDKTF code for importing a IdentityRoleAssignmentV3 resource upon running "cdktf plan <stack-name>".
func IdentityRoleAssignmentV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIdentityRoleAssignmentV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.identityRoleAssignmentV3.IdentityRoleAssignmentV3",
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
func IdentityRoleAssignmentV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityRoleAssignmentV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.identityRoleAssignmentV3.IdentityRoleAssignmentV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityRoleAssignmentV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityRoleAssignmentV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.identityRoleAssignmentV3.IdentityRoleAssignmentV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityRoleAssignmentV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityRoleAssignmentV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.identityRoleAssignmentV3.IdentityRoleAssignmentV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IdentityRoleAssignmentV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.identityRoleAssignmentV3.IdentityRoleAssignmentV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IdentityRoleAssignmentV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IdentityRoleAssignmentV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IdentityRoleAssignmentV3) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IdentityRoleAssignmentV3) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IdentityRoleAssignmentV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IdentityRoleAssignmentV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IdentityRoleAssignmentV3) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IdentityRoleAssignmentV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IdentityRoleAssignmentV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IdentityRoleAssignmentV3) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) ResetAllProjects() {
	_jsii_.InvokeVoid(
		i,
		"resetAllProjects",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) ResetDomainId() {
	_jsii_.InvokeVoid(
		i,
		"resetDomainId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) ResetProjectId() {
	_jsii_.InvokeVoid(
		i,
		"resetProjectId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityRoleAssignmentV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

