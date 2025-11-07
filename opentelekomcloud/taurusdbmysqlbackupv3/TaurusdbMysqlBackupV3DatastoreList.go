// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package taurusdbmysqlbackupv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/taurusdbmysqlbackupv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TaurusdbMysqlBackupV3DatastoreList interface {
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
	Get(index *float64) TaurusdbMysqlBackupV3DatastoreOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TaurusdbMysqlBackupV3DatastoreList
type jsiiProxy_TaurusdbMysqlBackupV3DatastoreList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewTaurusdbMysqlBackupV3DatastoreList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) TaurusdbMysqlBackupV3DatastoreList {
	_init_.Initialize()

	if err := validateNewTaurusdbMysqlBackupV3DatastoreListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TaurusdbMysqlBackupV3DatastoreList{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlBackupV3.TaurusdbMysqlBackupV3DatastoreList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewTaurusdbMysqlBackupV3DatastoreList_Override(t TaurusdbMysqlBackupV3DatastoreList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.taurusdbMysqlBackupV3.TaurusdbMysqlBackupV3DatastoreList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		t,
	)
}

func (j *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (t *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := t.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		t,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList) Get(index *float64) TaurusdbMysqlBackupV3DatastoreOutputReference {
	if err := t.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns TaurusdbMysqlBackupV3DatastoreOutputReference

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList) Resolve(context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaurusdbMysqlBackupV3DatastoreList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

