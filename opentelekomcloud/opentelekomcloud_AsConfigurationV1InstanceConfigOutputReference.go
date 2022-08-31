// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AsConfigurationV1InstanceConfigOutputReference interface {
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
	Disk() AsConfigurationV1InstanceConfigDiskList
	DiskInput() interface{}
	Flavor() *string
	SetFlavor(val *string)
	FlavorInput() *string
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	InternalValue() *AsConfigurationV1InstanceConfig
	SetInternalValue(val *AsConfigurationV1InstanceConfig)
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Personality() AsConfigurationV1InstanceConfigPersonalityList
	PersonalityInput() interface{}
	PublicIp() AsConfigurationV1InstanceConfigPublicIpOutputReference
	PublicIpInput() *AsConfigurationV1InstanceConfigPublicIp
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
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
	PutDisk(value interface{})
	PutPersonality(value interface{})
	PutPublicIp(value *AsConfigurationV1InstanceConfigPublicIp)
	ResetDisk()
	ResetFlavor()
	ResetImage()
	ResetInstanceId()
	ResetMetadata()
	ResetPersonality()
	ResetPublicIp()
	ResetSecurityGroups()
	ResetUserData()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AsConfigurationV1InstanceConfigOutputReference
type jsiiProxy_AsConfigurationV1InstanceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) Disk() AsConfigurationV1InstanceConfigDiskList {
	var returns AsConfigurationV1InstanceConfigDiskList
	_jsii_.Get(
		j,
		"disk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) DiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) Flavor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) FlavorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) InternalValue() *AsConfigurationV1InstanceConfig {
	var returns *AsConfigurationV1InstanceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) Personality() AsConfigurationV1InstanceConfigPersonalityList {
	var returns AsConfigurationV1InstanceConfigPersonalityList
	_jsii_.Get(
		j,
		"personality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) PersonalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"personalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) PublicIp() AsConfigurationV1InstanceConfigPublicIpOutputReference {
	var returns AsConfigurationV1InstanceConfigPublicIpOutputReference
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) PublicIpInput() *AsConfigurationV1InstanceConfigPublicIp {
	var returns *AsConfigurationV1InstanceConfigPublicIp
	_jsii_.Get(
		j,
		"publicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}


func NewAsConfigurationV1InstanceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AsConfigurationV1InstanceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAsConfigurationV1InstanceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AsConfigurationV1InstanceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.AsConfigurationV1InstanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAsConfigurationV1InstanceConfigOutputReference_Override(a AsConfigurationV1InstanceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.AsConfigurationV1InstanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference)SetFlavor(val *string) {
	if err := j.validateSetFlavorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flavor",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference)SetInternalValue(val *AsConfigurationV1InstanceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) PutDisk(value interface{}) {
	if err := a.validatePutDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDisk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) PutPersonality(value interface{}) {
	if err := a.validatePutPersonalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPersonality",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) PutPublicIp(value *AsConfigurationV1InstanceConfigPublicIp) {
	if err := a.validatePutPublicIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPublicIp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ResetDisk() {
	_jsii_.InvokeVoid(
		a,
		"resetDisk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ResetFlavor() {
	_jsii_.InvokeVoid(
		a,
		"resetFlavor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		a,
		"resetImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ResetInstanceId() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ResetPersonality() {
	_jsii_.InvokeVoid(
		a,
		"resetPersonality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ResetPublicIp() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ResetUserData() {
	_jsii_.InvokeVoid(
		a,
		"resetUserData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

