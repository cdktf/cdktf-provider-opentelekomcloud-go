package obsbucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v4/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v4/obsbucket/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObsBucketWebsiteOutputReference interface {
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
	ErrorDocument() *string
	SetErrorDocument(val *string)
	ErrorDocumentInput() *string
	// Experimental.
	Fqn() *string
	IndexDocument() *string
	SetIndexDocument(val *string)
	IndexDocumentInput() *string
	InternalValue() *ObsBucketWebsite
	SetInternalValue(val *ObsBucketWebsite)
	RedirectAllRequestsTo() *string
	SetRedirectAllRequestsTo(val *string)
	RedirectAllRequestsToInput() *string
	RoutingRules() *string
	SetRoutingRules(val *string)
	RoutingRulesInput() *string
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
	ResetErrorDocument()
	ResetIndexDocument()
	ResetRedirectAllRequestsTo()
	ResetRoutingRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObsBucketWebsiteOutputReference
type jsiiProxy_ObsBucketWebsiteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) ErrorDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) ErrorDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) IndexDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) IndexDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) InternalValue() *ObsBucketWebsite {
	var returns *ObsBucketWebsite
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) RedirectAllRequestsTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectAllRequestsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) RedirectAllRequestsToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectAllRequestsToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) RoutingRules() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) RoutingRulesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObsBucketWebsiteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ObsBucketWebsiteOutputReference {
	_init_.Initialize()

	if err := validateNewObsBucketWebsiteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObsBucketWebsiteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.obsBucket.ObsBucketWebsiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewObsBucketWebsiteOutputReference_Override(o ObsBucketWebsiteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.obsBucket.ObsBucketWebsiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference)SetErrorDocument(val *string) {
	if err := j.validateSetErrorDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorDocument",
		val,
	)
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference)SetIndexDocument(val *string) {
	if err := j.validateSetIndexDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexDocument",
		val,
	)
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference)SetInternalValue(val *ObsBucketWebsite) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference)SetRedirectAllRequestsTo(val *string) {
	if err := j.validateSetRedirectAllRequestsToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectAllRequestsTo",
		val,
	)
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference)SetRoutingRules(val *string) {
	if err := j.validateSetRoutingRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingRules",
		val,
	)
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObsBucketWebsiteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) ResetErrorDocument() {
	_jsii_.InvokeVoid(
		o,
		"resetErrorDocument",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) ResetIndexDocument() {
	_jsii_.InvokeVoid(
		o,
		"resetIndexDocument",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) ResetRedirectAllRequestsTo() {
	_jsii_.InvokeVoid(
		o,
		"resetRedirectAllRequestsTo",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) ResetRoutingRules() {
	_jsii_.InvokeVoid(
		o,
		"resetRoutingRules",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObsBucketWebsiteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

