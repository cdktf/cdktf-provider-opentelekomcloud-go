// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudhssintrusioneventsv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudhssintrusioneventsv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference interface {
	cdktf.ComplexObject
	AgentStatus() *string
	AssetValue() *string
	AttackPhase() *string
	AttackTag() *string
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
	ContainerName() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EventClassId() *string
	EventDetails() *string
	EventName() *string
	EventType() *float64
	FileInfoList() DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructList
	// Experimental.
	Fqn() *string
	HandleMethod() *string
	Handler() *string
	HandleStatus() *string
	HandleTime() *float64
	HostId() *string
	HostName() *string
	HostStatus() *string
	Id() *string
	ImageName() *string
	InternalValue() *DataOpentelekomcloudHssIntrusionEventsV5Events
	SetInternalValue(val *DataOpentelekomcloudHssIntrusionEventsV5Events)
	OccurTime() *float64
	OperateAcceptList() *[]*string
	OperateDetailList() DataOpentelekomcloudHssIntrusionEventsV5EventsOperateDetailListStructList
	OsType() *string
	PrivateIp() *string
	ProcessInfoList() DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructList
	ProtectStatus() *string
	PublicIp() *string
	Recommendation() *string
	ResourceInfo() DataOpentelekomcloudHssIntrusionEventsV5EventsResourceInfoList
	Severity() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserInfoList() DataOpentelekomcloudHssIntrusionEventsV5EventsUserInfoListStructList
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

// The jsii proxy struct for DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference
type jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) AgentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) AssetValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) AttackPhase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attackPhase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) AttackTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attackTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) EventClassId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventClassId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) EventDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) EventName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) EventType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) FileInfoList() DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructList {
	var returns DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructList
	_jsii_.Get(
		j,
		"fileInfoList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) HandleMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handleMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) HandleStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handleStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) HandleTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"handleTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) HostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) HostStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) InternalValue() *DataOpentelekomcloudHssIntrusionEventsV5Events {
	var returns *DataOpentelekomcloudHssIntrusionEventsV5Events
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) OccurTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"occurTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) OperateAcceptList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operateAcceptList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) OperateDetailList() DataOpentelekomcloudHssIntrusionEventsV5EventsOperateDetailListStructList {
	var returns DataOpentelekomcloudHssIntrusionEventsV5EventsOperateDetailListStructList
	_jsii_.Get(
		j,
		"operateDetailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) ProcessInfoList() DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructList {
	var returns DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructList
	_jsii_.Get(
		j,
		"processInfoList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) ProtectStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) Recommendation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) ResourceInfo() DataOpentelekomcloudHssIntrusionEventsV5EventsResourceInfoList {
	var returns DataOpentelekomcloudHssIntrusionEventsV5EventsResourceInfoList
	_jsii_.Get(
		j,
		"resourceInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) UserInfoList() DataOpentelekomcloudHssIntrusionEventsV5EventsUserInfoListStructList {
	var returns DataOpentelekomcloudHssIntrusionEventsV5EventsUserInfoListStructList
	_jsii_.Get(
		j,
		"userInfoList",
		&returns,
	)
	return returns
}


func NewDataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudHssIntrusionEventsV5EventsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference_Override(d DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference)SetInternalValue(val *DataOpentelekomcloudHssIntrusionEventsV5Events) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

