// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisevpngatewayv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/enterprisevpngatewayv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/enterprise_vpn_gateway_v5 opentelekomcloud_enterprise_vpn_gateway_v5}.
type EnterpriseVpnGatewayV5 interface {
	cdktf.TerraformResource
	AccessPrivateIp1() *string
	SetAccessPrivateIp1(val *string)
	AccessPrivateIp1Input() *string
	AccessPrivateIp2() *string
	SetAccessPrivateIp2(val *string)
	AccessPrivateIp2Input() *string
	AccessSubnetId() *string
	SetAccessSubnetId(val *string)
	AccessSubnetIdInput() *string
	AccessVpcId() *string
	SetAccessVpcId(val *string)
	AccessVpcIdInput() *string
	Asn() *float64
	SetAsn(val *float64)
	AsnInput() *float64
	AttachmentType() *string
	SetAttachmentType(val *string)
	AttachmentTypeInput() *string
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectSubnet() *string
	SetConnectSubnet(val *string)
	ConnectSubnetInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	DeleteEip() interface{}
	SetDeleteEip(val interface{})
	DeleteEipInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Eip1() EnterpriseVpnGatewayV5Eip1OutputReference
	Eip1Input() *EnterpriseVpnGatewayV5Eip1
	Eip2() EnterpriseVpnGatewayV5Eip2OutputReference
	Eip2Input() *EnterpriseVpnGatewayV5Eip2
	ErAttachmentId() *string
	ErId() *string
	SetErId(val *string)
	ErIdInput() *string
	Flavor() *string
	SetFlavor(val *string)
	FlavorInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HaMode() *string
	SetHaMode(val *string)
	HaModeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalSubnets() *[]*string
	SetLocalSubnets(val *[]*string)
	LocalSubnetsInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
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
	Timeouts() EnterpriseVpnGatewayV5TimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdatedAt() *string
	UsedConnectionGroup() *float64
	UsedConnectionNumber() *float64
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
	PutEip1(value *EnterpriseVpnGatewayV5Eip1)
	PutEip2(value *EnterpriseVpnGatewayV5Eip2)
	PutTimeouts(value *EnterpriseVpnGatewayV5Timeouts)
	ResetAccessPrivateIp1()
	ResetAccessPrivateIp2()
	ResetAccessSubnetId()
	ResetAccessVpcId()
	ResetAsn()
	ResetAttachmentType()
	ResetConnectSubnet()
	ResetDeleteEip()
	ResetEip1()
	ResetEip2()
	ResetErId()
	ResetFlavor()
	ResetHaMode()
	ResetId()
	ResetLocalSubnets()
	ResetNetworkType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTimeouts()
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

// The jsii proxy struct for EnterpriseVpnGatewayV5
type jsiiProxy_EnterpriseVpnGatewayV5 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AccessPrivateIp1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPrivateIp1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AccessPrivateIp1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPrivateIp1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AccessPrivateIp2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPrivateIp2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AccessPrivateIp2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPrivateIp2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AccessSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AccessSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AccessVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AccessVpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Asn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AttachmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) ConnectSubnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectSubnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) ConnectSubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectSubnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) DeleteEip() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteEip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) DeleteEipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteEipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Eip1() EnterpriseVpnGatewayV5Eip1OutputReference {
	var returns EnterpriseVpnGatewayV5Eip1OutputReference
	_jsii_.Get(
		j,
		"eip1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Eip1Input() *EnterpriseVpnGatewayV5Eip1 {
	var returns *EnterpriseVpnGatewayV5Eip1
	_jsii_.Get(
		j,
		"eip1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Eip2() EnterpriseVpnGatewayV5Eip2OutputReference {
	var returns EnterpriseVpnGatewayV5Eip2OutputReference
	_jsii_.Get(
		j,
		"eip2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Eip2Input() *EnterpriseVpnGatewayV5Eip2 {
	var returns *EnterpriseVpnGatewayV5Eip2
	_jsii_.Get(
		j,
		"eip2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) ErAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"erAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) ErId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"erId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) ErIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"erIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Flavor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) FlavorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) HaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) HaModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) LocalSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) LocalSubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localSubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) Timeouts() EnterpriseVpnGatewayV5TimeoutsOutputReference {
	var returns EnterpriseVpnGatewayV5TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) UsedConnectionGroup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedConnectionGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) UsedConnectionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedConnectionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/enterprise_vpn_gateway_v5 opentelekomcloud_enterprise_vpn_gateway_v5} Resource.
