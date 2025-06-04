// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltskeywordsalarmrulev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/ltskeywordsalarmrulev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsKeywordsAlarmRuleV2FrequencyOutputReference interface {
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
	CronExpression() *string
	SetCronExpression(val *string)
	CronExpressionInput() *string
	DayOfWeek() *float64
	SetDayOfWeek(val *float64)
	DayOfWeekInput() *float64
	FixedRate() *float64
	SetFixedRate(val *float64)
	FixedRateInput() *float64
	FixedRateUnit() *string
	SetFixedRateUnit(val *string)
	FixedRateUnitInput() *string
	// Experimental.
	Fqn() *string
	HourOfDay() *float64
	SetHourOfDay(val *float64)
	HourOfDayInput() *float64
	InternalValue() *LtsKeywordsAlarmRuleV2Frequency
	SetInternalValue(val *LtsKeywordsAlarmRuleV2Frequency)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetCronExpression()
	ResetDayOfWeek()
	ResetFixedRate()
	ResetFixedRateUnit()
	ResetHourOfDay()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LtsKeywordsAlarmRuleV2FrequencyOutputReference
type jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) CronExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) CronExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) DayOfWeek() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) DayOfWeekInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) FixedRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fixedRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) FixedRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fixedRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) FixedRateUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedRateUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) FixedRateUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedRateUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) HourOfDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) HourOfDayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) InternalValue() *LtsKeywordsAlarmRuleV2Frequency {
	var returns *LtsKeywordsAlarmRuleV2Frequency
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewLtsKeywordsAlarmRuleV2FrequencyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LtsKeywordsAlarmRuleV2FrequencyOutputReference {
	_init_.Initialize()

	if err := validateNewLtsKeywordsAlarmRuleV2FrequencyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2FrequencyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLtsKeywordsAlarmRuleV2FrequencyOutputReference_Override(l LtsKeywordsAlarmRuleV2FrequencyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2FrequencyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference)SetCronExpression(val *string) {
	if err := j.validateSetCronExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cronExpression",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference)SetDayOfWeek(val *float64) {
	if err := j.validateSetDayOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayOfWeek",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference)SetFixedRate(val *float64) {
	if err := j.validateSetFixedRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedRate",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference)SetFixedRateUnit(val *string) {
	if err := j.validateSetFixedRateUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedRateUnit",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference)SetHourOfDay(val *float64) {
	if err := j.validateSetHourOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hourOfDay",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference)SetInternalValue(val *LtsKeywordsAlarmRuleV2Frequency) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) ResetCronExpression() {
	_jsii_.InvokeVoid(
		l,
		"resetCronExpression",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) ResetDayOfWeek() {
	_jsii_.InvokeVoid(
		l,
		"resetDayOfWeek",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) ResetFixedRate() {
	_jsii_.InvokeVoid(
		l,
		"resetFixedRate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) ResetFixedRateUnit() {
	_jsii_.InvokeVoid(
		l,
		"resetFixedRateUnit",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) ResetHourOfDay() {
	_jsii_.InvokeVoid(
		l,
		"resetHourOfDay",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2FrequencyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

