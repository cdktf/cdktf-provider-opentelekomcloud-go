package obsbucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v4/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v4/obsbucket/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObsBucketLifecycleRuleOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Expiration() ObsBucketLifecycleRuleExpirationList
	ExpirationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	NoncurrentVersionExpiration() ObsBucketLifecycleRuleNoncurrentVersionExpirationList
	NoncurrentVersionExpirationInput() interface{}
	NoncurrentVersionTransition() ObsBucketLifecycleRuleNoncurrentVersionTransitionList
	NoncurrentVersionTransitionInput() interface{}
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transition() ObsBucketLifecycleRuleTransitionList
	TransitionInput() interface{}
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
	PutExpiration(value interface{})
	PutNoncurrentVersionExpiration(value interface{})
	PutNoncurrentVersionTransition(value interface{})
	PutTransition(value interface{})
	ResetExpiration()
	ResetNoncurrentVersionExpiration()
	ResetNoncurrentVersionTransition()
	ResetPrefix()
	ResetTransition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObsBucketLifecycleRuleOutputReference
type jsiiProxy_ObsBucketLifecycleRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) Expiration() ObsBucketLifecycleRuleExpirationList {
	var returns ObsBucketLifecycleRuleExpirationList
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) ExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) NoncurrentVersionExpiration() ObsBucketLifecycleRuleNoncurrentVersionExpirationList {
	var returns ObsBucketLifecycleRuleNoncurrentVersionExpirationList
	_jsii_.Get(
		j,
		"noncurrentVersionExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) NoncurrentVersionExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) NoncurrentVersionTransition() ObsBucketLifecycleRuleNoncurrentVersionTransitionList {
	var returns ObsBucketLifecycleRuleNoncurrentVersionTransitionList
	_jsii_.Get(
		j,
		"noncurrentVersionTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) NoncurrentVersionTransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) Transition() ObsBucketLifecycleRuleTransitionList {
	var returns ObsBucketLifecycleRuleTransitionList
	_jsii_.Get(
		j,
		"transition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference) TransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitionInput",
		&returns,
	)
	return returns
}


func NewObsBucketLifecycleRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObsBucketLifecycleRuleOutputReference {
	_init_.Initialize()

	if err := validateNewObsBucketLifecycleRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObsBucketLifecycleRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.obsBucket.ObsBucketLifecycleRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObsBucketLifecycleRuleOutputReference_Override(o ObsBucketLifecycleRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.obsBucket.ObsBucketLifecycleRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObsBucketLifecycleRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) PutExpiration(value interface{}) {
	if err := o.validatePutExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putExpiration",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) PutNoncurrentVersionExpiration(value interface{}) {
	if err := o.validatePutNoncurrentVersionExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNoncurrentVersionExpiration",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) PutNoncurrentVersionTransition(value interface{}) {
	if err := o.validatePutNoncurrentVersionTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNoncurrentVersionTransition",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) PutTransition(value interface{}) {
	if err := o.validatePutTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTransition",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) ResetExpiration() {
	_jsii_.InvokeVoid(
		o,
		"resetExpiration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) ResetNoncurrentVersionExpiration() {
	_jsii_.InvokeVoid(
		o,
		"resetNoncurrentVersionExpiration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) ResetNoncurrentVersionTransition() {
	_jsii_.InvokeVoid(
		o,
		"resetNoncurrentVersionTransition",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		o,
		"resetPrefix",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) ResetTransition() {
	_jsii_.InvokeVoid(
		o,
		"resetTransition",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketLifecycleRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

