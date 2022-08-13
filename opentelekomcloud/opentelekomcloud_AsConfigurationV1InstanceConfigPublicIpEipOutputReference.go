// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AsConfigurationV1InstanceConfigPublicIpEipOutputReference interface {
	cdktf.ComplexObject
	Bandwidth() AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference
	BandwidthInput() *AsConfigurationV1InstanceConfigPublicIpEipBandwidth
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
	InternalValue() *AsConfigurationV1InstanceConfigPublicIpEip
	SetInternalValue(val *AsConfigurationV1InstanceConfigPublicIpEip)
	IpType() *string
	SetIpType(val *string)
	IpTypeInput() *string
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
	PutBandwidth(value *AsConfigurationV1InstanceConfigPublicIpEipBandwidth)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AsConfigurationV1InstanceConfigPublicIpEipOutputReference
type jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) Bandwidth() AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference {
	var returns AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference
	_jsii_.Get(
		j,
		"bandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) BandwidthInput() *AsConfigurationV1InstanceConfigPublicIpEipBandwidth {
	var returns *AsConfigurationV1InstanceConfigPublicIpEipBandwidth
	_jsii_.Get(
		j,
		"bandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) InternalValue() *AsConfigurationV1InstanceConfigPublicIpEip {
	var returns *AsConfigurationV1InstanceConfigPublicIpEip
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) IpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) IpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAsConfigurationV1InstanceConfigPublicIpEipOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AsConfigurationV1InstanceConfigPublicIpEipOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.AsConfigurationV1InstanceConfigPublicIpEipOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAsConfigurationV1InstanceConfigPublicIpEipOutputReference_Override(a AsConfigurationV1InstanceConfigPublicIpEipOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.AsConfigurationV1InstanceConfigPublicIpEipOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) SetInternalValue(val *AsConfigurationV1InstanceConfigPublicIpEip) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) SetIpType(val *string) {
	_jsii_.Set(
		j,
		"ipType",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) PutBandwidth(value *AsConfigurationV1InstanceConfigPublicIpEipBandwidth) {
	_jsii_.InvokeVoid(
		a,
		"putBandwidth",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

