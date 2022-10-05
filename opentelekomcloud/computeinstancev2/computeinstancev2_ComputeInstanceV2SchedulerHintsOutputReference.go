package computeinstancev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v3/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v3/computeinstancev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceV2SchedulerHintsOutputReference interface {
	cdktf.ComplexObject
	BuildNearHostIp() *string
	SetBuildNearHostIp(val *string)
	BuildNearHostIpInput() *string
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
	DehId() *string
	SetDehId(val *string)
	DehIdInput() *string
	DifferentHost() *[]*string
	SetDifferentHost(val *[]*string)
	DifferentHostInput() *[]*string
	// Experimental.
	Fqn() *string
	Group() *string
	SetGroup(val *string)
	GroupInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Query() *[]*string
	SetQuery(val *[]*string)
	QueryInput() *[]*string
	SameHost() *[]*string
	SetSameHost(val *[]*string)
	SameHostInput() *[]*string
	TargetCell() *string
	SetTargetCell(val *string)
	TargetCellInput() *string
	Tenancy() *string
	SetTenancy(val *string)
	TenancyInput() *string
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
	ResetBuildNearHostIp()
	ResetDehId()
	ResetDifferentHost()
	ResetGroup()
	ResetQuery()
	ResetSameHost()
	ResetTargetCell()
	ResetTenancy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceV2SchedulerHintsOutputReference
type jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) BuildNearHostIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildNearHostIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) BuildNearHostIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildNearHostIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) DehId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dehId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) DehIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dehIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) DifferentHost() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"differentHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) DifferentHostInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"differentHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) Query() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) QueryInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) SameHost() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sameHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) SameHostInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sameHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) TargetCell() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCell",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) TargetCellInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCellInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeInstanceV2SchedulerHintsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeInstanceV2SchedulerHintsOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInstanceV2SchedulerHintsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.computeInstanceV2.ComputeInstanceV2SchedulerHintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeInstanceV2SchedulerHintsOutputReference_Override(c ComputeInstanceV2SchedulerHintsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.computeInstanceV2.ComputeInstanceV2SchedulerHintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetBuildNearHostIp(val *string) {
	if err := j.validateSetBuildNearHostIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildNearHostIp",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetDehId(val *string) {
	if err := j.validateSetDehIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dehId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetDifferentHost(val *[]*string) {
	if err := j.validateSetDifferentHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"differentHost",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetQuery(val *[]*string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetSameHost(val *[]*string) {
	if err := j.validateSetSameHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sameHost",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetTargetCell(val *string) {
	if err := j.validateSetTargetCellParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCell",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetTenancy(val *string) {
	if err := j.validateSetTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) ResetBuildNearHostIp() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildNearHostIp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) ResetDehId() {
	_jsii_.InvokeVoid(
		c,
		"resetDehId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) ResetDifferentHost() {
	_jsii_.InvokeVoid(
		c,
		"resetDifferentHost",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		c,
		"resetQuery",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) ResetSameHost() {
	_jsii_.InvokeVoid(
		c,
		"resetSameHost",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) ResetTargetCell() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetCell",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) ResetTenancy() {
	_jsii_.InvokeVoid(
		c,
		"resetTenancy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2SchedulerHintsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

