// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/d/lb_listener_v3 opentelekomcloud_lb_listener_v3}.
type DataOpentelekomcloudLbListenerV3 interface {
	cdktf.TerraformDataSource
	AdminStateUp() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCaTlsContainerRef() *string
	SetClientCaTlsContainerRef(val *string)
	ClientCaTlsContainerRefInput() *string
	ClientTimeout() *float64
	SetClientTimeout(val *float64)
	ClientTimeoutInput() *float64
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
	Http2Enable() cdktf.IResolvable
	Id() *string
	SetId(val *string)
	IdInput() *string
	InsertHeaders() DataOpentelekomcloudLbListenerV3InsertHeadersList
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
	MemberAddress() *string
	SetMemberAddress(val *string)
	MemberAddressInput() *string
	MemberDeviceId() *string
	SetMemberDeviceId(val *string)
	MemberDeviceIdInput() *string
	MemberRetryEnable() cdktf.IResolvable
	MemberTimeout() *float64
	SetMemberTimeout(val *float64)
	MemberTimeoutInput() *float64
	MemoryRetryEnable() cdktf.IResolvable
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProjectId() *string
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
	RawOverrides() interface{}
	SniContainerRefs() *[]*string
	Tags() cdktf.StringMap
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
	ResetClientCaTlsContainerRef()
	ResetClientTimeout()
	ResetDefaultPoolId()
	ResetDefaultTlsContainerRef()
	ResetDescription()
	ResetId()
	ResetKeepAliveTimeout()
	ResetLoadbalancerId()
	ResetMemberAddress()
	ResetMemberDeviceId()
	ResetMemberTimeout()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtocol()
	ResetProtocolPort()
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

// The jsii proxy struct for DataOpentelekomcloudLbListenerV3
type jsiiProxy_DataOpentelekomcloudLbListenerV3 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) AdminStateUp() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"adminStateUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) ClientCaTlsContainerRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCaTlsContainerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) ClientCaTlsContainerRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCaTlsContainerRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) ClientTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) ClientTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) DefaultPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) DefaultPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) DefaultTlsContainerRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTlsContainerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) DefaultTlsContainerRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTlsContainerRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) Http2Enable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"http2Enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) InsertHeaders() DataOpentelekomcloudLbListenerV3InsertHeadersList {
	var returns DataOpentelekomcloudLbListenerV3InsertHeadersList
	_jsii_.Get(
		j,
		"insertHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) KeepAliveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) KeepAliveTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) LoadbalancerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) LoadbalancerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) MemberAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) MemberAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) MemberDeviceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberDeviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) MemberDeviceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberDeviceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) MemberRetryEnable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"memberRetryEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) MemberTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memberTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) MemberTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memberTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) MemoryRetryEnable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"memoryRetryEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) ProtocolPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) ProtocolPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SniContainerRefs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sniContainerRefs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) TlsCiphersPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCiphersPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) TlsCiphersPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCiphersPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/d/lb_listener_v3 opentelekomcloud_lb_listener_v3} Data Source.
func NewDataOpentelekomcloudLbListenerV3(scope constructs.Construct, id *string, config *DataOpentelekomcloudLbListenerV3Config) DataOpentelekomcloudLbListenerV3 {
	_init_.Initialize()

	j := jsiiProxy_DataOpentelekomcloudLbListenerV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.DataOpentelekomcloudLbListenerV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/d/lb_listener_v3 opentelekomcloud_lb_listener_v3} Data Source.
func NewDataOpentelekomcloudLbListenerV3_Override(d DataOpentelekomcloudLbListenerV3, scope constructs.Construct, id *string, config *DataOpentelekomcloudLbListenerV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.DataOpentelekomcloudLbListenerV3",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetClientCaTlsContainerRef(val *string) {
	_jsii_.Set(
		j,
		"clientCaTlsContainerRef",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetClientTimeout(val *float64) {
	_jsii_.Set(
		j,
		"clientTimeout",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetDefaultPoolId(val *string) {
	_jsii_.Set(
		j,
		"defaultPoolId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetDefaultTlsContainerRef(val *string) {
	_jsii_.Set(
		j,
		"defaultTlsContainerRef",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetKeepAliveTimeout(val *float64) {
	_jsii_.Set(
		j,
		"keepAliveTimeout",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetLoadbalancerId(val *string) {
	_jsii_.Set(
		j,
		"loadbalancerId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetMemberAddress(val *string) {
	_jsii_.Set(
		j,
		"memberAddress",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetMemberDeviceId(val *string) {
	_jsii_.Set(
		j,
		"memberDeviceId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetMemberTimeout(val *float64) {
	_jsii_.Set(
		j,
		"memberTimeout",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetProtocolPort(val *float64) {
	_jsii_.Set(
		j,
		"protocolPort",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudLbListenerV3) SetTlsCiphersPolicy(val *string) {
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
func DataOpentelekomcloudLbListenerV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.DataOpentelekomcloudLbListenerV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudLbListenerV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.DataOpentelekomcloudLbListenerV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetClientCaTlsContainerRef() {
	_jsii_.InvokeVoid(
		d,
		"resetClientCaTlsContainerRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetClientTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetClientTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetDefaultPoolId() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultPoolId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetDefaultTlsContainerRef() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultTlsContainerRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetKeepAliveTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetKeepAliveTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetLoadbalancerId() {
	_jsii_.InvokeVoid(
		d,
		"resetLoadbalancerId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetMemberAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetMemberAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetMemberDeviceId() {
	_jsii_.InvokeVoid(
		d,
		"resetMemberDeviceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetMemberTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetMemberTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetProtocolPort() {
	_jsii_.InvokeVoid(
		d,
		"resetProtocolPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ResetTlsCiphersPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsCiphersPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudLbListenerV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
