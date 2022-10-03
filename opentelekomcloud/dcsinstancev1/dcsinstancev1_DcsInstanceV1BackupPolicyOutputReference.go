package dcsinstancev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v3/jsii"

	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v3/dcsinstancev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DcsInstanceV1BackupPolicyOutputReference interface {
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
	InternalValue() *DcsInstanceV1BackupPolicy
	SetInternalValue(val *DcsInstanceV1BackupPolicy)
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

// The jsii proxy struct for DcsInstanceV1BackupPolicyOutputReference
type jsiiProxy_DcsInstanceV1BackupPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) BackupAt() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"backupAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) BackupAtInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"backupAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) BackupType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) BackupTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) BeginAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) BeginAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) InternalValue() *DcsInstanceV1BackupPolicy {
	var returns *DcsInstanceV1BackupPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) PeriodType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) PeriodTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) SaveDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saveDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) SaveDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saveDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDcsInstanceV1BackupPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DcsInstanceV1BackupPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewDcsInstanceV1BackupPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DcsInstanceV1BackupPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV1.DcsInstanceV1BackupPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDcsInstanceV1BackupPolicyOutputReference_Override(d DcsInstanceV1BackupPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcsInstanceV1.DcsInstanceV1BackupPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference)SetBackupAt(val *[]*float64) {
	if err := j.validateSetBackupAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupAt",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference)SetBackupType(val *string) {
	if err := j.validateSetBackupTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupType",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference)SetBeginAt(val *string) {
	if err := j.validateSetBeginAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"beginAt",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference)SetInternalValue(val *DcsInstanceV1BackupPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference)SetPeriodType(val *string) {
	if err := j.validateSetPeriodTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodType",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference)SetSaveDays(val *float64) {
	if err := j.validateSetSaveDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saveDays",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) ResetBackupType() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) ResetSaveDays() {
	_jsii_.InvokeVoid(
		d,
		"resetSaveDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DcsInstanceV1BackupPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

