// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package drstaskv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v9/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v9/drstaskv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DrsTaskV3SourceDbOutputReference interface {
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
	EngineType() *string
	SetEngineType(val *string)
	EngineTypeInput() *string
	// Experimental.
	Fqn() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	InternalValue() *DrsTaskV3SourceDb
	SetInternalValue(val *DrsTaskV3SourceDb)
	Ip() *string
	SetIp(val *string)
	IpInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PrivateIp() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SslCertCheckSum() *string
	SetSslCertCheckSum(val *string)
	SslCertCheckSumInput() *string
	SslCertKey() *string
	SetSslCertKey(val *string)
	SslCertKeyInput() *string
	SslCertName() *string
	SetSslCertName(val *string)
	SslCertNameInput() *string
	SslCertPassword() *string
	SetSslCertPassword(val *string)
	SslCertPasswordInput() *string
	SslEnabled() interface{}
	SetSslEnabled(val interface{})
	SslEnabledInput() interface{}
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	User() *string
	SetUser(val *string)
	UserInput() *string
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
	ResetInstanceId()
	ResetName()
	ResetRegion()
	ResetSslCertCheckSum()
	ResetSslCertKey()
	ResetSslCertName()
	ResetSslCertPassword()
	ResetSslEnabled()
	ResetSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DrsTaskV3SourceDbOutputReference
type jsiiProxy_DrsTaskV3SourceDbOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) EngineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) InternalValue() *DrsTaskV3SourceDb {
	var returns *DrsTaskV3SourceDb
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) Ip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) IpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) SslCertCheckSum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertCheckSum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) SslCertCheckSumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertCheckSumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) SslCertKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) SslCertKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) SslCertName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) SslCertNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) SslCertPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) SslCertPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) SslEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) SslEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewDrsTaskV3SourceDbOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DrsTaskV3SourceDbOutputReference {
	_init_.Initialize()

	if err := validateNewDrsTaskV3SourceDbOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DrsTaskV3SourceDbOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.drsTaskV3.DrsTaskV3SourceDbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDrsTaskV3SourceDbOutputReference_Override(d DrsTaskV3SourceDbOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.drsTaskV3.DrsTaskV3SourceDbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetEngineType(val *string) {
	if err := j.validateSetEngineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineType",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetInternalValue(val *DrsTaskV3SourceDb) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetIp(val *string) {
	if err := j.validateSetIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ip",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetSslCertCheckSum(val *string) {
	if err := j.validateSetSslCertCheckSumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCertCheckSum",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetSslCertKey(val *string) {
	if err := j.validateSetSslCertKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCertKey",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetSslCertName(val *string) {
	if err := j.validateSetSslCertNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCertName",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetSslCertPassword(val *string) {
	if err := j.validateSetSslCertPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCertPassword",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetSslEnabled(val interface{}) {
	if err := j.validateSetSslEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslEnabled",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DrsTaskV3SourceDbOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) ResetInstanceId() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) ResetSslCertCheckSum() {
	_jsii_.InvokeVoid(
		d,
		"resetSslCertCheckSum",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) ResetSslCertKey() {
	_jsii_.InvokeVoid(
		d,
		"resetSslCertKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) ResetSslCertName() {
	_jsii_.InvokeVoid(
		d,
		"resetSslCertName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) ResetSslCertPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetSslCertPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) ResetSslEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetSslEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		d,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DrsTaskV3SourceDbOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

