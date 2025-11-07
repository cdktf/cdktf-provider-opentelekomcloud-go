// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudhssintrusioneventsv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudhssintrusioneventsv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference interface {
	cdktf.ComplexObject
	ChildProcessCmdline() *string
	ChildProcessEgid() *float64
	ChildProcessEuid() *float64
	ChildProcessFilename() *string
	ChildProcessGid() *float64
	ChildProcessName() *string
	ChildProcessPath() *string
	ChildProcessPid() *float64
	ChildProcessStartTime() *float64
	ChildProcessUid() *float64
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
	EscapeCmd() *string
	EscapeMode() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStruct
	SetInternalValue(val *DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStruct)
	ParentProcessCmdline() *string
	ParentProcessEgid() *float64
	ParentProcessEuid() *float64
	ParentProcessFilename() *string
	ParentProcessGid() *float64
	ParentProcessName() *string
	ParentProcessPath() *string
	ParentProcessPid() *float64
	ParentProcessStartTime() *float64
	ParentProcessUid() *float64
	ProcessCmdline() *string
	ProcessEgid() *float64
	ProcessEuid() *float64
	ProcessFilename() *string
	ProcessGid() *float64
	ProcessHash() *string
	ProcessName() *string
	ProcessPath() *string
	ProcessPid() *float64
	ProcessStartTime() *float64
	ProcessUid() *float64
	ProcessUsername() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtCmd() *string
	VirtProcessName() *string
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

// The jsii proxy struct for DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference
type jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ChildProcessCmdline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"childProcessCmdline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ChildProcessEgid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"childProcessEgid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ChildProcessEuid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"childProcessEuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ChildProcessFilename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"childProcessFilename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ChildProcessGid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"childProcessGid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ChildProcessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"childProcessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ChildProcessPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"childProcessPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ChildProcessPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"childProcessPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ChildProcessStartTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"childProcessStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ChildProcessUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"childProcessUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) EscapeCmd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeCmd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) EscapeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) InternalValue() *DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStruct {
	var returns *DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStruct
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ParentProcessCmdline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentProcessCmdline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ParentProcessEgid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parentProcessEgid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ParentProcessEuid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parentProcessEuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ParentProcessFilename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentProcessFilename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ParentProcessGid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parentProcessGid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ParentProcessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentProcessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ParentProcessPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentProcessPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ParentProcessPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parentProcessPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ParentProcessStartTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parentProcessStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ParentProcessUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parentProcessUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ProcessCmdline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processCmdline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ProcessEgid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"processEgid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ProcessEuid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"processEuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ProcessFilename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processFilename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ProcessGid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"processGid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ProcessHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ProcessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ProcessPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ProcessPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"processPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ProcessStartTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"processStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ProcessUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"processUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ProcessUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) VirtCmd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtCmd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) VirtProcessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtProcessName",
		&returns,
	)
	return returns
}


func NewDataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference_Override(d DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference)SetInternalValue(val *DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStruct) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

