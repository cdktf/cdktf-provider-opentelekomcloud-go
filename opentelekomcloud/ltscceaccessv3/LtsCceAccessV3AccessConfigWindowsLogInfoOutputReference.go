// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltscceaccessv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/ltscceaccessv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference interface {
	cdktf.ComplexObject
	Categories() *[]*string
	SetCategories(val *[]*string)
	CategoriesInput() *[]*string
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
	EventLevel() *[]*string
	SetEventLevel(val *[]*string)
	EventLevelInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *LtsCceAccessV3AccessConfigWindowsLogInfo
	SetInternalValue(val *LtsCceAccessV3AccessConfigWindowsLogInfo)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeOffset() *float64
	SetTimeOffset(val *float64)
	TimeOffsetInput() *float64
	TimeOffsetUnit() *string
	SetTimeOffsetUnit(val *string)
	TimeOffsetUnitInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference
type jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) Categories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"categories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) CategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"categoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) EventLevel() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) EventLevelInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) InternalValue() *LtsCceAccessV3AccessConfigWindowsLogInfo {
	var returns *LtsCceAccessV3AccessConfigWindowsLogInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) TimeOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) TimeOffsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) TimeOffsetUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOffsetUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) TimeOffsetUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOffsetUnitInput",
		&returns,
	)
	return returns
}


func NewLtsCceAccessV3AccessConfigWindowsLogInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference {
	_init_.Initialize()

	if err := validateNewLtsCceAccessV3AccessConfigWindowsLogInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsCceAccessV3.LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLtsCceAccessV3AccessConfigWindowsLogInfoOutputReference_Override(l LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsCceAccessV3.LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference)SetCategories(val *[]*string) {
	if err := j.validateSetCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"categories",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference)SetEventLevel(val *[]*string) {
	if err := j.validateSetEventLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventLevel",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference)SetInternalValue(val *LtsCceAccessV3AccessConfigWindowsLogInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference)SetTimeOffset(val *float64) {
	if err := j.validateSetTimeOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeOffset",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference)SetTimeOffsetUnit(val *string) {
	if err := j.validateSetTimeOffsetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeOffsetUnit",
		val,
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigWindowsLogInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

