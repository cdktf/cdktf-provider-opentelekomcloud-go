// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cssconfigurationv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/cssconfigurationv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1 opentelekomcloud_css_configuration_v1}.
type CssConfigurationV1 interface {
	cdktf.TerraformResource
	AutoCreateIndex() *string
	SetAutoCreateIndex(val *string)
	AutoCreateIndexInput() *string
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
	HttpCorsAllowCredentials() *string
	SetHttpCorsAllowCredentials(val *string)
	HttpCorsAllowCredentialsInput() *string
	HttpCorsAllowHeaders() *string
	SetHttpCorsAllowHeaders(val *string)
	HttpCorsAllowHeadersInput() *string
	HttpCorsAllowMethods() *string
	SetHttpCorsAllowMethods(val *string)
	HttpCorsAllowMethodsInput() *string
	HttpCorsAllowOrigin() *string
	SetHttpCorsAllowOrigin(val *string)
	HttpCorsAllowOriginInput() *string
	HttpCorsEnabled() *string
	SetHttpCorsEnabled(val *string)
	HttpCorsEnabledInput() *string
	HttpCorsMaxAge() *string
	SetHttpCorsMaxAge(val *string)
	HttpCorsMaxAgeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IndicesQueriesCacheSize() *string
	SetIndicesQueriesCacheSize(val *string)
	IndicesQueriesCacheSizeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	ReindexRemoteWhitelist() *string
	SetReindexRemoteWhitelist(val *string)
	ReindexRemoteWhitelistInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThreadPoolForceMergeSize() *string
	SetThreadPoolForceMergeSize(val *string)
	ThreadPoolForceMergeSizeInput() *string
	Timeouts() CssConfigurationV1TimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *CssConfigurationV1Timeouts)
	ResetAutoCreateIndex()
	ResetHttpCorsAllowCredentials()
	ResetHttpCorsAllowHeaders()
	ResetHttpCorsAllowMethods()
	ResetHttpCorsAllowOrigin()
	ResetHttpCorsEnabled()
	ResetHttpCorsMaxAge()
	ResetId()
	ResetIndicesQueriesCacheSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReindexRemoteWhitelist()
	ResetThreadPoolForceMergeSize()
	ResetTimeouts()
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

// The jsii proxy struct for CssConfigurationV1
type jsiiProxy_CssConfigurationV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CssConfigurationV1) AutoCreateIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoCreateIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) AutoCreateIndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoCreateIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) HttpCorsAllowCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorsAllowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) HttpCorsAllowCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorsAllowCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) HttpCorsAllowHeaders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorsAllowHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) HttpCorsAllowHeadersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorsAllowHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) HttpCorsAllowMethods() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorsAllowMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) HttpCorsAllowMethodsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorsAllowMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) HttpCorsAllowOrigin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorsAllowOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) HttpCorsAllowOriginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorsAllowOriginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) HttpCorsEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) HttpCorsEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) HttpCorsMaxAge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorsMaxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) HttpCorsMaxAgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorsMaxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) IndicesQueriesCacheSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indicesQueriesCacheSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) IndicesQueriesCacheSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indicesQueriesCacheSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) ReindexRemoteWhitelist() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reindexRemoteWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) ReindexRemoteWhitelistInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reindexRemoteWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) ThreadPoolForceMergeSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threadPoolForceMergeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) ThreadPoolForceMergeSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threadPoolForceMergeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) Timeouts() CssConfigurationV1TimeoutsOutputReference {
	var returns CssConfigurationV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CssConfigurationV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1 opentelekomcloud_css_configuration_v1} Resource.
func NewCssConfigurationV1(scope constructs.Construct, id *string, config *CssConfigurationV1Config) CssConfigurationV1 {
	_init_.Initialize()

	if err := validateNewCssConfigurationV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CssConfigurationV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cssConfigurationV1.CssConfigurationV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1 opentelekomcloud_css_configuration_v1} Resource.
func NewCssConfigurationV1_Override(c CssConfigurationV1, scope constructs.Construct, id *string, config *CssConfigurationV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cssConfigurationV1.CssConfigurationV1",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetAutoCreateIndex(val *string) {
	if err := j.validateSetAutoCreateIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoCreateIndex",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetHttpCorsAllowCredentials(val *string) {
	if err := j.validateSetHttpCorsAllowCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCorsAllowCredentials",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetHttpCorsAllowHeaders(val *string) {
	if err := j.validateSetHttpCorsAllowHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCorsAllowHeaders",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetHttpCorsAllowMethods(val *string) {
	if err := j.validateSetHttpCorsAllowMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCorsAllowMethods",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetHttpCorsAllowOrigin(val *string) {
	if err := j.validateSetHttpCorsAllowOriginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCorsAllowOrigin",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetHttpCorsEnabled(val *string) {
	if err := j.validateSetHttpCorsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCorsEnabled",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetHttpCorsMaxAge(val *string) {
	if err := j.validateSetHttpCorsMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCorsMaxAge",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetIndicesQueriesCacheSize(val *string) {
	if err := j.validateSetIndicesQueriesCacheSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesQueriesCacheSize",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetReindexRemoteWhitelist(val *string) {
	if err := j.validateSetReindexRemoteWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reindexRemoteWhitelist",
		val,
	)
}

func (j *jsiiProxy_CssConfigurationV1)SetThreadPoolForceMergeSize(val *string) {
	if err := j.validateSetThreadPoolForceMergeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolForceMergeSize",
		val,
	)
}

// Generates CDKTF code for importing a CssConfigurationV1 resource upon running "cdktf plan <stack-name>".
func CssConfigurationV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCssConfigurationV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cssConfigurationV1.CssConfigurationV1",
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
func CssConfigurationV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCssConfigurationV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cssConfigurationV1.CssConfigurationV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CssConfigurationV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCssConfigurationV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cssConfigurationV1.CssConfigurationV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CssConfigurationV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCssConfigurationV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cssConfigurationV1.CssConfigurationV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CssConfigurationV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.cssConfigurationV1.CssConfigurationV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CssConfigurationV1) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CssConfigurationV1) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CssConfigurationV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CssConfigurationV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CssConfigurationV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CssConfigurationV1) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CssConfigurationV1) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CssConfigurationV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CssConfigurationV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CssConfigurationV1) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CssConfigurationV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CssConfigurationV1) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CssConfigurationV1) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CssConfigurationV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CssConfigurationV1) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CssConfigurationV1) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CssConfigurationV1) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CssConfigurationV1) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CssConfigurationV1) PutTimeouts(value *CssConfigurationV1Timeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetAutoCreateIndex() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoCreateIndex",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetHttpCorsAllowCredentials() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpCorsAllowCredentials",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetHttpCorsAllowHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpCorsAllowHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetHttpCorsAllowMethods() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpCorsAllowMethods",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetHttpCorsAllowOrigin() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpCorsAllowOrigin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetHttpCorsEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpCorsEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetHttpCorsMaxAge() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpCorsMaxAge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetIndicesQueriesCacheSize() {
	_jsii_.InvokeVoid(
		c,
		"resetIndicesQueriesCacheSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetReindexRemoteWhitelist() {
	_jsii_.InvokeVoid(
		c,
		"resetReindexRemoteWhitelist",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetThreadPoolForceMergeSize() {
	_jsii_.InvokeVoid(
		c,
		"resetThreadPoolForceMergeSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CssConfigurationV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CssConfigurationV1) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CssConfigurationV1) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CssConfigurationV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CssConfigurationV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CssConfigurationV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

