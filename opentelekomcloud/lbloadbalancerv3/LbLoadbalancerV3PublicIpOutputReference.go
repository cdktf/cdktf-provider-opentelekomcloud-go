package lbloadbalancerv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v5/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v5/lbloadbalancerv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbLoadbalancerV3PublicIpOutputReference interface {
	cdktf.ComplexObject
	Address() *string
	BandwidthChargeMode() *string
	SetBandwidthChargeMode(val *string)
	BandwidthChargeModeInput() *string
	BandwidthName() *string
	SetBandwidthName(val *string)
	BandwidthNameInput() *string
	BandwidthShareType() *string
	SetBandwidthShareType(val *string)
	BandwidthShareTypeInput() *string
	BandwidthSize() *float64
	SetBandwidthSize(val *float64)
	BandwidthSizeInput() *float64
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() *LbLoadbalancerV3PublicIp
	SetInternalValue(val *LbLoadbalancerV3PublicIp)
	IpType() *string
	SetIpType(val *string)
	IpTypeInput() *string
	Managed() cdktf.IResolvable
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
	ResetBandwidthChargeMode()
	ResetBandwidthName()
	ResetBandwidthShareType()
	ResetBandwidthSize()
	ResetId()
	ResetIpType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbLoadbalancerV3PublicIpOutputReference
type jsiiProxy_LbLoadbalancerV3PublicIpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) BandwidthChargeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthChargeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) BandwidthChargeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthChargeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) BandwidthName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) BandwidthNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) BandwidthShareType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthShareType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) BandwidthShareTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthShareTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) BandwidthSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) BandwidthSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) InternalValue() *LbLoadbalancerV3PublicIp {
	var returns *LbLoadbalancerV3PublicIp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) IpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) IpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) Managed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"managed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLbLoadbalancerV3PublicIpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LbLoadbalancerV3PublicIpOutputReference {
	_init_.Initialize()

	if err := validateNewLbLoadbalancerV3PublicIpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbLoadbalancerV3PublicIpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.lbLoadbalancerV3.LbLoadbalancerV3PublicIpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLbLoadbalancerV3PublicIpOutputReference_Override(l LbLoadbalancerV3PublicIpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.lbLoadbalancerV3.LbLoadbalancerV3PublicIpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference)SetBandwidthChargeMode(val *string) {
	if err := j.validateSetBandwidthChargeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthChargeMode",
		val,
	)
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference)SetBandwidthName(val *string) {
	if err := j.validateSetBandwidthNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthName",
		val,
	)
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference)SetBandwidthShareType(val *string) {
	if err := j.validateSetBandwidthShareTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthShareType",
		val,
	)
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference)SetBandwidthSize(val *float64) {
	if err := j.validateSetBandwidthSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthSize",
		val,
	)
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference)SetInternalValue(val *LbLoadbalancerV3PublicIp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference)SetIpType(val *string) {
	if err := j.validateSetIpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipType",
		val,
	)
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) ResetBandwidthChargeMode() {
	_jsii_.InvokeVoid(
		l,
		"resetBandwidthChargeMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) ResetBandwidthName() {
	_jsii_.InvokeVoid(
		l,
		"resetBandwidthName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) ResetBandwidthShareType() {
	_jsii_.InvokeVoid(
		l,
		"resetBandwidthShareType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) ResetBandwidthSize() {
	_jsii_.InvokeVoid(
		l,
		"resetBandwidthSize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) ResetIpType() {
	_jsii_.InvokeVoid(
		l,
		"resetIpType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LbLoadbalancerV3PublicIpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

