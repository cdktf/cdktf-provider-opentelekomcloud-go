// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbListenerV3InsertHeadersOutputReference interface {
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
	ForwardedForPort() interface{}
	SetForwardedForPort(val interface{})
	ForwardedForPortInput() interface{}
	ForwardedHost() interface{}
	SetForwardedHost(val interface{})
	ForwardedHostInput() interface{}
	ForwardedPort() interface{}
	SetForwardedPort(val interface{})
	ForwardedPortInput() interface{}
	ForwardElbIp() interface{}
	SetForwardElbIp(val interface{})
	ForwardElbIpInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *LbListenerV3InsertHeaders
	SetInternalValue(val *LbListenerV3InsertHeaders)
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
	ResetForwardedForPort()
	ResetForwardedHost()
	ResetForwardedPort()
	ResetForwardElbIp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbListenerV3InsertHeadersOutputReference
type jsiiProxy_LbListenerV3InsertHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ForwardedForPort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardedForPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ForwardedForPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardedForPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ForwardedHost() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardedHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ForwardedHostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardedHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ForwardedPort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardedPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ForwardedPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardedPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ForwardElbIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardElbIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ForwardElbIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardElbIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) InternalValue() *LbListenerV3InsertHeaders {
	var returns *LbListenerV3InsertHeaders
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLbListenerV3InsertHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LbListenerV3InsertHeadersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LbListenerV3InsertHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.LbListenerV3InsertHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLbListenerV3InsertHeadersOutputReference_Override(l LbListenerV3InsertHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.LbListenerV3InsertHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) SetForwardedForPort(val interface{}) {
	_jsii_.Set(
		j,
		"forwardedForPort",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) SetForwardedHost(val interface{}) {
	_jsii_.Set(
		j,
		"forwardedHost",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) SetForwardedPort(val interface{}) {
	_jsii_.Set(
		j,
		"forwardedPort",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) SetForwardElbIp(val interface{}) {
	_jsii_.Set(
		j,
		"forwardElbIp",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) SetInternalValue(val *LbListenerV3InsertHeaders) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3InsertHeadersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ResetForwardedForPort() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardedForPort",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ResetForwardedHost() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardedHost",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ResetForwardedPort() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardedPort",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ResetForwardElbIp() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardElbIp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3InsertHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
