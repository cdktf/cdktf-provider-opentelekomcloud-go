// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltskeywordsalarmrulev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/ltskeywordsalarmrulev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsKeywordsAlarmRuleV2NotificationRuleOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *LtsKeywordsAlarmRuleV2NotificationRule
	SetInternalValue(val *LtsKeywordsAlarmRuleV2NotificationRule)
	Language() *string
	SetLanguage(val *string)
	LanguageInput() *string
	TemplateName() *string
	SetTemplateName(val *string)
	TemplateNameInput() *string
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
	Topics() LtsKeywordsAlarmRuleV2NotificationRuleTopicsList
	TopicsInput() interface{}
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
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
	PutTopics(value interface{})
	ResetLanguage()
	ResetTemplateName()
	ResetTimezone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LtsKeywordsAlarmRuleV2NotificationRuleOutputReference
type jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) InternalValue() *LtsKeywordsAlarmRuleV2NotificationRule {
	var returns *LtsKeywordsAlarmRuleV2NotificationRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) Language() *string {
	var returns *string
	_jsii_.Get(
		j,
		"language",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) LanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) TemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) TemplateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) Topics() LtsKeywordsAlarmRuleV2NotificationRuleTopicsList {
	var returns LtsKeywordsAlarmRuleV2NotificationRuleTopicsList
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) TopicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}


func NewLtsKeywordsAlarmRuleV2NotificationRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LtsKeywordsAlarmRuleV2NotificationRuleOutputReference {
	_init_.Initialize()

	if err := validateNewLtsKeywordsAlarmRuleV2NotificationRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2NotificationRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLtsKeywordsAlarmRuleV2NotificationRuleOutputReference_Override(l LtsKeywordsAlarmRuleV2NotificationRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2NotificationRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference)SetInternalValue(val *LtsKeywordsAlarmRuleV2NotificationRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference)SetLanguage(val *string) {
	if err := j.validateSetLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"language",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference)SetTemplateName(val *string) {
	if err := j.validateSetTemplateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateName",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference)SetUserName(val *string) {
	if err := j.validateSetUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) PutTopics(value interface{}) {
	if err := l.validatePutTopicsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTopics",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) ResetLanguage() {
	_jsii_.InvokeVoid(
		l,
		"resetLanguage",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) ResetTemplateName() {
	_jsii_.InvokeVoid(
		l,
		"resetTemplateName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		l,
		"resetTimezone",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

