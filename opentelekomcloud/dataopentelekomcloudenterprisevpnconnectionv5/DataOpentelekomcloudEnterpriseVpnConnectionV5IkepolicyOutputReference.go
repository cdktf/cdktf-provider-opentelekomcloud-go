// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudenterprisevpnconnectionv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudenterprisevpnconnectionv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference interface {
	cdktf.ComplexObject
	AuthenticationAlgorithm() *string
	AuthenticationMethod() *string
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
	DhGroup() *string
	Dpd() DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyDpdList
	EncryptionAlgorithm() *string
	// Experimental.
	Fqn() *string
	IkeVersion() *string
	InternalValue() *DataOpentelekomcloudEnterpriseVpnConnectionV5Ikepolicy
	SetInternalValue(val *DataOpentelekomcloudEnterpriseVpnConnectionV5Ikepolicy)
	LifetimeSeconds() *float64
	LocalId() *string
	LocalIdType() *string
	PeerId() *string
	PeerIdType() *string
	PhaseOneNegotiationMode() *string
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

// The jsii proxy struct for DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference
type jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) AuthenticationAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) DhGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) Dpd() DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyDpdList {
	var returns DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyDpdList
	_jsii_.Get(
		j,
		"dpd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) EncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) IkeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) InternalValue() *DataOpentelekomcloudEnterpriseVpnConnectionV5Ikepolicy {
	var returns *DataOpentelekomcloudEnterpriseVpnConnectionV5Ikepolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) LocalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) LocalIdType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIdType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) PeerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) PeerIdType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIdType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) PhaseOneNegotiationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phaseOneNegotiationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudEnterpriseVpnConnectionV5.DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference_Override(d DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudEnterpriseVpnConnectionV5.DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference)SetInternalValue(val *DataOpentelekomcloudEnterpriseVpnConnectionV5Ikepolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudEnterpriseVpnConnectionV5IkepolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

