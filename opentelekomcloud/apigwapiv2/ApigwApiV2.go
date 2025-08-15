// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwapiv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/apigwapiv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/apigw_api_v2 opentelekomcloud_apigw_api_v2}.
type ApigwApiV2 interface {
	cdktf.TerraformResource
	AuthorizerId() *string
	SetAuthorizerId(val *string)
	AuthorizerIdInput() *string
	BackendParams() ApigwApiV2BackendParamsList
	BackendParamsInput() interface{}
	BodyDescription() *string
	SetBodyDescription(val *string)
	BodyDescriptionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Cors() interface{}
	SetCors(val interface{})
	CorsInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	FailureResponse() *string
	SetFailureResponse(val *string)
	FailureResponseInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FuncGraph() ApigwApiV2FuncGraphOutputReference
	FuncGraphInput() *ApigwApiV2FuncGraph
	FuncGraphPolicy() ApigwApiV2FuncGraphPolicyList
	FuncGraphPolicyInput() interface{}
	GatewayId() *string
	SetGatewayId(val *string)
	GatewayIdInput() *string
	GroupId() *string
	SetGroupId(val *string)
	GroupIdInput() *string
	Http() ApigwApiV2HttpOutputReference
	HttpInput() *ApigwApiV2Http
	HttpPolicy() ApigwApiV2HttpPolicyList
	HttpPolicyInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MatchMode() *string
	SetMatchMode(val *string)
	MatchModeInput() *string
	Mock() ApigwApiV2MockOutputReference
	MockInput() *ApigwApiV2Mock
	MockPolicy() ApigwApiV2MockPolicyList
	MockPolicyInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	SetRegion(val *string)
	RegionInput() *string
	RegisteredAt() *string
	RequestMethod() *string
	SetRequestMethod(val *string)
	RequestMethodInput() *string
	RequestParams() ApigwApiV2RequestParamsList
	RequestParamsInput() interface{}
	RequestProtocol() *string
	SetRequestProtocol(val *string)
	RequestProtocolInput() *string
	RequestUri() *string
	SetRequestUri(val *string)
	RequestUriInput() *string
	ResponseId() *string
	SetResponseId(val *string)
	ResponseIdInput() *string
	SecurityAuthenticationEnabled() interface{}
	SetSecurityAuthenticationEnabled(val interface{})
	SecurityAuthenticationEnabledInput() interface{}
	SecurityAuthenticationType() *string
	SetSecurityAuthenticationType(val *string)
	SecurityAuthenticationTypeInput() *string
	SuccessResponse() *string
	SetSuccessResponse(val *string)
	SuccessResponseInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UpdatedAt() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutBackendParams(value interface{})
	PutFuncGraph(value *ApigwApiV2FuncGraph)
	PutFuncGraphPolicy(value interface{})
	PutHttp(value *ApigwApiV2Http)
	PutHttpPolicy(value interface{})
	PutMock(value *ApigwApiV2Mock)
	PutMockPolicy(value interface{})
	PutRequestParams(value interface{})
	ResetAuthorizerId()
	ResetBackendParams()
	ResetBodyDescription()
	ResetCors()
	ResetDescription()
	ResetFailureResponse()
	ResetFuncGraph()
	ResetFuncGraphPolicy()
	ResetHttp()
	ResetHttpPolicy()
	ResetId()
	ResetMatchMode()
	ResetMock()
	ResetMockPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetRequestParams()
	ResetResponseId()
	ResetSecurityAuthenticationEnabled()
	ResetSecurityAuthenticationType()
	ResetSuccessResponse()
	ResetTags()
	ResetVersion()
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

// The jsii proxy struct for ApigwApiV2
type jsiiProxy_ApigwApiV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApigwApiV2) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) AuthorizerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) BackendParams() ApigwApiV2BackendParamsList {
	var returns ApigwApiV2BackendParamsList
	_jsii_.Get(
		j,
		"backendParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) BackendParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) BodyDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) BodyDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Cors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) CorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) FailureResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) FailureResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) FuncGraph() ApigwApiV2FuncGraphOutputReference {
	var returns ApigwApiV2FuncGraphOutputReference
	_jsii_.Get(
		j,
		"funcGraph",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) FuncGraphInput() *ApigwApiV2FuncGraph {
	var returns *ApigwApiV2FuncGraph
	_jsii_.Get(
		j,
		"funcGraphInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) FuncGraphPolicy() ApigwApiV2FuncGraphPolicyList {
	var returns ApigwApiV2FuncGraphPolicyList
	_jsii_.Get(
		j,
		"funcGraphPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) FuncGraphPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"funcGraphPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) GatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) GroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Http() ApigwApiV2HttpOutputReference {
	var returns ApigwApiV2HttpOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) HttpInput() *ApigwApiV2Http {
	var returns *ApigwApiV2Http
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) HttpPolicy() ApigwApiV2HttpPolicyList {
	var returns ApigwApiV2HttpPolicyList
	_jsii_.Get(
		j,
		"httpPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) HttpPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) MatchMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) MatchModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Mock() ApigwApiV2MockOutputReference {
	var returns ApigwApiV2MockOutputReference
	_jsii_.Get(
		j,
		"mock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) MockInput() *ApigwApiV2Mock {
	var returns *ApigwApiV2Mock
	_jsii_.Get(
		j,
		"mockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) MockPolicy() ApigwApiV2MockPolicyList {
	var returns ApigwApiV2MockPolicyList
	_jsii_.Get(
		j,
		"mockPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) MockPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mockPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) RegisteredAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registeredAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) RequestMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) RequestMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) RequestParams() ApigwApiV2RequestParamsList {
	var returns ApigwApiV2RequestParamsList
	_jsii_.Get(
		j,
		"requestParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) RequestParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) RequestProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) RequestProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) RequestUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) RequestUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) ResponseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) ResponseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) SecurityAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) SecurityAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) SecurityAuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityAuthenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) SecurityAuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityAuthenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) SuccessResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) SuccessResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwApiV2) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/apigw_api_v2 opentelekomcloud_apigw_api_v2} Resource.
func NewApigwApiV2(scope constructs.Construct, id *string, config *ApigwApiV2Config) ApigwApiV2 {
	_init_.Initialize()

	if err := validateNewApigwApiV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigwApiV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/apigw_api_v2 opentelekomcloud_apigw_api_v2} Resource.
func NewApigwApiV2_Override(a ApigwApiV2, scope constructs.Construct, id *string, config *ApigwApiV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetAuthorizerId(val *string) {
	if err := j.validateSetAuthorizerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerId",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetBodyDescription(val *string) {
	if err := j.validateSetBodyDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bodyDescription",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetCors(val interface{}) {
	if err := j.validateSetCorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cors",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetFailureResponse(val *string) {
	if err := j.validateSetFailureResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureResponse",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetGatewayId(val *string) {
	if err := j.validateSetGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayId",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetGroupId(val *string) {
	if err := j.validateSetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetMatchMode(val *string) {
	if err := j.validateSetMatchModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchMode",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetRequestMethod(val *string) {
	if err := j.validateSetRequestMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMethod",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetRequestProtocol(val *string) {
	if err := j.validateSetRequestProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestProtocol",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetRequestUri(val *string) {
	if err := j.validateSetRequestUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestUri",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetResponseId(val *string) {
	if err := j.validateSetResponseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseId",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetSecurityAuthenticationEnabled(val interface{}) {
	if err := j.validateSetSecurityAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetSecurityAuthenticationType(val *string) {
	if err := j.validateSetSecurityAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityAuthenticationType",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetSuccessResponse(val *string) {
	if err := j.validateSetSuccessResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successResponse",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ApigwApiV2)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a ApigwApiV2 resource upon running "cdktf plan <stack-name>".
func ApigwApiV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApigwApiV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2",
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
func ApigwApiV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigwApiV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigwApiV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigwApiV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigwApiV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigwApiV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApigwApiV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.apigwApiV2.ApigwApiV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApigwApiV2) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApigwApiV2) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApigwApiV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApigwApiV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigwApiV2) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApigwApiV2) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigwApiV2) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApigwApiV2) PutBackendParams(value interface{}) {
	if err := a.validatePutBackendParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackendParams",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2) PutFuncGraph(value *ApigwApiV2FuncGraph) {
	if err := a.validatePutFuncGraphParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFuncGraph",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2) PutFuncGraphPolicy(value interface{}) {
	if err := a.validatePutFuncGraphPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFuncGraphPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2) PutHttp(value *ApigwApiV2Http) {
	if err := a.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2) PutHttpPolicy(value interface{}) {
	if err := a.validatePutHttpPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2) PutMock(value *ApigwApiV2Mock) {
	if err := a.validatePutMockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMock",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2) PutMockPolicy(value interface{}) {
	if err := a.validatePutMockPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMockPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2) PutRequestParams(value interface{}) {
	if err := a.validatePutRequestParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRequestParams",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetAuthorizerId() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetBackendParams() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetBodyDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetBodyDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetCors() {
	_jsii_.InvokeVoid(
		a,
		"resetCors",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetFailureResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetFuncGraph() {
	_jsii_.InvokeVoid(
		a,
		"resetFuncGraph",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetFuncGraphPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetFuncGraphPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetHttp() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetHttpPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetMatchMode() {
	_jsii_.InvokeVoid(
		a,
		"resetMatchMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetMock() {
	_jsii_.InvokeVoid(
		a,
		"resetMock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetMockPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetMockPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetRequestParams() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetResponseId() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetSecurityAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityAuthenticationEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetSecurityAuthenticationType() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityAuthenticationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetSuccessResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccessResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) ResetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwApiV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwApiV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

