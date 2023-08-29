// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cbrpolicyv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v9/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v9/cbrpolicyv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CbrPolicyV3OperationDefinitionOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *CbrPolicyV3OperationDefinition
	SetInternalValue(val *CbrPolicyV3OperationDefinition)
	MaxBackups() *float64
	SetMaxBackups(val *float64)
	MaxBackupsInput() *float64
	MonthBackups() *float64
	SetMonthBackups(val *float64)
	MonthBackupsInput() *float64
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
	ResetMaxBackups()
	ResetMonthBackups()
	ResetRetentionDurationDays()
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

// The jsii proxy struct for CbrPolicyV3OperationDefinitionOutputReference
type jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) DayBackups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) DayBackupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) InternalValue() *CbrPolicyV3OperationDefinition {
	var returns *CbrPolicyV3OperationDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) MaxBackups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) MaxBackupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) MonthBackups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) MonthBackupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) RetentionDurationDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDurationDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) RetentionDurationDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDurationDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) WeekBackups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) WeekBackupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) YearBackups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yearBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) YearBackupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yearBackupsInput",
		&returns,
	)
	return returns
}


func NewCbrPolicyV3OperationDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CbrPolicyV3OperationDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewCbrPolicyV3OperationDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cbrPolicyV3.CbrPolicyV3OperationDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCbrPolicyV3OperationDefinitionOutputReference_Override(c CbrPolicyV3OperationDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cbrPolicyV3.CbrPolicyV3OperationDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference)SetDayBackups(val *float64) {
	if err := j.validateSetDayBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayBackups",
		val,
	)
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference)SetInternalValue(val *CbrPolicyV3OperationDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference)SetMaxBackups(val *float64) {
	if err := j.validateSetMaxBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBackups",
		val,
	)
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference)SetMonthBackups(val *float64) {
	if err := j.validateSetMonthBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthBackups",
		val,
	)
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference)SetRetentionDurationDays(val *float64) {
	if err := j.validateSetRetentionDurationDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionDurationDays",
		val,
	)
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference)SetWeekBackups(val *float64) {
	if err := j.validateSetWeekBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekBackups",
		val,
	)
}

func (j *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference)SetYearBackups(val *float64) {
	if err := j.validateSetYearBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yearBackups",
		val,
	)
}

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) ResetDayBackups() {
	_jsii_.InvokeVoid(
		c,
		"resetDayBackups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) ResetMaxBackups() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxBackups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) ResetMonthBackups() {
	_jsii_.InvokeVoid(
		c,
		"resetMonthBackups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) ResetRetentionDurationDays() {
	_jsii_.InvokeVoid(
		c,
		"resetRetentionDurationDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) ResetWeekBackups() {
	_jsii_.InvokeVoid(
		c,
		"resetWeekBackups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) ResetYearBackups() {
	_jsii_.InvokeVoid(
		c,
		"resetYearBackups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CbrPolicyV3OperationDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

