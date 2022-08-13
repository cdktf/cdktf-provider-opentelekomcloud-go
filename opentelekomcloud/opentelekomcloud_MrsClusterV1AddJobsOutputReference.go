// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MrsClusterV1AddJobsOutputReference interface {
	cdktf.ComplexObject
	Arguments() *string
	SetArguments(val *string)
	ArgumentsInput() *string
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
	FileAction() *string
	SetFileAction(val *string)
	FileActionInput() *string
	// Experimental.
	Fqn() *string
	HiveScriptPath() *string
	SetHiveScriptPath(val *string)
	HiveScriptPathInput() *string
	Hql() *string
	SetHql(val *string)
	HqlInput() *string
	Input() *string
	SetInput(val *string)
	InputInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JarPath() *string
	SetJarPath(val *string)
	JarPathInput() *string
	JobLog() *string
	SetJobLog(val *string)
	JobLogInput() *string
	JobName() *string
	SetJobName(val *string)
	JobNameInput() *string
	JobType() *float64
	SetJobType(val *float64)
	JobTypeInput() *float64
	Output() *string
	SetOutput(val *string)
	OutputInput() *string
	ShutdownCluster() interface{}
	SetShutdownCluster(val interface{})
	ShutdownClusterInput() interface{}
	SubmitJobOnceClusterRun() interface{}
	SetSubmitJobOnceClusterRun(val interface{})
	SubmitJobOnceClusterRunInput() interface{}
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
	ResetArguments()
	ResetFileAction()
	ResetHiveScriptPath()
	ResetHql()
	ResetInput()
	ResetJobLog()
	ResetOutput()
	ResetShutdownCluster()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MrsClusterV1AddJobsOutputReference
type jsiiProxy_MrsClusterV1AddJobsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) Arguments() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) ArgumentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) FileAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) FileActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) HiveScriptPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hiveScriptPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) HiveScriptPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hiveScriptPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) Hql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) HqlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) JarPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jarPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) JarPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jarPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) JobLog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) JobLogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) JobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) JobType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jobType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) JobTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jobTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) Output() *string {
	var returns *string
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) OutputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) ShutdownCluster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shutdownCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) ShutdownClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shutdownClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SubmitJobOnceClusterRun() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"submitJobOnceClusterRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SubmitJobOnceClusterRunInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"submitJobOnceClusterRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMrsClusterV1AddJobsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MrsClusterV1AddJobsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MrsClusterV1AddJobsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.MrsClusterV1AddJobsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMrsClusterV1AddJobsOutputReference_Override(m MrsClusterV1AddJobsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.MrsClusterV1AddJobsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetArguments(val *string) {
	_jsii_.Set(
		j,
		"arguments",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetFileAction(val *string) {
	_jsii_.Set(
		j,
		"fileAction",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetHiveScriptPath(val *string) {
	_jsii_.Set(
		j,
		"hiveScriptPath",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetHql(val *string) {
	_jsii_.Set(
		j,
		"hql",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetInput(val *string) {
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetJarPath(val *string) {
	_jsii_.Set(
		j,
		"jarPath",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetJobLog(val *string) {
	_jsii_.Set(
		j,
		"jobLog",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetJobName(val *string) {
	_jsii_.Set(
		j,
		"jobName",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetJobType(val *float64) {
	_jsii_.Set(
		j,
		"jobType",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetOutput(val *string) {
	_jsii_.Set(
		j,
		"output",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetShutdownCluster(val interface{}) {
	_jsii_.Set(
		j,
		"shutdownCluster",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetSubmitJobOnceClusterRun(val interface{}) {
	_jsii_.Set(
		j,
		"submitJobOnceClusterRun",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1AddJobsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) ResetArguments() {
	_jsii_.InvokeVoid(
		m,
		"resetArguments",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) ResetFileAction() {
	_jsii_.InvokeVoid(
		m,
		"resetFileAction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) ResetHiveScriptPath() {
	_jsii_.InvokeVoid(
		m,
		"resetHiveScriptPath",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) ResetHql() {
	_jsii_.InvokeVoid(
		m,
		"resetHql",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) ResetInput() {
	_jsii_.InvokeVoid(
		m,
		"resetInput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) ResetJobLog() {
	_jsii_.InvokeVoid(
		m,
		"resetJobLog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) ResetOutput() {
	_jsii_.InvokeVoid(
		m,
		"resetOutput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) ResetShutdownCluster() {
	_jsii_.InvokeVoid(
		m,
		"resetShutdownCluster",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1AddJobsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

