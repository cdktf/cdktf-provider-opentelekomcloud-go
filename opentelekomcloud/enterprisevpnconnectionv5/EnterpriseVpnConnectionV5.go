// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisevpnconnectionv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/enterprisevpnconnectionv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5 opentelekomcloud_enterprise_vpn_connection_v5}.
type EnterpriseVpnConnectionV5 interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreatedAt() *string
	CustomerGatewayId() *string
	SetCustomerGatewayId(val *string)
	CustomerGatewayIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableNqa() interface{}
	SetEnableNqa(val interface{})
	EnableNqaInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayId() *string
	SetGatewayId(val *string)
	GatewayIdInput() *string
	GatewayIp() *string
	SetGatewayIp(val *string)
	GatewayIpInput() *string
	HaRole() *string
	SetHaRole(val *string)
	HaRoleInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Ikepolicy() EnterpriseVpnConnectionV5IkepolicyOutputReference
	IkepolicyInput() *EnterpriseVpnConnectionV5Ikepolicy
	Ipsecpolicy() EnterpriseVpnConnectionV5IpsecpolicyOutputReference
	IpsecpolicyInput() *EnterpriseVpnConnectionV5Ipsecpolicy
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PeerSubnets() *[]*string
	SetPeerSubnets(val *[]*string)
	PeerSubnetsInput() *[]*string
	PolicyRules() EnterpriseVpnConnectionV5PolicyRulesList
	PolicyRulesInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Psk() *string
	SetPsk(val *string)
	PskInput() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	Status() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() EnterpriseVpnConnectionV5TimeoutsOutputReference
	TimeoutsInput() interface{}
	TunnelLocalAddress() *string
	SetTunnelLocalAddress(val *string)
	TunnelLocalAddressInput() *string
	TunnelPeerAddress() *string
	SetTunnelPeerAddress(val *string)
	TunnelPeerAddressInput() *string
	UpdatedAt() *string
	VpnType() *string
	SetVpnType(val *string)
	VpnTypeInput() *string
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
	PutIkepolicy(value *EnterpriseVpnConnectionV5Ikepolicy)
	PutIpsecpolicy(value *EnterpriseVpnConnectionV5Ipsecpolicy)
	PutPolicyRules(value interface{})
	PutTimeouts(value *EnterpriseVpnConnectionV5Timeouts)
	ResetEnableNqa()
	ResetHaRole()
	ResetId()
	ResetIkepolicy()
	ResetIpsecpolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeerSubnets()
	ResetPolicyRules()
	ResetTags()
	ResetTimeouts()
	ResetTunnelLocalAddress()
	ResetTunnelPeerAddress()
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

// The jsii proxy struct for EnterpriseVpnConnectionV5
type jsiiProxy_EnterpriseVpnConnectionV5 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) CustomerGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) CustomerGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) EnableNqa() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNqa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) EnableNqaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNqaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) GatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) GatewayIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) GatewayIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) HaRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) HaRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Ikepolicy() EnterpriseVpnConnectionV5IkepolicyOutputReference {
	var returns EnterpriseVpnConnectionV5IkepolicyOutputReference
	_jsii_.Get(
		j,
		"ikepolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) IkepolicyInput() *EnterpriseVpnConnectionV5Ikepolicy {
	var returns *EnterpriseVpnConnectionV5Ikepolicy
	_jsii_.Get(
		j,
		"ikepolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Ipsecpolicy() EnterpriseVpnConnectionV5IpsecpolicyOutputReference {
	var returns EnterpriseVpnConnectionV5IpsecpolicyOutputReference
	_jsii_.Get(
		j,
		"ipsecpolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) IpsecpolicyInput() *EnterpriseVpnConnectionV5Ipsecpolicy {
	var returns *EnterpriseVpnConnectionV5Ipsecpolicy
	_jsii_.Get(
		j,
		"ipsecpolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) PeerSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) PeerSubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerSubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) PolicyRules() EnterpriseVpnConnectionV5PolicyRulesList {
	var returns EnterpriseVpnConnectionV5PolicyRulesList
	_jsii_.Get(
		j,
		"policyRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) PolicyRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Psk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"psk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) PskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) Timeouts() EnterpriseVpnConnectionV5TimeoutsOutputReference {
	var returns EnterpriseVpnConnectionV5TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) TunnelLocalAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelLocalAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) TunnelLocalAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelLocalAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) TunnelPeerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelPeerAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) TunnelPeerAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelPeerAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) VpnType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5) VpnTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5 opentelekomcloud_enterprise_vpn_connection_v5} Resource.
