// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1 opentelekomcloud_cts_tracker_v1}.
type CtsTrackerV1 interface {
	cdktf.TerraformResource
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FilePrefixName() *string
	SetFilePrefixName(val *string)
	FilePrefixNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsSendAllKeyOperation() interface{}
	SetIsSendAllKeyOperation(val interface{})
	IsSendAllKeyOperationInput() interface{}
	IsSupportSmn() interface{}
	SetIsSupportSmn(val interface{})
	IsSupportSmnInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NeedNotifyUserList() *[]*string
	SetNeedNotifyUserList(val *[]*string)
	NeedNotifyUserListInput() *[]*string
	// The tree node.
	Node() constructs.Node
	Operations() *[]*string
	SetOperations(val *[]*string)
	OperationsInput() *[]*string
	ProjectName() *string
	SetProjectName(val *string)
	ProjectNameInput() *string
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
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CtsTrackerV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	TopicId() *string
	SetTopicId(val *string)
	TopicIdInput() *string
	TrackerName() *string
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
	PutTimeouts(value *CtsTrackerV1Timeouts)
	ResetFilePrefixName()
	ResetId()
	ResetIsSendAllKeyOperation()
	ResetNeedNotifyUserList()
	ResetOperations()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectName()
	ResetRegion()
	ResetStatus()
	ResetTimeouts()
	ResetTopicId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CtsTrackerV1
type jsiiProxy_CtsTrackerV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CtsTrackerV1) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) FilePrefixName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePrefixName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) FilePrefixNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePrefixNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) IsSendAllKeyOperation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSendAllKeyOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) IsSendAllKeyOperationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSendAllKeyOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) IsSupportSmn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSupportSmn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) IsSupportSmnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSupportSmnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) NeedNotifyUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"needNotifyUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) NeedNotifyUserListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"needNotifyUserListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) Operations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) OperationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) ProjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) Timeouts() CtsTrackerV1TimeoutsOutputReference {
	var returns CtsTrackerV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) TopicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) TopicIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CtsTrackerV1) TrackerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackerName",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1 opentelekomcloud_cts_tracker_v1} Resource.
func NewCtsTrackerV1(scope constructs.Construct, id *string, config *CtsTrackerV1Config) CtsTrackerV1 {
	_init_.Initialize()

	j := jsiiProxy_CtsTrackerV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.CtsTrackerV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1 opentelekomcloud_cts_tracker_v1} Resource.
func NewCtsTrackerV1_Override(c CtsTrackerV1, scope constructs.Construct, id *string, config *CtsTrackerV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.CtsTrackerV1",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetBucketName(val *string) {
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetFilePrefixName(val *string) {
	_jsii_.Set(
		j,
		"filePrefixName",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetIsSendAllKeyOperation(val interface{}) {
	_jsii_.Set(
		j,
		"isSendAllKeyOperation",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetIsSupportSmn(val interface{}) {
	_jsii_.Set(
		j,
		"isSupportSmn",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetNeedNotifyUserList(val *[]*string) {
	_jsii_.Set(
		j,
		"needNotifyUserList",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetOperations(val *[]*string) {
	_jsii_.Set(
		j,
		"operations",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetProjectName(val *string) {
	_jsii_.Set(
		j,
		"projectName",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CtsTrackerV1) SetTopicId(val *string) {
	_jsii_.Set(
		j,
		"topicId",
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
func CtsTrackerV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.CtsTrackerV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CtsTrackerV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.CtsTrackerV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CtsTrackerV1) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CtsTrackerV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CtsTrackerV1) PutTimeouts(value *CtsTrackerV1Timeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CtsTrackerV1) ResetFilePrefixName() {
	_jsii_.InvokeVoid(
		c,
		"resetFilePrefixName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CtsTrackerV1) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CtsTrackerV1) ResetIsSendAllKeyOperation() {
	_jsii_.InvokeVoid(
		c,
		"resetIsSendAllKeyOperation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CtsTrackerV1) ResetNeedNotifyUserList() {
	_jsii_.InvokeVoid(
		c,
		"resetNeedNotifyUserList",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CtsTrackerV1) ResetOperations() {
	_jsii_.InvokeVoid(
		c,
		"resetOperations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CtsTrackerV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CtsTrackerV1) ResetProjectName() {
	_jsii_.InvokeVoid(
		c,
		"resetProjectName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CtsTrackerV1) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CtsTrackerV1) ResetStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetStatus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CtsTrackerV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CtsTrackerV1) ResetTopicId() {
	_jsii_.InvokeVoid(
		c,
		"resetTopicId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CtsTrackerV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CtsTrackerV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

