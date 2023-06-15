package asconfigurationv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v7/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v7/asconfigurationv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference interface {
	cdktf.ComplexObject
	ChargingMode() *string
	SetChargingMode(val *string)
	ChargingModeInput() *string
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
	InternalValue() *AsConfigurationV1InstanceConfigPublicIpEipBandwidth
	SetInternalValue(val *AsConfigurationV1InstanceConfigPublicIpEipBandwidth)
	ShareType() *string
	SetShareType(val *string)
	ShareTypeInput() *string
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
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

// The jsii proxy struct for AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference
type jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) ChargingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) ChargingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) InternalValue() *AsConfigurationV1InstanceConfigPublicIpEipBandwidth {
	var returns *AsConfigurationV1InstanceConfigPublicIpEipBandwidth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) ShareType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) ShareTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference {
	_init_.Initialize()

	if err := validateNewAsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.asConfigurationV1.AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference_Override(a AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.asConfigurationV1.AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference)SetChargingMode(val *string) {
	if err := j.validateSetChargingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chargingMode",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference)SetInternalValue(val *AsConfigurationV1InstanceConfigPublicIpEipBandwidth) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference)SetShareType(val *string) {
	if err := j.validateSetShareTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareType",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AsConfigurationV1InstanceConfigPublicIpEipBandwidthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

