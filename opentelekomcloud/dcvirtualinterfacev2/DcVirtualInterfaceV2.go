// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcvirtualinterfacev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/dcvirtualinterfacev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/dc_virtual_interface_v2 opentelekomcloud_dc_virtual_interface_v2}.
type DcVirtualInterfaceV2 interface {
	cdktf.TerraformResource
	Asn() *float64
	SetAsn(val *float64)
	AsnInput() *float64
	Bandwidth() *float64
	SetBandwidth(val *float64)
	BandwidthInput() *float64
	BgpMd5() *string
	SetBgpMd5(val *string)
	BgpMd5Input() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DirectConnectId() *string
	SetDirectConnectId(val *string)
	DirectConnectIdInput() *string
	EnableBfd() interface{}
	SetEnableBfd(val interface{})
	EnableBfdInput() interface{}
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	LagId() *string
	SetLagId(val *string)
	LagIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalGatewayV4Ip() *string
	SetLocalGatewayV4Ip(val *string)
	LocalGatewayV4IpInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
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
	RemoteEpGroup() DcVirtualInterfaceV2RemoteEpGroupOutputReference
	RemoteEpGroupId() *string
	RemoteEpGroupInput() *DcVirtualInterfaceV2RemoteEpGroup
	RemoteGatewayV4Ip() *string
	SetRemoteGatewayV4Ip(val *string)
	RemoteGatewayV4IpInput() *string
	RouteMode() *string
	SetRouteMode(val *string)
	RouteModeInput() *string
	ServiceType() *string
	SetServiceType(val *string)
	ServiceTypeInput() *string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VirtualGatewayId() *string
	SetVirtualGatewayId(val *string)
	VirtualGatewayIdInput() *string
	Vlan() *float64
	SetVlan(val *float64)
	VlanInput() *float64
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
	PutRemoteEpGroup(value *DcVirtualInterfaceV2RemoteEpGroup)
	ResetAsn()
	ResetBgpMd5()
	ResetDescription()
	ResetEnableBfd()
	ResetEnableNqa()
	ResetId()
	ResetLagId()
	ResetLocalGatewayV4Ip()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectId()
	ResetRemoteGatewayV4Ip()
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

// The jsii proxy struct for DcVirtualInterfaceV2
type jsiiProxy_DcVirtualInterfaceV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Asn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) AsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Bandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) BandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) BgpMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) BgpMd5Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpMd5Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) DirectConnectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directConnectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) DirectConnectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directConnectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) EnableBfd() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBfd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) EnableBfdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBfdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) EnableNqa() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNqa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) EnableNqaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNqaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) LagId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lagId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) LagIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lagIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) LocalGatewayV4Ip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayV4Ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) LocalGatewayV4IpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayV4IpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) RemoteEpGroup() DcVirtualInterfaceV2RemoteEpGroupOutputReference {
	var returns DcVirtualInterfaceV2RemoteEpGroupOutputReference
	_jsii_.Get(
		j,
		"remoteEpGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) RemoteEpGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteEpGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) RemoteEpGroupInput() *DcVirtualInterfaceV2RemoteEpGroup {
	var returns *DcVirtualInterfaceV2RemoteEpGroup
	_jsii_.Get(
		j,
		"remoteEpGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) RemoteGatewayV4Ip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteGatewayV4Ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) RemoteGatewayV4IpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteGatewayV4IpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) RouteMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) RouteModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) ServiceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) ServiceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) VirtualGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) VirtualGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) Vlan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfaceV2) VlanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/dc_virtual_interface_v2 opentelekomcloud_dc_virtual_interface_v2} Resource.
func NewDcVirtualInterfaceV2(scope constructs.Construct, id *string, config *DcVirtualInterfaceV2Config) DcVirtualInterfaceV2 {
	_init_.Initialize()

	if err := validateNewDcVirtualInterfaceV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DcVirtualInterfaceV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfaceV2.DcVirtualInterfaceV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/dc_virtual_interface_v2 opentelekomcloud_dc_virtual_interface_v2} Resource.
func NewDcVirtualInterfaceV2_Override(d DcVirtualInterfaceV2, scope constructs.Construct, id *string, config *DcVirtualInterfaceV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfaceV2.DcVirtualInterfaceV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetAsn(val *float64) {
	if err := j.validateSetAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asn",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetBandwidth(val *float64) {
	if err := j.validateSetBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidth",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetBgpMd5(val *string) {
	if err := j.validateSetBgpMd5Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpMd5",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetDirectConnectId(val *string) {
	if err := j.validateSetDirectConnectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directConnectId",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetEnableBfd(val interface{}) {
	if err := j.validateSetEnableBfdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBfd",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetEnableNqa(val interface{}) {
	if err := j.validateSetEnableNqaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNqa",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetLagId(val *string) {
	if err := j.validateSetLagIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lagId",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetLocalGatewayV4Ip(val *string) {
	if err := j.validateSetLocalGatewayV4IpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localGatewayV4Ip",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetRemoteGatewayV4Ip(val *string) {
	if err := j.validateSetRemoteGatewayV4IpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteGatewayV4Ip",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetRouteMode(val *string) {
	if err := j.validateSetRouteModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeMode",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetServiceType(val *string) {
	if err := j.validateSetServiceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceType",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetVirtualGatewayId(val *string) {
	if err := j.validateSetVirtualGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualGatewayId",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfaceV2)SetVlan(val *float64) {
	if err := j.validateSetVlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vlan",
		val,
	)
}

// Generates CDKTF code for importing a DcVirtualInterfaceV2 resource upon running "cdktf plan <stack-name>".
func DcVirtualInterfaceV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDcVirtualInterfaceV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfaceV2.DcVirtualInterfaceV2",
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
func DcVirtualInterfaceV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDcVirtualInterfaceV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfaceV2.DcVirtualInterfaceV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DcVirtualInterfaceV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDcVirtualInterfaceV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfaceV2.DcVirtualInterfaceV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DcVirtualInterfaceV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDcVirtualInterfaceV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfaceV2.DcVirtualInterfaceV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DcVirtualInterfaceV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfaceV2.DcVirtualInterfaceV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) PutRemoteEpGroup(value *DcVirtualInterfaceV2RemoteEpGroup) {
	if err := d.validatePutRemoteEpGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRemoteEpGroup",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ResetAsn() {
	_jsii_.InvokeVoid(
		d,
		"resetAsn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ResetBgpMd5() {
	_jsii_.InvokeVoid(
		d,
		"resetBgpMd5",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ResetEnableBfd() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableBfd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ResetEnableNqa() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableNqa",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ResetLagId() {
	_jsii_.InvokeVoid(
		d,
		"resetLagId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ResetLocalGatewayV4Ip() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalGatewayV4Ip",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ResetProjectId() {
	_jsii_.InvokeVoid(
		d,
		"resetProjectId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ResetRemoteGatewayV4Ip() {
	_jsii_.InvokeVoid(
		d,
		"resetRemoteGatewayV4Ip",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfaceV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfaceV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

