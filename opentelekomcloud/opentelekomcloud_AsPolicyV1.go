// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v1 opentelekomcloud_as_policy_v1}.
type AsPolicyV1 interface {
	cdktf.TerraformResource
	AlarmId() *string
	SetAlarmId(val *string)
	AlarmIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CoolDownTime() *float64
	SetCoolDownTime(val *float64)
	CoolDownTimeInput() *float64
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ScalingGroupId() *string
	SetScalingGroupId(val *string)
	ScalingGroupIdInput() *string
	ScalingPolicyAction() AsPolicyV1ScalingPolicyActionOutputReference
	ScalingPolicyActionInput() *AsPolicyV1ScalingPolicyAction
	ScalingPolicyName() *string
	SetScalingPolicyName(val *string)
	ScalingPolicyNameInput() *string
	ScalingPolicyType() *string
	SetScalingPolicyType(val *string)
	ScalingPolicyTypeInput() *string
	ScheduledPolicy() AsPolicyV1ScheduledPolicyOutputReference
	ScheduledPolicyInput() *AsPolicyV1ScheduledPolicy
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutScalingPolicyAction(value *AsPolicyV1ScalingPolicyAction)
	PutScheduledPolicy(value *AsPolicyV1ScheduledPolicy)
	ResetAlarmId()
	ResetCoolDownTime()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetScalingPolicyAction()
	ResetScheduledPolicy()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AsPolicyV1
type jsiiProxy_AsPolicyV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AsPolicyV1) AlarmId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) AlarmIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) CoolDownTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coolDownTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) CoolDownTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coolDownTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) ScalingGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) ScalingGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) ScalingPolicyAction() AsPolicyV1ScalingPolicyActionOutputReference {
	var returns AsPolicyV1ScalingPolicyActionOutputReference
	_jsii_.Get(
		j,
		"scalingPolicyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) ScalingPolicyActionInput() *AsPolicyV1ScalingPolicyAction {
	var returns *AsPolicyV1ScalingPolicyAction
	_jsii_.Get(
		j,
		"scalingPolicyActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) ScalingPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) ScalingPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) ScalingPolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) ScalingPolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) ScheduledPolicy() AsPolicyV1ScheduledPolicyOutputReference {
	var returns AsPolicyV1ScheduledPolicyOutputReference
	_jsii_.Get(
		j,
		"scheduledPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) ScheduledPolicyInput() *AsPolicyV1ScheduledPolicy {
	var returns *AsPolicyV1ScheduledPolicy
	_jsii_.Get(
		j,
		"scheduledPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v1 opentelekomcloud_as_policy_v1} Resource.
func NewAsPolicyV1(scope constructs.Construct, id *string, config *AsPolicyV1Config) AsPolicyV1 {
	_init_.Initialize()

	j := jsiiProxy_AsPolicyV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.AsPolicyV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v1 opentelekomcloud_as_policy_v1} Resource.
func NewAsPolicyV1_Override(a AsPolicyV1, scope constructs.Construct, id *string, config *AsPolicyV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.AsPolicyV1",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetAlarmId(val *string) {
	_jsii_.Set(
		j,
		"alarmId",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetCoolDownTime(val *float64) {
	_jsii_.Set(
		j,
		"coolDownTime",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetScalingGroupId(val *string) {
	_jsii_.Set(
		j,
		"scalingGroupId",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetScalingPolicyName(val *string) {
	_jsii_.Set(
		j,
		"scalingPolicyName",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV1) SetScalingPolicyType(val *string) {
	_jsii_.Set(
		j,
		"scalingPolicyType",
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
func AsPolicyV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.AsPolicyV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AsPolicyV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.AsPolicyV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AsPolicyV1) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AsPolicyV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AsPolicyV1) PutScalingPolicyAction(value *AsPolicyV1ScalingPolicyAction) {
	_jsii_.InvokeVoid(
		a,
		"putScalingPolicyAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AsPolicyV1) PutScheduledPolicy(value *AsPolicyV1ScheduledPolicy) {
	_jsii_.InvokeVoid(
		a,
		"putScheduledPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AsPolicyV1) ResetAlarmId() {
	_jsii_.InvokeVoid(
		a,
		"resetAlarmId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsPolicyV1) ResetCoolDownTime() {
	_jsii_.InvokeVoid(
		a,
		"resetCoolDownTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsPolicyV1) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsPolicyV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsPolicyV1) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsPolicyV1) ResetScalingPolicyAction() {
	_jsii_.InvokeVoid(
		a,
		"resetScalingPolicyAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsPolicyV1) ResetScheduledPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduledPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsPolicyV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
