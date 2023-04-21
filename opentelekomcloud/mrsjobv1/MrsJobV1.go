package mrsjobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v6/mrsjobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/mrs_job_v1 opentelekomcloud_mrs_job_v1}.
type MrsJobV1 interface {
	cdktf.TerraformResource
	Arguments() *string
	SetArguments(val *string)
	ArgumentsInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HiveScriptPath() *string
	SetHiveScriptPath(val *string)
	HiveScriptPathInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Input() *string
	SetInput(val *string)
	InputInput() *string
	IsProtected() interface{}
	SetIsProtected(val interface{})
	IsProtectedInput() interface{}
	IsPublic() interface{}
	SetIsPublic(val interface{})
	IsPublicInput() interface{}
	JarPath() *string
	SetJarPath(val *string)
	JarPathInput() *string
	JobLog() *string
	SetJobLog(val *string)
	JobLogInput() *string
	JobName() *string
	SetJobName(val *string)
	JobNameInput() *string
	JobState() *string
	JobType() *float64
	SetJobType(val *float64)
	JobTypeInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Output() *string
	SetOutput(val *string)
	OutputInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MrsJobV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *MrsJobV1Timeouts)
	ResetArguments()
	ResetHiveScriptPath()
	ResetId()
	ResetInput()
	ResetIsProtected()
	ResetIsPublic()
	ResetJobLog()
	ResetOutput()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MrsJobV1
type jsiiProxy_MrsJobV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MrsJobV1) Arguments() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) ArgumentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) HiveScriptPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hiveScriptPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) HiveScriptPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hiveScriptPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) IsProtected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isProtected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) IsProtectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isProtectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) IsPublic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) IsPublicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPublicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) JarPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jarPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) JarPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jarPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) JobLog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) JobLogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) JobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) JobState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) JobType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jobType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) JobTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jobTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) Output() *string {
	var returns *string
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) OutputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) Timeouts() MrsJobV1TimeoutsOutputReference {
	var returns MrsJobV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsJobV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/mrs_job_v1 opentelekomcloud_mrs_job_v1} Resource.
func NewMrsJobV1(scope constructs.Construct, id *string, config *MrsJobV1Config) MrsJobV1 {
	_init_.Initialize()

	if err := validateNewMrsJobV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MrsJobV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.mrsJobV1.MrsJobV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/mrs_job_v1 opentelekomcloud_mrs_job_v1} Resource.
func NewMrsJobV1_Override(m MrsJobV1, scope constructs.Construct, id *string, config *MrsJobV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.mrsJobV1.MrsJobV1",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MrsJobV1)SetArguments(val *string) {
	if err := j.validateSetArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arguments",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetHiveScriptPath(val *string) {
	if err := j.validateSetHiveScriptPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiveScriptPath",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetInput(val *string) {
	if err := j.validateSetInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetIsProtected(val interface{}) {
	if err := j.validateSetIsProtectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isProtected",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetIsPublic(val interface{}) {
	if err := j.validateSetIsPublicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPublic",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetJarPath(val *string) {
	if err := j.validateSetJarPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jarPath",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetJobLog(val *string) {
	if err := j.validateSetJobLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobLog",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetJobName(val *string) {
	if err := j.validateSetJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobName",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetJobType(val *float64) {
	if err := j.validateSetJobTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobType",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetOutput(val *string) {
	if err := j.validateSetOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"output",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MrsJobV1)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func MrsJobV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMrsJobV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.mrsJobV1.MrsJobV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MrsJobV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMrsJobV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.mrsJobV1.MrsJobV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MrsJobV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMrsJobV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.mrsJobV1.MrsJobV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MrsJobV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.mrsJobV1.MrsJobV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MrsJobV1) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MrsJobV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MrsJobV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MrsJobV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MrsJobV1) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MrsJobV1) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MrsJobV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MrsJobV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MrsJobV1) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MrsJobV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MrsJobV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsJobV1) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MrsJobV1) PutTimeouts(value *MrsJobV1Timeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrsJobV1) ResetArguments() {
	_jsii_.InvokeVoid(
		m,
		"resetArguments",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsJobV1) ResetHiveScriptPath() {
	_jsii_.InvokeVoid(
		m,
		"resetHiveScriptPath",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsJobV1) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsJobV1) ResetInput() {
	_jsii_.InvokeVoid(
		m,
		"resetInput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsJobV1) ResetIsProtected() {
	_jsii_.InvokeVoid(
		m,
		"resetIsProtected",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsJobV1) ResetIsPublic() {
	_jsii_.InvokeVoid(
		m,
		"resetIsPublic",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsJobV1) ResetJobLog() {
	_jsii_.InvokeVoid(
		m,
		"resetJobLog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsJobV1) ResetOutput() {
	_jsii_.InvokeVoid(
		m,
		"resetOutput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsJobV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsJobV1) ResetRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsJobV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsJobV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsJobV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsJobV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsJobV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

