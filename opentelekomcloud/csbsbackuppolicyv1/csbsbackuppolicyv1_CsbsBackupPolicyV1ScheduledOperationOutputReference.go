package csbsbackuppolicyv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v3/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v3/csbsbackuppolicyv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CsbsBackupPolicyV1ScheduledOperationOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DayBackups() *float64
	SetDayBackups(val *float64)
	DayBackupsInput() *float64
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *CsbsBackupPolicyV1ScheduledOperation
	SetInternalValue(val *CsbsBackupPolicyV1ScheduledOperation)
	MaxBackups() *float64
	SetMaxBackups(val *float64)
	MaxBackupsInput() *float64
	MonthBackups() *float64
	SetMonthBackups(val *float64)
	MonthBackupsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	OperationType() *string
	SetOperationType(val *string)
	OperationTypeInput() *string
	Permanent() interface{}
	SetPermanent(val interface{})
	PermanentInput() interface{}
	RetentionDurationDays() *float64
	SetRetentionDurationDays(val *float64)
	RetentionDurationDaysInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	TriggerId() *string
	TriggerName() *string
	TriggerPattern() *string
	SetTriggerPattern(val *string)
	TriggerPatternInput() *string
	TriggerType() *string
	WeekBackups() *float64
	SetWeekBackups(val *float64)
	WeekBackupsInput() *float64
	YearBackups() *float64
	SetYearBackups(val *float64)
	YearBackupsInput() *float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDayBackups()
	ResetDescription()
	ResetEnabled()
	ResetMaxBackups()
	ResetMonthBackups()
	ResetName()
	ResetPermanent()
	ResetRetentionDurationDays()
	ResetTimezone()
	ResetWeekBackups()
	ResetYearBackups()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CsbsBackupPolicyV1ScheduledOperationOutputReference
type jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) DayBackups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) DayBackupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) InternalValue() *CsbsBackupPolicyV1ScheduledOperation {
	var returns *CsbsBackupPolicyV1ScheduledOperation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) MaxBackups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) MaxBackupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) MonthBackups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) MonthBackupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) OperationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) OperationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) Permanent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permanent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) PermanentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permanentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) RetentionDurationDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDurationDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) RetentionDurationDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDurationDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) TriggerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) TriggerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) TriggerPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) TriggerPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) TriggerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) WeekBackups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) WeekBackupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) YearBackups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yearBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) YearBackupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yearBackupsInput",
		&returns,
	)
	return returns
}


func NewCsbsBackupPolicyV1ScheduledOperationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CsbsBackupPolicyV1ScheduledOperationOutputReference {
	_init_.Initialize()

	if err := validateNewCsbsBackupPolicyV1ScheduledOperationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.csbsBackupPolicyV1.CsbsBackupPolicyV1ScheduledOperationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCsbsBackupPolicyV1ScheduledOperationOutputReference_Override(c CsbsBackupPolicyV1ScheduledOperationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.csbsBackupPolicyV1.CsbsBackupPolicyV1ScheduledOperationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetDayBackups(val *float64) {
	if err := j.validateSetDayBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayBackups",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetInternalValue(val *CsbsBackupPolicyV1ScheduledOperation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetMaxBackups(val *float64) {
	if err := j.validateSetMaxBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBackups",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetMonthBackups(val *float64) {
	if err := j.validateSetMonthBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthBackups",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetOperationType(val *string) {
	if err := j.validateSetOperationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationType",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetPermanent(val interface{}) {
	if err := j.validateSetPermanentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permanent",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetRetentionDurationDays(val *float64) {
	if err := j.validateSetRetentionDurationDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionDurationDays",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetTriggerPattern(val *string) {
	if err := j.validateSetTriggerPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerPattern",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetWeekBackups(val *float64) {
	if err := j.validateSetWeekBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekBackups",
		val,
	)
}

func (j *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference)SetYearBackups(val *float64) {
	if err := j.validateSetYearBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yearBackups",
		val,
	)
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ResetDayBackups() {
	_jsii_.InvokeVoid(
		c,
		"resetDayBackups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ResetMaxBackups() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxBackups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ResetMonthBackups() {
	_jsii_.InvokeVoid(
		c,
		"resetMonthBackups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ResetPermanent() {
	_jsii_.InvokeVoid(
		c,
		"resetPermanent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ResetRetentionDurationDays() {
	_jsii_.InvokeVoid(
		c,
		"resetRetentionDurationDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		c,
		"resetTimezone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ResetWeekBackups() {
	_jsii_.InvokeVoid(
		c,
		"resetWeekBackups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ResetYearBackups() {
	_jsii_.InvokeVoid(
		c,
		"resetYearBackups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsbsBackupPolicyV1ScheduledOperationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

