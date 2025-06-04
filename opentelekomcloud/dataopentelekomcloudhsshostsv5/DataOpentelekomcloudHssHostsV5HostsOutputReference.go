// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudhsshostsv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudhsshostsv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudHssHostsV5HostsOutputReference interface {
	cdktf.ComplexObject
	AgentId() *string
	AgentStatus() *string
	AssetRiskNum() *float64
	AssetValue() *string
	BaselineRiskNum() *float64
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
	DetectResult() *string
	// Experimental.
	Fqn() *string
	GroupId() *string
	Id() *string
	InternalValue() *DataOpentelekomcloudHssHostsV5Hosts
	SetInternalValue(val *DataOpentelekomcloudHssHostsV5Hosts)
	IntrusionRiskNum() *float64
	Name() *string
	OsType() *string
	PolicyGroupId() *string
	PrivateIp() *string
	ProtectChargingMode() *string
	ProtectStatus() *string
	ProtectVersion() *string
	PublicIp() *string
	ResourceId() *string
	Status() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VulnerabilityRiskNum() *float64
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

// The jsii proxy struct for DataOpentelekomcloudHssHostsV5HostsOutputReference
type jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) AgentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) AssetRiskNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"assetRiskNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) AssetValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) BaselineRiskNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baselineRiskNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) DetectResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) InternalValue() *DataOpentelekomcloudHssHostsV5Hosts {
	var returns *DataOpentelekomcloudHssHostsV5Hosts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) IntrusionRiskNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intrusionRiskNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) PolicyGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) ProtectChargingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectChargingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) ProtectStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) ProtectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) VulnerabilityRiskNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vulnerabilityRiskNum",
		&returns,
	)
	return returns
}


func NewDataOpentelekomcloudHssHostsV5HostsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataOpentelekomcloudHssHostsV5HostsOutputReference {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudHssHostsV5HostsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssHostsV5.DataOpentelekomcloudHssHostsV5HostsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataOpentelekomcloudHssHostsV5HostsOutputReference_Override(d DataOpentelekomcloudHssHostsV5HostsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssHostsV5.DataOpentelekomcloudHssHostsV5HostsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference)SetInternalValue(val *DataOpentelekomcloudHssHostsV5Hosts) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudHssHostsV5HostsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

