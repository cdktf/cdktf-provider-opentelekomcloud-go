// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltskeywordsalarmrulev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/ltskeywordsalarmrulev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference interface {
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	PushPolicy() *float64
	SetPushPolicy(val *float64)
	PushPolicyInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TopicUrn() *string
	SetTopicUrn(val *string)
	TopicUrnInput() *string
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
	ResetDisplayName()
	ResetPushPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference
type jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) PushPolicy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pushPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) PushPolicyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pushPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) TopicUrn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicUrn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) TopicUrnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicUrnInput",
		&returns,
	)
	return returns
}


func NewLtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference {
	_init_.Initialize()

	if err := validateNewLtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference_Override(l LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsKeywordsAlarmRuleV2.LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference)SetPushPolicy(val *float64) {
	if err := j.validateSetPushPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushPolicy",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference)SetTopicUrn(val *string) {
	if err := j.validateSetTopicUrnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicUrn",
		val,
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		l,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) ResetPushPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetPushPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LtsKeywordsAlarmRuleV2NotificationRuleTopicsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

