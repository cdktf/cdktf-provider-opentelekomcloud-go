// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/lblistenerv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_listener_v2 opentelekomcloud_lb_listener_v2}.
type LbListenerV2 interface {
	cdktf.TerraformResource
	AdminStateUp() interface{}
	SetAdminStateUp(val interface{})
	AdminStateUpInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCaTlsContainerRef() *string
	SetClientCaTlsContainerRef(val *string)
	ClientCaTlsContainerRefInput() *string
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
	IpGroup() LbListenerV2IpGroupOutputReference
	IpGroupInput() *LbListenerV2IpGroup
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadbalancerId() *string
	SetLoadbalancerId(val *string)
	LoadbalancerIdInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SniContainerRefs() *[]*string
	SetSniContainerRefs(val *[]*string)
	SniContainerRefsInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LbListenerV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	TlsCiphersPolicy() *string
	SetTlsCiphersPolicy(val *string)
	TlsCiphersPolicyInput() *string
	TransparentClientIpEnable() interface{}
	SetTransparentClientIpEnable(val interface{})
	TransparentClientIpEnableInput() interface{}
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
	PutIpGroup(value *LbListenerV2IpGroup)
	PutTimeouts(value *LbListenerV2Timeouts)
	ResetAdminStateUp()
	ResetClientCaTlsContainerRef()
	ResetDefaultPoolId()
	ResetDefaultTlsContainerRef()
	ResetDescription()
	ResetHttp2Enable()
	ResetId()
	ResetIpGroup()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSniContainerRefs()
	ResetTags()
	ResetTenantId()
	ResetTimeouts()
	ResetTlsCiphersPolicy()
	ResetTransparentClientIpEnable()
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

// The jsii proxy struct for LbListenerV2
type jsiiProxy_LbListenerV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LbListenerV2) AdminStateUp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminStateUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) AdminStateUpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminStateUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) ClientCaTlsContainerRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCaTlsContainerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) ClientCaTlsContainerRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCaTlsContainerRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) DefaultPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) DefaultPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) DefaultTlsContainerRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTlsContainerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) DefaultTlsContainerRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTlsContainerRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Http2Enable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Http2EnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2EnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) IpGroup() LbListenerV2IpGroupOutputReference {
	var returns LbListenerV2IpGroupOutputReference
	_jsii_.Get(
		j,
		"ipGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) IpGroupInput() *LbListenerV2IpGroup {
	var returns *LbListenerV2IpGroup
	_jsii_.Get(
		j,
		"ipGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) LoadbalancerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) LoadbalancerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) ProtocolPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) ProtocolPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) SniContainerRefs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sniContainerRefs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) SniContainerRefsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sniContainerRefsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) Timeouts() LbListenerV2TimeoutsOutputReference {
	var returns LbListenerV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) TlsCiphersPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCiphersPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) TlsCiphersPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCiphersPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) TransparentClientIpEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transparentClientIpEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerV2) TransparentClientIpEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transparentClientIpEnableInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_listener_v2 opentelekomcloud_lb_listener_v2} Resource.
func NewLbListenerV2(scope constructs.Construct, id *string, config *LbListenerV2Config) LbListenerV2 {
	_init_.Initialize()

	if err := validateNewLbListenerV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbListenerV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.lbListenerV2.LbListenerV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_listener_v2 opentelekomcloud_lb_listener_v2} Resource.
func NewLbListenerV2_Override(l LbListenerV2, scope constructs.Construct, id *string, config *LbListenerV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.lbListenerV2.LbListenerV2",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LbListenerV2)SetAdminStateUp(val interface{}) {
	if err := j.validateSetAdminStateUpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminStateUp",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetClientCaTlsContainerRef(val *string) {
	if err := j.validateSetClientCaTlsContainerRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCaTlsContainerRef",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetDefaultPoolId(val *string) {
	if err := j.validateSetDefaultPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPoolId",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetDefaultTlsContainerRef(val *string) {
	if err := j.validateSetDefaultTlsContainerRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTlsContainerRef",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetHttp2Enable(val interface{}) {
	if err := j.validateSetHttp2EnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http2Enable",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetLoadbalancerId(val *string) {
	if err := j.validateSetLoadbalancerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadbalancerId",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetProtocolPort(val *float64) {
	if err := j.validateSetProtocolPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolPort",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetSniContainerRefs(val *[]*string) {
	if err := j.validateSetSniContainerRefsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sniContainerRefs",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetTlsCiphersPolicy(val *string) {
	if err := j.validateSetTlsCiphersPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCiphersPolicy",
		val,
	)
}

func (j *jsiiProxy_LbListenerV2)SetTransparentClientIpEnable(val interface{}) {
	if err := j.validateSetTransparentClientIpEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentClientIpEnable",
		val,
	)
}

// Generates CDKTF code for importing a LbListenerV2 resource upon running "cdktf plan <stack-name>".
func LbListenerV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLbListenerV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.lbListenerV2.LbListenerV2",
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
func LbListenerV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbListenerV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.lbListenerV2.LbListenerV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbListenerV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbListenerV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.lbListenerV2.LbListenerV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbListenerV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbListenerV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.lbListenerV2.LbListenerV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LbListenerV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.lbListenerV2.LbListenerV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LbListenerV2) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LbListenerV2) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LbListenerV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LbListenerV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbListenerV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LbListenerV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LbListenerV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LbListenerV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LbListenerV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LbListenerV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LbListenerV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LbListenerV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LbListenerV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbListenerV2) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LbListenerV2) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LbListenerV2) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LbListenerV2) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LbListenerV2) PutIpGroup(value *LbListenerV2IpGroup) {
	if err := l.validatePutIpGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIpGroup",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbListenerV2) PutTimeouts(value *LbListenerV2Timeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbListenerV2) ResetAdminStateUp() {
	_jsii_.InvokeVoid(
		l,
		"resetAdminStateUp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetClientCaTlsContainerRef() {
	_jsii_.InvokeVoid(
		l,
		"resetClientCaTlsContainerRef",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetDefaultPoolId() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultPoolId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetDefaultTlsContainerRef() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultTlsContainerRef",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetHttp2Enable() {
	_jsii_.InvokeVoid(
		l,
		"resetHttp2Enable",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetIpGroup() {
	_jsii_.InvokeVoid(
		l,
		"resetIpGroup",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetName() {
	_jsii_.InvokeVoid(
		l,
		"resetName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetRegion() {
	_jsii_.InvokeVoid(
		l,
		"resetRegion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetSniContainerRefs() {
	_jsii_.InvokeVoid(
		l,
		"resetSniContainerRefs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetTenantId() {
	_jsii_.InvokeVoid(
		l,
		"resetTenantId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetTlsCiphersPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetTlsCiphersPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) ResetTransparentClientIpEnable() {
	_jsii_.InvokeVoid(
		l,
		"resetTransparentClientIpEnable",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

