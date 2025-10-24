// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudtaurusdbmysqlproxiesv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudtaurusdbmysqlproxiesv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference interface {
	cdktf.ComplexObject
	Address() *string
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
	DelayThresholdInSeconds() *float64
	ElbVip() *string
	Flavor() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStruct
	SetInternalValue(val *DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStruct)
	MasterNodeWeight() DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListMasterNodeWeightList
	Mode() *string
	Name() *string
	NodeNum() *float64
	Nodes() DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListNodesList
	Port() *float64
	Ram() *string
	ReadonlyNodesWeight() DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListReadonlyNodesWeightList
	Status() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransactionSplit() *string
	Vcpus() *string
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

// The jsii proxy struct for DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference
type jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) DelayThresholdInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delayThresholdInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) ElbVip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elbVip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) Flavor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) InternalValue() *DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStruct {
	var returns *DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStruct
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) MasterNodeWeight() DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListMasterNodeWeightList {
	var returns DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListMasterNodeWeightList
	_jsii_.Get(
		j,
		"masterNodeWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) NodeNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) Nodes() DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListNodesList {
	var returns DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListNodesList
	_jsii_.Get(
		j,
		"nodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) Ram() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) ReadonlyNodesWeight() DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListReadonlyNodesWeightList {
	var returns DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListReadonlyNodesWeightList
	_jsii_.Get(
		j,
		"readonlyNodesWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) TransactionSplit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionSplit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) Vcpus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcpus",
		&returns,
	)
	return returns
}


func NewDataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudTaurusdbMysqlProxiesV3.DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference_Override(d DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudTaurusdbMysqlProxiesV3.DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference)SetInternalValue(val *DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStruct) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudTaurusdbMysqlProxiesV3ProxyListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

