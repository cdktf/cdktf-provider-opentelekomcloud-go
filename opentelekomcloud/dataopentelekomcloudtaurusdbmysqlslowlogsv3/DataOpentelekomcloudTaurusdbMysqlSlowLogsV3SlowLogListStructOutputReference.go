// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudtaurusdbmysqlslowlogsv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudtaurusdbmysqlslowlogsv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference interface {
	cdktf.ComplexObject
	ClientIp() *string
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
	Count() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Database() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStruct
	SetInternalValue(val *DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStruct)
	LockTime() *string
	NodeId() *string
	QuerySample() *string
	RowsExamined() *string
	RowsSent() *string
	StartTime() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Time() *string
	Type() *string
	Users() *string
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

// The jsii proxy struct for DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference
type jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) ClientIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) Count() *string {
	var returns *string
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) InternalValue() *DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStruct {
	var returns *DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStruct
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) LockTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) NodeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) QuerySample() *string {
	var returns *string
	_jsii_.Get(
		j,
		"querySample",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) RowsExamined() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowsExamined",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) RowsSent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowsSent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) Time() *string {
	var returns *string
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) Users() *string {
	var returns *string
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}


func NewDataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudTaurusdbMysqlSlowLogsV3.DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference_Override(d DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudTaurusdbMysqlSlowLogsV3.DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference)SetInternalValue(val *DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStruct) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlSlowLogsV3SlowLogListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

