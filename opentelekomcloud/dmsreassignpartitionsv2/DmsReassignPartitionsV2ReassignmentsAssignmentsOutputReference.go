// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsreassignpartitionsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/dmsreassignpartitionsv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Partition() *float64
	SetPartition(val *float64)
	PartitionBrokers() *[]*float64
	SetPartitionBrokers(val *[]*float64)
	PartitionBrokersInput() *[]*float64
	PartitionInput() *float64
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
	ResetPartition()
	ResetPartitionBrokers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference
type jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) Partition() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) PartitionBrokers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"partitionBrokers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) PartitionBrokersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"partitionBrokersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) PartitionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsReassignPartitionsV2ReassignmentsAssignmentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dmsReassignPartitionsV2.DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference_Override(d DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dmsReassignPartitionsV2.DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference)SetPartition(val *float64) {
	if err := j.validateSetPartitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partition",
		val,
	)
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference)SetPartitionBrokers(val *[]*float64) {
	if err := j.validateSetPartitionBrokersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionBrokers",
		val,
	)
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) ResetPartition() {
	_jsii_.InvokeVoid(
		d,
		"resetPartition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) ResetPartitionBrokers() {
	_jsii_.InvokeVoid(
		d,
		"resetPartitionBrokers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsReassignPartitionsV2ReassignmentsAssignmentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

