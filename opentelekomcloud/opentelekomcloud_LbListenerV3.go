// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_listener_v3 opentelekomcloud_lb_listener_v3}.
type LbListenerV3 interface {
	cdktf.TerraformResource
	AdminStateUp() interface{}
	SetAdminStateUp(val interface{})
	AdminStateUpInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCaTlsContainerRef() *string
	SetClientCaTlsContainerRef(val *string)
	ClientCaTlsContainerRefInput() *string
	ClientTimeout() *float64
	SetClientTimeout(val *float64)
	ClientTimeoutInput() *float64
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
	CreatedAt() *string
	DefaultPoolId() *string
	SetDefaultPoolId(val *string)
	DefaultPoolIdInput() *string
	DefaultTlsContainerRef() *string
	SetDefaultTlsContainerRef(val *string)
	DefaultTlsContainerRefInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Http2Enable() interface{}
	SetHttp2Enable(val interface{})
	Http2EnableInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	InsertHeaders() LbListenerV3InsertHeadersOutputReference
	InsertHeadersInput() *LbListenerV3InsertHeaders
	KeepAliveTimeout() *float64
	SetKeepAliveTimeout(val *float64)
	KeepAliveTimeoutInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadbalancerId() *string
	SetLoadbalancerId(val *string)
	LoadbalancerIdInput() *string
	MemberRetryEnable() interface{}
	SetMemberRetryEnable(val interface{})
	MemberRetryEnableInput() interface{}
	MemberTimeout() *float64
	SetMemberTimeout(val *float64)
	MemberTimeoutInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	ProtocolPort() *float64
	SetProtocolPort(val *float64)
	ProtocolPortInput() *float64
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
	SniContainerRefs() *[]*string
	SetSniContainerRefs(val *[]*string)
	SniContainerRefsInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TlsCiphersPolicy() *string
	SetTlsCiphersPolicy(val *string)
	TlsCiphersPolicyInput() *string
	UpdatedAt() *string
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
	PutInsertHeaders(value *LbListenerV3InsertHeaders)
	ResetAdminStateUp()
	ResetClientCaTlsContainerRef()
	ResetClientTimeout()
	ResetDefaultPoolId()
	ResetDefaultTlsContainerRef()
	ResetDescription()
	ResetHttp2Enable()
	ResetId()
	ResetInsertHeaders()
	ResetKeepAliveTimeout()
	ResetMemberRetryEnable()
	ResetMemberTimeout()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSniContainerRefs()
	ResetTags()
	ResetTlsCiphersPolicy()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LbListenerV3
type jsiiProxy_LbListenerV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LbListenerV3) AdminStateUp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminStateUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) AdminStateUpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminStateUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) ClientCaTlsContainerRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCaTlsContainerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) ClientCaTlsContainerRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCaTlsContainerRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) ClientTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) ClientTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) DefaultPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) DefaultPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) DefaultTlsContainerRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTlsContainerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) DefaultTlsContainerRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTlsContainerRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Http2Enable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Http2EnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2EnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) InsertHeaders() LbListenerV3InsertHeadersOutputReference {
	var returns LbListenerV3InsertHeadersOutputReference
	_jsii_.Get(
		j,
		"insertHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) InsertHeadersInput() *LbListenerV3InsertHeaders {
	var returns *LbListenerV3InsertHeaders
	_jsii_.Get(
		j,
		"insertHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) KeepAliveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) KeepAliveTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) LoadbalancerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) LoadbalancerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) MemberRetryEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memberRetryEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) MemberRetryEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memberRetryEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) MemberTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memberTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) MemberTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memberTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) ProtocolPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) ProtocolPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) SniContainerRefs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sniContainerRefs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) SniContainerRefsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sniContainerRefsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) TlsCiphersPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCiphersPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) TlsCiphersPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCiphersPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV3) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_listener_v3 opentelekomcloud_lb_listener_v3} Resource.
func NewLbListenerV3(scope constructs.Construct, id *string, config *LbListenerV3Config) LbListenerV3 {
	_init_.Initialize()

	j := jsiiProxy_LbListenerV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.LbListenerV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_listener_v3 opentelekomcloud_lb_listener_v3} Resource.
func NewLbListenerV3_Override(l LbListenerV3, scope constructs.Construct, id *string, config *LbListenerV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.LbListenerV3",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LbListenerV3) SetAdminStateUp(val interface{}) {
	_jsii_.Set(
		j,
		"adminStateUp",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetClientCaTlsContainerRef(val *string) {
	_jsii_.Set(
		j,
		"clientCaTlsContainerRef",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetClientTimeout(val *float64) {
	_jsii_.Set(
		j,
		"clientTimeout",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetDefaultPoolId(val *string) {
	_jsii_.Set(
		j,
		"defaultPoolId",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetDefaultTlsContainerRef(val *string) {
	_jsii_.Set(
		j,
		"defaultTlsContainerRef",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetHttp2Enable(val interface{}) {
	_jsii_.Set(
		j,
		"http2Enable",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetKeepAliveTimeout(val *float64) {
	_jsii_.Set(
		j,
		"keepAliveTimeout",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetLoadbalancerId(val *string) {
	_jsii_.Set(
		j,
		"loadbalancerId",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetMemberRetryEnable(val interface{}) {
	_jsii_.Set(
		j,
		"memberRetryEnable",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetMemberTimeout(val *float64) {
	_jsii_.Set(
		j,
		"memberTimeout",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetProtocolPort(val *float64) {
	_jsii_.Set(
		j,
		"protocolPort",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetSniContainerRefs(val *[]*string) {
	_jsii_.Set(
		j,
		"sniContainerRefs",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LbListenerV3) SetTlsCiphersPolicy(val *string) {
	_jsii_.Set(
		j,
		"tlsCiphersPolicy",
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
func LbListenerV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.LbListenerV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LbListenerV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.LbListenerV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LbListenerV3) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LbListenerV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LbListenerV3) PutInsertHeaders(value *LbListenerV3InsertHeaders) {
	_jsii_.InvokeVoid(
		l,
		"putInsertHeaders",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbListenerV3) ResetAdminStateUp() {
	_jsii_.InvokeVoid(
		l,
		"resetAdminStateUp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetClientCaTlsContainerRef() {
	_jsii_.InvokeVoid(
		l,
		"resetClientCaTlsContainerRef",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetClientTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetClientTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetDefaultPoolId() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultPoolId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetDefaultTlsContainerRef() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultTlsContainerRef",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetHttp2Enable() {
	_jsii_.InvokeVoid(
		l,
		"resetHttp2Enable",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetInsertHeaders() {
	_jsii_.InvokeVoid(
		l,
		"resetInsertHeaders",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetKeepAliveTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetKeepAliveTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetMemberRetryEnable() {
	_jsii_.InvokeVoid(
		l,
		"resetMemberRetryEnable",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetMemberTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetMemberTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetName() {
	_jsii_.InvokeVoid(
		l,
		"resetName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetSniContainerRefs() {
	_jsii_.InvokeVoid(
		l,
		"resetSniContainerRefs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) ResetTlsCiphersPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetTlsCiphersPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

