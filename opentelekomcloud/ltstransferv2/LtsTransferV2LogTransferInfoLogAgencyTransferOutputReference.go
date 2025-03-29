// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltstransferv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/ltstransferv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference interface {
	cdktf.ComplexObject
	AgencyDomainId() *string
	SetAgencyDomainId(val *string)
	AgencyDomainIdInput() *string
	AgencyDomainName() *string
	SetAgencyDomainName(val *string)
	AgencyDomainNameInput() *string
	AgencyName() *string
	SetAgencyName(val *string)
	AgencyNameInput() *string
	AgencyProjectId() *string
	SetAgencyProjectId(val *string)
	AgencyProjectIdInput() *string
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
	InternalValue() *LtsTransferV2LogTransferInfoLogAgencyTransfer
	SetInternalValue(val *LtsTransferV2LogTransferInfoLogAgencyTransfer)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference
type jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) AgencyDomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyDomainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) AgencyDomainIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyDomainIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) AgencyDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) AgencyDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) AgencyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) AgencyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) AgencyProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) AgencyProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) InternalValue() *LtsTransferV2LogTransferInfoLogAgencyTransfer {
	var returns *LtsTransferV2LogTransferInfoLogAgencyTransfer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLtsTransferV2LogTransferInfoLogAgencyTransferOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference {
	_init_.Initialize()

	if err := validateNewLtsTransferV2LogTransferInfoLogAgencyTransferOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLtsTransferV2LogTransferInfoLogAgencyTransferOutputReference_Override(l LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference)SetAgencyDomainId(val *string) {
	if err := j.validateSetAgencyDomainIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agencyDomainId",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference)SetAgencyDomainName(val *string) {
	if err := j.validateSetAgencyDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agencyDomainName",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference)SetAgencyName(val *string) {
	if err := j.validateSetAgencyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agencyName",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference)SetAgencyProjectId(val *string) {
	if err := j.validateSetAgencyProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agencyProjectId",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference)SetInternalValue(val *LtsTransferV2LogTransferInfoLogAgencyTransfer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

