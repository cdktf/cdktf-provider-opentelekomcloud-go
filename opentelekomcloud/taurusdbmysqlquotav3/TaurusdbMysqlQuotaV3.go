// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package taurusdbmysqlquotav3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/taurusdbmysqlquotav3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/taurusdb_mysql_quota_v3 opentelekomcloud_taurusdb_mysql_quota_v3}.
type TaurusdbMysqlQuotaV3 interface {
	cdktf.TerraformResource
	AvailabilityInstanceQuota() *float64
	AvailabilityRamQuota() *float64
	AvailabilityVcpusQuota() *float64
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
	EnterpriseProjectId() *string
	SetEnterpriseProjectId(val *string)
	EnterpriseProjectIdInput() *string
	EnterpriseProjectName() *string
	SetEnterpriseProjectName(val *string)
	EnterpriseProjectNameInput() *string
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
	InstanceQuota() *float64
	SetInstanceQuota(val *float64)
	InstanceQuotaInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	RamQuota() *float64
	SetRamQuota(val *float64)
	RamQuotaInput() *float64
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VcpusQuota() *float64
	SetVcpusQuota(val *float64)
	VcpusQuotaInput() *float64
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
	ResetInstanceQuota()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRamQuota()
	ResetVcpusQuota()
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

// The jsii proxy struct for TaurusdbMysqlQuotaV3
type jsiiProxy_TaurusdbMysqlQuotaV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) AvailabilityInstanceQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityInstanceQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) AvailabilityRamQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityRamQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) AvailabilityVcpusQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityVcpusQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) EnterpriseProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) EnterpriseProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) EnterpriseProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseProjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) EnterpriseProjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseProjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) InstanceQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) InstanceQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) RamQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ramQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) RamQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ramQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) VcpusQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpusQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3) VcpusQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpusQuotaInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/taurusdb_mysql_quota_v3 opentelekomcloud_taurusdb_mysql_quota_v3} Resource.
func NewTaurusdbMysqlQuotaV3(scope constructs.Construct, id *string, config *TaurusdbMysqlQuotaV3Config) TaurusdbMysqlQuotaV3 {
	_init_.Initialize()

	if err := validateNewTaurusdbMysqlQuotaV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TaurusdbMysqlQuotaV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlQuotaV3.TaurusdbMysqlQuotaV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/taurusdb_mysql_quota_v3 opentelekomcloud_taurusdb_mysql_quota_v3} Resource.
func NewTaurusdbMysqlQuotaV3_Override(t TaurusdbMysqlQuotaV3, scope constructs.Construct, id *string, config *TaurusdbMysqlQuotaV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlQuotaV3.TaurusdbMysqlQuotaV3",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetEnterpriseProjectId(val *string) {
	if err := j.validateSetEnterpriseProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enterpriseProjectId",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetEnterpriseProjectName(val *string) {
	if err := j.validateSetEnterpriseProjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enterpriseProjectName",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetInstanceQuota(val *float64) {
	if err := j.validateSetInstanceQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceQuota",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetRamQuota(val *float64) {
	if err := j.validateSetRamQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ramQuota",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlQuotaV3)SetVcpusQuota(val *float64) {
	if err := j.validateSetVcpusQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcpusQuota",
		val,
	)
}

// Generates CDKTF code for importing a TaurusdbMysqlQuotaV3 resource upon running "cdktf plan <stack-name>".
func TaurusdbMysqlQuotaV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateTaurusdbMysqlQuotaV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlQuotaV3.TaurusdbMysqlQuotaV3",
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
func TaurusdbMysqlQuotaV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTaurusdbMysqlQuotaV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlQuotaV3.TaurusdbMysqlQuotaV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TaurusdbMysqlQuotaV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTaurusdbMysqlQuotaV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlQuotaV3.TaurusdbMysqlQuotaV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TaurusdbMysqlQuotaV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTaurusdbMysqlQuotaV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlQuotaV3.TaurusdbMysqlQuotaV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TaurusdbMysqlQuotaV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlQuotaV3.TaurusdbMysqlQuotaV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) ResetInstanceQuota() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceQuota",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) ResetRamQuota() {
	_jsii_.InvokeVoid(
		t,
		"resetRamQuota",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) ResetVcpusQuota() {
	_jsii_.InvokeVoid(
		t,
		"resetVcpusQuota",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlQuotaV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

