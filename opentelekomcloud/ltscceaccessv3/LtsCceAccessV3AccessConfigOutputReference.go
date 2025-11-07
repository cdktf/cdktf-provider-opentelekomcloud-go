// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltscceaccessv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/ltscceaccessv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsCceAccessV3AccessConfigOutputReference interface {
	cdktf.ComplexObject
	BlackPaths() *[]*string
	SetBlackPaths(val *[]*string)
	BlackPathsInput() *[]*string
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
	ContainerNameRegex() *string
	SetContainerNameRegex(val *string)
	ContainerNameRegexInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludeEnvs() *map[string]*string
	SetExcludeEnvs(val *map[string]*string)
	ExcludeEnvsInput() *map[string]*string
	ExcludeK8SLabels() *map[string]*string
	SetExcludeK8SLabels(val *map[string]*string)
	ExcludeK8SLabelsInput() *map[string]*string
	ExcludeLabels() *map[string]*string
	SetExcludeLabels(val *map[string]*string)
	ExcludeLabelsInput() *map[string]*string
	// Experimental.
	Fqn() *string
	IncludeEnvs() *map[string]*string
	SetIncludeEnvs(val *map[string]*string)
	IncludeEnvsInput() *map[string]*string
	IncludeK8SLabels() *map[string]*string
	SetIncludeK8SLabels(val *map[string]*string)
	IncludeK8SLabelsInput() *map[string]*string
	IncludeLabels() *map[string]*string
	SetIncludeLabels(val *map[string]*string)
	IncludeLabelsInput() *map[string]*string
	InternalValue() *LtsCceAccessV3AccessConfig
	SetInternalValue(val *LtsCceAccessV3AccessConfig)
	LogEnvs() *map[string]*string
	SetLogEnvs(val *map[string]*string)
	LogEnvsInput() *map[string]*string
	LogK8S() *map[string]*string
	SetLogK8S(val *map[string]*string)
	LogK8SInput() *map[string]*string
	LogLabels() *map[string]*string
	SetLogLabels(val *map[string]*string)
	LogLabelsInput() *map[string]*string
	MultiLogFormat() LtsCceAccessV3AccessConfigMultiLogFormatOutputReference
	MultiLogFormatInput() *LtsCceAccessV3AccessConfigMultiLogFormat
	NameSpaceRegex() *string
	SetNameSpaceRegex(val *string)
	NameSpaceRegexInput() *string
	Paths() *[]*string
	SetPaths(val *[]*string)
	PathsInput() *[]*string
	PathType() *string
	SetPathType(val *string)
	PathTypeInput() *string
	PodNameRegex() *string
	SetPodNameRegex(val *string)
	PodNameRegexInput() *string
	SingleLogFormat() LtsCceAccessV3AccessConfigSingleLogFormatOutputReference
	SingleLogFormatInput() *LtsCceAccessV3AccessConfigSingleLogFormat
	Stderr() interface{}
	SetStderr(val interface{})
	StderrInput() interface{}
	Stdout() interface{}
	SetStdout(val interface{})
	StdoutInput() interface{}
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
	PutMultiLogFormat(value *LtsCceAccessV3AccessConfigMultiLogFormat)
	PutSingleLogFormat(value *LtsCceAccessV3AccessConfigSingleLogFormat)
	ResetBlackPaths()
	ResetContainerNameRegex()
	ResetExcludeEnvs()
	ResetExcludeK8SLabels()
	ResetExcludeLabels()
	ResetIncludeEnvs()
	ResetIncludeK8SLabels()
	ResetIncludeLabels()
	ResetLogEnvs()
	ResetLogK8S()
	ResetLogLabels()
	ResetMultiLogFormat()
	ResetNameSpaceRegex()
	ResetPaths()
	ResetPodNameRegex()
	ResetSingleLogFormat()
	ResetStderr()
	ResetStdout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LtsCceAccessV3AccessConfigOutputReference
