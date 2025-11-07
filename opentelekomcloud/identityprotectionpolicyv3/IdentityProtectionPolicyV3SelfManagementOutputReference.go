// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityprotectionpolicyv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/identityprotectionpolicyv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityProtectionPolicyV3SelfManagementOutputReference interface {
	cdktf.ComplexObject
	AccessKey() interface{}
	SetAccessKey(val interface{})
	AccessKeyInput() interface{}
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
	Email() interface{}
	SetEmail(val interface{})
	EmailInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *IdentityProtectionPolicyV3SelfManagement
	SetInternalValue(val *IdentityProtectionPolicyV3SelfManagement)
	Mobile() interface{}
	SetMobile(val interface{})
	MobileInput() interface{}
	Password() interface{}
	SetPassword(val interface{})
	PasswordInput() interface{}
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
	ResetAccessKey()
	ResetEmail()
	ResetMobile()
	ResetPassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IdentityProtectionPolicyV3SelfManagementOutputReference
type jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) AccessKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) AccessKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) Email() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) EmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) InternalValue() *IdentityProtectionPolicyV3SelfManagement {
	var returns *IdentityProtectionPolicyV3SelfManagement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) Mobile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mobile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) MobileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mobileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) Password() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) PasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIdentityProtectionPolicyV3SelfManagementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IdentityProtectionPolicyV3SelfManagementOutputReference {
	_init_.Initialize()

	if err := validateNewIdentityProtectionPolicyV3SelfManagementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.identityProtectionPolicyV3.IdentityProtectionPolicyV3SelfManagementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIdentityProtectionPolicyV3SelfManagementOutputReference_Override(i IdentityProtectionPolicyV3SelfManagementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.identityProtectionPolicyV3.IdentityProtectionPolicyV3SelfManagementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference)SetAccessKey(val interface{}) {
	if err := j.validateSetAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference)SetEmail(val interface{}) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference)SetInternalValue(val *IdentityProtectionPolicyV3SelfManagement) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference)SetMobile(val interface{}) {
	if err := j.validateSetMobileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobile",
		val,
	)
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference)SetPassword(val interface{}) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) ResetAccessKey() {
	_jsii_.InvokeVoid(
		i,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		i,
		"resetEmail",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) ResetMobile() {
	_jsii_.InvokeVoid(
		i,
		"resetMobile",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		i,
		"resetPassword",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityProtectionPolicyV3SelfManagementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

