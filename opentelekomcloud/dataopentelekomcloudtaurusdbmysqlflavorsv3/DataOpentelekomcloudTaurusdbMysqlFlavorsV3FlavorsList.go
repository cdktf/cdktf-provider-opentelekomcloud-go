// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudtaurusdbmysqlflavorsv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudtaurusdbmysqlflavorsv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList
type jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudTaurusdbMysqlFlavorsV3.DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList_Override(d DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudTaurusdbMysqlFlavorsV3.DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList) Get(index *float64) DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlFlavorsV3FlavorsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

