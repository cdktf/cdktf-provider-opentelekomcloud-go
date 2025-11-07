// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ceseventreportv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/ceseventreportv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CesEventReportV1DetailOutputReference interface {
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
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EventLevel() *string
	SetEventLevel(val *string)
	EventLevelInput() *string
	EventState() *string
	SetEventState(val *string)
	EventStateInput() *string
	EventType() *string
	SetEventType(val *string)
	EventTypeInput() *string
	EventUser() *string
	SetEventUser(val *string)
	EventUserInput() *string
	// Experimental.
	Fqn() *string
	GroupId() *string
	SetGroupId(val *string)
	GroupIdInput() *string
	InternalValue() *CesEventReportV1Detail
	SetInternalValue(val *CesEventReportV1Detail)
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ResourceName() *string
	SetResourceName(val *string)
	ResourceNameInput() *string
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
	ResetContent()
	ResetEventLevel()
	ResetEventState()
	ResetEventType()
	ResetEventUser()
	ResetGroupId()
	ResetResourceId()
	ResetResourceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesEventReportV1DetailOutputReference
type jsiiProxy_CesEventReportV1DetailOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) EventLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) EventLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) EventState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) EventStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) EventType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) EventTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) EventUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) EventUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) GroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) InternalValue() *CesEventReportV1Detail {
	var returns *CesEventReportV1Detail
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) ResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesEventReportV1DetailOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesEventReportV1DetailOutputReference {
	_init_.Initialize()

	if err := validateNewCesEventReportV1DetailOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesEventReportV1DetailOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cesEventReportV1.CesEventReportV1DetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesEventReportV1DetailOutputReference_Override(c CesEventReportV1DetailOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cesEventReportV1.CesEventReportV1DetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetEventLevel(val *string) {
	if err := j.validateSetEventLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventLevel",
		val,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetEventState(val *string) {
	if err := j.validateSetEventStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventState",
		val,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetEventType(val *string) {
	if err := j.validateSetEventTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventType",
		val,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetEventUser(val *string) {
	if err := j.validateSetEventUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventUser",
		val,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetGroupId(val *string) {
	if err := j.validateSetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetInternalValue(val *CesEventReportV1Detail) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetResourceName(val *string) {
	if err := j.validateSetResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceName",
		val,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesEventReportV1DetailOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) ResetContent() {
	_jsii_.InvokeVoid(
		c,
		"resetContent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) ResetEventLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetEventLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) ResetEventState() {
	_jsii_.InvokeVoid(
		c,
		"resetEventState",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) ResetEventType() {
	_jsii_.InvokeVoid(
		c,
		"resetEventType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) ResetEventUser() {
	_jsii_.InvokeVoid(
		c,
		"resetEventUser",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) ResetGroupId() {
	_jsii_.InvokeVoid(
		c,
		"resetGroupId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) ResetResourceName() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesEventReportV1DetailOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

