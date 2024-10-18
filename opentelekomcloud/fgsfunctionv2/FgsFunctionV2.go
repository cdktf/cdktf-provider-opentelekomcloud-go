// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsfunctionv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/fgsfunctionv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/fgs_function_v2 opentelekomcloud_fgs_function_v2}.
type FgsFunctionV2 interface {
	cdktf.TerraformResource
	Agency() *string
	SetAgency(val *string)
	AgencyInput() *string
	App() *string
	SetApp(val *string)
	AppAgency() *string
	SetAppAgency(val *string)
	AppAgencyInput() *string
	AppInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CodeFilename() *string
	SetCodeFilename(val *string)
	CodeFilenameInput() *string
	CodeType() *string
	SetCodeType(val *string)
	CodeTypeInput() *string
	CodeUrl() *string
	SetCodeUrl(val *string)
	CodeUrlInput() *string
	ConcurrencyNum() *float64
	SetConcurrencyNum(val *float64)
	ConcurrencyNumInput() *float64
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
	CustomImage() FgsFunctionV2CustomImageOutputReference
	CustomImageInput() *FgsFunctionV2CustomImage
	DependList() *[]*string
	SetDependList(val *[]*string)
	DependListInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DnsList() *string
	EncryptedUserData() *string
	SetEncryptedUserData(val *string)
	EncryptedUserDataInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FuncCode() *string
	SetFuncCode(val *string)
	FuncCodeInput() *string
	FuncMounts() FgsFunctionV2FuncMountsList
	FuncMountsInput() interface{}
	FunctiongraphVersion() *string
	SetFunctiongraphVersion(val *string)
	FunctiongraphVersionInput() *string
	GpuMemory() *float64
	SetGpuMemory(val *float64)
	GpuMemoryInput() *float64
	GpuType() *string
	Handler() *string
	SetHandler(val *string)
	HandlerInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InitializerHandler() *string
	SetInitializerHandler(val *string)
	InitializerHandlerInput() *string
	InitializerTimeout() *float64
	SetInitializerTimeout(val *float64)
	InitializerTimeoutInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogGroupId() *string
	SetLogGroupId(val *string)
	LogGroupIdInput() *string
	LogGroupName() *string
	SetLogGroupName(val *string)
	LogGroupNameInput() *string
	LogTopicId() *string
	SetLogTopicId(val *string)
	LogTopicIdInput() *string
	LogTopicName() *string
	SetLogTopicName(val *string)
	LogTopicNameInput() *string
	MaxInstanceNum() *string
	SetMaxInstanceNum(val *string)
	MaxInstanceNumInput() *string
	MemorySize() *float64
	SetMemorySize(val *float64)
	MemorySizeInput() *float64
	MountUserGroupId() *float64
	SetMountUserGroupId(val *float64)
	MountUserGroupIdInput() *float64
	MountUserId() *float64
	SetMountUserId(val *float64)
	MountUserIdInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
	// The tree node.
	Node() constructs.Node
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
	ReservedInstances() FgsFunctionV2ReservedInstancesList
	ReservedInstancesInput() interface{}
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Timeouts() FgsFunctionV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	Urn() *string
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	Version() *string
	Versions() FgsFunctionV2VersionsList
	VersionsInput() interface{}
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCustomImage(value *FgsFunctionV2CustomImage)
	PutFuncMounts(value interface{})
	PutReservedInstances(value interface{})
	PutTimeouts(value *FgsFunctionV2Timeouts)
	PutVersions(value interface{})
	ResetAgency()
	ResetApp()
	ResetAppAgency()
	ResetCodeFilename()
	ResetCodeType()
	ResetCodeUrl()
	ResetConcurrencyNum()
	ResetCustomImage()
	ResetDependList()
	ResetDescription()
	ResetEncryptedUserData()
	ResetFuncCode()
	ResetFuncMounts()
	ResetFunctiongraphVersion()
	ResetGpuMemory()
	ResetHandler()
	ResetId()
	ResetInitializerHandler()
	ResetInitializerTimeout()
	ResetLogGroupId()
	ResetLogGroupName()
	ResetLogTopicId()
	ResetLogTopicName()
	ResetMaxInstanceNum()
	ResetMountUserGroupId()
	ResetMountUserId()
	ResetNetworkId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReservedInstances()
	ResetTags()
	ResetTimeouts()
	ResetUserData()
	ResetVersions()
	ResetVpcId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FgsFunctionV2
type jsiiProxy_FgsFunctionV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FgsFunctionV2) Agency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) AgencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) App() *string {
	var returns *string
	_jsii_.Get(
		j,
		"app",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) AppAgency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appAgency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) AppAgencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appAgencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) AppInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) CodeFilename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeFilename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) CodeFilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeFilenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) CodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) CodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) CodeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) CodeUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) ConcurrencyNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrencyNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) ConcurrencyNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrencyNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) CustomImage() FgsFunctionV2CustomImageOutputReference {
	var returns FgsFunctionV2CustomImageOutputReference
	_jsii_.Get(
		j,
		"customImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) CustomImageInput() *FgsFunctionV2CustomImage {
	var returns *FgsFunctionV2CustomImage
	_jsii_.Get(
		j,
		"customImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) DependList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) DependListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) DnsList() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) EncryptedUserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptedUserData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) EncryptedUserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptedUserDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) FuncCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"funcCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) FuncCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"funcCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) FuncMounts() FgsFunctionV2FuncMountsList {
	var returns FgsFunctionV2FuncMountsList
	_jsii_.Get(
		j,
		"funcMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) FuncMountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"funcMountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) FunctiongraphVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functiongraphVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) FunctiongraphVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functiongraphVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) GpuMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpuMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) GpuMemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpuMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) GpuType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) InitializerHandler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initializerHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) InitializerHandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initializerHandlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) InitializerTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initializerTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) InitializerTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initializerTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) LogGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) LogGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) LogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) LogTopicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTopicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) LogTopicIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTopicIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) LogTopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTopicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) LogTopicNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTopicNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) MaxInstanceNum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxInstanceNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) MaxInstanceNumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxInstanceNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) MemorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) MountUserGroupId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mountUserGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) MountUserGroupIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mountUserGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) MountUserId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mountUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) MountUserIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mountUserIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) ReservedInstances() FgsFunctionV2ReservedInstancesList {
	var returns FgsFunctionV2ReservedInstancesList
	_jsii_.Get(
		j,
		"reservedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) ReservedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reservedInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Timeouts() FgsFunctionV2TimeoutsOutputReference {
	var returns FgsFunctionV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Urn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) Versions() FgsFunctionV2VersionsList {
	var returns FgsFunctionV2VersionsList
	_jsii_.Get(
		j,
		"versions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) VersionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FgsFunctionV2) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/fgs_function_v2 opentelekomcloud_fgs_function_v2} Resource.
func NewFgsFunctionV2(scope constructs.Construct, id *string, config *FgsFunctionV2Config) FgsFunctionV2 {
	_init_.Initialize()

	if err := validateNewFgsFunctionV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FgsFunctionV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/fgs_function_v2 opentelekomcloud_fgs_function_v2} Resource.
func NewFgsFunctionV2_Override(f FgsFunctionV2, scope constructs.Construct, id *string, config *FgsFunctionV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetAgency(val *string) {
	if err := j.validateSetAgencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agency",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetApp(val *string) {
	if err := j.validateSetAppParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"app",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetAppAgency(val *string) {
	if err := j.validateSetAppAgencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appAgency",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetCodeFilename(val *string) {
	if err := j.validateSetCodeFilenameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeFilename",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetCodeType(val *string) {
	if err := j.validateSetCodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeType",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetCodeUrl(val *string) {
	if err := j.validateSetCodeUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeUrl",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetConcurrencyNum(val *float64) {
	if err := j.validateSetConcurrencyNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"concurrencyNum",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetDependList(val *[]*string) {
	if err := j.validateSetDependListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependList",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetEncryptedUserData(val *string) {
	if err := j.validateSetEncryptedUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptedUserData",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetFuncCode(val *string) {
	if err := j.validateSetFuncCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"funcCode",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetFunctiongraphVersion(val *string) {
	if err := j.validateSetFunctiongraphVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functiongraphVersion",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetGpuMemory(val *float64) {
	if err := j.validateSetGpuMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpuMemory",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetInitializerHandler(val *string) {
	if err := j.validateSetInitializerHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initializerHandler",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetInitializerTimeout(val *float64) {
	if err := j.validateSetInitializerTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initializerTimeout",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetLogGroupId(val *string) {
	if err := j.validateSetLogGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupId",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetLogGroupName(val *string) {
	if err := j.validateSetLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetLogTopicId(val *string) {
	if err := j.validateSetLogTopicIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logTopicId",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetLogTopicName(val *string) {
	if err := j.validateSetLogTopicNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logTopicName",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetMaxInstanceNum(val *string) {
	if err := j.validateSetMaxInstanceNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstanceNum",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetMemorySize(val *float64) {
	if err := j.validateSetMemorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySize",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetMountUserGroupId(val *float64) {
	if err := j.validateSetMountUserGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountUserGroupId",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetMountUserId(val *float64) {
	if err := j.validateSetMountUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountUserId",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetNetworkId(val *string) {
	if err := j.validateSetNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_FgsFunctionV2)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a FgsFunctionV2 resource upon running "cdktf plan <stack-name>".
func FgsFunctionV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFgsFunctionV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func FgsFunctionV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFgsFunctionV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FgsFunctionV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFgsFunctionV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FgsFunctionV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFgsFunctionV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FgsFunctionV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.fgsFunctionV2.FgsFunctionV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FgsFunctionV2) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FgsFunctionV2) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FgsFunctionV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FgsFunctionV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FgsFunctionV2) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FgsFunctionV2) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FgsFunctionV2) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FgsFunctionV2) PutCustomImage(value *FgsFunctionV2CustomImage) {
	if err := f.validatePutCustomImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putCustomImage",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FgsFunctionV2) PutFuncMounts(value interface{}) {
	if err := f.validatePutFuncMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putFuncMounts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FgsFunctionV2) PutReservedInstances(value interface{}) {
	if err := f.validatePutReservedInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putReservedInstances",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FgsFunctionV2) PutTimeouts(value *FgsFunctionV2Timeouts) {
	if err := f.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FgsFunctionV2) PutVersions(value interface{}) {
	if err := f.validatePutVersionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putVersions",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetAgency() {
	_jsii_.InvokeVoid(
		f,
		"resetAgency",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetApp() {
	_jsii_.InvokeVoid(
		f,
		"resetApp",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetAppAgency() {
	_jsii_.InvokeVoid(
		f,
		"resetAppAgency",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetCodeFilename() {
	_jsii_.InvokeVoid(
		f,
		"resetCodeFilename",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetCodeType() {
	_jsii_.InvokeVoid(
		f,
		"resetCodeType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetCodeUrl() {
	_jsii_.InvokeVoid(
		f,
		"resetCodeUrl",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetConcurrencyNum() {
	_jsii_.InvokeVoid(
		f,
		"resetConcurrencyNum",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetCustomImage() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomImage",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetDependList() {
	_jsii_.InvokeVoid(
		f,
		"resetDependList",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetEncryptedUserData() {
	_jsii_.InvokeVoid(
		f,
		"resetEncryptedUserData",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetFuncCode() {
	_jsii_.InvokeVoid(
		f,
		"resetFuncCode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetFuncMounts() {
	_jsii_.InvokeVoid(
		f,
		"resetFuncMounts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetFunctiongraphVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetFunctiongraphVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetGpuMemory() {
	_jsii_.InvokeVoid(
		f,
		"resetGpuMemory",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetHandler() {
	_jsii_.InvokeVoid(
		f,
		"resetHandler",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetInitializerHandler() {
	_jsii_.InvokeVoid(
		f,
		"resetInitializerHandler",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetInitializerTimeout() {
	_jsii_.InvokeVoid(
		f,
		"resetInitializerTimeout",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetLogGroupId() {
	_jsii_.InvokeVoid(
		f,
		"resetLogGroupId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetLogGroupName() {
	_jsii_.InvokeVoid(
		f,
		"resetLogGroupName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetLogTopicId() {
	_jsii_.InvokeVoid(
		f,
		"resetLogTopicId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetLogTopicName() {
	_jsii_.InvokeVoid(
		f,
		"resetLogTopicName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetMaxInstanceNum() {
	_jsii_.InvokeVoid(
		f,
		"resetMaxInstanceNum",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetMountUserGroupId() {
	_jsii_.InvokeVoid(
		f,
		"resetMountUserGroupId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetMountUserId() {
	_jsii_.InvokeVoid(
		f,
		"resetMountUserId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetNetworkId() {
	_jsii_.InvokeVoid(
		f,
		"resetNetworkId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetReservedInstances() {
	_jsii_.InvokeVoid(
		f,
		"resetReservedInstances",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetUserData() {
	_jsii_.InvokeVoid(
		f,
		"resetUserData",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetVersions() {
	_jsii_.InvokeVoid(
		f,
		"resetVersions",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) ResetVpcId() {
	_jsii_.InvokeVoid(
		f,
		"resetVpcId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FgsFunctionV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FgsFunctionV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

