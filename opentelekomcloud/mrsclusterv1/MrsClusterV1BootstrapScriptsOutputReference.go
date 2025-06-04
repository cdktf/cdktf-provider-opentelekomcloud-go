// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mrsclusterv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/mrsclusterv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MrsClusterV1BootstrapScriptsOutputReference interface {
	cdktf.ComplexObject
	ActiveMaster() interface{}
	SetActiveMaster(val interface{})
	ActiveMasterInput() interface{}
	BeforeComponentStart() interface{}
	SetBeforeComponentStart(val interface{})
	BeforeComponentStartInput() interface{}
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
	FailAction() *string
	SetFailAction(val *string)
	FailActionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nodes() *[]*string
	SetNodes(val *[]*string)
	NodesInput() *[]*string
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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
	ResetActiveMaster()
	ResetBeforeComponentStart()
	ResetParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MrsClusterV1BootstrapScriptsOutputReference
type jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) ActiveMaster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeMaster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) ActiveMasterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeMasterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) BeforeComponentStart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"beforeComponentStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) BeforeComponentStartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"beforeComponentStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) FailAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) FailActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) Nodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) NodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewMrsClusterV1BootstrapScriptsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MrsClusterV1BootstrapScriptsOutputReference {
	_init_.Initialize()

	if err := validateNewMrsClusterV1BootstrapScriptsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.mrsClusterV1.MrsClusterV1BootstrapScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMrsClusterV1BootstrapScriptsOutputReference_Override(m MrsClusterV1BootstrapScriptsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.mrsClusterV1.MrsClusterV1BootstrapScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference)SetActiveMaster(val interface{}) {
	if err := j.validateSetActiveMasterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeMaster",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference)SetBeforeComponentStart(val interface{}) {
	if err := j.validateSetBeforeComponentStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"beforeComponentStart",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference)SetFailAction(val *string) {
	if err := j.validateSetFailActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failAction",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference)SetNodes(val *[]*string) {
	if err := j.validateSetNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodes",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference)SetParameters(val *string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) ResetActiveMaster() {
	_jsii_.InvokeVoid(
		m,
		"resetActiveMaster",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) ResetBeforeComponentStart() {
	_jsii_.InvokeVoid(
		m,
		"resetBeforeComponentStart",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		m,
		"resetParameters",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1BootstrapScriptsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

