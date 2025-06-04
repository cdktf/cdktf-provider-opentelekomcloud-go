// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcsinstancev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dcsinstancev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DcsInstanceV2BackupPolicyOutputReference interface {
	cdktf.ComplexObject
	BackupAt() *[]*float64
	SetBackupAt(val *[]*float64)
	BackupAtInput() *[]*float64
	BackupType() *string
	SetBackupType(val *string)
	BackupTypeInput() *string
	BeginAt() *string
	SetBeginAt(val *string)
	BeginAtInput() *string
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
	InternalValue() *DcsInstanceV2BackupPolicy
	SetInternalValue(val *DcsInstanceV2BackupPolicy)
	PeriodType() *string
	SetPeriodType(val *string)
	PeriodTypeInput() *string
	SaveDays() *float64
	SetSaveDays(val *float64)
	SaveDaysInput() *float64
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
	ResetBackupType()
	ResetPeriodType()
	ResetSaveDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DcsInstanceV2BackupPolicyOutputReference
type jsiiProxy_DcsInstanceV2BackupPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) BackupAt() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"backupAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) BackupAtInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"backupAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) BackupType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) BackupTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) BeginAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) BeginAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) InternalValue() *DcsInstanceV2BackupPolicy {
	var returns *DcsInstanceV2BackupPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) PeriodType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) PeriodTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) SaveDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saveDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) SaveDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saveDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDcsInstanceV2BackupPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DcsInstanceV2BackupPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewDcsInstanceV2BackupPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DcsInstanceV2BackupPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV2.DcsInstanceV2BackupPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDcsInstanceV2BackupPolicyOutputReference_Override(d DcsInstanceV2BackupPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV2.DcsInstanceV2BackupPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference)SetBackupAt(val *[]*float64) {
	if err := j.validateSetBackupAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupAt",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference)SetBackupType(val *string) {
	if err := j.validateSetBackupTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupType",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference)SetBeginAt(val *string) {
	if err := j.validateSetBeginAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"beginAt",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference)SetInternalValue(val *DcsInstanceV2BackupPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference)SetPeriodType(val *string) {
	if err := j.validateSetPeriodTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodType",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference)SetSaveDays(val *float64) {
	if err := j.validateSetSaveDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saveDays",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) ResetBackupType() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) ResetPeriodType() {
	_jsii_.InvokeVoid(
		d,
		"resetPeriodType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) ResetSaveDays() {
	_jsii_.InvokeVoid(
		d,
		"resetSaveDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV2BackupPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