type jsiiProxy_LtsCceAccessV3AccessConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) BlackPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blackPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) BlackPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blackPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ContainerNameRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ContainerNameRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ExcludeEnvs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"excludeEnvs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ExcludeEnvsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"excludeEnvsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ExcludeK8SLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"excludeK8SLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ExcludeK8SLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"excludeK8SLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ExcludeLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"excludeLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ExcludeLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"excludeLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) IncludeEnvs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"includeEnvs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) IncludeEnvsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"includeEnvsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) IncludeK8SLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"includeK8SLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) IncludeK8SLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"includeK8SLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) IncludeLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"includeLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) IncludeLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"includeLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) InternalValue() *LtsCceAccessV3AccessConfig {
	var returns *LtsCceAccessV3AccessConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) LogEnvs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"logEnvs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) LogEnvsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"logEnvsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) LogK8S() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"logK8S",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) LogK8SInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"logK8SInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) LogLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"logLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) LogLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"logLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) MultiLogFormat() LtsCceAccessV3AccessConfigMultiLogFormatOutputReference {
	var returns LtsCceAccessV3AccessConfigMultiLogFormatOutputReference
	_jsii_.Get(
		j,
		"multiLogFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) MultiLogFormatInput() *LtsCceAccessV3AccessConfigMultiLogFormat {
	var returns *LtsCceAccessV3AccessConfigMultiLogFormat
	_jsii_.Get(
		j,
		"multiLogFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) NameSpaceRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameSpaceRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) NameSpaceRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameSpaceRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) Paths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"paths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) PathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) PathType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) PathTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) PodNameRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podNameRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) PodNameRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podNameRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) SingleLogFormat() LtsCceAccessV3AccessConfigSingleLogFormatOutputReference {
	var returns LtsCceAccessV3AccessConfigSingleLogFormatOutputReference
	_jsii_.Get(
		j,
		"singleLogFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) SingleLogFormatInput() *LtsCceAccessV3AccessConfigSingleLogFormat {
	var returns *LtsCceAccessV3AccessConfigSingleLogFormat
	_jsii_.Get(
		j,
		"singleLogFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) Stderr() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stderr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) StderrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stderrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) Stdout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) StdoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLtsCceAccessV3AccessConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LtsCceAccessV3AccessConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLtsCceAccessV3AccessConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LtsCceAccessV3AccessConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsCceAccessV3.LtsCceAccessV3AccessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLtsCceAccessV3AccessConfigOutputReference_Override(l LtsCceAccessV3AccessConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ltsCceAccessV3.LtsCceAccessV3AccessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetBlackPaths(val *[]*string) {
	if err := j.validateSetBlackPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blackPaths",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetContainerNameRegex(val *string) {
	if err := j.validateSetContainerNameRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerNameRegex",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetExcludeEnvs(val *map[string]*string) {
	if err := j.validateSetExcludeEnvsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeEnvs",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetExcludeK8SLabels(val *map[string]*string) {
	if err := j.validateSetExcludeK8SLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeK8SLabels",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetExcludeLabels(val *map[string]*string) {
	if err := j.validateSetExcludeLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeLabels",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetIncludeEnvs(val *map[string]*string) {
	if err := j.validateSetIncludeEnvsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeEnvs",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetIncludeK8SLabels(val *map[string]*string) {
	if err := j.validateSetIncludeK8SLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeK8SLabels",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetIncludeLabels(val *map[string]*string) {
	if err := j.validateSetIncludeLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeLabels",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetInternalValue(val *LtsCceAccessV3AccessConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetLogEnvs(val *map[string]*string) {
	if err := j.validateSetLogEnvsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logEnvs",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetLogK8S(val *map[string]*string) {
	if err := j.validateSetLogK8SParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logK8S",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetLogLabels(val *map[string]*string) {
	if err := j.validateSetLogLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLabels",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetNameSpaceRegex(val *string) {
	if err := j.validateSetNameSpaceRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameSpaceRegex",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetPaths(val *[]*string) {
	if err := j.validateSetPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paths",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetPathType(val *string) {
	if err := j.validateSetPathTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathType",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetPodNameRegex(val *string) {
	if err := j.validateSetPodNameRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podNameRegex",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetStderr(val interface{}) {
	if err := j.validateSetStderrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stderr",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetStdout(val interface{}) {
	if err := j.validateSetStdoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdout",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) PutMultiLogFormat(value *LtsCceAccessV3AccessConfigMultiLogFormat) {
	if err := l.validatePutMultiLogFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMultiLogFormat",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) PutSingleLogFormat(value *LtsCceAccessV3AccessConfigSingleLogFormat) {
	if err := l.validatePutSingleLogFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSingleLogFormat",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetBlackPaths() {
	_jsii_.InvokeVoid(
		l,
		"resetBlackPaths",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetContainerNameRegex() {
	_jsii_.InvokeVoid(
		l,
		"resetContainerNameRegex",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetExcludeEnvs() {
	_jsii_.InvokeVoid(
		l,
		"resetExcludeEnvs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetExcludeK8SLabels() {
	_jsii_.InvokeVoid(
		l,
		"resetExcludeK8SLabels",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetExcludeLabels() {
	_jsii_.InvokeVoid(
		l,
		"resetExcludeLabels",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetIncludeEnvs() {
	_jsii_.InvokeVoid(
		l,
		"resetIncludeEnvs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetIncludeK8SLabels() {
	_jsii_.InvokeVoid(
		l,
		"resetIncludeK8SLabels",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetIncludeLabels() {
	_jsii_.InvokeVoid(
		l,
		"resetIncludeLabels",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetLogEnvs() {
	_jsii_.InvokeVoid(
		l,
		"resetLogEnvs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetLogK8S() {
	_jsii_.InvokeVoid(
		l,
		"resetLogK8S",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetLogLabels() {
	_jsii_.InvokeVoid(
		l,
		"resetLogLabels",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetMultiLogFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetMultiLogFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetNameSpaceRegex() {
	_jsii_.InvokeVoid(
		l,
		"resetNameSpaceRegex",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetPaths() {
	_jsii_.InvokeVoid(
		l,
		"resetPaths",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetPodNameRegex() {
	_jsii_.InvokeVoid(
		l,
		"resetPodNameRegex",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetSingleLogFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetSingleLogFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetStderr() {
	_jsii_.InvokeVoid(
		l,
		"resetStderr",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ResetStdout() {
	_jsii_.InvokeVoid(
		l,
		"resetStdout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LtsCceAccessV3AccessConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

