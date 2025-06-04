// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltstransferv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/ltstransferv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsTransferV2LogTransferInfoLogTransferDetailOutputReference interface {
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
	InternalValue() *LtsTransferV2LogTransferInfoLogTransferDetail
	SetInternalValue(val *LtsTransferV2LogTransferInfoLogTransferDetail)
	ObsBucketName() *string
	SetObsBucketName(val *string)
	ObsBucketNameInput() *string
	ObsDirPrefixName() *string
	SetObsDirPrefixName(val *string)
	ObsDirPrefixNameInput() *string
	ObsEncryptedEnable() interface{}
	SetObsEncryptedEnable(val interface{})
	ObsEncryptedEnableInput() interface{}
	ObsEncryptedId() *string
	SetObsEncryptedId(val *string)
	ObsEncryptedIdInput() *string
	ObsEpsId() *string
	SetObsEpsId(val *string)
	ObsEpsIdInput() *string
	ObsPeriod() *float64
	SetObsPeriod(val *float64)
	ObsPeriodInput() *float64
	ObsPeriodUnit() *string
	SetObsPeriodUnit(val *string)
	ObsPeriodUnitInput() *string
	ObsPrefixName() *string
	SetObsPrefixName(val *string)
	ObsPrefixNameInput() *string
	ObsTimeZone() *string
	SetObsTimeZone(val *string)
	ObsTimeZoneId() *string
	SetObsTimeZoneId(val *string)
	ObsTimeZoneIdInput() *string
	ObsTimeZoneInput() *string
	ObsTransferPath() *string
	SetObsTransferPath(val *string)
	ObsTransferPathInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
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
	ResetObsBucketName()
	ResetObsDirPrefixName()
	ResetObsEncryptedEnable()
	ResetObsEncryptedId()
	ResetObsEpsId()
	ResetObsPeriod()
	ResetObsPeriodUnit()
	ResetObsPrefixName()
	ResetObsTimeZone()
	ResetObsTimeZoneId()
	ResetObsTransferPath()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LtsTransferV2LogTransferInfoLogTransferDetailOutputReference
type jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) InternalValue() *LtsTransferV2LogTransferInfoLogTransferDetail {
	var returns *LtsTransferV2LogTransferInfoLogTransferDetail
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsBucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsBucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsBucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsBucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsDirPrefixName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsDirPrefixName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsDirPrefixNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsDirPrefixNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsEncryptedEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"obsEncryptedEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsEncryptedEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"obsEncryptedEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsEncryptedId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsEncryptedId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsEncryptedIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsEncryptedIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsEpsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsEpsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsEpsIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsEpsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"obsPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"obsPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsPeriodUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsPeriodUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsPeriodUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsPeriodUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsPrefixName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsPrefixName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsPrefixNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsPrefixNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsTimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsTimeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsTimeZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsTimeZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsTimeZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsTimeZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsTimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsTimeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsTransferPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsTransferPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ObsTransferPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"obsTransferPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLtsTransferV2LogTransferInfoLogTransferDetailOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LtsTransferV2LogTransferInfoLogTransferDetailOutputReference {
	_init_.Initialize()

	if err := validateNewLtsTransferV2LogTransferInfoLogTransferDetailOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogTransferInfoLogTransferDetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLtsTransferV2LogTransferInfoLogTransferDetailOutputReference_Override(l LtsTransferV2LogTransferInfoLogTransferDetailOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogTransferInfoLogTransferDetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetInternalValue(val *LtsTransferV2LogTransferInfoLogTransferDetail) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetObsBucketName(val *string) {
	if err := j.validateSetObsBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsBucketName",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetObsDirPrefixName(val *string) {
	if err := j.validateSetObsDirPrefixNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsDirPrefixName",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetObsEncryptedEnable(val interface{}) {
	if err := j.validateSetObsEncryptedEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsEncryptedEnable",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetObsEncryptedId(val *string) {
	if err := j.validateSetObsEncryptedIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsEncryptedId",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetObsEpsId(val *string) {
	if err := j.validateSetObsEpsIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsEpsId",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetObsPeriod(val *float64) {
	if err := j.validateSetObsPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsPeriod",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetObsPeriodUnit(val *string) {
	if err := j.validateSetObsPeriodUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsPeriodUnit",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetObsPrefixName(val *string) {
	if err := j.validateSetObsPrefixNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsPrefixName",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetObsTimeZone(val *string) {
	if err := j.validateSetObsTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsTimeZone",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetObsTimeZoneId(val *string) {
	if err := j.validateSetObsTimeZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsTimeZoneId",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetObsTransferPath(val *string) {
	if err := j.validateSetObsTransferPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"obsTransferPath",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ResetObsBucketName() {
	_jsii_.InvokeVoid(
		l,
		"resetObsBucketName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ResetObsDirPrefixName() {
	_jsii_.InvokeVoid(
		l,
		"resetObsDirPrefixName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ResetObsEncryptedEnable() {
	_jsii_.InvokeVoid(
		l,
		"resetObsEncryptedEnable",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ResetObsEncryptedId() {
	_jsii_.InvokeVoid(
		l,
		"resetObsEncryptedId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ResetObsEpsId() {
	_jsii_.InvokeVoid(
		l,
		"resetObsEpsId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ResetObsPeriod() {
	_jsii_.InvokeVoid(
		l,
		"resetObsPeriod",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ResetObsPeriodUnit() {
	_jsii_.InvokeVoid(
		l,
		"resetObsPeriodUnit",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ResetObsPrefixName() {
	_jsii_.InvokeVoid(
		l,
		"resetObsPrefixName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ResetObsTimeZone() {
	_jsii_.InvokeVoid(
		l,
		"resetObsTimeZone",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ResetObsTimeZoneId() {
	_jsii_.InvokeVoid(
		l,
		"resetObsTimeZoneId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ResetObsTransferPath() {
	_jsii_.InvokeVoid(
		l,
		"resetObsTransferPath",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

