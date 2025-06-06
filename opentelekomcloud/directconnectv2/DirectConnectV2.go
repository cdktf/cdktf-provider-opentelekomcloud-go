// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directconnectv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/directconnectv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2 opentelekomcloud_direct_connect_v2}.
type DirectConnectV2 interface {
	cdktf.TerraformResource
	AdminStateUp() interface{}
	SetAdminStateUp(val interface{})
	AdminStateUpInput() interface{}
	Applicant() *string
	ApplyTime() *string
	Bandwidth() *float64
	SetBandwidth(val *float64)
	BandwidthInput() *float64
	BuildingLineProductId() *string
	CableLabel() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChargeMode() *string
	SetChargeMode(val *string)
	ChargeModeInput() *string
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
	CreateTime() *string
	DeleteTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DeviceId() *string
	SetDeviceId(val *string)
	DeviceIdInput() *string
	Email() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostingId() *string
	SetHostingId(val *string)
	HostingIdInput() *string
	Id() *string
	InterfaceName() *string
	SetInterfaceName(val *string)
	InterfaceNameInput() *string
	LagId() *string
	LastOnestopProductId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Mobile() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OnestopProductId() *string
	OrderId() *string
	SetOrderId(val *string)
	OrderIdInput() *string
	PeerLocation() *string
	SetPeerLocation(val *string)
	PeerLocationInput() *string
	PeerPortType() *string
	PeerProvider() *string
	PeriodNum() *float64
	PeriodType() *float64
	PortType() *string
	SetPortType(val *string)
	PortTypeInput() *string
	ProductId() *string
	SetProductId(val *string)
	ProductIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderName() *string
	SetProviderName(val *string)
	ProviderNameInput() *string
	ProviderStatus() *string
	SetProviderStatus(val *string)
	ProviderStatusInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Reason() *string
	RedundantId() *string
	SetRedundantId(val *string)
	RedundantIdInput() *string
	RegionId() *string
	ServiceKey() *string
	SpecCode() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DirectConnectV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VgwType() *string
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
	PutTimeouts(value *DirectConnectV2Timeouts)
	ResetAdminStateUp()
	ResetBandwidth()
	ResetChargeMode()
	ResetDescription()
	ResetDeviceId()
	ResetHostingId()
	ResetInterfaceName()
	ResetLocation()
	ResetName()
	ResetOrderId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeerLocation()
	ResetPortType()
	ResetProductId()
	ResetProviderStatus()
	ResetRedundantId()
	ResetStatus()
	ResetTenantId()
	ResetTimeouts()
	ResetType()
	ResetVlan()
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

// The jsii proxy struct for DirectConnectV2
type jsiiProxy_DirectConnectV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DirectConnectV2) AdminStateUp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminStateUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) AdminStateUpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminStateUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Applicant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) ApplyTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applyTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Bandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) BandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) BuildingLineProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildingLineProductId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) CableLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cableLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) ChargeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) ChargeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) DeleteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) DeviceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) DeviceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) HostingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) HostingIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostingIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) InterfaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) InterfaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) LagId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lagId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) LastOnestopProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastOnestopProductId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Mobile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) OnestopProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onestopProductId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) OrderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) OrderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) PeerLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) PeerLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) PeerPortType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerPortType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) PeerProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) PeriodNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) PeriodType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) PortType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) PortTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) ProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) ProviderStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) ProviderStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Reason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) RedundantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redundantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) RedundantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redundantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) RegionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) ServiceKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) SpecCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Timeouts() DirectConnectV2TimeoutsOutputReference {
	var returns DirectConnectV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) VgwType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vgwType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) Vlan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectConnectV2) VlanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2 opentelekomcloud_direct_connect_v2} Resource.
func NewDirectConnectV2(scope constructs.Construct, id *string, config *DirectConnectV2Config) DirectConnectV2 {
	_init_.Initialize()

	if err := validateNewDirectConnectV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DirectConnectV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.directConnectV2.DirectConnectV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2 opentelekomcloud_direct_connect_v2} Resource.
func NewDirectConnectV2_Override(d DirectConnectV2, scope constructs.Construct, id *string, config *DirectConnectV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.directConnectV2.DirectConnectV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetAdminStateUp(val interface{}) {
	if err := j.validateSetAdminStateUpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminStateUp",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetBandwidth(val *float64) {
	if err := j.validateSetBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidth",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetChargeMode(val *string) {
	if err := j.validateSetChargeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chargeMode",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetDeviceId(val *string) {
	if err := j.validateSetDeviceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceId",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetHostingId(val *string) {
	if err := j.validateSetHostingIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostingId",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetInterfaceName(val *string) {
	if err := j.validateSetInterfaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceName",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetOrderId(val *string) {
	if err := j.validateSetOrderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orderId",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetPeerLocation(val *string) {
	if err := j.validateSetPeerLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerLocation",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetPortType(val *string) {
	if err := j.validateSetPortTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portType",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetProductId(val *string) {
	if err := j.validateSetProductIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetProviderName(val *string) {
	if err := j.validateSetProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerName",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetProviderStatus(val *string) {
	if err := j.validateSetProviderStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerStatus",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetRedundantId(val *string) {
	if err := j.validateSetRedundantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redundantId",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DirectConnectV2)SetVlan(val *float64) {
	if err := j.validateSetVlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vlan",
		val,
	)
}

// Generates CDKTF code for importing a DirectConnectV2 resource upon running "cdktf plan <stack-name>".
func DirectConnectV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDirectConnectV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.directConnectV2.DirectConnectV2",
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
func DirectConnectV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectConnectV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.directConnectV2.DirectConnectV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DirectConnectV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectConnectV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.directConnectV2.DirectConnectV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DirectConnectV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectConnectV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.directConnectV2.DirectConnectV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DirectConnectV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.directConnectV2.DirectConnectV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DirectConnectV2) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DirectConnectV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DirectConnectV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DirectConnectV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DirectConnectV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DirectConnectV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DirectConnectV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DirectConnectV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DirectConnectV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DirectConnectV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DirectConnectV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DirectConnectV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectConnectV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DirectConnectV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DirectConnectV2) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DirectConnectV2) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DirectConnectV2) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DirectConnectV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DirectConnectV2) PutTimeouts(value *DirectConnectV2Timeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetAdminStateUp() {
	_jsii_.InvokeVoid(
		d,
		"resetAdminStateUp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetBandwidth() {
	_jsii_.InvokeVoid(
		d,
		"resetBandwidth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetChargeMode() {
	_jsii_.InvokeVoid(
		d,
		"resetChargeMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetDeviceId() {
	_jsii_.InvokeVoid(
		d,
		"resetDeviceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetHostingId() {
	_jsii_.InvokeVoid(
		d,
		"resetHostingId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetInterfaceName() {
	_jsii_.InvokeVoid(
		d,
		"resetInterfaceName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetOrderId() {
	_jsii_.InvokeVoid(
		d,
		"resetOrderId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetPeerLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetPeerLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetPortType() {
	_jsii_.InvokeVoid(
		d,
		"resetPortType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetProductId() {
	_jsii_.InvokeVoid(
		d,
		"resetProductId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetProviderStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetProviderStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetRedundantId() {
	_jsii_.InvokeVoid(
		d,
		"resetRedundantId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetTenantId() {
	_jsii_.InvokeVoid(
		d,
		"resetTenantId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) ResetVlan() {
	_jsii_.InvokeVoid(
		d,
		"resetVlan",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectConnectV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectConnectV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectConnectV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectConnectV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectConnectV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectConnectV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

