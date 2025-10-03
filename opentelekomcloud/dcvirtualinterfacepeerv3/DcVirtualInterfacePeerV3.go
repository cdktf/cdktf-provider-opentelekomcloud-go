// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcvirtualinterfacepeerv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dcvirtualinterfacepeerv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/dc_virtual_interface_peer_v3 opentelekomcloud_dc_virtual_interface_peer_v3}.
type DcVirtualInterfacePeerV3 interface {
	cdktf.TerraformResource
	AddressFamily() *string
	SetAddressFamily(val *string)
	AddressFamilyInput() *string
	BgpAsn() *float64
	SetBgpAsn(val *float64)
	BgpAsnInput() *float64
	BgpMd5() *string
	SetBgpMd5(val *string)
	BgpMd5Input() *string
	BgpRouteLimit() *float64
	BgpStatus() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DeviceId() *string
	EnableBfd() cdktf.IResolvable
	EnableNqa() cdktf.IResolvable
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalGatewayIp() *string
	SetLocalGatewayIp(val *string)
	LocalGatewayIpInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProjectId() *string
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
	ReceiveRouteNum() *float64
	Region() *string
	RemoteEpGroup() *[]*string
	SetRemoteEpGroup(val *[]*string)
	RemoteEpGroupInput() *[]*string
	RemoteGatewayIp() *string
	SetRemoteGatewayIp(val *string)
	RemoteGatewayIpInput() *string
	RouteMode() *string
	SetRouteMode(val *string)
	RouteModeInput() *string
	ServiceEpGroup() *[]*string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VifId() *string
	SetVifId(val *string)
	VifIdInput() *string
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
	ResetBgpAsn()
	ResetBgpMd5()
	ResetDescription()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRouteMode()
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

// The jsii proxy struct for DcVirtualInterfacePeerV3
type jsiiProxy_DcVirtualInterfacePeerV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) AddressFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) AddressFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) BgpAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgpAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) BgpAsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgpAsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) BgpMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) BgpMd5Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpMd5Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) BgpRouteLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgpRouteLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) BgpStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) DeviceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) EnableBfd() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableBfd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) EnableNqa() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableNqa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) LocalGatewayIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) LocalGatewayIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) ReceiveRouteNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"receiveRouteNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) RemoteEpGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteEpGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) RemoteEpGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteEpGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) RemoteGatewayIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteGatewayIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) RemoteGatewayIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteGatewayIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) RouteMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) RouteModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) ServiceEpGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceEpGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) VifId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vifId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3) VifIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vifIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/dc_virtual_interface_peer_v3 opentelekomcloud_dc_virtual_interface_peer_v3} Resource.
func NewDcVirtualInterfacePeerV3(scope constructs.Construct, id *string, config *DcVirtualInterfacePeerV3Config) DcVirtualInterfacePeerV3 {
	_init_.Initialize()

	if err := validateNewDcVirtualInterfacePeerV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DcVirtualInterfacePeerV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfacePeerV3.DcVirtualInterfacePeerV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/dc_virtual_interface_peer_v3 opentelekomcloud_dc_virtual_interface_peer_v3} Resource.
func NewDcVirtualInterfacePeerV3_Override(d DcVirtualInterfacePeerV3, scope constructs.Construct, id *string, config *DcVirtualInterfacePeerV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfacePeerV3.DcVirtualInterfacePeerV3",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetAddressFamily(val *string) {
	if err := j.validateSetAddressFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressFamily",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetBgpAsn(val *float64) {
	if err := j.validateSetBgpAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpAsn",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetBgpMd5(val *string) {
	if err := j.validateSetBgpMd5Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpMd5",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetLocalGatewayIp(val *string) {
	if err := j.validateSetLocalGatewayIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localGatewayIp",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetRemoteEpGroup(val *[]*string) {
	if err := j.validateSetRemoteEpGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteEpGroup",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetRemoteGatewayIp(val *string) {
	if err := j.validateSetRemoteGatewayIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteGatewayIp",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetRouteMode(val *string) {
	if err := j.validateSetRouteModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeMode",
		val,
	)
}

func (j *jsiiProxy_DcVirtualInterfacePeerV3)SetVifId(val *string) {
	if err := j.validateSetVifIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vifId",
		val,
	)
}

// Generates CDKTF code for importing a DcVirtualInterfacePeerV3 resource upon running "cdktf plan <stack-name>".
func DcVirtualInterfacePeerV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDcVirtualInterfacePeerV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfacePeerV3.DcVirtualInterfacePeerV3",
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
func DcVirtualInterfacePeerV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDcVirtualInterfacePeerV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfacePeerV3.DcVirtualInterfacePeerV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DcVirtualInterfacePeerV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDcVirtualInterfacePeerV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfacePeerV3.DcVirtualInterfacePeerV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DcVirtualInterfacePeerV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDcVirtualInterfacePeerV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfacePeerV3.DcVirtualInterfacePeerV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DcVirtualInterfacePeerV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dcVirtualInterfacePeerV3.DcVirtualInterfacePeerV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DcVirtualInterfacePeerV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DcVirtualInterfacePeerV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DcVirtualInterfacePeerV3) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DcVirtualInterfacePeerV3) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DcVirtualInterfacePeerV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DcVirtualInterfacePeerV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DcVirtualInterfacePeerV3) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DcVirtualInterfacePeerV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DcVirtualInterfacePeerV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DcVirtualInterfacePeerV3) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) ResetBgpAsn() {
	_jsii_.InvokeVoid(
		d,
		"resetBgpAsn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) ResetBgpMd5() {
	_jsii_.InvokeVoid(
		d,
		"resetBgpMd5",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) ResetRouteMode() {
	_jsii_.InvokeVoid(
		d,
		"resetRouteMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DcVirtualInterfacePeerV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

