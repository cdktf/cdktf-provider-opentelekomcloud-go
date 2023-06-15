package aspolicyv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v7/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v7/aspolicyv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AsPolicyV2ScheduledPolicyOutputReference interface {
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
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LaunchTime() *string
	SetLaunchTime(val *string)
	LaunchTimeInput() *string
	RecurrenceType() *string
	SetRecurrenceType(val *string)
	RecurrenceTypeInput() *string
	RecurrenceValue() *string
	SetRecurrenceValue(val *string)
	RecurrenceValueInput() *string
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetEndTime()
	ResetRecurrenceType()
	ResetRecurrenceValue()
	ResetStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AsPolicyV2ScheduledPolicyOutputReference
type jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) LaunchTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) LaunchTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) RecurrenceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) RecurrenceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) RecurrenceValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) RecurrenceValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAsPolicyV2ScheduledPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AsPolicyV2ScheduledPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewAsPolicyV2ScheduledPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.asPolicyV2.AsPolicyV2ScheduledPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAsPolicyV2ScheduledPolicyOutputReference_Override(a AsPolicyV2ScheduledPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.asPolicyV2.AsPolicyV2ScheduledPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference)SetEndTime(val *string) {
	if err := j.validateSetEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference)SetLaunchTime(val *string) {
	if err := j.validateSetLaunchTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchTime",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference)SetRecurrenceType(val *string) {
	if err := j.validateSetRecurrenceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrenceType",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference)SetRecurrenceValue(val *string) {
	if err := j.validateSetRecurrenceValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrenceValue",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) ResetEndTime() {
	_jsii_.InvokeVoid(
		a,
		"resetEndTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) ResetRecurrenceType() {
	_jsii_.InvokeVoid(
		a,
		"resetRecurrenceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) ResetRecurrenceValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRecurrenceValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsPolicyV2ScheduledPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

