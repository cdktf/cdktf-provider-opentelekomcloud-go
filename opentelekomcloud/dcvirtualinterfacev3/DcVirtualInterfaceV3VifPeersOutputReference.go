// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcvirtualinterfacev3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/dcvirtualinterfacev3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DcVirtualInterfaceV3VifPeersOutputReference interface {
	cdktf.ComplexObject
	AddressFamily() *string
	BgpAsn() *float64
	BgpMd5() *string
	BgpRouteLimit() *float64
	BgpStatus() *string
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
	DeviceId() *string
	EnableBfd() cdktf.IResolvable
	EnableNqa() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DcVirtualInterfaceV3VifPeers
	SetInternalValue(val *DcVirtualInterfaceV3VifPeers)
	LocalGatewayIp() *string
	Name() *string
	ReceiveRouteNum() *float64
	RemoteEpGroup() *[]*string
	RemoteGatewayIp() *string
	RouteMode() *string
	ServiceEpGroup() *[]*string
	Status() *string
	TenantId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VifId() *string
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

// The jsii proxy struct for DcVirtualInterfaceV3VifPeersOutputReference
type jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) AddressFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) BgpAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgpAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) BgpMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) BgpRouteLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgpRouteLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) BgpStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) DeviceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) EnableBfd() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableBfd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) EnableNqa() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableNqa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) InternalValue() *DcVirtualInterfaceV3VifPeers {
	var returns *DcVirtualInterfaceV3VifPeers
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) LocalGatewayIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) ReceiveRouteNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"receiveRouteNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) RemoteEpGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteEpGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) RemoteGatewayIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteGatewayIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) RouteMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) ServiceEpGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceEpGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) VifId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vifId",
		&returns,
	)
	return returns
}


func NewDcVirtualInterfaceV3VifPeersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DcVirtualInterfaceV3VifPeersOutputReference {
	_init_.Initialize()

	if err := validateNewDcVirtualInterfaceV3VifPeersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfaceV3.DcVirtualInterfaceV3VifPeersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDcVirtualInterfaceV3VifPeersOutputReference_Override(d DcVirtualInterfaceV3VifPeersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfaceV3.DcVirtualInterfaceV3VifPeersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference)SetInternalValue(val *DcVirtualInterfaceV3VifPeers) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DcVirtualInterfaceV3VifPeersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