func NewEnterpriseVpnGatewayV5(scope constructs.Construct, id *string, config *EnterpriseVpnGatewayV5Config) EnterpriseVpnGatewayV5 {
	_init_.Initialize()

	if err := validateNewEnterpriseVpnGatewayV5Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnterpriseVpnGatewayV5{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnGatewayV5.EnterpriseVpnGatewayV5",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/enterprise_vpn_gateway_v5 opentelekomcloud_enterprise_vpn_gateway_v5} Resource.
func NewEnterpriseVpnGatewayV5_Override(e EnterpriseVpnGatewayV5, scope constructs.Construct, id *string, config *EnterpriseVpnGatewayV5Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnGatewayV5.EnterpriseVpnGatewayV5",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetAccessPrivateIp1(val *string) {
	if err := j.validateSetAccessPrivateIp1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPrivateIp1",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetAccessPrivateIp2(val *string) {
	if err := j.validateSetAccessPrivateIp2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPrivateIp2",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetAccessSubnetId(val *string) {
	if err := j.validateSetAccessSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessSubnetId",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetAccessVpcId(val *string) {
	if err := j.validateSetAccessVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessVpcId",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetAsn(val *float64) {
	if err := j.validateSetAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asn",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetAttachmentType(val *string) {
	if err := j.validateSetAttachmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentType",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetConnectSubnet(val *string) {
	if err := j.validateSetConnectSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectSubnet",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetDeleteEip(val interface{}) {
	if err := j.validateSetDeleteEipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteEip",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetErId(val *string) {
	if err := j.validateSetErIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"erId",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetFlavor(val *string) {
	if err := j.validateSetFlavorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flavor",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetHaMode(val *string) {
	if err := j.validateSetHaModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haMode",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetLocalSubnets(val *[]*string) {
	if err := j.validateSetLocalSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSubnets",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EnterpriseVpnGatewayV5)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a EnterpriseVpnGatewayV5 resource upon running "cdktf plan <stack-name>".
func EnterpriseVpnGatewayV5_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEnterpriseVpnGatewayV5_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnGatewayV5.EnterpriseVpnGatewayV5",
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
func EnterpriseVpnGatewayV5_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseVpnGatewayV5_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnGatewayV5.EnterpriseVpnGatewayV5",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EnterpriseVpnGatewayV5_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseVpnGatewayV5_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnGatewayV5.EnterpriseVpnGatewayV5",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EnterpriseVpnGatewayV5_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseVpnGatewayV5_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnGatewayV5.EnterpriseVpnGatewayV5",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EnterpriseVpnGatewayV5_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.enterpriseVpnGatewayV5.EnterpriseVpnGatewayV5",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EnterpriseVpnGatewayV5) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) PutEip1(value *EnterpriseVpnGatewayV5Eip1) {
	if err := e.validatePutEip1Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEip1",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) PutEip2(value *EnterpriseVpnGatewayV5Eip2) {
	if err := e.validatePutEip2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEip2",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) PutTimeouts(value *EnterpriseVpnGatewayV5Timeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetAccessPrivateIp1() {
	_jsii_.InvokeVoid(
		e,
		"resetAccessPrivateIp1",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetAccessPrivateIp2() {
	_jsii_.InvokeVoid(
		e,
		"resetAccessPrivateIp2",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetAccessSubnetId() {
	_jsii_.InvokeVoid(
		e,
		"resetAccessSubnetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetAccessVpcId() {
	_jsii_.InvokeVoid(
		e,
		"resetAccessVpcId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetAsn() {
	_jsii_.InvokeVoid(
		e,
		"resetAsn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetAttachmentType() {
	_jsii_.InvokeVoid(
		e,
		"resetAttachmentType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetConnectSubnet() {
	_jsii_.InvokeVoid(
		e,
		"resetConnectSubnet",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetDeleteEip() {
	_jsii_.InvokeVoid(
		e,
		"resetDeleteEip",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetEip1() {
	_jsii_.InvokeVoid(
		e,
		"resetEip1",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetEip2() {
	_jsii_.InvokeVoid(
		e,
		"resetEip2",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetErId() {
	_jsii_.InvokeVoid(
		e,
		"resetErId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetFlavor() {
	_jsii_.InvokeVoid(
		e,
		"resetFlavor",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetHaMode() {
	_jsii_.InvokeVoid(
		e,
		"resetHaMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetLocalSubnets() {
	_jsii_.InvokeVoid(
		e,
		"resetLocalSubnets",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetNetworkType() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ResetVpcId() {
	_jsii_.InvokeVoid(
		e,
		"resetVpcId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseVpnGatewayV5) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

