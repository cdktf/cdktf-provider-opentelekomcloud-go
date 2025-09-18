// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logtanktransferv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/logtanktransferv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/logtank_transfer_v2 opentelekomcloud_logtank_transfer_v2}.
type LogtankTransferV2 interface {
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
	CreateTime() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DirPrefixName() *string
	SetDirPrefixName(val *string)
	DirPrefixNameInput() *string
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
	LogGroupId() *string
	SetLogGroupId(val *string)
	LogGroupIdInput() *string
	LogGroupName() *string
	LogStreamIds() *[]*string
	SetLogStreamIds(val *[]*string)
	LogStreamIdsInput() *[]*string
	LogTransferMode() *string
	LogTransferType() *string
	// The tree node.
	Node() constructs.Node
	ObsBucketName() *string
	SetObsBucketName(val *string)
	ObsBucketNameInput() *string
	ObsEncryptionEnable() cdktf.IResolvable
	ObsEncryptionId() *string
	Period() *float64
	SetPeriod(val *float64)
	PeriodInput() *float64
	PeriodUnit() *string
	SetPeriodUnit(val *string)
	PeriodUnitInput() *string
	PrefixName() *string
	SetPrefixName(val *string)
	PrefixNameInput() *string
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
	Status() *string
	StorageFormat() *string
	SetStorageFormat(val *string)
	StorageFormatInput() *string
	SwitchOn() interface{}
	SetSwitchOn(val interface{})
	SwitchOnInput() interface{}
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
	ResetDirPrefixName()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrefixName()
	ResetSwitchOn()
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

// The jsii proxy struct for LogtankTransferV2
type jsiiProxy_LogtankTransferV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LogtankTransferV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) CreateTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) DirPrefixName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dirPrefixName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) DirPrefixNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dirPrefixNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) LogGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) LogGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) LogStreamIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logStreamIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) LogStreamIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logStreamIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) LogTransferMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTransferMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) LogTransferType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTransferType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) ObsBucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsBucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) ObsBucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsBucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) ObsEncryptionEnable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"obsEncryptionEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) ObsEncryptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsEncryptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) PeriodUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) PeriodUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) PrefixName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) PrefixNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) StorageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) StorageFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) SwitchOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"switchOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) SwitchOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"switchOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogtankTransferV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/logtank_transfer_v2 opentelekomcloud_logtank_transfer_v2} Resource.
func NewLogtankTransferV2(scope constructs.Construct, id *string, config *LogtankTransferV2Config) LogtankTransferV2 {
	_init_.Initialize()

	if err := validateNewLogtankTransferV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogtankTransferV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.logtankTransferV2.LogtankTransferV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/logtank_transfer_v2 opentelekomcloud_logtank_transfer_v2} Resource.
func NewLogtankTransferV2_Override(l LogtankTransferV2, scope constructs.Construct, id *string, config *LogtankTransferV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.logtankTransferV2.LogtankTransferV2",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetDirPrefixName(val *string) {
	if err := j.validateSetDirPrefixNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dirPrefixName",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetLogGroupId(val *string) {
	if err := j.validateSetLogGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupId",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetLogStreamIds(val *[]*string) {
	if err := j.validateSetLogStreamIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamIds",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetObsBucketName(val *string) {
	if err := j.validateSetObsBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsBucketName",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetPeriod(val *float64) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetPeriodUnit(val *string) {
	if err := j.validateSetPeriodUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodUnit",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetPrefixName(val *string) {
	if err := j.validateSetPrefixNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixName",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetStorageFormat(val *string) {
	if err := j.validateSetStorageFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageFormat",
		val,
	)
}

func (j *jsiiProxy_LogtankTransferV2)SetSwitchOn(val interface{}) {
	if err := j.validateSetSwitchOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"switchOn",
		val,
	)
}

// Generates CDKTF code for importing a LogtankTransferV2 resource upon running "cdktf plan <stack-name>".
func LogtankTransferV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLogtankTransferV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.logtankTransferV2.LogtankTransferV2",
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
func LogtankTransferV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogtankTransferV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.logtankTransferV2.LogtankTransferV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogtankTransferV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogtankTransferV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.logtankTransferV2.LogtankTransferV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogtankTransferV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogtankTransferV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.logtankTransferV2.LogtankTransferV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LogtankTransferV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.logtankTransferV2.LogtankTransferV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LogtankTransferV2) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LogtankTransferV2) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LogtankTransferV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LogtankTransferV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogtankTransferV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LogtankTransferV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LogtankTransferV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LogtankTransferV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LogtankTransferV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LogtankTransferV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LogtankTransferV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LogtankTransferV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogtankTransferV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LogtankTransferV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogtankTransferV2) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogtankTransferV2) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LogtankTransferV2) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogtankTransferV2) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LogtankTransferV2) ResetDirPrefixName() {
	_jsii_.InvokeVoid(
		l,
		"resetDirPrefixName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogtankTransferV2) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogtankTransferV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogtankTransferV2) ResetPrefixName() {
	_jsii_.InvokeVoid(
		l,
		"resetPrefixName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogtankTransferV2) ResetSwitchOn() {
	_jsii_.InvokeVoid(
		l,
		"resetSwitchOn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogtankTransferV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogtankTransferV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogtankTransferV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogtankTransferV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogtankTransferV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogtankTransferV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

