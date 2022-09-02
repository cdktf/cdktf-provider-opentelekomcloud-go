// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/ces_alarmrule opentelekomcloud_ces_alarmrule}.
type CesAlarmrule interface {
	cdktf.TerraformResource
	AlarmActionEnabled() interface{}
	SetAlarmActionEnabled(val interface{})
	AlarmActionEnabledInput() interface{}
	AlarmActions() CesAlarmruleAlarmActionsList
	AlarmActionsInput() interface{}
	AlarmDescription() *string
	SetAlarmDescription(val *string)
	AlarmDescriptionInput() *string
	AlarmEnabled() interface{}
	SetAlarmEnabled(val interface{})
	AlarmEnabledInput() interface{}
	AlarmLevel() *float64
	SetAlarmLevel(val *float64)
	AlarmLevelInput() *float64
	AlarmName() *string
	SetAlarmName(val *string)
	AlarmNameInput() *string
	AlarmState() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Condition() CesAlarmruleConditionOutputReference
	ConditionInput() *CesAlarmruleCondition
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
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
	Metric() CesAlarmruleMetricOutputReference
	MetricInput() *CesAlarmruleMetric
	// The tree node.
	Node() constructs.Node
	OkActions() CesAlarmruleOkActionsList
	OkActionsInput() interface{}
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
	Timeouts() CesAlarmruleTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *float64
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
	PutAlarmActions(value interface{})
	PutCondition(value *CesAlarmruleCondition)
	PutMetric(value *CesAlarmruleMetric)
	PutOkActions(value interface{})
	PutTimeouts(value *CesAlarmruleTimeouts)
	ResetAlarmActionEnabled()
	ResetAlarmActions()
	ResetAlarmDescription()
	ResetAlarmEnabled()
	ResetAlarmLevel()
	ResetId()
	ResetOkActions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CesAlarmrule
type jsiiProxy_CesAlarmrule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CesAlarmrule) AlarmActionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmActionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) AlarmActionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmActionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) AlarmActions() CesAlarmruleAlarmActionsList {
	var returns CesAlarmruleAlarmActionsList
	_jsii_.Get(
		j,
		"alarmActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) AlarmActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) AlarmDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) AlarmDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) AlarmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) AlarmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) AlarmLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"alarmLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) AlarmLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"alarmLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) AlarmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) AlarmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) AlarmState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) Condition() CesAlarmruleConditionOutputReference {
	var returns CesAlarmruleConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) ConditionInput() *CesAlarmruleCondition {
	var returns *CesAlarmruleCondition
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) Metric() CesAlarmruleMetricOutputReference {
	var returns CesAlarmruleMetricOutputReference
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) MetricInput() *CesAlarmruleMetric {
	var returns *CesAlarmruleMetric
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) OkActions() CesAlarmruleOkActionsList {
	var returns CesAlarmruleOkActionsList
	_jsii_.Get(
		j,
		"okActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) OkActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"okActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) Timeouts() CesAlarmruleTimeoutsOutputReference {
	var returns CesAlarmruleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAlarmrule) UpdateTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/ces_alarmrule opentelekomcloud_ces_alarmrule} Resource.
func NewCesAlarmrule(scope constructs.Construct, id *string, config *CesAlarmruleConfig) CesAlarmrule {
	_init_.Initialize()

	if err := validateNewCesAlarmruleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAlarmrule{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.CesAlarmrule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/ces_alarmrule opentelekomcloud_ces_alarmrule} Resource.
func NewCesAlarmrule_Override(c CesAlarmrule, scope constructs.Construct, id *string, config *CesAlarmruleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.CesAlarmrule",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetAlarmActionEnabled(val interface{}) {
	if err := j.validateSetAlarmActionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmActionEnabled",
		val,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetAlarmDescription(val *string) {
	if err := j.validateSetAlarmDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmDescription",
		val,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetAlarmEnabled(val interface{}) {
	if err := j.validateSetAlarmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmEnabled",
		val,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetAlarmLevel(val *float64) {
	if err := j.validateSetAlarmLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmLevel",
		val,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetAlarmName(val *string) {
	if err := j.validateSetAlarmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmName",
		val,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CesAlarmrule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func CesAlarmrule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCesAlarmrule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.CesAlarmrule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CesAlarmrule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.CesAlarmrule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CesAlarmrule) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CesAlarmrule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CesAlarmrule) PutAlarmActions(value interface{}) {
	if err := c.validatePutAlarmActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAlarmActions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAlarmrule) PutCondition(value *CesAlarmruleCondition) {
	if err := c.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAlarmrule) PutMetric(value *CesAlarmruleMetric) {
	if err := c.validatePutMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMetric",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAlarmrule) PutOkActions(value interface{}) {
	if err := c.validatePutOkActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOkActions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAlarmrule) PutTimeouts(value *CesAlarmruleTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAlarmrule) ResetAlarmActionEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmActionEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAlarmrule) ResetAlarmActions() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmActions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAlarmrule) ResetAlarmDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAlarmrule) ResetAlarmEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAlarmrule) ResetAlarmLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAlarmrule) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAlarmrule) ResetOkActions() {
	_jsii_.InvokeVoid(
		c,
		"resetOkActions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAlarmrule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAlarmrule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAlarmrule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAlarmrule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

