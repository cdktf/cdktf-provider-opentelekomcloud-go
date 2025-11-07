// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltskeywordsalarmrulev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/ltskeywordsalarmrulev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference interface {
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
	Condition() *string
	SetCondition(val *string)
	ConditionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Keyword() *string
	SetKeyword(val *string)
	KeywordInput() *string
	LogGroupId() *string
	SetLogGroupId(val *string)
	LogGroupIdInput() *string
	LogStreamId() *string
	SetLogStreamId(val *string)
	LogStreamIdInput() *string
	Number() *float64
	SetNumber(val *float64)
	NumberInput() *float64
	SearchTimeRange() *float64
	SetSearchTimeRange(val *float64)
	SearchTimeRangeInput() *float64
	SearchTimeRangeUnit() *string
	SetSearchTimeRangeUnit(val *string)
	SearchTimeRangeUnitInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference
type jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) ConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) Keyword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) KeywordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keywordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) LogGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) LogGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) LogStreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) LogStreamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) Number() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"number",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) NumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) SearchTimeRange() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"searchTimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) SearchTimeRangeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"searchTimeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) SearchTimeRangeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchTimeRangeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) SearchTimeRangeUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchTimeRangeUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference {
	_init_.Initialize()

	if err := validateNewLtsKeywordsAlarmRuleV2KeywordsRequestsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference_Override(l LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference)SetCondition(val *string) {
	if err := j.validateSetConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference)SetKeyword(val *string) {
	if err := j.validateSetKeywordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyword",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference)SetLogGroupId(val *string) {
	if err := j.validateSetLogGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupId",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference)SetLogStreamId(val *string) {
	if err := j.validateSetLogStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamId",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference)SetNumber(val *float64) {
	if err := j.validateSetNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"number",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference)SetSearchTimeRange(val *float64) {
	if err := j.validateSetSearchTimeRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchTimeRange",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference)SetSearchTimeRangeUnit(val *string) {
	if err := j.validateSetSearchTimeRangeUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchTimeRangeUnit",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2KeywordsRequestsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

