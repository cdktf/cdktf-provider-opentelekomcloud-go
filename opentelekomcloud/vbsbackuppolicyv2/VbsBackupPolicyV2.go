package vbsbackuppolicyv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v6/vbsbackuppolicyv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/vbs_backup_policy_v2 opentelekomcloud_vbs_backup_policy_v2}.
type VbsBackupPolicyV2 interface {
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
	Frequency() *float64
	SetFrequency(val *float64)
	FrequencyInput() *float64
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PolicyResourceCount() *float64
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
	SetRegion(val *string)
	RegionInput() *string
	RententionDay() *float64
	SetRententionDay(val *float64)
	RententionDayInput() *float64
	RententionNum() *float64
	SetRententionNum(val *float64)
	RententionNumInput() *float64
	Resources() *[]*string
	SetResources(val *[]*string)
	ResourcesInput() *[]*string
	RetainFirstBackup() *string
	SetRetainFirstBackup(val *string)
	RetainFirstBackupInput() *string
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tags() VbsBackupPolicyV2TagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VbsBackupPolicyV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	WeekFrequency() *[]*string
	SetWeekFrequency(val *[]*string)
	WeekFrequencyInput() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTags(value interface{})
	PutTimeouts(value *VbsBackupPolicyV2Timeouts)
	ResetFrequency()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetRententionDay()
	ResetRententionNum()
	ResetResources()
	ResetStatus()
	ResetTags()
	ResetTimeouts()
	ResetWeekFrequency()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VbsBackupPolicyV2
type jsiiProxy_VbsBackupPolicyV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VbsBackupPolicyV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Frequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) FrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) PolicyResourceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyResourceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) RententionDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rententionDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) RententionDayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rententionDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) RententionNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rententionNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) RententionNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rententionNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Resources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) ResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) RetainFirstBackup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retainFirstBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) RetainFirstBackupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retainFirstBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Tags() VbsBackupPolicyV2TagsList {
	var returns VbsBackupPolicyV2TagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) Timeouts() VbsBackupPolicyV2TimeoutsOutputReference {
	var returns VbsBackupPolicyV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) WeekFrequency() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VbsBackupPolicyV2) WeekFrequencyInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekFrequencyInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/vbs_backup_policy_v2 opentelekomcloud_vbs_backup_policy_v2} Resource.
func NewVbsBackupPolicyV2(scope constructs.Construct, id *string, config *VbsBackupPolicyV2Config) VbsBackupPolicyV2 {
	_init_.Initialize()

	if err := validateNewVbsBackupPolicyV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VbsBackupPolicyV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.vbsBackupPolicyV2.VbsBackupPolicyV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/vbs_backup_policy_v2 opentelekomcloud_vbs_backup_policy_v2} Resource.
func NewVbsBackupPolicyV2_Override(v VbsBackupPolicyV2, scope constructs.Construct, id *string, config *VbsBackupPolicyV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.vbsBackupPolicyV2.VbsBackupPolicyV2",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetFrequency(val *float64) {
	if err := j.validateSetFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetRententionDay(val *float64) {
	if err := j.validateSetRententionDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rententionDay",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetRententionNum(val *float64) {
	if err := j.validateSetRententionNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rententionNum",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetResources(val *[]*string) {
	if err := j.validateSetResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resources",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetRetainFirstBackup(val *string) {
	if err := j.validateSetRetainFirstBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainFirstBackup",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_VbsBackupPolicyV2)SetWeekFrequency(val *[]*string) {
	if err := j.validateSetWeekFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekFrequency",
		val,
	)
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
func VbsBackupPolicyV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVbsBackupPolicyV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.vbsBackupPolicyV2.VbsBackupPolicyV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VbsBackupPolicyV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVbsBackupPolicyV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.vbsBackupPolicyV2.VbsBackupPolicyV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VbsBackupPolicyV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVbsBackupPolicyV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.vbsBackupPolicyV2.VbsBackupPolicyV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VbsBackupPolicyV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.vbsBackupPolicyV2.VbsBackupPolicyV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) PutTags(value interface{}) {
	if err := v.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTags",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) PutTimeouts(value *VbsBackupPolicyV2Timeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) ResetFrequency() {
	_jsii_.InvokeVoid(
		v,
		"resetFrequency",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) ResetRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) ResetRententionDay() {
	_jsii_.InvokeVoid(
		v,
		"resetRententionDay",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) ResetRententionNum() {
	_jsii_.InvokeVoid(
		v,
		"resetRententionNum",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) ResetResources() {
	_jsii_.InvokeVoid(
		v,
		"resetResources",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) ResetStatus() {
	_jsii_.InvokeVoid(
		v,
		"resetStatus",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) ResetWeekFrequency() {
	_jsii_.InvokeVoid(
		v,
		"resetWeekFrequency",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VbsBackupPolicyV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VbsBackupPolicyV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

