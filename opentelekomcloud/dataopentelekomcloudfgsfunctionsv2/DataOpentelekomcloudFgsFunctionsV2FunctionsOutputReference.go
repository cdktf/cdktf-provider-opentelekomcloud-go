// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudfgsfunctionsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudfgsfunctionsv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference interface {
	cdktf.ComplexObject
	AppAgency() *string
	CodeFilename() *string
	CodeType() *string
	CodeUrl() *string
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
	Description() *string
	EncryptedUserData() *string
	EnterpriseProjectId() *string
	// Experimental.
	Fqn() *string
	FunctiongraphVersion() *string
	Handler() *string
	InitializerHandler() *string
	InitializerTimeout() *float64
	InternalValue() *DataOpentelekomcloudFgsFunctionsV2Functions
	SetInternalValue(val *DataOpentelekomcloudFgsFunctionsV2Functions)
	LogGroupId() *string
	LogStreamId() *string
	MaxInstanceNum() *float64
	MemorySize() *float64
	Name() *string
	NetworkId() *string
	Package() *string
	Runtime() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *float64
	Urn() *string
	UserData() *string
	Version() *string
	VpcId() *string
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

// The jsii proxy struct for DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference
type jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) AppAgency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appAgency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) CodeFilename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeFilename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) CodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) CodeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) EncryptedUserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptedUserData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) EnterpriseProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) FunctiongraphVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functiongraphVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) InitializerHandler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initializerHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) InitializerTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initializerTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) InternalValue() *DataOpentelekomcloudFgsFunctionsV2Functions {
	var returns *DataOpentelekomcloudFgsFunctionsV2Functions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) LogGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) LogStreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) MaxInstanceNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) Package() *string {
	var returns *string
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) Urn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


func NewDataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudFgsFunctionsV2FunctionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudFgsFunctionsV2.DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference_Override(d DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudFgsFunctionsV2.DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference)SetInternalValue(val *DataOpentelekomcloudFgsFunctionsV2Functions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudFgsFunctionsV2FunctionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

