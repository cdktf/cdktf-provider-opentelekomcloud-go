// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package disdumptaskv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v9/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v9/disdumptaskv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DisDumpTaskV2ObsDestinationDescriptorList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DisDumpTaskV2ObsDestinationDescriptorOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DisDumpTaskV2ObsDestinationDescriptorList
type jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDisDumpTaskV2ObsDestinationDescriptorList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DisDumpTaskV2ObsDestinationDescriptorList {
	_init_.Initialize()

	if err := validateNewDisDumpTaskV2ObsDestinationDescriptorListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.disDumpTaskV2.DisDumpTaskV2ObsDestinationDescriptorList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDisDumpTaskV2ObsDestinationDescriptorList_Override(d DisDumpTaskV2ObsDestinationDescriptorList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.disDumpTaskV2.DisDumpTaskV2ObsDestinationDescriptorList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList) Get(index *float64) DisDumpTaskV2ObsDestinationDescriptorOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DisDumpTaskV2ObsDestinationDescriptorOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DisDumpTaskV2ObsDestinationDescriptorList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