func NewEnterpriseVpnConnectionV5(scope constructs.Construct, id *string, config *EnterpriseVpnConnectionV5Config) EnterpriseVpnConnectionV5 {
	_init_.Initialize()

	if err := validateNewEnterpriseVpnConnectionV5Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnterpriseVpnConnectionV5{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnConnectionV5.EnterpriseVpnConnectionV5",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5 opentelekomcloud_enterprise_vpn_connection_v5} Resource.
func NewEnterpriseVpnConnectionV5_Override(e EnterpriseVpnConnectionV5, scope constructs.Construct, id *string, config *EnterpriseVpnConnectionV5Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnConnectionV5.EnterpriseVpnConnectionV5",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetCustomerGatewayId(val *string) {
	if err := j.validateSetCustomerGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerGatewayId",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetEnableNqa(val interface{}) {
	if err := j.validateSetEnableNqaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNqa",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetGatewayId(val *string) {
	if err := j.validateSetGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayId",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetGatewayIp(val *string) {
	if err := j.validateSetGatewayIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayIp",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetHaRole(val *string) {
	if err := j.validateSetHaRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haRole",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetPeerSubnets(val *[]*string) {
	if err := j.validateSetPeerSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerSubnets",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetPsk(val *string) {
	if err := j.validateSetPskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"psk",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetTunnelLocalAddress(val *string) {
	if err := j.validateSetTunnelLocalAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelLocalAddress",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetTunnelPeerAddress(val *string) {
	if err := j.validateSetTunnelPeerAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelPeerAddress",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnConnectionV5)SetVpnType(val *string) {
	if err := j.validateSetVpnTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnType",
		val,
	)
}

// Generates CDKTF code for importing a EnterpriseVpnConnectionV5 resource upon running "cdktf plan <stack-name>".
func EnterpriseVpnConnectionV5_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEnterpriseVpnConnectionV5_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnConnectionV5.EnterpriseVpnConnectionV5",
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
func EnterpriseVpnConnectionV5_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseVpnConnectionV5_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnConnectionV5.EnterpriseVpnConnectionV5",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EnterpriseVpnConnectionV5_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseVpnConnectionV5_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnConnectionV5.EnterpriseVpnConnectionV5",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EnterpriseVpnConnectionV5_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseVpnConnectionV5_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnConnectionV5.EnterpriseVpnConnectionV5",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EnterpriseVpnConnectionV5_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnConnectionV5.EnterpriseVpnConnectionV5",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) PutIkepolicy(value *EnterpriseVpnConnectionV5Ikepolicy) {
	if err := e.validatePutIkepolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIkepolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) PutIpsecpolicy(value *EnterpriseVpnConnectionV5Ipsecpolicy) {
	if err := e.validatePutIpsecpolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIpsecpolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) PutPolicyRules(value interface{}) {
	if err := e.validatePutPolicyRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPolicyRules",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) PutTimeouts(value *EnterpriseVpnConnectionV5Timeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ResetEnableNqa() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableNqa",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ResetHaRole() {
	_jsii_.InvokeVoid(
		e,
		"resetHaRole",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ResetIkepolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetIkepolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ResetIpsecpolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetIpsecpolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ResetPeerSubnets() {
	_jsii_.InvokeVoid(
		e,
		"resetPeerSubnets",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ResetPolicyRules() {
	_jsii_.InvokeVoid(
		e,
		"resetPolicyRules",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ResetTunnelLocalAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetTunnelLocalAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ResetTunnelPeerAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetTunnelPeerAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnConnectionV5) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

