// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltstransferv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/ltstransferv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsTransferV2LogTransferInfoOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *LtsTransferV2LogTransferInfo
	SetInternalValue(val *LtsTransferV2LogTransferInfo)
	LogAgencyTransfer() LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference
	LogAgencyTransferInput() *LtsTransferV2LogTransferInfoLogAgencyTransfer
	LogCreatedAt() *string
	LogStorageFormat() *string
	SetLogStorageFormat(val *string)
	LogStorageFormatInput() *string
	LogTransferDetail() LtsTransferV2LogTransferInfoLogTransferDetailOutputReference
	LogTransferDetailInput() *LtsTransferV2LogTransferInfoLogTransferDetail
	LogTransferMode() *string
	SetLogTransferMode(val *string)
	LogTransferModeInput() *string
	LogTransferStatus() *string
	SetLogTransferStatus(val *string)
	LogTransferStatusInput() *string
	LogTransferType() *string
	SetLogTransferType(val *string)
	LogTransferTypeInput() *string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutLogAgencyTransfer(value *LtsTransferV2LogTransferInfoLogAgencyTransfer)
	PutLogTransferDetail(value *LtsTransferV2LogTransferInfoLogTransferDetail)
	ResetLogAgencyTransfer()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LtsTransferV2LogTransferInfoOutputReference
type jsiiProxy_LtsTransferV2LogTransferInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) InternalValue() *LtsTransferV2LogTransferInfo {
	var returns *LtsTransferV2LogTransferInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogAgencyTransfer() LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference {
	var returns LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference
	_jsii_.Get(
		j,
		"logAgencyTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogAgencyTransferInput() *LtsTransferV2LogTransferInfoLogAgencyTransfer {
	var returns *LtsTransferV2LogTransferInfoLogAgencyTransfer
	_jsii_.Get(
		j,
		"logAgencyTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogStorageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStorageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogStorageFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStorageFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogTransferDetail() LtsTransferV2LogTransferInfoLogTransferDetailOutputReference {
	var returns LtsTransferV2LogTransferInfoLogTransferDetailOutputReference
	_jsii_.Get(
		j,
		"logTransferDetail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogTransferDetailInput() *LtsTransferV2LogTransferInfoLogTransferDetail {
	var returns *LtsTransferV2LogTransferInfoLogTransferDetail
	_jsii_.Get(
		j,
		"logTransferDetailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogTransferMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTransferMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogTransferModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTransferModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogTransferStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTransferStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogTransferStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTransferStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogTransferType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTransferType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) LogTransferTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTransferTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLtsTransferV2LogTransferInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LtsTransferV2LogTransferInfoOutputReference {
	_init_.Initialize()

	if err := validateNewLtsTransferV2LogTransferInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LtsTransferV2LogTransferInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogTransferInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLtsTransferV2LogTransferInfoOutputReference_Override(l LtsTransferV2LogTransferInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogTransferInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference)SetInternalValue(val *LtsTransferV2LogTransferInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference)SetLogStorageFormat(val *string) {
	if err := j.validateSetLogStorageFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStorageFormat",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference)SetLogTransferMode(val *string) {
	if err := j.validateSetLogTransferModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logTransferMode",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference)SetLogTransferStatus(val *string) {
	if err := j.validateSetLogTransferStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logTransferStatus",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference)SetLogTransferType(val *string) {
	if err := j.validateSetLogTransferTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logTransferType",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) PutLogAgencyTransfer(value *LtsTransferV2LogTransferInfoLogAgencyTransfer) {
	if err := l.validatePutLogAgencyTransferParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLogAgencyTransfer",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) PutLogTransferDetail(value *LtsTransferV2LogTransferInfoLogTransferDetail) {
	if err := l.validatePutLogTransferDetailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLogTransferDetail",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) ResetLogAgencyTransfer() {
	_jsii_.InvokeVoid(
		l,
		"resetLogAgencyTransfer",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

